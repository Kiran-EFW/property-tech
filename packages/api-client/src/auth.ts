import type { ApiClient } from './client';
import type {
  User,
  RegisterData,
  LoginResponse,
  AuthTokens,
  ProfileUpdateData,
} from './types';

/**
 * API client for authentication endpoints.
 *
 * Auth flow:
 *   1. User submits phone number -> OTP is sent via SMS
 *   2. User submits phone + OTP -> receives access + refresh tokens
 *   3. Access token is sent with every authenticated request
 *   4. When access token expires, use refresh token to get new pair
 *   5. On logout, clear tokens client-side
 *
 * The API client instance handles token storage internally.
 * After login/refresh, tokens are automatically set on the client.
 */
export class AuthAPI {
  constructor(private client: ApiClient) {}

  /**
   * Registers a new user account.
   * Triggers an OTP to the provided phone number for verification.
   *
   * POST /auth/register
   */
  async register(data: RegisterData): Promise<User> {
    return this.client.post<User>('/auth/register', data);
  }

  /**
   * Requests an OTP for login.
   * Call this first, then call login() with the OTP.
   *
   * POST /auth/login
   */
  async requestOTP(phone: string): Promise<{ message: string }> {
    return this.client.post<{ message: string }>('/auth/login', {
      phone,
    });
  }

  /**
   * Logs in with phone number and OTP.
   * Returns user data and auth tokens.
   * Automatically sets tokens on the API client instance.
   *
   * POST /auth/verify
   */
  async login(phone: string, otp: string): Promise<LoginResponse> {
    const response = await this.client.post<LoginResponse>('/auth/verify', {
      phone,
      otp,
    });

    // Automatically set tokens on the client
    this.client.setTokens(
      response.tokens.accessToken,
      response.tokens.refreshToken,
    );

    return response;
  }

  /**
   * Refreshes the access token using the stored refresh token.
   * Automatically updates tokens on the API client instance.
   *
   * POST /auth/refresh
   */
  async refreshToken(): Promise<AuthTokens> {
    const currentRefreshToken = this.client.getRefreshToken();

    if (!currentRefreshToken) {
      throw new Error('No refresh token available. Please log in again.');
    }

    const tokens = await this.client.post<AuthTokens>('/auth/refresh', {
      refresh_token: currentRefreshToken,
    });

    // Update tokens on the client
    this.client.setTokens(tokens.accessToken, tokens.refreshToken);

    return tokens;
  }

  /**
   * Gets the currently authenticated user's profile.
   * Requires a valid access token.
   *
   * GET /auth/me
   */
  async getMe(): Promise<User> {
    return this.client.get<User>('/auth/me');
  }

  /**
   * Updates the current user's profile.
   *
   * PUT /auth/profile
   */
  async updateProfile(data: ProfileUpdateData): Promise<User> {
    return this.client.put<User>('/auth/profile', data);
  }

  /**
   * Logs out the current user.
   * Clears tokens from the API client.
   * Optionally calls the server to invalidate the refresh token.
   */
  async logout(): Promise<void> {
    try {
      // Try to invalidate the refresh token server-side
      await this.client.post('/auth/logout');
    } catch {
      // Ignore errors — we still clear tokens locally
    } finally {
      this.client.clearTokens();
    }
  }
}
