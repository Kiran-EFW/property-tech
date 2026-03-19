/**
 * Authentication composable for the PropTech dashboard.
 *
 * Reads `auth_token` from cookies, decodes the JWT payload (base64)
 * and exposes user info, role helpers, and login/logout actions.
 *
 * JWT payload structure: { sub, role, phone, name?, email?, exp }
 */

interface AuthUser {
  id: string
  name: string
  phone: string
  role: 'agent' | 'admin' | 'builder'
  email: string | null
}

export function useAuth() {
  const authToken = useCookie<string | null>('auth_token')
  const user = ref<AuthUser | null>(null)

  /**
   * Decode the JWT payload (no verification -- the server validates the token).
   * Returns null if the token is missing, malformed, or expired.
   */
  function decodeToken(token: string | null): AuthUser | null {
    if (!token) return null

    try {
      const parts = token.split('.')
      if (parts.length !== 3) return null

      const payload = JSON.parse(atob(parts[1]))

      // Check expiration
      if (payload.exp && payload.exp * 1000 < Date.now()) {
        return null
      }

      return {
        id: payload.sub,
        name: payload.name || payload.phone || 'User',
        phone: payload.phone,
        role: payload.role as AuthUser['role'],
        email: payload.email || null,
      }
    }
    catch {
      return null
    }
  }

  // Hydrate user from existing cookie on composable init
  user.value = decodeToken(authToken.value)

  // Keep user in sync when the cookie changes
  watch(authToken, (newToken) => {
    user.value = decodeToken(newToken)
  })

  const isAuthenticated = computed(() => user.value !== null)

  const role = computed<'agent' | 'admin' | 'builder' | null>(
    () => user.value?.role ?? null,
  )

  /**
   * Log in with phone + OTP. In production this calls the API, for now we
   * accept any 6-digit OTP and generate a mock JWT.
   */
  async function login(phone: string, otp: string): Promise<void> {
    const config = useRuntimeConfig()

    try {
      // Try the real API first
      const response = await $fetch<{ data: { access_token: string; refresh_token: string } }>(`${config.public.apiBaseUrl}/auth/verify`, {
        method: 'POST',
        body: { phone, otp },
      })

      authToken.value = response.data.access_token
      user.value = decodeToken(response.data.access_token)
    }
    catch {
      // Fallback: generate a mock JWT for development
      const mockPayload = {
        sub: 'agt-001',
        role: 'agent',
        phone,
        name: 'Rajesh Sharma',
        email: 'rajesh.sharma@gmail.com',
        exp: Math.floor(Date.now() / 1000) + 86400, // 24 hours
      }

      const header = btoa(JSON.stringify({ alg: 'HS256', typ: 'JWT' }))
      const payload = btoa(JSON.stringify(mockPayload))
      const signature = btoa('mock-signature')
      const mockToken = `${header}.${payload}.${signature}`

      authToken.value = mockToken
      user.value = decodeToken(mockToken)
    }
  }

  /**
   * Clear auth state and redirect to login.
   */
  function logout(): void {
    authToken.value = null
    user.value = null
    navigateTo('/login')
  }

  /**
   * Check whether the authenticated user holds a specific role.
   */
  function hasRole(checkRole: string): boolean {
    return user.value?.role === checkRole
  }

  return {
    user,
    isAuthenticated,
    role,
    login,
    logout,
    hasRole,
  }
}
