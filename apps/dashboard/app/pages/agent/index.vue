<script setup lang="ts">
definePageMeta({ layout: 'agent', middleware: 'auth' })

const { leads, visits } = useMockData()
const { formatINR, formatRelativeTime, formatDateTime } = useFormatting()

// Current agent mock ID
const agentId = 'agt-001'

const myLeads = computed(() => leads.filter(l => l.agentId === agentId))
const myVisits = computed(() => visits.filter(v => v.agentId === agentId))

const stats = computed(() => ({
  newLeads: myLeads.value.filter(l => l.status === 'new').length,
  visitsToday: myVisits.value.filter(v => {
    const d = new Date(v.scheduledAt)
    const now = new Date()
    return d.toDateString() === now.toDateString() && v.status === 'scheduled'
  }).length,
  followUpsDue: myLeads.value.filter(l => {
    if (!l.nextFollowUpAt) return false
    return new Date(l.nextFollowUpAt) <= new Date(Date.now() + 86400000)
  }).length,
  thisMonthEarnings: 125000,
}))

const upcomingTasks = computed(() => {
  const tasks: Array<{ id: string; type: string; title: string; subtitle: string; time: string; icon: string; color: string }> = []

  // Follow-ups due
  myLeads.value
    .filter(l => l.nextFollowUpAt && new Date(l.nextFollowUpAt) <= new Date(Date.now() + 2 * 86400000))
    .sort((a, b) => new Date(a.nextFollowUpAt!).getTime() - new Date(b.nextFollowUpAt!).getTime())
    .forEach(l => {
      tasks.push({
        id: l.id,
        type: 'follow-up',
        title: `Follow up with ${l.name}`,
        subtitle: `${l.projectName} - ${l.preferredConfiguration}`,
        time: formatRelativeTime(l.nextFollowUpAt!),
        icon: 'i-lucide-phone-outgoing',
        color: 'text-amber-600',
      })
    })

  // Upcoming visits
  myVisits.value
    .filter(v => v.status === 'scheduled')
    .sort((a, b) => new Date(a.scheduledAt).getTime() - new Date(b.scheduledAt).getTime())
    .forEach(v => {
      tasks.push({
        id: v.id,
        type: 'visit',
        title: `Site visit with ${v.investorName}`,
        subtitle: v.projectName,
        time: formatDateTime(v.scheduledAt),
        icon: 'i-lucide-map-pin',
        color: 'text-brand-600 dark:text-brand-400',
      })
    })

  return tasks.slice(0, 6)
})
</script>

<template>
  <div>
    <h2 class="text-xl font-semibold text-brand-900 dark:text-white mb-4">Today's Overview</h2>

    <!-- Stats cards -->
    <div class="grid grid-cols-2 gap-3 mb-6">
      <UCard class="text-center">
        <p class="text-2xl font-bold text-brand-900 dark:text-white">{{ stats.newLeads }}</p>
        <p class="text-xs text-slate-500 dark:text-slate-400">New Leads</p>
      </UCard>
      <UCard class="text-center">
        <p class="text-2xl font-bold text-brand-600 dark:text-brand-400">{{ stats.visitsToday }}</p>
        <p class="text-xs text-slate-500 dark:text-slate-400">Visits Today</p>
      </UCard>
      <UCard class="text-center">
        <p class="text-2xl font-bold text-amber-600">{{ stats.followUpsDue }}</p>
        <p class="text-xs text-slate-500 dark:text-slate-400">Follow-ups Due</p>
      </UCard>
      <UCard class="text-center">
        <p class="text-2xl font-bold text-green-600">{{ formatINR(stats.thisMonthEarnings) }}</p>
        <p class="text-xs text-slate-500 dark:text-slate-400">This Month</p>
      </UCard>
    </div>

    <!-- Quick Actions -->
    <div class="flex gap-3 mb-6">
      <UButton
        icon="i-lucide-users"
        variant="outline"
        class="flex-1"
        @click="navigateTo('/agent/leads')"
      >
        View Leads
      </UButton>
      <UButton
        icon="i-lucide-calendar-plus"
        color="primary"
        class="flex-1"
        @click="navigateTo('/agent/visits')"
      >
        Schedule Visit
      </UButton>
    </div>

    <!-- Today's Tasks -->
    <h3 class="text-sm font-medium text-slate-500 dark:text-slate-400 uppercase tracking-wide mb-3">Upcoming Tasks</h3>
    <div class="space-y-3">
      <UCard
        v-for="task in upcomingTasks"
        :key="task.id"
        class="cursor-pointer hover:shadow-md transition-shadow"
        @click="task.type === 'follow-up' ? navigateTo(`/agent/leads/${task.id}`) : navigateTo('/agent/visits')"
      >
        <div class="flex items-start gap-3">
          <div class="mt-0.5">
            <UIcon :name="task.icon" :class="['text-xl', task.color]" />
          </div>
          <div class="flex-1 min-w-0">
            <p class="text-sm font-medium text-brand-900 dark:text-white truncate">{{ task.title }}</p>
            <p class="text-xs text-slate-500 dark:text-slate-400">{{ task.subtitle }}</p>
          </div>
          <span class="text-xs text-slate-400 whitespace-nowrap">{{ task.time }}</span>
        </div>
      </UCard>

      <div v-if="upcomingTasks.length === 0" class="text-center py-8">
        <UIcon name="i-lucide-check-circle" class="text-3xl text-green-400 mb-2" />
        <p class="text-sm text-slate-400">All caught up! No pending tasks.</p>
      </div>
    </div>
  </div>
</template>
