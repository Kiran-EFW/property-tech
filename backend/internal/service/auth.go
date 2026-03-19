package service

import (
	"context"
	"crypto/rand"
	"fmt"
	"math/big"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog/log"

	"github.com/proptech/backend/internal/config"
	"github.com/proptech/backend/internal/domain"
)

// TokenPair holds JWT access and refresh tokens returned after authentication.
type TokenPair struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresAt    int64  `json:"expires_at"`
}

// AuthRepository defines the database operations required by AuthService.
type AuthRepository interface {
	CreateUser(ctx context.Context, user *domain.User) error
	GetUserByPhone(ctx context.Context, phone string) (*domain.User, error)
	GetUserByID(ctx context.Context, id uuid.UUID) (*domain.User, error)
	UpdateUser(ctx context.Context, user *domain.User) error
}

// AuthService handles authentication and user management.
type AuthService struct {
	repo  AuthRepository
	redis *redis.Client
	cfg   *config.Config
}

// NewAuthService creates a new AuthService.
func NewAuthService(repo AuthRepository, redisClient *redis.Client, cfg *config.Config) *AuthService {
	return &AuthService{
		repo:  repo,
		redis: redisClient,
		cfg:   cfg,
	}
}

// Register creates a new user account.
func (s *AuthService) Register(ctx context.Context, phone, name, email string, role domain.UserRole) (*domain.User, error) {
	// Check if user already exists.
	existing, err := s.repo.GetUserByPhone(ctx, phone)
	if err == nil && existing != nil {
		return nil, fmt.Errorf("user with phone %s already exists", phone)
	}

	user := &domain.User{
		ID:        uuid.New(),
		Phone:     phone,
		Name:      name,
		Role:      role,
		IsActive:  true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	if email != "" {
		user.Email = &email
	}

	if err := s.repo.CreateUser(ctx, user); err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	log.Info().Str("user_id", user.ID.String()).Str("phone", phone).Msg("user registered")
	return user, nil
}

// RequestOTP generates a 6-digit OTP for the given phone number and stores it
// in Redis with a 5-minute TTL.
func (s *AuthService) RequestOTP(ctx context.Context, phone string) error {
	otp, err := generateOTP(6)
	if err != nil {
		return fmt.Errorf("failed to generate OTP: %w", err)
	}

	key := otpRedisKey(phone)
	if err := s.redis.Set(ctx, key, otp, 5*time.Minute).Err(); err != nil {
		return fmt.Errorf("failed to store OTP: %w", err)
	}

	// In production, send the OTP via SMS. For now, log it in non-prod environments.
	if !s.cfg.IsProd() {
		log.Info().Str("phone", phone).Str("otp", otp).Msg("OTP generated (dev mode)")
	}

	return nil
}

// VerifyOTP validates the OTP for the given phone number and returns a JWT token pair.
func (s *AuthService) VerifyOTP(ctx context.Context, phone, otp string) (*TokenPair, error) {
	key := otpRedisKey(phone)
	storedOTP, err := s.redis.Get(ctx, key).Result()
	if err == redis.Nil {
		return nil, fmt.Errorf("OTP expired or not found")
	}
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve OTP: %w", err)
	}

	if storedOTP != otp {
		return nil, fmt.Errorf("invalid OTP")
	}

	// Delete OTP after successful verification.
	s.redis.Del(ctx, key)

	// Look up or create the user.
	user, err := s.repo.GetUserByPhone(ctx, phone)
	if err != nil {
		return nil, fmt.Errorf("user not found; please register first")
	}

	tokens, err := s.generateTokenPair(user)
	if err != nil {
		return nil, fmt.Errorf("failed to generate tokens: %w", err)
	}

	log.Info().Str("user_id", user.ID.String()).Msg("user authenticated via OTP")
	return tokens, nil
}

// RefreshToken validates a refresh token and issues a new token pair.
func (s *AuthService) RefreshToken(ctx context.Context, refreshToken string) (*TokenPair, error) {
	claims := &jwt.RegisteredClaims{}
	token, err := jwt.ParseWithClaims(refreshToken, claims, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return []byte(s.cfg.JWTSecret), nil
	})
	if err != nil || !token.Valid {
		return nil, fmt.Errorf("invalid refresh token")
	}

	userID, err := uuid.Parse(claims.Subject)
	if err != nil {
		return nil, fmt.Errorf("invalid token subject")
	}

	user, err := s.repo.GetUserByID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("user not found: %w", err)
	}

	tokens, err := s.generateTokenPair(user)
	if err != nil {
		return nil, fmt.Errorf("failed to generate tokens: %w", err)
	}

	return tokens, nil
}

// GetUser returns a user by ID.
func (s *AuthService) GetUser(ctx context.Context, userID uuid.UUID) (*domain.User, error) {
	user, err := s.repo.GetUserByID(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("user not found: %w", err)
	}
	return user, nil
}

// UpdateProfile updates the user's name and email.
func (s *AuthService) UpdateProfile(ctx context.Context, userID uuid.UUID, name, email string) error {
	user, err := s.repo.GetUserByID(ctx, userID)
	if err != nil {
		return fmt.Errorf("user not found: %w", err)
	}

	if name != "" {
		user.Name = name
	}
	if email != "" {
		user.Email = &email
	}
	user.UpdatedAt = time.Now()

	if err := s.repo.UpdateUser(ctx, user); err != nil {
		return fmt.Errorf("failed to update profile: %w", err)
	}

	return nil
}

// CustomClaims extends RegisteredClaims to include the user's role.
type CustomClaims struct {
	jwt.RegisteredClaims
	Role string `json:"role"`
}

// generateTokenPair creates JWT access and refresh tokens for the given user.
func (s *AuthService) generateTokenPair(user *domain.User) (*TokenPair, error) {
	now := time.Now()
	accessExpiry := now.Add(s.cfg.JWTExpiry)

	accessClaims := CustomClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   user.ID.String(),
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(accessExpiry),
			Issuer:    "proptech-api",
		},
		Role: string(user.Role),
	}
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	accessTokenStr, err := accessToken.SignedString([]byte(s.cfg.JWTSecret))
	if err != nil {
		return nil, fmt.Errorf("failed to sign access token: %w", err)
	}

	// Refresh token has a longer expiry (7 days).
	refreshExpiry := now.Add(7 * 24 * time.Hour)
	refreshClaims := CustomClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   user.ID.String(),
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(refreshExpiry),
			Issuer:    "proptech-api",
			ID:        uuid.New().String(), // unique jti for refresh tokens
		},
		Role: string(user.Role),
	}
	refreshTokenObj := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refreshTokenStr, err := refreshTokenObj.SignedString([]byte(s.cfg.JWTSecret))
	if err != nil {
		return nil, fmt.Errorf("failed to sign refresh token: %w", err)
	}

	return &TokenPair{
		AccessToken:  accessTokenStr,
		RefreshToken: refreshTokenStr,
		ExpiresAt:    accessExpiry.Unix(),
	}, nil
}

// generateOTP generates a cryptographically random numeric OTP of the given length.
func generateOTP(length int) (string, error) {
	otp := ""
	for i := 0; i < length; i++ {
		n, err := rand.Int(rand.Reader, big.NewInt(10))
		if err != nil {
			return "", err
		}
		otp += fmt.Sprintf("%d", n.Int64())
	}
	return otp, nil
}

// otpRedisKey returns the Redis key for storing an OTP.
func otpRedisKey(phone string) string {
	return fmt.Sprintf("otp:%s", phone)
}
