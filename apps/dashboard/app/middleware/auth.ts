export default defineNuxtRouteMiddleware((to) => {
  // Skip auth for login page
  if (to.path === '/login') return

  const token = useCookie('auth_token')

  if (!token.value) {
    return navigateTo('/login')
  }

  // Decode JWT to get role (basic decode, no verification — server validates)
  try {
    const payload = JSON.parse(atob(token.value.split('.')[1]))
    const role = payload.role as string

    // Enforce role-based route access
    if (to.path.startsWith('/agent') && role !== 'agent') {
      return navigateTo('/login')
    }
    if (to.path.startsWith('/admin') && role !== 'admin') {
      return navigateTo('/login')
    }
    if (to.path.startsWith('/builder') && role !== 'builder') {
      return navigateTo('/login')
    }
  } catch {
    return navigateTo('/login')
  }
})
