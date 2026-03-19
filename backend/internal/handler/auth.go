package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

	"github.com/proptech/backend/internal/domain"
	"github.com/proptech/backend/internal/middleware"
	"github.com/proptech/backend/internal/service"
	"github.com/proptech/backend/pkg/response"
	"github.com/proptech/backend/pkg/validator"
)

// AuthHandler handles authentication HTTP endpoints.
type AuthHandler struct {
	svc *service.AuthService
}

// NewAuthHandler creates a new AuthHandler.
func NewAuthHandler(svc *service.AuthService) *AuthHandler {
	return &AuthHandler{svc: svc}
}

// registerRequest is the request body for POST /auth/register.
type registerRequest struct {
	Phone string `json:"phone" validate:"required"`
	Name  string `json:"name" validate:"required"`
	Email string `json:"email"`
	Role  string `json:"role" validate:"required,oneof=investor agent builder"`
}

// Register handles POST /auth/register.
func (h *AuthHandler) Register(c *fiber.Ctx) error {
	var req registerRequest
	if err := c.BodyParser(&req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "invalid_body", "invalid request body")
	}

	if err := validator.Validate(req); err != nil {
		return response.ValidationError(c, validator.FormatValidationErrors(err))
	}

	if err := validator.ValidatePhone(req.Phone); err != nil {
		return response.ValidationError(c, validator.FormatValidationErrors(err))
	}

	user, err := h.svc.Register(c.Context(), req.Phone, req.Name, req.Email, domain.UserRole(req.Role))
	if err != nil {
		return response.Error(c, fiber.StatusConflict, "registration_failed", err.Error())
	}

	return response.Created(c, user)
}

// loginRequest is the request body for POST /auth/login.
type loginRequest struct {
	Phone string `json:"phone" validate:"required"`
}

// RequestOTP handles POST /auth/login.
func (h *AuthHandler) RequestOTP(c *fiber.Ctx) error {
	var req loginRequest
	if err := c.BodyParser(&req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "invalid_body", "invalid request body")
	}

	if err := validator.ValidatePhone(req.Phone); err != nil {
		return response.ValidationError(c, validator.FormatValidationErrors(err))
	}

	if err := h.svc.RequestOTP(c.Context(), req.Phone); err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "otp_failed", err.Error())
	}

	return response.Success(c, fiber.Map{"message": "OTP sent successfully"})
}

// verifyOTPRequest is the request body for POST /auth/verify.
type verifyOTPRequest struct {
	Phone string `json:"phone" validate:"required"`
	OTP   string `json:"otp" validate:"required,len=6"`
}

// VerifyOTP handles POST /auth/verify.
func (h *AuthHandler) VerifyOTP(c *fiber.Ctx) error {
	var req verifyOTPRequest
	if err := c.BodyParser(&req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "invalid_body", "invalid request body")
	}

	if err := validator.Validate(req); err != nil {
		return response.ValidationError(c, validator.FormatValidationErrors(err))
	}

	tokens, err := h.svc.VerifyOTP(c.Context(), req.Phone, req.OTP)
	if err != nil {
		return response.Error(c, fiber.StatusUnauthorized, "verification_failed", err.Error())
	}

	return response.Success(c, tokens)
}

// refreshTokenRequest is the request body for POST /auth/refresh.
type refreshTokenRequest struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
}

// RefreshToken handles POST /auth/refresh.
func (h *AuthHandler) RefreshToken(c *fiber.Ctx) error {
	var req refreshTokenRequest
	if err := c.BodyParser(&req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "invalid_body", "invalid request body")
	}

	if err := validator.Validate(req); err != nil {
		return response.ValidationError(c, validator.FormatValidationErrors(err))
	}

	tokens, err := h.svc.RefreshToken(c.Context(), req.RefreshToken)
	if err != nil {
		return response.Error(c, fiber.StatusUnauthorized, "refresh_failed", err.Error())
	}

	return response.Success(c, tokens)
}

// GetMe handles GET /auth/me.
func (h *AuthHandler) GetMe(c *fiber.Ctx) error {
	userID := middleware.GetUserID(c)
	if userID == uuid.Nil {
		return response.Unauthorized(c)
	}

	user, err := h.svc.GetUser(c.Context(), userID)
	if err != nil {
		return response.NotFound(c, "user")
	}

	return response.Success(c, user)
}

// updateProfileRequest is the request body for PUT /auth/profile.
type updateProfileRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

// UpdateProfile handles PUT /auth/profile.
func (h *AuthHandler) UpdateProfile(c *fiber.Ctx) error {
	userID := middleware.GetUserID(c)
	if userID == uuid.Nil {
		return response.Unauthorized(c)
	}

	var req updateProfileRequest
	if err := c.BodyParser(&req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "invalid_body", "invalid request body")
	}

	if req.Email != "" {
		if err := validator.ValidateEmail(req.Email); err != nil {
			return response.ValidationError(c, validator.FormatValidationErrors(err))
		}
	}

	if err := h.svc.UpdateProfile(c.Context(), userID, req.Name, req.Email); err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "update_failed", err.Error())
	}

	return response.Success(c, fiber.Map{"message": "profile updated"})
}
