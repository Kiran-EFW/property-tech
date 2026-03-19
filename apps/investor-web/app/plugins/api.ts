import { createPropTechAPI } from '@proptech/api-client'

export default defineNuxtPlugin(() => {
  const config = useRuntimeConfig()

  const api = createPropTechAPI({
    baseURL: config.public.apiBaseUrl as string,
  })

  return {
    provide: {
      api,
    },
  }
})
