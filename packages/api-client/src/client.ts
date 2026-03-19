import type { ApiError, ApiResponse } from './types';

// ---------------------------------------------------------------------------
// Configuration
// ---------------------------------------------------------------------------

export interface ApiClientConfig {
  /** Base URL of the API (e.g., "https://api.example.com/v1") */
  baseURL: string;

  /** Default request timeout in milliseconds. Defaults to 30000 (30s) */
  timeout?: number;

  /** Custom headers to include in every request */
  defaultHeaders?: Record<string, string>;

  /** Called when a 401 is received. Use this to trigger re-auth or redirect */
  onUnauthorized?: () => void;
}

// ---------------------------------------------------------------------------
// API Client
// ---------------------------------------------------------------------------

/**
 * Base HTTP client for the PropTech Go backend API.
 *
 * Uses the Fetch API (no axios dependency). Handles:
 *   - Auth token management (get/set/clear)
 *   - Automatic Authorization header injection
 *   - Typed JSON responses
 *   - Structured error handling
 *   - Request timeouts via AbortController
 */
export class ApiClient {
  private baseURL: string;
  private timeout: number;
  private defaultHeaders: Record<string, string>;
  private accessToken: string | null = null;
  private refreshToken: string | null = null;
  private onUnauthorized?: () => void;

  constructor(config: ApiClientConfig) {
    // Strip trailing slash
    this.baseURL = config.baseURL.replace(/\/+$/, '');
    this.timeout = config.timeout ?? 30_000;
    this.defaultHeaders = config.defaultHeaders ?? {};
    this.onUnauthorized = config.onUnauthorized;
  }

  // -------------------------------------------------------------------------
  // Token Management
  // -------------------------------------------------------------------------

  /** Returns the current access token, or null if not set */
  getAccessToken(): string | null {
    return this.accessToken;
  }

  /** Returns the current refresh token, or null if not set */
  getRefreshToken(): string | null {
    return this.refreshToken;
  }

  /** Sets both access and refresh tokens */
  setTokens(accessToken: string, refreshToken?: string): void {
    this.accessToken = accessToken;
    if (refreshToken !== undefined) {
      this.refreshToken = refreshToken;
    }
  }

  /** Clears all stored tokens (logout) */
  clearTokens(): void {
    this.accessToken = null;
    this.refreshToken = null;
  }

  /** Returns true if an access token is currently set */
  isAuthenticated(): boolean {
    return this.accessToken !== null;
  }

  // -------------------------------------------------------------------------
  // HTTP Methods
  // -------------------------------------------------------------------------

  /**
   * Sends a GET request and returns the typed response data.
   */
  async get<T>(
    path: string,
    params?: Record<string, string | number | boolean | undefined>,
  ): Promise<T> {
    const url = this.buildURL(path, params);
    return this.request<T>(url, { method: 'GET' });
  }

  /**
   * Sends a POST request with a JSON body and returns the typed response data.
   */
  async post<T>(path: string, body?: unknown): Promise<T> {
    const url = this.buildURL(path);
    return this.request<T>(url, {
      method: 'POST',
      body: body !== undefined ? JSON.stringify(body) : undefined,
    });
  }

  /**
   * Sends a PUT request with a JSON body and returns the typed response data.
   */
  async put<T>(path: string, body?: unknown): Promise<T> {
    const url = this.buildURL(path);
    return this.request<T>(url, {
      method: 'PUT',
      body: body !== undefined ? JSON.stringify(body) : undefined,
    });
  }

  /**
   * Sends a DELETE request and returns the typed response data.
   */
  async delete<T>(path: string): Promise<T> {
    const url = this.buildURL(path);
    return this.request<T>(url, { method: 'DELETE' });
  }

  // -------------------------------------------------------------------------
  // Internal
  // -------------------------------------------------------------------------

  private buildURL(
    path: string,
    params?: Record<string, string | number | boolean | undefined>,
  ): string {
    const url = new URL(`${this.baseURL}${path}`);

    if (params) {
      for (const [key, value] of Object.entries(params)) {
        if (value !== undefined && value !== null) {
          url.searchParams.set(key, String(value));
        }
      }
    }

    return url.toString();
  }

  private async request<T>(url: string, init: RequestInit): Promise<T> {
    const controller = new AbortController();
    const timeoutId = setTimeout(() => controller.abort(), this.timeout);

    const headers: Record<string, string> = {
      'Content-Type': 'application/json',
      Accept: 'application/json',
      ...this.defaultHeaders,
    };

    if (this.accessToken) {
      headers['Authorization'] = `Bearer ${this.accessToken}`;
    }

    try {
      const response = await fetch(url, {
        ...init,
        headers,
        signal: controller.signal,
      });

      clearTimeout(timeoutId);

      // Handle 204 No Content
      if (response.status === 204) {
        return undefined as T;
      }

      // Parse JSON response
      const json = await response.json();

      if (!response.ok) {
        const error = json as ApiError;

        // Handle 401 Unauthorized
        if (response.status === 401) {
          this.onUnauthorized?.();
        }

        throw new ApiClientError(
          error.message || `Request failed with status ${response.status}`,
          response.status,
          error.errors,
        );
      }

      // The API wraps responses in { data: T, message?: string }
      // Return the data field if it exists, otherwise the whole response
      if (json && typeof json === 'object' && 'data' in json) {
        return (json as ApiResponse<T>).data;
      }

      return json as T;
    } catch (err) {
      clearTimeout(timeoutId);

      if (err instanceof ApiClientError) {
        throw err;
      }

      if (err instanceof DOMException && err.name === 'AbortError') {
        throw new ApiClientError(
          `Request timed out after ${this.timeout}ms`,
          408,
        );
      }

      throw new ApiClientError(
        err instanceof Error ? err.message : 'Network error',
        0,
      );
    }
  }
}

// ---------------------------------------------------------------------------
// Error Class
// ---------------------------------------------------------------------------

/**
 * Custom error class for API client errors.
 * Carries HTTP status code and optional field-level validation errors.
 */
export class ApiClientError extends Error {
  readonly statusCode: number;
  readonly errors?: Record<string, string[]>;

  constructor(
    message: string,
    statusCode: number,
    errors?: Record<string, string[]>,
  ) {
    super(message);
    this.name = 'ApiClientError';
    this.statusCode = statusCode;
    this.errors = errors;
  }

  /** Returns true if this is a 401 Unauthorized error */
  isUnauthorized(): boolean {
    return this.statusCode === 401;
  }

  /** Returns true if this is a 403 Forbidden error */
  isForbidden(): boolean {
    return this.statusCode === 403;
  }

  /** Returns true if this is a 404 Not Found error */
  isNotFound(): boolean {
    return this.statusCode === 404;
  }

  /** Returns true if this is a 422 Validation error */
  isValidationError(): boolean {
    return this.statusCode === 422;
  }

  /** Returns true if this is a network/timeout error (status 0 or 408) */
  isNetworkError(): boolean {
    return this.statusCode === 0 || this.statusCode === 408;
  }
}
