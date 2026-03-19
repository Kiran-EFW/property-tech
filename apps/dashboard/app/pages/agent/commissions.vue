<script setup lang="ts">
definePageMeta({ layout: 'agent', middleware: 'auth' })

const { commissions } = useMockData()
const { formatINR, formatDate } = useFormatting()

const agentId = 'agt-001'

const myCommissions = computed(() => commissions.filter(c => c.agentId === agentId))

const totalEarned = computed(() =>
  myCommissions.value.filter(c => c.status === 'paid').reduce((sum, c) => sum + c.netPayableToAgent, 0)
)
const totalPending = computed(() =>
  myCommissions.value.filter(c => c.status === 'pending' || c.status === 'approved').reduce((sum, c) => sum + c.netPayableToAgent, 0)
)
const totalPaid = computed(() =>
  myCommissions.value.filter(c => c.status === 'paid').reduce((sum, c) => sum + c.netPayableToAgent, 0)
)

function getStatusColor(status: string) {
  const map: Record<string, string> = { paid: 'success', approved: 'info', pending: 'warning', rejected: 'error' }
  return map[status] || 'neutral'
}

const columns = [
  { key: 'bookingRef', label: 'Booking' },
  { key: 'projectName', label: 'Project' },
  { key: 'investorName', label: 'Investor' },
  { key: 'agentCommission', label: 'Amount' },
  { key: 'tdsAmount', label: 'TDS' },
  { key: 'netPayableToAgent', label: 'Net' },
  { key: 'status', label: 'Status' },
  { key: 'createdAt', label: 'Date' },
]
</script>

<template>
  <div>
    <h1 class="text-xl font-bold text-brand-900 dark:text-white">Earnings</h1>
    <p class="text-slate-500 dark:text-slate-400 text-sm mt-1 mb-6">Track your commissions and payouts</p>

    <!-- Summary stats -->
    <div class="grid grid-cols-3 gap-3 mb-6">
      <UCard>
        <p class="text-xs text-slate-500 dark:text-slate-400">Total Earned</p>
        <p class="text-lg font-bold text-brand-900 dark:text-white mt-1">{{ formatINR(totalEarned) }}</p>
      </UCard>
      <UCard>
        <p class="text-xs text-slate-500 dark:text-slate-400">Pending</p>
        <p class="text-lg font-bold text-amber-600 mt-1">{{ formatINR(totalPending) }}</p>
      </UCard>
      <UCard>
        <p class="text-xs text-slate-500 dark:text-slate-400">Paid Out</p>
        <p class="text-lg font-bold text-green-600 mt-1">{{ formatINR(totalPaid) }}</p>
      </UCard>
    </div>

    <!-- Transaction history -->
    <h2 class="text-sm font-semibold text-brand-900 dark:text-white mb-3">Transaction History</h2>

    <!-- Mobile card view -->
    <div class="space-y-3">
      <UCard v-for="c in myCommissions" :key="c.id">
        <div class="flex items-start justify-between mb-2">
          <div>
            <p class="text-sm font-medium text-brand-900 dark:text-white">{{ c.bookingRef }}</p>
            <p class="text-xs text-slate-500 dark:text-slate-400">{{ c.projectName }}</p>
          </div>
          <UBadge :color="getStatusColor(c.status)" variant="subtle" size="xs">
            {{ c.status }}
          </UBadge>
        </div>

        <div class="grid grid-cols-2 gap-2 text-xs">
          <div>
            <span class="text-slate-400">Investor:</span>
            <span class="ml-1 text-slate-600 dark:text-slate-300">{{ c.investorName }}</span>
          </div>
          <div>
            <span class="text-slate-400">Date:</span>
            <span class="ml-1 text-slate-600 dark:text-slate-300">{{ formatDate(c.createdAt) }}</span>
          </div>
          <div>
            <span class="text-slate-400">Commission:</span>
            <span class="ml-1 font-medium text-brand-900 dark:text-white">{{ formatINR(c.agentCommission) }}</span>
          </div>
          <div>
            <span class="text-slate-400">TDS (5%):</span>
            <span class="ml-1 text-red-500">-{{ formatINR(c.tdsAmount) }}</span>
          </div>
        </div>

        <USeparator class="my-2" />

        <div class="flex items-center justify-between">
          <span class="text-xs text-slate-500 dark:text-slate-400">Net Payable</span>
          <span class="text-sm font-bold text-green-600">{{ formatINR(c.netPayableToAgent) }}</span>
        </div>
      </UCard>

      <div v-if="myCommissions.length === 0" class="text-center py-12">
        <UIcon name="i-lucide-wallet" class="text-4xl text-slate-300 mb-2" />
        <p class="text-slate-400">No commission records yet</p>
      </div>
    </div>
  </div>
</template>
