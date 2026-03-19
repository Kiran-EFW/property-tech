<script setup lang="ts">
definePageMeta({ layout: 'builder', middleware: 'auth' })

const { projects, leads, units, commissions } = useMockData()
const { formatINR, formatRelativeTime, formatINRShort } = useFormatting()

// Builder mock data - assume this builder owns projects prj-001 and prj-002
const builderProjectIds = ['prj-001', 'prj-002']

const builderProjects = computed(() => projects.filter(p => builderProjectIds.includes(p.id)))
const builderUnits = computed(() => units.filter(u => builderProjectIds.includes(u.projectId)))
const builderLeads = computed(() => leads.filter(l => l.projectId && builderProjectIds.includes(l.projectId)))
const builderCommissions = computed(() => commissions.filter(c => builderProjectIds.includes(c.projectId)))

const activeProjects = computed(() => builderProjects.value.filter(p => p.status === 'active').length)
const totalAvailableUnits = computed(() => builderUnits.value.filter(u => u.status === 'available').length)
const totalLeadsReceived = computed(() => builderLeads.value.length)
const unitsSoldViaAgent = computed(() => builderUnits.value.filter(u => u.status === 'sold').length)

// Recent leads for the builder's projects
const recentLeads = computed(() =>
  [...builderLeads.value]
    .sort((a, b) => new Date(b.createdAt).getTime() - new Date(a.createdAt).getTime())
    .slice(0, 5)
)

// Inventory summary
const inventorySummary = computed(() => {
  const total = builderUnits.value.length
  const available = builderUnits.value.filter(u => u.status === 'available').length
  const booked = builderUnits.value.filter(u => u.status === 'booked').length
  const sold = builderUnits.value.filter(u => u.status === 'sold').length
  const blocked = builderUnits.value.filter(u => u.status === 'blocked').length
  return { total, available, booked, sold, blocked }
})

function getStatusColor(status: string) {
  const map: Record<string, string> = {
    new: 'info', contacted: 'warning', site_visit: 'primary', booked: 'success',
  }
  return map[status] || 'neutral'
}

function getStatusLabel(status: string) {
  const map: Record<string, string> = {
    new: 'New', contacted: 'Contacted', site_visit: 'Visit Done', booked: 'Booked',
  }
  return map[status] || status
}
</script>

<template>
  <div>
    <h1 class="text-2xl font-bold text-brand-900 dark:text-white mb-6">Builder Overview</h1>

    <!-- Stats -->
    <div class="grid md:grid-cols-4 gap-4 mb-8">
      <UCard>
        <div class="flex items-center gap-3">
          <div class="p-2 bg-blue-50 rounded-lg">
            <UIcon name="i-lucide-building-2" class="text-xl text-blue-500" />
          </div>
          <div>
            <p class="text-xs text-slate-500 dark:text-slate-400">Active Projects</p>
            <p class="text-2xl font-bold text-brand-900 dark:text-white">{{ activeProjects }}</p>
          </div>
        </div>
      </UCard>
      <UCard>
        <div class="flex items-center gap-3">
          <div class="p-2 bg-green-50 rounded-lg">
            <UIcon name="i-lucide-home" class="text-xl text-green-500" />
          </div>
          <div>
            <p class="text-xs text-slate-500 dark:text-slate-400">Units Available</p>
            <p class="text-2xl font-bold text-accent-600 dark:text-accent-400">{{ totalAvailableUnits }}</p>
          </div>
        </div>
      </UCard>
      <UCard>
        <div class="flex items-center gap-3">
          <div class="p-2 bg-purple-50 rounded-lg">
            <UIcon name="i-lucide-users" class="text-xl text-purple-500" />
          </div>
          <div>
            <p class="text-xs text-slate-500 dark:text-slate-400">Leads Received</p>
            <p class="text-2xl font-bold text-brand-900 dark:text-white">{{ totalLeadsReceived }}</p>
          </div>
        </div>
      </UCard>
      <UCard>
        <div class="flex items-center gap-3">
          <div class="p-2 bg-amber-50 rounded-lg">
            <UIcon name="i-lucide-check-circle" class="text-xl text-amber-500" />
          </div>
          <div>
            <p class="text-xs text-slate-500 dark:text-slate-400">Units Sold via PropTech</p>
            <p class="text-2xl font-bold text-brand-900 dark:text-white">{{ unitsSoldViaAgent }}</p>
          </div>
        </div>
      </UCard>
    </div>

    <div class="grid md:grid-cols-2 gap-6">
      <!-- Recent Leads -->
      <UCard>
        <template #header>
          <div class="flex items-center justify-between">
            <h2 class="font-semibold text-brand-900 dark:text-white">Recent Leads</h2>
            <UButton variant="link" size="xs" @click="navigateTo('/builder/leads')">View All</UButton>
          </div>
        </template>
        <div class="space-y-3">
          <div
            v-for="lead in recentLeads"
            :key="lead.id"
            class="flex items-center justify-between p-2 rounded-lg hover:bg-slate-50 dark:hover:bg-brand-800"
          >
            <div>
              <p class="text-sm font-medium text-brand-900 dark:text-white">{{ lead.name }}</p>
              <p class="text-xs text-slate-500 dark:text-slate-400">
                {{ lead.projectName }} - {{ lead.preferredConfiguration }}
              </p>
            </div>
            <div class="text-right">
              <UBadge :color="getStatusColor(lead.status)" variant="subtle" size="xs">
                {{ getStatusLabel(lead.status) }}
              </UBadge>
              <p class="text-xs text-slate-400 dark:text-slate-500 mt-1">{{ formatRelativeTime(lead.createdAt) }}</p>
            </div>
          </div>
        </div>
      </UCard>

      <!-- Inventory Status -->
      <UCard>
        <template #header>
          <div class="flex items-center justify-between">
            <h2 class="font-semibold text-brand-900 dark:text-white">Inventory Status</h2>
            <UButton variant="link" size="xs" @click="navigateTo('/builder/inventory')">Manage</UButton>
          </div>
        </template>

        <div class="space-y-4">
          <div class="flex items-center justify-between">
            <span class="text-sm text-slate-600 dark:text-slate-300">Total Units</span>
            <span class="font-semibold text-brand-900 dark:text-white">{{ inventorySummary.total }}</span>
          </div>

          <div>
            <div class="flex items-center justify-between text-sm mb-1">
              <span class="text-green-600">Available</span>
              <span class="font-medium">{{ inventorySummary.available }}</span>
            </div>
            <div class="w-full bg-slate-100 dark:bg-brand-800 rounded-full h-3">
              <div
                class="bg-green-500 h-3 rounded-full"
                :style="{ width: `${(inventorySummary.available / inventorySummary.total) * 100}%` }"
              />
            </div>
          </div>

          <div>
            <div class="flex items-center justify-between text-sm mb-1">
              <span class="text-blue-600">Booked</span>
              <span class="font-medium">{{ inventorySummary.booked }}</span>
            </div>
            <div class="w-full bg-slate-100 dark:bg-brand-800 rounded-full h-3">
              <div
                class="bg-blue-500 h-3 rounded-full"
                :style="{ width: `${(inventorySummary.booked / inventorySummary.total) * 100}%` }"
              />
            </div>
          </div>

          <div>
            <div class="flex items-center justify-between text-sm mb-1">
              <span class="text-red-600">Sold</span>
              <span class="font-medium">{{ inventorySummary.sold }}</span>
            </div>
            <div class="w-full bg-slate-100 dark:bg-brand-800 rounded-full h-3">
              <div
                class="bg-red-500 h-3 rounded-full"
                :style="{ width: `${(inventorySummary.sold / inventorySummary.total) * 100}%` }"
              />
            </div>
          </div>

          <div>
            <div class="flex items-center justify-between text-sm mb-1">
              <span class="text-amber-600">Blocked</span>
              <span class="font-medium">{{ inventorySummary.blocked }}</span>
            </div>
            <div class="w-full bg-slate-100 dark:bg-brand-800 rounded-full h-3">
              <div
                class="bg-amber-500 h-3 rounded-full"
                :style="{ width: `${(inventorySummary.blocked / inventorySummary.total) * 100}%` }"
              />
            </div>
          </div>
        </div>
      </UCard>
    </div>
  </div>
</template>
