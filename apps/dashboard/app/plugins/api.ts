import { createPropTechAPI } from '@proptech/api-client'

export default defineNuxtPlugin(() => {
  const config = useRuntimeConfig()
  const authToken = useCookie<string | null>('auth_token')

  const api = createPropTechAPI({
    baseURL: config.public.apiBaseUrl as string,
    onUnauthorized: () => {
      authToken.value = null
      navigateTo('/login')
    },
  })

  // If there's an existing auth token, set it on the API client
  if (authToken.value) {
    api.client.setTokens(authToken.value)
  }

  // Watch for cookie changes to keep the client in sync
  watch(authToken, (newToken) => {
    if (newToken) {
      api.client.setTokens(newToken)
    } else {
      api.client.clearTokens()
    }
  })

  return {
    provide: {
      api,
    },
  }
})
