<script setup lang="ts">
definePageMeta({ layout: 'admin', middleware: 'auth' })

const { commissions } = useMockData()
const { formatINR, formatDate } = useFormatting()

const totalPaid = computed(() =>
  commissions.filter(c => c.status === 'paid').reduce((sum, c) => sum + c.netPayableToAgent, 0)
)
const totalPending = computed(() =>
  commissions.filter(c => c.status === 'pending' || c.status === 'approved').reduce((sum, c) => sum + c.netPayableToAgent, 0)
)
const thisMonth = computed(() =>
  commissions
    .filter(c => {
      const d = new Date(c.createdAt)
      const now = new Date()
      return d.getMonth() === now.getMonth() && d.getFullYear() === now.getFullYear()
    })
    .reduce((sum, c) => sum + c.netPayableToAgent, 0)
)
const avgPerDeal = computed(() => {
  const paid = commissions.filter(c => c.status === 'paid')
  return paid.length > 0 ? Math.round(paid.reduce((sum, c) => sum + c.netPayableToAgent, 0) / paid.length) : 0
})

function getStatusColor(status: string) {
  const map: Record<string, string> = {
    paid: 'success', approved: 'info', pending: 'warning', rejected: 'error',
  }
  return map[status] || 'neutral'
}

function approveCommission(id: string) {
  console.log('[Commission] Approved:', id)
}

function rejectCommission(id: string) {
  console.log('[Commission] Rejected:', id)
}

function markPaid(id: string) {
  console.log('[Commission] Marked as paid:', id)
}
</script>

<template>
  <div>
    <div class="flex items-center justify-between mb-6">
      <div>
        <h1 class="text-2xl font-bold text-brand-900 dark:text-white">Commissions</h1>
        <p class="text-slate-500 dark:text-slate-400 text-sm mt-1">Track payouts and commission structure</p>
      </div>
      <UButton icon="i-lucide-settings" variant="outline">Commission Rules</UButton>
    </div>

    <!-- Stats -->
    <div class="grid md:grid-cols-4 gap-4 mb-6">
      <UCard>
        <div class="flex items-center gap-3">
          <div class="p-2 bg-green-50 rounded-lg">
            <UIcon name="i-lucide-check-circle" class="text-xl text-green-500" />
          </div>
          <div>
            <p class="text-xs text-slate-500 dark:text-slate-400">Total Paid Out</p>
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
            <p class="text-xs text-slate-500 dark:text-slate-400">Pending Payouts</p>
            <p class="text-xl font-bold text-amber-600">{{ formatINR(totalPending) }}</p>
          </div>
        </div>
      </UCard>
      <UCard>
        <div class="flex items-center gap-3">
          <div class="p-2 bg-blue-50 rounded-lg">
            <UIcon name="i-lucide-calendar" class="text-xl text-blue-500" />
          </div>
          <div>
            <p class="text-xs text-slate-500 dark:text-slate-400">This Month</p>
            <p class="text-xl font-bold text-accent-600 dark:text-accent-400">{{ formatINR(thisMonth) }}</p>
          </div>
        </div>
      </UCard>
      <UCard>
        <div class="flex items-center gap-3">
          <div class="p-2 bg-purple-50 rounded-lg">
            <UIcon name="i-lucide-bar-chart-2" class="text-xl text-purple-500" />
          </div>
          <div>
            <p class="text-xs text-slate-500 dark:text-slate-400">Avg per Deal</p>
            <p class="text-xl font-bold text-brand-900 dark:text-white">{{ formatINR(avgPerDeal) }}</p>
          </div>
        </div>
      </UCard>
    </div>

    <!-- Commission Table -->
    <UCard>
      <div class="overflow-x-auto">
        <table class="w-full text-sm">
          <thead>
            <tr class="border-b border-slate-200 dark:border-slate-700">
              <th class="text-left py-3 px-3 text-xs font-semibold text-slate-500 dark:text-slate-400 uppercase">Agent</th>
              <th class="text-left py-3 px-3 text-xs font-semibold text-slate-500 dark:text-slate-400 uppercase">Booking Ref</th>
              <th class="text-left py-3 px-3 text-xs font-semibold text-slate-500 dark:text-slate-400 uppercase">Project</th>
              <th class="text-left py-3 px-3 text-xs font-semibold text-slate-500 dark:text-slate-400 uppercase">Investor</th>
              <th class="text-right py-3 px-3 text-xs font-semibold text-slate-500 dark:text-slate-400 uppercase">Amount</th>
              <th class="text-right py-3 px-3 text-xs font-semibold text-slate-500 dark:text-slate-400 uppercase">TDS (5%)</th>
              <th class="text-right py-3 px-3 text-xs font-semibold text-slate-500 dark:text-slate-400 uppercase">Net Amount</th>
              <th class="text-left py-3 px-3 text-xs font-semibold text-slate-500 dark:text-slate-400 uppercase">Status</th>
              <th class="text-left py-3 px-3 text-xs font-semibold text-slate-500 dark:text-slate-400 uppercase">Date</th>
              <th class="text-left py-3 px-3 text-xs font-semibold text-slate-500 dark:text-slate-400 uppercase">Actions</th>
            </tr>
          </thead>
          <tbody>
            <tr
              v-for="c in commissions"
              :key="c.id"
              class="border-b border-slate-100 dark:border-slate-800 hover:bg-slate-50 dark:hover:bg-brand-800"
            >
              <td class="py-3 px-3 font-medium text-brand-900 dark:text-white">{{ c.agentName }}</td>
              <td class="py-3 px-3 text-xs font-mono text-slate-500 dark:text-slate-400">{{ c.bookingRef }}</td>
              <td class="py-3 px-3 text-slate-600 dark:text-slate-300">{{ c.projectName }}</td>
              <td class="py-3 px-3 text-slate-600 dark:text-slate-300">{{ c.investorName }}</td>
              <td class="py-3 px-3 text-right font-medium text-brand-900 dark:text-white">{{ formatINR(c.agentCommission) }}</td>
              <td class="py-3 px-3 text-right text-red-500">-{{ formatINR(c.tdsAmount) }}</td>
              <td class="py-3 px-3 text-right font-semibold text-green-600">{{ formatINR(c.netPayableToAgent) }}</td>
              <td class="py-3 px-3">
                <UBadge :color="getStatusColor(c.status)" variant="subtle" size="xs">
                  {{ c.status }}
                </UBadge>
              </td>
              <td class="py-3 px-3 text-xs text-slate-500 dark:text-slate-400">{{ formatDate(c.createdAt) }}</td>
              <td class="py-3 px-3">
                <div class="flex items-center gap-1">
                  <template v-if="c.status === 'pending'">
                    <UButton
                      icon="i-lucide-check"
                      variant="ghost"
                      color="success"
                      size="xs"
                      title="Approve"
                      @click="approveCommission(c.id)"
                    />
                    <UButton
                      icon="i-lucide-x"
                      variant="ghost"
                      color="error"
                      size="xs"
                      title="Reject"
                      @click="rejectCommission(c.id)"
                    />
                  </template>
                  <template v-else-if="c.status === 'approved'">
                    <UButton
                      icon="i-lucide-indian-rupee"
                      variant="ghost"
                      color="success"
                      size="xs"
                      title="Mark as Paid"
                      @click="markPaid(c.id)"
                    >
                      Pay
                    </UButton>
                  </template>
                  <span v-else-if="c.status === 'paid'" class="text-xs text-green-500">
                    Paid {{ c.paidAt ? formatDate(c.paidAt) : '' }}
                  </span>
                </div>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </UCard>
  </div>
</template>
