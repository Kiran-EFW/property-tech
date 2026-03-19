<script setup lang="ts">
definePageMeta({ layout: 'agent', middleware: 'auth' })

const { visits, leads, projects } = useMockData()
const { formatDateTime, formatRelativeTime } = useFormatting()

const agentId = 'agt-001'

const myVisits = computed(() => visits.filter(v => v.agentId === agentId))

const activeTab = ref('upcoming')

const tabItems = computed(() => [
  { label: `Upcoming (${myVisits.value.filter(v => v.status === 'scheduled').length})`, value: 'upcoming' },
  { label: `Completed (${myVisits.value.filter(v => v.status === 'completed').length})`, value: 'completed' },
  { label: `Cancelled (${myVisits.value.filter(v => v.status === 'cancelled').length})`, value: 'cancelled' },
])

const filteredVisits = computed(() => {
  if (activeTab.value === 'upcoming') return myVisits.value.filter(v => v.status === 'scheduled')
  if (activeTab.value === 'completed') return myVisits.value.filter(v => v.status === 'completed')
  if (activeTab.value === 'cancelled') return myVisits.value.filter(v => v.status === 'cancelled')
  return myVisits.value
})

function getInterestColor(level: string | null) {
  if (level === 'high') return 'success'
  if (level === 'medium') return 'warning'
  if (level === 'low') return 'error'
  return 'neutral'
}

function getOutcomeLabel(outcome: string | null) {
  const map: Record<string, string> = {
    interested: 'Interested',
    follow_up: 'Follow Up',
    not_interested: 'Not Interested',
    booked: 'Booked',
  }
  return outcome ? map[outcome] || outcome : '-'
}

function getOutcomeColor(outcome: string | null) {
  const map: Record<string, string> = {
    interested: 'success',
    follow_up: 'warning',
    not_interested: 'error',
    booked: 'primary',
  }
  return outcome ? map[outcome] || 'neutral' : 'neutral'
}

// Add visit modal
const showAddModal = ref(false)
const newVisit = reactive({
  leadId: '',
  projectId: '',
  date: '',
  time: '',
})

const agentLeads = computed(() =>
  leads
    .filter(l => l.agentId === agentId)
    .map(l => ({ label: `${l.name} - ${l.phone}`, value: l.id }))
)

const projectOptions = computed(() =>
  projects
    .filter(p => p.status === 'active')
    .map(p => ({ label: `${p.name} (${p.location})`, value: p.id }))
)

function submitVisit() {
  console.log('[Visit] New visit scheduled:', { ...newVisit })
  showAddModal.value = false
  Object.assign(newVisit, { leadId: '', projectId: '', date: '', time: '' })
}
</script>

<template>
  <div>
    <div class="flex items-center justify-between mb-4">
      <div>
        <h1 class="text-xl font-bold text-brand-900 dark:text-white">Site Visits</h1>
        <p class="text-slate-500 dark:text-slate-400 text-sm mt-1">Manage your scheduled site visits</p>
      </div>
      <UButton icon="i-lucide-plus" size="sm" @click="showAddModal = true">
        Add Visit
      </UButton>
    </div>

    <UTabs :items="tabItems" v-model="activeTab" class="mb-4" />

    <div class="space-y-3">
      <UCard
        v-for="visit in filteredVisits"
        :key="visit.id"
        class="hover:shadow-md transition-shadow"
      >
        <div class="flex items-start justify-between mb-2">
          <div>
            <p class="font-medium text-brand-900 dark:text-white">{{ visit.investorName }}</p>
            <p class="text-xs text-slate-500 dark:text-slate-400">{{ visit.projectName }}</p>
          </div>
          <UBadge
            :color="visit.status === 'scheduled' ? 'info' : visit.status === 'completed' ? 'success' : 'error'"
            variant="subtle"
            size="xs"
          >
            {{ visit.status === 'scheduled' ? 'Upcoming' : visit.status === 'completed' ? 'Completed' : 'Cancelled' }}
          </UBadge>
        </div>

        <div class="flex items-center gap-2 text-xs text-slate-500 dark:text-slate-400 mb-2">
          <UIcon name="i-lucide-calendar" />
          <span>{{ formatDateTime(visit.scheduledAt) }}</span>
        </div>

        <!-- Completed visit details -->
        <template v-if="visit.status === 'completed'">
          <USeparator class="my-3" />

          <div v-if="visit.investorInterestLevel" class="flex items-center gap-2 mb-2">
            <span class="text-xs text-slate-500 dark:text-slate-400">Interest:</span>
            <UBadge :color="getInterestColor(visit.investorInterestLevel)" variant="subtle" size="xs">
              {{ visit.investorInterestLevel }}
            </UBadge>
          </div>

          <div v-if="visit.outcome" class="flex items-center gap-2 mb-2">
            <span class="text-xs text-slate-500 dark:text-slate-400">Outcome:</span>
            <UBadge :color="getOutcomeColor(visit.outcome)" variant="subtle" size="xs">
              {{ getOutcomeLabel(visit.outcome) }}
            </UBadge>
          </div>

          <div v-if="visit.investorFeedback" class="mt-2">
            <p class="text-xs text-slate-500 dark:text-slate-400 mb-1">Feedback:</p>
            <p class="text-sm text-slate-700 dark:text-slate-300 bg-slate-50 dark:bg-brand-900/50 rounded p-2 italic">"{{ visit.investorFeedback }}"</p>
          </div>

          <div v-if="visit.agentNotes" class="mt-2">
            <p class="text-xs text-slate-500 dark:text-slate-400 mb-1">Agent Notes:</p>
            <p class="text-sm text-slate-600 dark:text-slate-300">{{ visit.agentNotes }}</p>
          </div>

          <div v-if="visit.nextSteps" class="mt-2">
            <p class="text-xs text-slate-500 dark:text-slate-400 mb-1">Next Steps:</p>
            <p class="text-sm text-brand-600 dark:text-brand-400 font-medium">{{ visit.nextSteps }}</p>
          </div>
        </template>

        <!-- Cancelled visit details -->
        <template v-if="visit.status === 'cancelled' && visit.agentNotes">
          <USeparator class="my-3" />
          <p class="text-sm text-slate-500 dark:text-slate-400 italic">{{ visit.agentNotes }}</p>
        </template>

        <!-- Upcoming visit actions -->
        <template v-if="visit.status === 'scheduled'">
          <div class="flex gap-2 mt-3">
            <UButton
              icon="i-lucide-phone"
              variant="outline"
              size="xs"
              @click="() => window.open(`tel:+91${visit.investorPhone}`, '_self')"
            >
              Call
            </UButton>
            <UButton
              icon="i-lucide-navigation"
              variant="outline"
              color="primary"
              size="xs"
              @click="console.log('[Visit] Get directions for:', visit.projectName)"
            >
              Directions
            </UButton>
          </div>
        </template>
      </UCard>

      <div v-if="filteredVisits.length === 0" class="text-center py-12">
        <UIcon name="i-lucide-calendar" class="text-4xl text-slate-300 mb-2" />
        <p class="text-slate-400">No {{ activeTab }} visits</p>
      </div>
    </div>

    <!-- Add Visit Modal -->
    <UModal v-model:open="showAddModal">
      <template #content>
        <div class="p-6">
          <h3 class="text-lg font-semibold text-brand-900 dark:text-white mb-4">Schedule New Visit</h3>

          <div class="space-y-4">
            <UFormField label="Select Lead">
              <USelect
                v-model="newVisit.leadId"
                :items="agentLeads"
                placeholder="Choose a lead"
                class="w-full"
              />
            </UFormField>

            <UFormField label="Project">
              <USelect
                v-model="newVisit.projectId"
                :items="projectOptions"
                placeholder="Choose a project"
                class="w-full"
              />
            </UFormField>

            <UFormField label="Date">
              <UInput v-model="newVisit.date" type="date" class="w-full" />
            </UFormField>

            <UFormField label="Time">
              <UInput v-model="newVisit.time" type="time" class="w-full" />
            </UFormField>

            <div class="flex gap-3 pt-2">
              <UButton variant="outline" class="flex-1" @click="showAddModal = false">
                Cancel
              </UButton>
              <UButton
                class="flex-1"
                :disabled="!newVisit.leadId || !newVisit.projectId || !newVisit.date || !newVisit.time"
                @click="submitVisit"
              >
                Schedule Visit
              </UButton>
            </div>
          </div>
        </div>
      </template>
    </UModal>
  </div>
</template>
