<script setup lang="ts">
definePageMeta({ layout: 'admin', middleware: 'auth' })

const { agents } = useMockData()
const { formatPhone, formatDate, formatPercent } = useFormatting()

const searchQuery = ref('')
const statusFilter = ref('')

const statusOptions = [
  { label: 'All Status', value: '' },
  { label: 'Active', value: 'active' },
  { label: 'Pending', value: 'pending' },
  { label: 'Inactive', value: 'inactive' },
]

const filteredAgents = computed(() => {
  let result = [...agents]
  if (statusFilter.value) {
    result = result.filter(a => a.status === statusFilter.value)
  }
  if (searchQuery.value) {
    const q = searchQuery.value.toLowerCase()
    result = result.filter(a =>
      a.name.toLowerCase().includes(q) ||
      a.phone.includes(q) ||
      a.reraNumber.toLowerCase().includes(q)
    )
  }
  return result
})

const totalAgents = agents.length
const activeAgents = agents.filter(a => a.status === 'active').length
const pendingAgents = agents.filter(a => a.status === 'pending').length

function getTierColor(tier: string) {
  const map: Record<string, string> = {
    platinum: 'primary', gold: 'warning', silver: 'neutral', bronze: 'error',
  }
  return map[tier] || 'neutral'
}

function getStatusColor(status: string) {
  const map: Record<string, string> = {
    active: 'success', pending: 'warning', inactive: 'error',
  }
  return map[status] || 'neutral'
}

// Invite agent modal
const showInviteModal = ref(false)
const inviteData = reactive({
  phone: '',
  reraNumber: '',
  name: '',
})

function inviteAgent() {
  console.log('[Agent] Invite sent:', { ...inviteData })
  showInviteModal.value = false
  Object.assign(inviteData, { phone: '', reraNumber: '', name: '' })
}

// Performance modal
const showPerformanceModal = ref(false)
const selectedAgent = ref<typeof agents[0] | null>(null)

function viewPerformance(agent: typeof agents[0]) {
  selectedAgent.value = agent
  showPerformanceModal.value = true
}

function changeTier(agentId: string, tier: string) {
  console.log('[Agent] Tier change:', agentId, '->', tier)
}

function toggleActive(agentId: string) {
  console.log('[Agent] Toggle active:', agentId)
}
</script>

<template>
  <div>
    <div class="flex items-center justify-between mb-6">
      <div>
        <h1 class="text-2xl font-bold text-brand-900 dark:text-white">Agents</h1>
        <p class="text-slate-500 dark:text-slate-400 text-sm mt-1">Manage channel partner agents</p>
      </div>
      <UButton icon="i-lucide-user-plus" @click="showInviteModal = true">Invite Agent</UButton>
    </div>

    <!-- Stats row -->
    <div class="grid md:grid-cols-3 gap-4 mb-6">
      <UCard>
        <div class="flex items-center gap-3">
          <div class="p-2 bg-blue-50 rounded-lg">
            <UIcon name="i-lucide-users" class="text-xl text-blue-500" />
          </div>
          <div>
            <p class="text-xs text-slate-500 dark:text-slate-400">Total Agents</p>
            <p class="text-2xl font-bold text-brand-900 dark:text-white">{{ totalAgents }}</p>
          </div>
        </div>
      </UCard>
      <UCard>
        <div class="flex items-center gap-3">
          <div class="p-2 bg-green-50 rounded-lg">
            <UIcon name="i-lucide-user-check" class="text-xl text-green-500" />
          </div>
          <div>
            <p class="text-xs text-slate-500 dark:text-slate-400">Active This Month</p>
            <p class="text-2xl font-bold text-accent-600 dark:text-accent-400">{{ activeAgents }}</p>
          </div>
        </div>
      </UCard>
      <UCard>
        <div class="flex items-center gap-3">
          <div class="p-2 bg-amber-50 rounded-lg">
            <UIcon name="i-lucide-clock" class="text-xl text-amber-500" />
          </div>
          <div>
            <p class="text-xs text-slate-500 dark:text-slate-400">Pending Approval</p>
            <p class="text-2xl font-bold text-amber-600">{{ pendingAgents }}</p>
          </div>
        </div>
      </UCard>
    </div>

    <UCard>
      <div class="flex items-center gap-3 mb-4">
        <UInput v-model="searchQuery" placeholder="Search agents..." icon="i-lucide-search" class="flex-1" />
        <USelect v-model="statusFilter" :items="statusOptions" placeholder="Status" />
      </div>

      <div class="overflow-x-auto">
        <table class="w-full text-sm">
          <thead>
            <tr class="border-b border-slate-200 dark:border-slate-700">
              <th class="text-left py-3 px-3 text-xs font-semibold text-slate-500 dark:text-slate-400 uppercase">Agent</th>
              <th class="text-left py-3 px-3 text-xs font-semibold text-slate-500 dark:text-slate-400 uppercase">Phone</th>
              <th class="text-left py-3 px-3 text-xs font-semibold text-slate-500 dark:text-slate-400 uppercase">RERA</th>
              <th class="text-left py-3 px-3 text-xs font-semibold text-slate-500 dark:text-slate-400 uppercase">Tier</th>
              <th class="text-left py-3 px-3 text-xs font-semibold text-slate-500 dark:text-slate-400 uppercase">Leads</th>
              <th class="text-left py-3 px-3 text-xs font-semibold text-slate-500 dark:text-slate-400 uppercase">Bookings</th>
              <th class="text-left py-3 px-3 text-xs font-semibold text-slate-500 dark:text-slate-400 uppercase">Conv. %</th>
              <th class="text-left py-3 px-3 text-xs font-semibold text-slate-500 dark:text-slate-400 uppercase">Status</th>
              <th class="text-left py-3 px-3 text-xs font-semibold text-slate-500 dark:text-slate-400 uppercase">Actions</th>
            </tr>
          </thead>
          <tbody>
            <tr
              v-for="agent in filteredAgents"
              :key="agent.id"
              class="border-b border-slate-100 dark:border-slate-800 hover:bg-slate-50 dark:hover:bg-brand-800"
            >
              <td class="py-3 px-3">
                <div class="flex items-center gap-2">
                  <div class="w-8 h-8 rounded-full bg-brand-800 dark:bg-brand-700 flex items-center justify-center text-white text-xs font-bold flex-shrink-0">
                    {{ agent.name.split(' ').map((n: string) => n[0]).join('') }}
                  </div>
                  <div>
                    <p class="font-medium text-brand-900 dark:text-white">{{ agent.name }}</p>
                    <p class="text-xs text-slate-400">{{ agent.operatingAreas.join(', ') }}</p>
                  </div>
                </div>
              </td>
              <td class="py-3 px-3">
                <a :href="`tel:+91${agent.phone}`" class="text-brand-600 dark:text-brand-400 text-xs">
                  {{ formatPhone(agent.phone) }}
                </a>
              </td>
              <td class="py-3 px-3">
                <div class="flex items-center gap-1">
                  <span class="text-xs font-mono text-slate-500 dark:text-slate-400">{{ agent.reraNumber }}</span>
                  <UIcon
                    :name="agent.reraVerified ? 'i-lucide-check-circle' : 'i-lucide-alert-circle'"
                    :class="agent.reraVerified ? 'text-green-500' : 'text-amber-500'"
                    class="text-xs"
                  />
                </div>
              </td>
              <td class="py-3 px-3">
                <UBadge :color="getTierColor(agent.tier)" variant="subtle" size="xs">
                  {{ agent.tier }}
                </UBadge>
              </td>
              <td class="py-3 px-3 text-slate-600 dark:text-slate-300">{{ agent.totalLeads }}</td>
              <td class="py-3 px-3 text-slate-600 dark:text-slate-300">{{ agent.totalBookings }}</td>
              <td class="py-3 px-3">
                <span :class="agent.conversionRate > 15 ? 'text-green-600' : agent.conversionRate > 0 ? 'text-amber-600' : 'text-slate-400'">
                  {{ agent.conversionRate }}%
                </span>
              </td>
              <td class="py-3 px-3">
                <UBadge :color="getStatusColor(agent.status)" variant="subtle" size="xs">
                  {{ agent.status }}
                </UBadge>
              </td>
              <td class="py-3 px-3">
                <div class="flex items-center gap-1">
                  <UButton
                    icon="i-lucide-bar-chart-2"
                    variant="ghost"
                    color="neutral"
                    size="xs"
                    title="View Performance"
                    @click="viewPerformance(agent)"
                  />
                  <USelect
                    :model-value="agent.tier"
                    :items="[
                      { label: 'Bronze', value: 'bronze' },
                      { label: 'Silver', value: 'silver' },
                      { label: 'Gold', value: 'gold' },
                      { label: 'Platinum', value: 'platinum' },
                    ]"
                    size="xs"
                    class="w-24"
                    @update:model-value="(val: string) => changeTier(agent.id, val)"
                  />
                  <UButton
                    :icon="agent.isActive ? 'i-lucide-user-x' : 'i-lucide-user-check'"
                    variant="ghost"
                    :color="agent.isActive ? 'error' : 'success'"
                    size="xs"
                    :title="agent.isActive ? 'Deactivate' : 'Activate'"
                    @click="toggleActive(agent.id)"
                  />
                </div>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </UCard>

    <!-- Invite Agent Modal -->
    <UModal v-model:open="showInviteModal">
      <template #content>
        <div class="p-6">
          <h3 class="text-lg font-semibold text-brand-900 dark:text-white mb-4">Invite New Agent</h3>

          <div class="space-y-4">
            <UFormField label="Agent Name">
              <UInput v-model="inviteData.name" placeholder="Full name" class="w-full" />
            </UFormField>
            <UFormField label="Phone Number">
              <UInput v-model="inviteData.phone" placeholder="+91 98765 43210" type="tel" class="w-full" />
            </UFormField>
            <UFormField label="RERA Number">
              <UInput v-model="inviteData.reraNumber" placeholder="A51900XXXXXX" class="w-full" />
            </UFormField>

            <div class="flex gap-3 pt-2">
              <UButton variant="outline" class="flex-1" @click="showInviteModal = false">Cancel</UButton>
              <UButton
                class="flex-1"
                :disabled="!inviteData.phone || !inviteData.reraNumber"
                @click="inviteAgent"
              >
                Send Invite
              </UButton>
            </div>
          </div>
        </div>
      </template>
    </UModal>

    <!-- Performance Modal -->
    <UModal v-model:open="showPerformanceModal">
      <template #content>
        <div v-if="selectedAgent" class="p-6">
          <h3 class="text-lg font-semibold text-brand-900 dark:text-white mb-4">{{ selectedAgent.name }} - Performance</h3>

          <div class="grid grid-cols-2 gap-4 mb-4">
            <div class="p-3 bg-slate-50 dark:bg-brand-900/50 rounded-lg text-center">
              <p class="text-2xl font-bold text-brand-900 dark:text-white">{{ selectedAgent.totalLeads }}</p>
              <p class="text-xs text-slate-500 dark:text-slate-400">Total Leads</p>
            </div>
            <div class="p-3 bg-slate-50 dark:bg-brand-900/50 rounded-lg text-center">
              <p class="text-2xl font-bold text-accent-600 dark:text-accent-400">{{ selectedAgent.totalBookings }}</p>
              <p class="text-xs text-slate-500 dark:text-slate-400">Total Bookings</p>
            </div>
            <div class="p-3 bg-slate-50 dark:bg-brand-900/50 rounded-lg text-center">
              <p class="text-2xl font-bold text-green-600">{{ selectedAgent.conversionRate }}%</p>
              <p class="text-xs text-slate-500 dark:text-slate-400">Conversion Rate</p>
            </div>
            <div class="p-3 bg-slate-50 dark:bg-brand-900/50 rounded-lg text-center">
              <p class="text-2xl font-bold text-brand-900 dark:text-white">{{ selectedAgent.experienceYears }} yrs</p>
              <p class="text-xs text-slate-500 dark:text-slate-400">Experience</p>
            </div>
          </div>

          <div class="mb-4">
            <p class="text-sm font-medium text-brand-900 dark:text-white mb-2">Operating Areas</p>
            <div class="flex flex-wrap gap-2">
              <UBadge v-for="area in selectedAgent.operatingAreas" :key="area" color="primary" variant="subtle" size="sm">
                {{ area }}
              </UBadge>
            </div>
          </div>

          <div class="flex items-center gap-2 mb-4">
            <p class="text-sm text-slate-500 dark:text-slate-400">RERA:</p>
            <span class="text-sm font-mono">{{ selectedAgent.reraNumber }}</span>
            <UIcon
              :name="selectedAgent.reraVerified ? 'i-lucide-check-circle' : 'i-lucide-alert-circle'"
              :class="selectedAgent.reraVerified ? 'text-green-500' : 'text-amber-500'"
            />
          </div>

          <div class="flex items-center gap-2">
            <p class="text-sm text-slate-500 dark:text-slate-400">Current Tier:</p>
            <UBadge :color="getTierColor(selectedAgent.tier)" variant="solid" size="sm">
              {{ selectedAgent.tier.toUpperCase() }}
            </UBadge>
          </div>

          <div class="mt-6">
            <UButton class="w-full" variant="outline" @click="showPerformanceModal = false">Close</UButton>
          </div>
        </div>
      </template>
    </UModal>
  </div>
</template>
