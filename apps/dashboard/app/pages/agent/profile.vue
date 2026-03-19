<script setup lang="ts">
definePageMeta({ layout: 'agent', middleware: 'auth' })

const { agents } = useMockData()
const { formatINR, formatPhone, formatDate, formatPercent } = useFormatting()

const agent = agents[0] // Current logged-in agent mock

const profile = reactive({
  name: agent.name,
  phone: agent.phone,
  email: agent.email || '',
  reraId: agent.reraNumber,
})

const saving = ref(false)

async function saveProfile() {
  saving.value = true
  await new Promise(r => setTimeout(r, 800))
  console.log('[Profile] Saved:', { ...profile })
  saving.value = false
}

function signOut() {
  const authCookie = useCookie('auth_token')
  authCookie.value = null
  navigateTo('/login')
}

function getTierColor(tier: string) {
  const map: Record<string, string> = {
    platinum: 'primary',
    gold: 'warning',
    silver: 'neutral',
    bronze: 'error',
  }
  return map[tier] || 'neutral'
}

const performanceStats = [
  { label: 'Total Leads', value: agent.totalLeads, icon: 'i-lucide-users' },
  { label: 'Bookings', value: agent.totalBookings, icon: 'i-lucide-check-circle' },
  { label: 'Conversion', value: `${agent.conversionRate}%`, icon: 'i-lucide-trending-up' },
  { label: 'Experience', value: `${agent.experienceYears} yrs`, icon: 'i-lucide-award' },
]
</script>

<template>
  <div>
    <h1 class="text-xl font-bold text-brand-900 dark:text-white">Profile</h1>
    <p class="text-slate-500 dark:text-slate-400 text-sm mt-1 mb-6">Manage your account details</p>

    <!-- Tier & RERA Badge -->
    <UCard class="mb-4">
      <div class="flex items-center justify-between">
        <div class="flex items-center gap-3">
          <div class="w-12 h-12 rounded-full bg-brand-800 dark:bg-brand-700 flex items-center justify-center text-white text-lg font-bold">
            {{ agent.name.split(' ').map((n: string) => n[0]).join('') }}
          </div>
          <div>
            <p class="font-semibold text-brand-900 dark:text-white">{{ agent.name }}</p>
            <p class="text-xs text-slate-500 dark:text-slate-400">{{ formatPhone(agent.phone) }}</p>
          </div>
        </div>
        <div class="flex flex-col items-end gap-1">
          <UBadge :color="getTierColor(agent.tier)" variant="solid" size="sm">
            {{ agent.tier.toUpperCase() }}
          </UBadge>
          <div class="flex items-center gap-1">
            <UIcon
              :name="agent.reraVerified ? 'i-lucide-shield-check' : 'i-lucide-shield-alert'"
              :class="agent.reraVerified ? 'text-green-500' : 'text-amber-500'"
              class="text-sm"
            />
            <span class="text-xs" :class="agent.reraVerified ? 'text-green-600' : 'text-amber-600'">
              {{ agent.reraVerified ? 'RERA Verified' : 'RERA Pending' }}
            </span>
          </div>
        </div>
      </div>
    </UCard>

    <!-- Performance Stats -->
    <div class="grid grid-cols-2 gap-3 mb-4">
      <UCard v-for="stat in performanceStats" :key="stat.label" class="text-center">
        <UIcon :name="stat.icon" class="text-brand-600 dark:text-brand-400 text-lg mb-1" />
        <p class="text-lg font-bold text-brand-900 dark:text-white">{{ stat.value }}</p>
        <p class="text-xs text-slate-500 dark:text-slate-400">{{ stat.label }}</p>
      </UCard>
    </div>

    <!-- Operating Areas -->
    <UCard class="mb-4">
      <p class="text-sm font-medium text-brand-900 dark:text-white mb-2">Operating Areas</p>
      <div class="flex flex-wrap gap-2">
        <UBadge
          v-for="area in agent.operatingAreas"
          :key="area"
          color="primary"
          variant="subtle"
          size="sm"
        >
          {{ area }}
        </UBadge>
      </div>
    </UCard>

    <!-- Edit Profile Form -->
    <UCard class="mb-4">
      <p class="text-sm font-medium text-brand-900 dark:text-white mb-4">Edit Profile</p>
      <div class="space-y-4">
        <UFormField label="Full Name">
          <UInput v-model="profile.name" placeholder="Your full name" />
        </UFormField>
        <UFormField label="Phone Number">
          <UInput v-model="profile.phone" type="tel" disabled placeholder="+91 98765 43210" />
        </UFormField>
        <UFormField label="Email">
          <UInput v-model="profile.email" type="email" placeholder="agent@example.com" />
        </UFormField>
        <UFormField label="RERA Agent ID">
          <UInput v-model="profile.reraId" placeholder="A51900XXXXXX" disabled />
        </UFormField>
        <UButton class="w-full" :loading="saving" @click="saveProfile">
          Save Changes
        </UButton>
      </div>
    </UCard>

    <UButton
      class="w-full"
      variant="ghost"
      color="error"
      @click="signOut"
    >
      Sign Out
    </UButton>
  </div>
</template>
