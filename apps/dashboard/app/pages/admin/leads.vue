<script setup lang="ts">
definePageMeta({ layout: 'admin', middleware: 'auth' })

const { leads, agents, projects } = useMockData()
const { formatINR, formatPhone, formatDate, formatRelativeTime } = useFormatting()

const searchQuery = ref('')
const sourceFilter = ref('')
const statusFilter = ref('')
const agentFilter = ref('')
const projectFilter = ref('')

const sourceOptions = [
  { label: 'All Sources', value: '' },
  { label: 'Website', value: 'website' },
  { label: 'WhatsApp', value: 'whatsapp' },
  { label: 'Referral', value: 'referral' },
]

const statusOptions = [
  { label: 'All Status', value: '' },
  { label: 'New', value: 'new' },
  { label: 'Contacted', value: 'contacted' },
  { label: 'Visit Done', value: 'site_visit' },
  { label: 'Booked', value: 'booked' },
  { label: 'Lost', value: 'lost' },
]

const agentOptions = computed(() => [
  { label: 'All Agents', value: '' },
  ...agents.filter(a => a.isActive).map(a => ({ label: a.name, value: a.id })),
])

const projectOptions = computed(() => [
  { label: 'All Projects', value: '' },
  ...projects.filter(p => p.status === 'active').map(p => ({ label: p.name, value: p.id })),
])

const filteredLeads = computed(() => {
  let result = [...leads]
  if (searchQuery.value) {
    const q = searchQuery.value.toLowerCase()
    result = result.filter(l =>
      l.name.toLowerCase().includes(q) ||
      l.phone.includes(q)
    )
  }
  if (sourceFilter.value) result = result.filter(l => l.source === sourceFilter.value)
  if (statusFilter.value) result = result.filter(l => l.status === statusFilter.value)
  if (agentFilter.value) result = result.filter(l => l.agentId === agentFilter.value)
  if (projectFilter.value) result = result.filter(l => l.projectId === projectFilter.value)
  return result
})

function getStatusColor(status: string) {
  const map: Record<string, string> = {
    new: 'info', contacted: 'warning', site_visit: 'primary', booked: 'success', lost: 'error',
  }
  return map[status] || 'neutral'
}

function getStatusLabel(status: string) {
  const map: Record<string, string> = {
    new: 'New', contacted: 'Contacted', site_visit: 'Visit Done', booked: 'Booked', lost: 'Lost',
  }
  return map[status] || status
}

// Bulk assign
const selectedLeads = ref<string[]>([])
const showAssignModal = ref(false)
const assignAgentId = ref('')

function toggleSelect(leadId: string) {
  const idx = selectedLeads.value.indexOf(leadId)
  if (idx >= 0) {
    selectedLeads.value.splice(idx, 1)
  } else {
    selectedLeads.value.push(leadId)
  }
}

function toggleSelectAll() {
  if (selectedLeads.value.length === filteredLeads.value.length) {
    selectedLeads.value = []
  } else {
    selectedLeads.value = filteredLeads.value.map(l => l.id)
  }
}

function bulkAssign() {
  console.log('[Lead] Bulk assign:', selectedLeads.value, 'to agent:', assignAgentId.value)
  showAssignModal.value = false
  selectedLeads.value = []
  assignAgentId.value = ''
}

function exportLeads() {
  console.log('[Lead] Export requested for', filteredLeads.value.length, 'leads')
}
</script>

<template>
  <div>
    <div class="flex items-center justify-between mb-6">
      <div>
        <h1 class="text-2xl font-bold text-brand-900 dark:text-white">Leads</h1>
        <p class="text-slate-500 dark:text-slate-400 text-sm mt-1">All investor enquiries across channels</p>
      </div>
      <div class="flex gap-2">
        <UButton
          v-if="selectedLeads.length > 0"
          icon="i-lucide-user-plus"
          variant="outline"
          @click="showAssignModal = true"
        >
          Assign ({{ selectedLeads.length }})
        </UButton>
        <UButton icon="i-lucide-download" variant="outline" @click="exportLeads">Export</UButton>
      </div>
    </div>

    <UCard>
      <!-- Filters -->
      <div class="flex items-center gap-3 mb-4 flex-wrap">
        <UInput v-model="searchQuery" placeholder="Search by name or phone..." icon="i-lucide-search" class="flex-1 min-w-48" />
        <USelect v-model="sourceFilter" :items="sourceOptions" placeholder="Source" />
        <USelect v-model="statusFilter" :items="statusOptions" placeholder="Status" />
        <USelect v-model="agentFilter" :items="agentOptions" placeholder="Agent" />
        <USelect v-model="projectFilter" :items="projectOptions" placeholder="Project" />
      </div>

      <!-- Results count -->
      <p class="text-xs text-slate-400 dark:text-slate-500 mb-3">{{ filteredLeads.length }} leads found</p>

      <!-- Table -->
      <div class="overflow-x-auto">
        <table class="w-full text-sm">
          <thead>
            <tr class="border-b border-slate-200 dark:border-slate-700">
              <th class="py-3 px-3 w-8">
                <input
                  type="checkbox"
                  :checked="selectedLeads.length === filteredLeads.length && filteredLeads.length > 0"
                  class="rounded"
                  @change="toggleSelectAll"
                />
              </th>
              <th class="text-left py-3 px-3 text-xs font-semibold text-slate-500 dark:text-slate-400 uppercase">Name</th>
              <th class="text-left py-3 px-3 text-xs font-semibold text-slate-500 dark:text-slate-400 uppercase">Phone</th>
              <th class="text-left py-3 px-3 text-xs font-semibold text-slate-500 dark:text-slate-400 uppercase">Project</th>
              <th class="text-left py-3 px-3 text-xs font-semibold text-slate-500 dark:text-slate-400 uppercase">Agent</th>
              <th class="text-left py-3 px-3 text-xs font-semibold text-slate-500 dark:text-slate-400 uppercase">Source</th>
              <th class="text-left py-3 px-3 text-xs font-semibold text-slate-500 dark:text-slate-400 uppercase">Status</th>
              <th class="text-left py-3 px-3 text-xs font-semibold text-slate-500 dark:text-slate-400 uppercase">Budget</th>
              <th class="text-left py-3 px-3 text-xs font-semibold text-slate-500 dark:text-slate-400 uppercase">Created</th>
            </tr>
          </thead>
          <tbody>
            <tr
              v-for="lead in filteredLeads"
              :key="lead.id"
              class="border-b border-slate-100 dark:border-slate-800 hover:bg-slate-50 dark:hover:bg-brand-800"
            >
              <td class="py-3 px-3">
                <input
                  type="checkbox"
                  :checked="selectedLeads.includes(lead.id)"
                  class="rounded"
                  @change="toggleSelect(lead.id)"
                />
              </td>
              <td class="py-3 px-3">
                <div class="flex items-center gap-1">
                  <p class="font-medium text-brand-900 dark:text-white">{{ lead.name }}</p>
                  <UIcon v-if="lead.isHot" name="i-lucide-flame" class="text-red-500 text-xs" />
                  <UBadge v-if="lead.isNRI" color="info" variant="subtle" size="xs">NRI</UBadge>
                </div>
              </td>
              <td class="py-3 px-3">
                <a :href="`tel:+91${lead.phone}`" class="text-brand-600 dark:text-brand-400 text-xs">
                  {{ formatPhone(lead.phone) }}
                </a>
              </td>
              <td class="py-3 px-3 text-slate-600 dark:text-slate-300 text-xs">{{ lead.projectName }}</td>
              <td class="py-3 px-3 text-slate-600 dark:text-slate-300 text-xs">{{ lead.agentName }}</td>
              <td class="py-3 px-3">
                <span class="text-xs text-slate-500 dark:text-slate-400 capitalize">{{ lead.source }}</span>
              </td>
              <td class="py-3 px-3">
                <UBadge :color="getStatusColor(lead.status)" variant="subtle" size="xs">
                  {{ getStatusLabel(lead.status) }}
                </UBadge>
              </td>
              <td class="py-3 px-3 text-slate-600 dark:text-slate-300 text-xs">{{ formatINR(lead.budget || 0) }}</td>
              <td class="py-3 px-3 text-slate-500 dark:text-slate-400 text-xs">{{ formatDate(lead.createdAt) }}</td>
            </tr>
          </tbody>
        </table>
      </div>

      <div v-if="filteredLeads.length === 0" class="text-center py-12">
        <UIcon name="i-lucide-users" class="text-4xl text-slate-300 mb-2" />
        <p class="text-slate-400">No leads match your filters</p>
      </div>
    </UCard>

    <!-- Bulk Assign Modal -->
    <UModal v-model:open="showAssignModal">
      <template #content>
        <div class="p-6">
          <h3 class="text-lg font-semibold text-brand-900 dark:text-white mb-4">
            Assign {{ selectedLeads.length }} Lead{{ selectedLeads.length > 1 ? 's' : '' }} to Agent
          </h3>

          <UFormField label="Select Agent">
            <USelect
              v-model="assignAgentId"
              :items="agents.filter(a => a.isActive).map(a => ({ label: `${a.name} (${a.totalLeads} leads)`, value: a.id }))"
              placeholder="Choose an agent"
              class="w-full"
            />
          </UFormField>

          <div class="flex gap-3 mt-6">
            <UButton variant="outline" class="flex-1" @click="showAssignModal = false">Cancel</UButton>
            <UButton class="flex-1" :disabled="!assignAgentId" @click="bulkAssign">
              Assign Leads
            </UButton>
          </div>
        </div>
      </template>
    </UModal>
  </div>
</template>
