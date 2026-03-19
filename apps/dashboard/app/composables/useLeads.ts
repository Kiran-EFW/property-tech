import type { LeadFilters, Lead, LeadCreateData, PaginatedResponse } from '@proptech/api-client'

/**
 * Fetches a paginated list of leads with reactive filters.
 */
export function useLeads(filters?: Ref<LeadFilters> | LeadFilters) {
  const api = useApi()
  const resolvedFilters = isRef(filters) ? filters : ref(filters ?? {})

  return useAsyncData(
    'leads',
    () => api.leads.listLeads(resolvedFilters.value) as Promise<PaginatedResponse<Lead>>,
    { watch: [resolvedFilters] },
  )
}

/**
 * Fetches a single lead by ID.
 */
export function useLead(id: Ref<string> | string) {
  const api = useApi()
  const resolvedId = isRef(id) ? id : ref(id)

  return useAsyncData(
    `lead-${resolvedId.value}`,
    () => api.leads.getLead(resolvedId.value),
    { watch: [resolvedId] },
  )
}

/**
 * Creates a new lead and returns it.
 */
export function useCreateLead() {
  const api = useApi()
  const loading = ref(false)
  const error = ref<string | null>(null)

  async function create(data: LeadCreateData): Promise<Lead | null> {
    loading.value = true
    error.value = null
    try {
      const lead = await api.leads.createLead(data)
      return lead
    } catch (err: any) {
      error.value = err.message || 'Failed to create lead'
      return null
    } finally {
      loading.value = false
    }
  }

  return { create, loading, error }
}

/**
 * Updates lead status.
 */
export function useUpdateLeadStatus() {
  const api = useApi()
  const loading = ref(false)

  async function updateStatus(leadId: string, status: string): Promise<boolean> {
    loading.value = true
    try {
      await api.leads.updateLeadStatus(leadId, status)
      return true
    } catch {
      return false
    } finally {
      loading.value = false
    }
  }

  return { updateStatus, loading }
}
