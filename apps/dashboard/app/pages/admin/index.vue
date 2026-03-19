<script setup lang="ts">
definePageMeta({ layout: 'admin', middleware: 'auth' })

const { leads, agents, projects, recentActivity } = useMockData()
const { formatINR, formatRelativeTime } = useFormatting()

// KPI stats
const totalLeads = leads.length
const activeAgents = agents.filter(a => a.isActive).length
const projectsListed = projects.filter(p => p.status === 'active').length
const revenueThisMonth = 450000

// Conversion funnel
const funnelData = computed(() => {
  const total = leads.length
  const contacted = leads.filter(l => ['contacted', 'site_visit', 'booked'].includes(l.status)).length
  const visited = leads.filter(l => ['site_visit', 'booked'].includes(l.status)).length
  const booked = leads.filter(l => l.status === 'booked').length

  return [
    { label: 'Total Leads', count: total, width: 100, color: 'bg-blue-500' },
    { label: 'Contacted', count: contacted, width: Math.round((contacted / total) * 100), color: 'bg-amber-500' },
    { label: 'Site Visits', count: visited, width: Math.round((visited / total) * 100), color: 'bg-accent-500 dark:bg-accent-400' },
    { label: 'Bookings', count: booked, width: Math.round((booked / total) * 100), color: 'bg-green-500' },
  ]
})

// Top performing agents
const topAgents = computed(() =>
  [...agents]
    .filter(a => a.isActive && a.totalBookings > 0)
    .sort((a, b) => b.conversionRate - a.conversionRate)
    .slice(0, 5)
)

function getActivityIcon(action: string) {
  const map: Record<string, string> = {
    lead_created: 'i-lucide-user-plus',
    site_visit_completed: 'i-lucide-map-pin',
    booking_created: 'i-lucide-check-circle',
    commission_paid: 'i-lucide-indian-rupee',
    agent_onboarded: 'i-lucide-user-check',
    lead_status_changed: 'i-lucide-arrow-right',
    site_visit_scheduled: 'i-lucide-calendar',
  }
  return map[action] || 'i-lucide-activity'
}

function getActivityColor(action: string) {
  const map: Record<string, string> = {
    lead_created: 'text-blue-500',
    site_visit_completed: 'text-brand-600 dark:text-brand-400',
    booking_created: 'text-green-500',
    commission_paid: 'text-green-600',
    agent_onboarded: 'text-purple-500',
    lead_status_changed: 'text-amber-500',
    site_visit_scheduled: 'text-blue-400',
  }
  return map[action] || 'text-slate-400'
}
</script>

<template>
  <div>
    <h1 class="text-2xl font-bold text-brand-900 dark:text-white mb-6">Dashboard Overview</h1>

    <!-- KPI Cards -->
    <div class="grid grid-cols-4 gap-4 mb-8">
      <UCard>
        <div class="flex items-center gap-3">
          <div class="p-2 bg-blue-50 rounded-lg">
            <UIcon name="i-lucide-users" class="text-xl text-blue-500" />
          </div>
          <div>
            <p class="text-sm text-slate-500 dark:text-slate-400">Total Leads</p>
            <p class="text-2xl font-bold text-brand-900 dark:text-white">{{ totalLeads }}</p>
          </div>
        </div>
      </UCard>
      <UCard>
        <div class="flex items-center gap-3">
          <div class="p-2 bg-green-50 rounded-lg">
            <UIcon name="i-lucide-user-check" class="text-xl text-green-500" />
          </div>
          <div>
            <p class="text-sm text-slate-500 dark:text-slate-400">Active Agents</p>
            <p class="text-2xl font-bold text-accent-600 dark:text-accent-400">{{ activeAgents }}</p>
          </div>
        </div>
      </UCard>
      <UCard>
        <div class="flex items-center gap-3">
          <div class="p-2 bg-purple-50 rounded-lg">
            <UIcon name="i-lucide-building-2" class="text-xl text-purple-500" />
          </div>
          <div>
            <p class="text-sm text-slate-500 dark:text-slate-400">Projects Listed</p>
            <p class="text-2xl font-bold text-brand-900 dark:text-white">{{ projectsListed }}</p>
          </div>
        </div>
      </UCard>
      <UCard>
        <div class="flex items-center gap-3">
          <div class="p-2 bg-amber-50 rounded-lg">
            <UIcon name="i-lucide-indian-rupee" class="text-xl text-amber-500" />
          </div>
          <div>
            <p class="text-sm text-slate-500 dark:text-slate-400">Revenue This Month</p>
            <p class="text-2xl font-bold text-accent-600 dark:text-accent-400">{{ formatINR(revenueThisMonth) }}</p>
          </div>
        </div>
      </UCard>
    </div>

    <div class="grid grid-cols-2 gap-6 mb-8">
      <!-- Conversion Funnel -->
      <UCard>
        <template #header>
          <h2 class="text-lg font-semibold text-brand-900 dark:text-white">Conversion Funnel</h2>
        </template>
        <div class="space-y-4">
          <div v-for="stage in funnelData" :key="stage.label">
            <div class="flex items-center justify-between text-sm mb-1">
              <span class="text-slate-600 dark:text-slate-300">{{ stage.label }}</span>
              <span class="font-semibold text-brand-900 dark:text-white">{{ stage.count }}</span>
            </div>
            <div class="w-full bg-slate-100 dark:bg-brand-800 rounded-full h-6 flex items-center">
              <div
                :class="[stage.color, 'h-6 rounded-full transition-all flex items-center justify-end pr-2']"
                :style="{ width: `${Math.max(stage.width, 8)}%` }"
              >
                <span v-if="stage.width > 15" class="text-xs text-white font-medium">{{ stage.width }}%</span>
              </div>
            </div>
          </div>
        </div>
      </UCard>

      <!-- Top Performing Agents -->
      <UCard>
        <template #header>
          <h2 class="text-lg font-semibold text-brand-900 dark:text-white">Top Performing Agents</h2>
        </template>
        <div class="space-y-3">
          <div
            v-for="(agent, index) in topAgents"
            :key="agent.id"
            class="flex items-center gap-3 p-2 rounded-lg hover:bg-slate-50 dark:hover:bg-brand-800"
          >
            <span class="text-sm font-bold text-slate-400 w-6">{{ index + 1 }}</span>
            <div class="w-8 h-8 rounded-full bg-brand-800 dark:bg-brand-700 flex items-center justify-center text-white text-xs font-bold">
              {{ agent.name.split(' ').map((n: string) => n[0]).join('') }}
            </div>
            <div class="flex-1 min-w-0">
              <p class="text-sm font-medium text-brand-900 dark:text-white truncate">{{ agent.name }}</p>
              <p class="text-xs text-slate-500 dark:text-slate-400">{{ agent.totalBookings }} bookings</p>
            </div>
            <div class="text-right">
              <p class="text-sm font-semibold text-accent-600 dark:text-accent-400">{{ agent.conversionRate }}%</p>
              <p class="text-xs text-slate-400">conversion</p>
            </div>
          </div>
        </div>
      </UCard>
    </div>

    <!-- Recent Activity Feed -->
    <UCard>
      <template #header>
        <h2 class="text-lg font-semibold text-brand-900 dark:text-white">Recent Activity</h2>
      </template>
      <div class="space-y-3">
        <div
          v-for="activity in recentActivity"
          :key="activity.id"
          class="flex items-start gap-3 p-2 rounded-lg hover:bg-slate-50 dark:hover:bg-brand-800"
        >
          <div class="mt-0.5">
            <UIcon :name="getActivityIcon(activity.action)" :class="['text-lg', getActivityColor(activity.action)]" />
          </div>
          <div class="flex-1 min-w-0">
            <p class="text-sm text-slate-700 dark:text-slate-300">{{ activity.description }}</p>
            <p class="text-xs text-slate-400">by {{ activity.actor }}</p>
          </div>
          <span class="text-xs text-slate-400 whitespace-nowrap">{{ formatRelativeTime(activity.time) }}</span>
        </div>
      </div>
    </UCard>
  </div>
</template>
