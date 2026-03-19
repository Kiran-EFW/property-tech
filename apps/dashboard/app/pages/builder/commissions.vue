<script setup lang="ts">
definePageMeta({ layout: 'builder', middleware: 'auth' })

const { commissions, units } = useMockData()
const { formatINR, formatDate } = useFormatting()

// Builder's projects
const builderProjectIds = ['prj-001', 'prj-002']

const builderCommissions = computed(() => commissions.filter(c => builderProjectIds.includes(c.projectId)))
const unitsSoldViaAgents = computed(() => units.filter(u => builderProjectIds.includes(u.projectId) && u.status === 'sold').length)

const totalPaid = computed(() =>
  builderCommissions.value.filter(c => c.status === 'paid').reduce((sum, c) => sum + c.totalBrokerage, 0)
)
const totalPending = computed(() =>
  builderCommissions.value.filter(c => c.status === 'pending' || c.status === 'approved').reduce((sum, c) => sum + c.totalBrokerage, 0)
)

function getStatusColor(status: string) {
  const map: Record<string, string> = {
    paid: 'success', approved: 'info', pending: 'warning', rejected: 'error',
  }
  return map[status] || 'neutral'
}
</script>

<template>
  <div>
    <div class="flex items-center justify-between mb-6">
      <div>
        <h1 class="text-2xl font-bold text-brand-900 dark:text-white">Commissions</h1>
        <p class="text-slate-500 dark:text-slate-400 text-sm mt-1">Commission structure and payment history</p>
      </div>
    </div>

    <!-- Stats -->
    <div class="grid md:grid-cols-3 gap-4 mb-6">
      <UCard>
        <div class="flex items-center gap-3">
          <div class="p-2 bg-green-50 rounded-lg">
            <UIcon name="i-lucide-check-circle" class="text-xl text-green-500" />
          </div>
          <div>
            <p class="text-xs text-slate-500 dark:text-slate-400">Total Commission Paid</p>
            <p class="text-xl font-bold text-brand-900 dark:text-white">{{ formatINR(totalPaid) }}</p>
          </div>
        </div>
      </UCard>
      <UCard>
        <div class="flex items-center gap-3">
          <div class="p-2 bg-amber-50 rounded-lg">
            <UIcon name="i-lucide-clock" class="text-xl text-amber-500" />
          </div>
          <div>
            <p class="text-xs text-slate-500 dark:text-slate-400">Pending</p>
            <p class="text-xl font-bold text-amber-600">{{ formatINR(totalPending) }}</p>
          </div>
        </div>
      </UCard>
      <UCard>
        <div class="flex items-center gap-3">
          <div class="p-2 bg-blue-50 rounded-lg">
            <UIcon name="i-lucide-home" class="text-xl text-blue-500" />
          </div>
          <div>
            <p class="text-xs text-slate-500 dark:text-slate-400">Units Sold via Agents</p>
            <p class="text-xl font-bold text-accent-600 dark:text-accent-400">{{ unitsSoldViaAgents }}</p>
          </div>
        </div>
      </UCard>
    </div>

    <!-- Payment History -->
    <UCard>
      <template #header>
        <h2 class="font-semibold text-brand-900 dark:text-white">Payment History</h2>
      </template>

      <div class="overflow-x-auto">
        <table class="w-full text-sm">
          <thead>
            <tr class="border-b border-slate-200 dark:border-slate-700">
              <th class="text-left py-3 px-3 text-xs font-semibold text-slate-500 dark:text-slate-400 uppercase">Booking Ref</th>
              <th class="text-left py-3 px-3 text-xs font-semibold text-slate-500 dark:text-slate-400 uppercase">Project</th>
              <th class="text-left py-3 px-3 text-xs font-semibold text-slate-500 dark:text-slate-400 uppercase">Agent</th>
              <th class="text-left py-3 px-3 text-xs font-semibold text-slate-500 dark:text-slate-400 uppercase">Investor</th>
              <th class="text-right py-3 px-3 text-xs font-semibold text-slate-500 dark:text-slate-400 uppercase">Agreement Value</th>
              <th class="text-right py-3 px-3 text-xs font-semibold text-slate-500 dark:text-slate-400 uppercase">Brokerage Rate</th>
              <th class="text-right py-3 px-3 text-xs font-semibold text-slate-500 dark:text-slate-400 uppercase">Total Brokerage</th>
              <th class="text-left py-3 px-3 text-xs font-semibold text-slate-500 dark:text-slate-400 uppercase">Status</th>
              <th class="text-left py-3 px-3 text-xs font-semibold text-slate-500 dark:text-slate-400 uppercase">Date</th>
            </tr>
          </thead>
          <tbody>
            <tr
              v-for="c in builderCommissions"
              :key="c.id"
              class="border-b border-slate-100 dark:border-slate-800 hover:bg-slate-50 dark:hover:bg-brand-800"
            >
              <td class="py-3 px-3 text-xs font-mono text-slate-500 dark:text-slate-400">{{ c.bookingRef }}</td>
              <td class="py-3 px-3 text-slate-600 dark:text-slate-300">{{ c.projectName }}</td>
              <td class="py-3 px-3 text-slate-600 dark:text-slate-300">{{ c.agentName }}</td>
              <td class="py-3 px-3 text-slate-600 dark:text-slate-300">{{ c.investorName }}</td>
              <td class="py-3 px-3 text-right font-medium text-brand-900 dark:text-white">{{ formatINR(c.agreementValue) }}</td>
              <td class="py-3 px-3 text-right text-slate-600 dark:text-slate-300">{{ c.brokerageRate }}%</td>
              <td class="py-3 px-3 text-right font-semibold text-brand-900 dark:text-white">{{ formatINR(c.totalBrokerage) }}</td>
              <td class="py-3 px-3">
                <UBadge :color="getStatusColor(c.status)" variant="subtle" size="xs">
                  {{ c.status }}
                </UBadge>
              </td>
              <td class="py-3 px-3 text-xs text-slate-500 dark:text-slate-400">{{ formatDate(c.createdAt) }}</td>
            </tr>
          </tbody>
        </table>
      </div>

      <div v-if="builderCommissions.length === 0" class="text-center py-12">
        <UIcon name="i-lucide-indian-rupee" class="text-4xl text-slate-300 mb-2" />
        <p class="text-slate-400 dark:text-slate-500">No commission payments recorded yet</p>
      </div>
    </UCard>
  </div>
</template>
