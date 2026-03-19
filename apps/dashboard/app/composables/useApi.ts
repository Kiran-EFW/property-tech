import type { PropTechAPI } from '@proptech/api-client'

/**
 * Returns the PropTech API client instance provided by the api plugin.
 *
 * @example
 * ```ts
 * const api = useApi()
 * const projects = await api.projects.listProjects({ status: 'active' })
 * ```
 */
export function useApi(): PropTechAPI {
  return useNuxtApp().$api
}
