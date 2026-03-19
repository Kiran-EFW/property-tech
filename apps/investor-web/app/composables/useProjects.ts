import type { ProjectFilters, Project, PaginatedResponse } from '@proptech/api-client'

/**
 * Fetches a paginated list of projects. SSR-safe via useAsyncData.
 */
export function useProjects(filters?: Ref<ProjectFilters> | ProjectFilters) {
  const api = useApi()
  const resolvedFilters = isRef(filters) ? filters : ref(filters ?? {})

  return useAsyncData(
    'projects',
    () => api.projects.listProjects(resolvedFilters.value) as Promise<PaginatedResponse<Project>>,
    { watch: [resolvedFilters] },
  )
}

/**
 * Fetches a single project by slug. SSR-safe.
 */
export function useProject(slug: Ref<string> | string) {
  const api = useApi()
  const resolvedSlug = isRef(slug) ? slug : ref(slug)

  return useAsyncData(
    `project-${resolvedSlug.value}`,
    () => api.projects.getProject(resolvedSlug.value),
    { watch: [resolvedSlug] },
  )
}
