<script setup lang="ts">
definePageMeta({ layout: 'builder', middleware: 'auth' })

const { leads, projects } = useMockData()
const { formatINR, formatPhone, formatDate, formatRelativeTime } = useFormatting()

// Builder's projects
const builderProjectIds = ['prj-001', 'prj-002']
const builderLeads = computed(() => leads.filter(l => l.projectId && builderProjectIds.includes(l.projectId)))
const builderProjects = computed(() => projects.filter(p => builderProjectIds.includes(p.id)))

const searchQuery = ref('')
const projectFilter = ref('')
const statusFilter = ref('')

const projectOptions = computed(() => [
  { label: 'All Projects', value: '' },
  ...builderProjects.value.map(p => ({ label: p.name, value: p.id })),
])

const statusOptions = [
  { label: 'All Status', value: '' },
  { label: 'New', value: 'new' },
  { label: 'Contacted', value: 'contacted' },
  { label: 'Visit Scheduled', value: 'site_visit' },
  { label: 'Booked', value: 'booked' },
]

const filteredLeads = computed(() => {
  let result = [...builderLeads.value]
  if (searchQuery.value) {
    const q = searchQuery.value.toLowerCase()
    result = result.filter(l =>
      l.name.toLowerCase().includes(q) ||
      l.phone.includes(q)
    )
  }
  if (projectFilter.value) result = result.filter(l => l.projectId === projectFilter.value)
  if (statusFilter.value) result = result.filter(l => l.status === statusFilter.value)
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
</script>

<template>
  <div>
    <div class="flex items-center justify-between mb-6">
      <div>
        <h1 class="text-2xl font-bold text-brand-900 dark:text-white">Leads</h1>
        <p class="text-slate-500 dark:text-slate-400 text-sm mt-1">Investor enquiries for your projects (read-only)</p>
      </div>
    </div>

    <UCard>
      <div class="flex items-center gap-3 mb-4">
        <UInput v-model="searchQuery" placeholder="Search leads..." icon="i-lucide-search" class="flex-1" />
        <USelect v-model="projectFilter" :items="projectOptions" placeholder="Project" />
        <USelect v-model="statusFilter" :items="statusOptions" placeholder="Status" />
      </div>

      <p class="text-xs text-slate-400 dark:text-slate-500 mb-3">{{ filteredLeads.length }} leads for your projects</p>

      <div class="overflow-x-auto">
        <table class="w-full text-sm">
          <thead>
            <tr class="border-b border-slate-200 dark:border-slate-700">
              <th class="text-left py-3 px-3 text-xs font-semibold text-slate-500 dark:text-slate-400 uppercase">Investor</th>
              <th class="text-left py-3 px-3 text-xs font-semibold text-slate-500 dark:text-slate-400 uppercase">Phone</th>
              <th class="text-left py-3 px-3 text-xs font-semibold text-slate-500 dark:text-slate-400 uppercase">Project</th>
              <th class="text-left py-3 px-3 text-xs font-semibold text-slate-500 dark:text-slate-400 uppercase">Config</th>
              <th class="text-left py-3 px-3 text-xs font-semibold text-slate-500 dark:text-slate-400 uppercase">Agent</th>
              <th class="text-left py-3 px-3 text-xs font-semibold text-slate-500 dark:text-slate-400 uppercase">Source</th>
              <th class="text-left py-3 px-3 text-xs font-semibold text-slate-500 dark:text-slate-400 uppercase">Status</th>
              <th class="text-left py-3 px-3 text-xs font-semibold text-slate-500 dark:text-slate-400 uppercase">Budget</th>
              <th class="text-left py-3 px-3 text-xs font-semibold text-slate-500 dark:text-slate-400 uppercase">Date</th>
            </tr>
          </thead>
          <tbody>
            <tr
              v-for="lead in filteredLeads"
              :key="lead.id"
              class="border-b border-slate-100 dark:border-slate-800 hover:bg-slate-50 dark:hover:bg-brand-800"
            >
              <td class="py-3 px-3">
                <div class="flex items-center gap-1">
                  <p class="font-medium text-brand-900 dark:text-white">{{ lead.name }}</p>
                  <UIcon v-if="lead.isHot" name="i-lucide-flame" class="text-red-500 text-xs" />
                </div>
              </td>
              <td class="py-3 px-3 text-xs text-slate-500 dark:text-slate-400">{{ formatPhone(lead.phone) }}</td>
              <td class="py-3 px-3 text-slate-600 dark:text-slate-300 text-xs">{{ lead.projectName }}</td>
              <td class="py-3 px-3 text-slate-600 dark:text-slate-300 text-xs">{{ lead.preferredConfiguration }}</td>
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
              <td class="py-3 px-3 text-xs text-slate-500 dark:text-slate-400">{{ formatRelativeTime(lead.createdAt) }}</td>
            </tr>
          </tbody>
        </table>
      </div>

      <div v-if="filteredLeads.length === 0" class="text-center py-12">
        <UIcon name="i-lucide-users" class="text-4xl text-slate-300 mb-2" />
        <p class="text-slate-400 dark:text-slate-500">No leads found</p>
      </div>
    </UCard>
  </div>
</template>
