import type { Area } from '@proptech/api-client'

/**
 * Fetches all area guides. SSR-safe.
 */
export function useAreas() {
  const api = useApi()

  return useAsyncData(
    'areas',
    () => api.client.get<{ data: Area[] }>('/areas').then(res => res),
  )
}

/**
 * Fetches a single area guide by slug. SSR-safe.
 */
export function useArea(slug: Ref<string> | string) {
  const api = useApi()
  const resolvedSlug = isRef(slug) ? slug : ref(slug)

  return useAsyncData(
    `area-${resolvedSlug.value}`,
    () => api.client.get<Area>(`/areas/${resolvedSlug.value}`),
    { watch: [resolvedSlug] },
  )
}
