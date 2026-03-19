<script setup lang="ts">
definePageMeta({ layout: 'agent', middleware: 'auth' })

const route = useRoute()
const { leads, visits } = useMockData()
const { formatINR, formatPhone, formatDate, formatDateTime, formatRelativeTime } = useFormatting()

const leadId = route.params.id as string

const lead = computed(() => leads.find(l => l.id === leadId))
const leadVisits = computed(() => visits.filter(v => v.leadId === leadId))

// Activity timeline mock
const timeline = computed(() => {
  if (!lead.value) return []
  const items: Array<{ id: string; type: string; title: string; description: string; time: string; icon: string; color: string }> = []

  items.push({
    id: 'tl-1',
    type: 'created',
    title: 'Lead Created',
    description: `${lead.value.name} enquired via ${lead.value.source}`,
    time: lead.value.createdAt,
    icon: 'i-lucide-plus-circle',
    color: 'text-blue-500',
  })

  if (lead.value.lastContactedAt) {
    items.push({
      id: 'tl-2',
      type: 'contacted',
      title: 'First Contact Made',
      description: `Agent called ${lead.value.name}`,
      time: lead.value.lastContactedAt,
      icon: 'i-lucide-phone',
      color: 'text-green-500',
    })
  }

  leadVisits.value
    .filter(v => v.status === 'completed')
    .forEach((v, i) => {
      items.push({
        id: `tl-visit-${i}`,
        type: 'visit',
        title: 'Site Visit Completed',
        description: `Visited ${v.projectName}. Outcome: ${v.outcome || 'N/A'}`,
        time: v.scheduledAt,
        icon: 'i-lucide-map-pin',
        color: 'text-brand-600 dark:text-brand-400',
      })
    })

  if (lead.value.status === 'booked') {
    items.push({
      id: 'tl-booked',
      type: 'booked',
      title: 'Booking Confirmed',
      description: `${lead.value.name} booked at ${lead.value.projectName}`,
      time: lead.value.updatedAt,
      icon: 'i-lucide-check-circle',
      color: 'text-green-600',
    })
  }

  return items.sort((a, b) => new Date(b.time).getTime() - new Date(a.time).getTime())
})

// Add note
const newNote = ref('')
const notes = ref<Array<{ id: string; content: string; time: string }>>([
  { id: 'n-1', content: 'Initial call done. Client is interested but wants to compare with other projects.', time: '2026-03-16T10:00:00Z' },
  { id: 'n-2', content: 'Shared brochure and price list via WhatsApp.', time: '2026-03-15T14:00:00Z' },
])

function addNote() {
  if (!newNote.value.trim()) return
  notes.value.unshift({
    id: `n-${Date.now()}`,
    content: newNote.value,
    time: new Date().toISOString(),
  })
  console.log('[Lead] Note added:', newNote.value)
  newNote.value = ''
}

// Status update
function updateStatus(newStatus: string) {
  console.log('[Lead] Status updated:', leadId, '->', newStatus)
}

// Follow-up date
const followUpDate = ref(lead.value?.nextFollowUpAt ? lead.value.nextFollowUpAt.slice(0, 10) : '')

function scheduleFollowUp() {
  console.log('[Lead] Follow-up scheduled for:', followUpDate.value)
}

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
  <div v-if="lead">
    <!-- Back button -->
    <div class="flex items-center gap-2 mb-4">
      <UButton
        icon="i-lucide-arrow-left"
        variant="ghost"
        color="neutral"
        size="sm"
        @click="navigateTo('/agent/leads')"
      />
      <h2 class="text-lg font-semibold text-brand-900 dark:text-white">Lead Details</h2>
    </div>

    <!-- Lead header -->
    <UCard class="mb-4">
      <div class="flex items-start justify-between mb-3">
        <div>
          <h3 class="text-lg font-bold text-brand-900 dark:text-white">{{ lead.name }}</h3>
          <div class="flex items-center gap-1 mt-1">
            <UBadge :color="getStatusColor(lead.status)" variant="subtle" size="xs">
              {{ getStatusLabel(lead.status) }}
            </UBadge>
            <UBadge v-if="lead.isHot" color="error" variant="subtle" size="xs">
              Hot
            </UBadge>
            <UBadge v-if="lead.isNRI" color="info" variant="subtle" size="xs">
              NRI
            </UBadge>
          </div>
        </div>
        <span class="text-xs text-slate-400">Score: {{ lead.score }}/100</span>
      </div>

      <!-- Contact info -->
      <div class="space-y-2 mb-4">
        <a :href="`tel:+91${lead.phone}`" class="flex items-center gap-2 text-sm text-brand-600 dark:text-brand-400">
          <UIcon name="i-lucide-phone" />
          {{ formatPhone(lead.phone) }}
        </a>
        <a v-if="lead.email" :href="`mailto:${lead.email}`" class="flex items-center gap-2 text-sm text-slate-600 dark:text-slate-300">
          <UIcon name="i-lucide-mail" />
          {{ lead.email }}
        </a>
      </div>

      <!-- Quick contact actions -->
      <div class="flex gap-2 mb-4">
        <UButton
          icon="i-lucide-phone"
          color="primary"
          size="sm"
          class="flex-1"
          @click="() => window.open(`tel:+91${lead.phone}`, '_self')"
        >
          Call
        </UButton>
        <UButton
          icon="i-lucide-message-circle"
          color="success"
          size="sm"
          class="flex-1"
          @click="() => {
            const msg = encodeURIComponent(`Hi ${lead.name}, this is regarding your property enquiry on PropTech.`)
            window.open(`https://wa.me/91${lead.phone}?text=${msg}`, '_blank')
          }"
        >
          WhatsApp
        </UButton>
      </div>

      <!-- Lead details grid -->
      <div class="grid grid-cols-2 gap-3 text-sm">
        <div>
          <p class="text-xs text-slate-400">Project</p>
          <p class="font-medium text-brand-900 dark:text-white">{{ lead.projectName }}</p>
        </div>
        <div>
          <p class="text-xs text-slate-400">Configuration</p>
          <p class="font-medium text-brand-900 dark:text-white">{{ lead.preferredConfiguration }}</p>
        </div>
        <div>
          <p class="text-xs text-slate-400">Budget</p>
          <p class="font-medium text-brand-900 dark:text-white">{{ formatINR(lead.budget || 0) }}</p>
        </div>
        <div>
          <p class="text-xs text-slate-400">Source</p>
          <p class="font-medium text-brand-900 dark:text-white capitalize">{{ lead.source }}</p>
        </div>
        <div>
          <p class="text-xs text-slate-400">Created</p>
          <p class="font-medium text-brand-900 dark:text-white">{{ formatDate(lead.createdAt) }}</p>
        </div>
        <div>
          <p class="text-xs text-slate-400">Last Contact</p>
          <p class="font-medium text-brand-900 dark:text-white">
            {{ lead.lastContactedAt ? formatRelativeTime(lead.lastContactedAt) : 'Not yet' }}
          </p>
        </div>
      </div>

      <div v-if="lead.remarks" class="mt-3 p-2 bg-slate-50 dark:bg-brand-900/50 rounded text-sm text-slate-600 dark:text-slate-300">
        <p class="text-xs text-slate-400 mb-1">Remarks</p>
        {{ lead.remarks }}
      </div>
    </UCard>

    <!-- Status Update Buttons -->
    <UCard class="mb-4">
      <p class="text-sm font-medium text-brand-900 dark:text-white mb-3">Update Status</p>
      <div class="flex flex-wrap gap-2">
        <UButton
          v-for="s in ['new', 'contacted', 'site_visit', 'booked', 'lost']"
          :key="s"
          size="xs"
          :variant="lead.status === s ? 'solid' : 'outline'"
          :color="lead.status === s ? 'primary' : 'neutral'"
          @click="updateStatus(s)"
        >
          {{ getStatusLabel(s) }}
        </UButton>
      </div>
    </UCard>

    <!-- Schedule Follow-up -->
    <UCard class="mb-4">
      <p class="text-sm font-medium text-brand-900 dark:text-white mb-3">Schedule Follow-up</p>
      <div class="flex gap-2">
        <UInput v-model="followUpDate" type="date" class="flex-1" />
        <UButton size="sm" @click="scheduleFollowUp" :disabled="!followUpDate">
          Schedule
        </UButton>
      </div>
      <p v-if="lead.nextFollowUpAt" class="text-xs text-slate-500 dark:text-slate-400 mt-2">
        Current follow-up: {{ formatDateTime(lead.nextFollowUpAt) }}
      </p>
    </UCard>

    <!-- Add Note -->
    <UCard class="mb-4">
      <p class="text-sm font-medium text-brand-900 dark:text-white mb-3">Notes</p>
      <div class="flex gap-2 mb-4">
        <UInput
          v-model="newNote"
          placeholder="Add a note..."
          class="flex-1"
          @keyup.enter="addNote"
        />
        <UButton size="sm" @click="addNote" :disabled="!newNote.trim()">
          Add
        </UButton>
      </div>
      <div class="space-y-3">
        <div v-for="note in notes" :key="note.id" class="p-2 bg-slate-50 dark:bg-brand-900/50 rounded">
          <p class="text-sm text-slate-700">{{ note.content }}</p>
          <p class="text-xs text-slate-400 mt-1">{{ formatRelativeTime(note.time) }}</p>
        </div>
      </div>
    </UCard>

    <!-- Activity Timeline -->
    <UCard>
      <p class="text-sm font-medium text-brand-900 dark:text-white mb-3">Activity Timeline</p>
      <div class="space-y-4">
        <div v-for="item in timeline" :key="item.id" class="flex gap-3">
          <div class="flex flex-col items-center">
            <UIcon :name="item.icon" :class="['text-lg', item.color]" />
            <div class="w-px flex-1 bg-slate-200 dark:bg-brand-700 mt-1" />
          </div>
          <div class="pb-4">
            <p class="text-sm font-medium text-brand-900 dark:text-white">{{ item.title }}</p>
            <p class="text-xs text-slate-500 dark:text-slate-400">{{ item.description }}</p>
            <p class="text-xs text-slate-400 mt-1">{{ formatRelativeTime(item.time) }}</p>
          </div>
        </div>
      </div>
    </UCard>
  </div>

  <div v-else class="text-center py-12">
    <UIcon name="i-lucide-user-x" class="text-4xl text-slate-300 mb-2" />
    <p class="text-slate-400">Lead not found</p>
    <UButton variant="link" @click="navigateTo('/agent/leads')" class="mt-2">
      Back to Leads
    </UButton>
  </div>
</template>
