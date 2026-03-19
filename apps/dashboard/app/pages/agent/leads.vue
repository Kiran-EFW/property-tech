<script setup lang="ts">
definePageMeta({ layout: 'agent', middleware: 'auth' })

const { leads } = useMockData()
const { formatINR, formatRelativeTime, formatPhone } = useFormatting()

const agentId = 'agt-001'

const myLeads = computed(() => leads.filter(l => l.agentId === agentId))

const searchQuery = ref('')
const activeTab = ref('all')

const tabItems = computed(() => [
  { label: `All (${myLeads.value.length})`, value: 'all' },
  { label: `New (${myLeads.value.filter(l => l.status === 'new').length})`, value: 'new' },
  { label: `Contacted (${myLeads.value.filter(l => l.status === 'contacted').length})`, value: 'contacted' },
  { label: `Visit Done (${myLeads.value.filter(l => l.status === 'site_visit').length})`, value: 'site_visit' },
  { label: `Booked (${myLeads.value.filter(l => l.status === 'booked').length})`, value: 'booked' },
])

const filteredLeads = computed(() => {
  let result = myLeads.value
  if (activeTab.value !== 'all') {
    result = result.filter(l => l.status === activeTab.value)
  }
  if (searchQuery.value) {
    const q = searchQuery.value.toLowerCase()
    result = result.filter(l =>
      l.name.toLowerCase().includes(q) ||
      l.phone.includes(q) ||
      l.projectName.toLowerCase().includes(q)
    )
  }
  return result
})

function getStatusColor(status: string) {
  const map: Record<string, string> = {
    new: 'info',
    contacted: 'warning',
    site_visit: 'primary',
    booked: 'success',
    lost: 'error',
  }
  return map[status] || 'neutral'
}

function getStatusLabel(status: string) {
  const map: Record<string, string> = {
    new: 'New',
    contacted: 'Contacted',
    site_visit: 'Visit Done',
    booked: 'Booked',
    lost: 'Lost',
  }
  return map[status] || status
}

function updateStatus(leadId: string, newStatus: string) {
  console.log('[Lead] Status update:', leadId, '->', newStatus)
}

function callLead(phone: string) {
  window.open(`tel:+91${phone}`, '_self')
}

function whatsappLead(phone: string, name: string) {
  const msg = encodeURIComponent(`Hi ${name}, this is regarding your property enquiry on PropTech. How can I help you?`)
  window.open(`https://wa.me/91${phone}?text=${msg}`, '_blank')
}
</script>

<template>
  <div>
    <h2 class="text-xl font-semibold text-brand-900 dark:text-white mb-4">My Leads</h2>

    <!-- Search -->
    <UInput
      v-model="searchQuery"
      placeholder="Search by name, phone, or project..."
      icon="i-lucide-search"
      size="lg"
      class="mb-4"
    />

    <!-- Filter tabs -->
    <UTabs
      :items="tabItems"
      v-model="activeTab"
      class="mb-4"
    />

    <!-- Lead list -->
    <div class="space-y-3">
      <UCard
        v-for="lead in filteredLeads"
        :key="lead.id"
        class="hover:shadow-md transition-shadow"
      >
        <div class="flex items-start justify-between mb-2">
          <div class="flex-1 min-w-0">
            <NuxtLink :to="`/agent/leads/${lead.id}`" class="hover:underline">
              <p class="font-medium text-brand-900 dark:text-white truncate">{{ lead.name }}</p>
            </NuxtLink>
            <a :href="`tel:+91${lead.phone}`" class="text-xs text-brand-600 dark:text-brand-400">
              {{ formatPhone(lead.phone) }}
            </a>
          </div>
          <UBadge :color="getStatusColor(lead.status)" variant="subtle" size="xs">
            {{ getStatusLabel(lead.status) }}
          </UBadge>
        </div>

        <div class="flex items-center gap-2 text-xs text-slate-500 dark:text-slate-400 mb-3">
          <UIcon name="i-lucide-building-2" class="text-slate-400" />
          <span>{{ lead.projectName }}</span>
          <span class="text-slate-300">|</span>
          <span>{{ lead.preferredConfiguration }}</span>
          <span class="text-slate-300">|</span>
          <span class="font-medium">{{ formatINR(lead.budget || 0) }}</span>
        </div>

        <div v-if="lead.isHot" class="flex items-center gap-1 text-xs text-red-500 mb-2">
          <UIcon name="i-lucide-flame" />
          <span>Hot Lead</span>
        </div>

        <div class="flex items-center gap-1 text-xs text-slate-400 mb-3">
          <UIcon name="i-lucide-clock" />
          <span>{{ lead.lastContactedAt ? `Last contact: ${formatRelativeTime(lead.lastContactedAt)}` : 'Not contacted yet' }}</span>
        </div>

        <!-- Action buttons -->
        <div class="flex items-center gap-2">
          <UButton
            icon="i-lucide-phone"
            variant="outline"
            color="primary"
            size="xs"
            @click="callLead(lead.phone)"
          >
            Call
          </UButton>
          <UButton
            icon="i-lucide-message-circle"
            variant="outline"
            color="success"
            size="xs"
            @click="whatsappLead(lead.phone, lead.name)"
          >
            WhatsApp
          </UButton>
          <div class="flex-1" />
          <USelect
            :model-value="lead.status"
            :items="[
              { label: 'New', value: 'new' },
              { label: 'Contacted', value: 'contacted' },
              { label: 'Visit Done', value: 'site_visit' },
              { label: 'Booked', value: 'booked' },
              { label: 'Lost', value: 'lost' },
            ]"
            size="xs"
            placeholder="Status"
            @update:model-value="(val: string) => updateStatus(lead.id, val)"
          />
        </div>
      </UCard>

      <div v-if="filteredLeads.length === 0" class="text-center py-12">
        <UIcon name="i-lucide-search-x" class="text-3xl text-slate-300 mb-2" />
        <p class="text-slate-400 text-sm">No leads found matching your search</p>
      </div>
    </div>
  </div>
</template>
