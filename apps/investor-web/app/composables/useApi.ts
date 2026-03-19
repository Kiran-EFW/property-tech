import type { PropTechAPI } from '@proptech/api-client'

/**
 * Returns the PropTech API client instance.
 */
export function useApi(): PropTechAPI {
  return useNuxtApp().$api
}
