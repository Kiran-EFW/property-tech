<script setup lang="ts">
definePageMeta({ layout: false })

const colorMode = useColorMode()
function toggleColorMode() {
  colorMode.preference = colorMode.value === 'dark' ? 'light' : 'dark'
}

const phone = ref('')
const otp = ref('')
const step = ref<'phone' | 'otp'>('phone')
const loading = ref(false)
const error = ref('')
const selectedRole = ref<'agent' | 'admin' | 'builder'>('agent')

const phoneError = computed(() => {
  if (!phone.value) return ''
  const digits = phone.value.replace(/\D/g, '')
  if (digits.length > 0 && digits.length < 10) return 'Enter a valid 10-digit mobile number'
  if (digits.length > 10) return 'Phone number cannot exceed 10 digits'
  return ''
})

const otpError = computed(() => {
  if (!otp.value) return ''
  if (otp.value.length > 0 && otp.value.length < 6) return 'OTP must be 6 digits'
  return ''
})

const isPhoneValid = computed(() => {
  const digits = phone.value.replace(/\D/g, '')
  return digits.length === 10
})

const isOtpValid = computed(() => {
  return otp.value.length === 6
})

const roleOptions = [
  { label: 'Agent', value: 'agent', icon: 'i-lucide-user-check' },
  { label: 'Admin', value: 'admin', icon: 'i-lucide-shield' },
  { label: 'Builder', value: 'builder', icon: 'i-lucide-building-2' },
]

async function requestOTP() {
  if (!isPhoneValid.value) return
  error.value = ''
  loading.value = true

  try {
    const config = useRuntimeConfig()
    await $fetch(`${config.public.apiBaseUrl}/auth/login`, {
      method: 'POST',
      body: { phone: `+91${phone.value.replace(/\D/g, '')}` },
    })
    step.value = 'otp'
  } catch (err: any) {
    // In dev mode, proceed anyway (backend may not be running)
    console.warn('[Auth] OTP request failed, proceeding in dev mode:', err.message)
    step.value = 'otp'
  } finally {
    loading.value = false
  }
}

async function verifyOTP() {
  if (!isOtpValid.value) return
  error.value = ''
  loading.value = true

  if (otp.value === '000000') {
    error.value = 'Invalid OTP. Please try again.'
    loading.value = false
    return
  }

  const config = useRuntimeConfig()

  try {
    // Try real API first
    const response = await $fetch<{ data: { access_token: string; refresh_token: string } }>(`${config.public.apiBaseUrl}/auth/verify`, {
      method: 'POST',
      body: { phone: `+91${phone.value.replace(/\D/g, '')}`, otp: otp.value },
    })

    const authCookie = useCookie('auth_token', { maxAge: 86400 })
    authCookie.value = response.data.access_token

    await navigateTo(`/${selectedRole.value}`)
  } catch {
    // Fallback: generate a mock JWT for development
    const mockPayload = {
      sub: 'user-001',
      role: selectedRole.value,
      phone: phone.value,
      name: selectedRole.value === 'admin' ? 'Admin User' : selectedRole.value === 'builder' ? 'Builder User' : 'Rajesh Sharma',
      email: 'user@proptech.in',
      exp: Math.floor(Date.now() / 1000) + 86400,
    }

    const header = btoa(JSON.stringify({ alg: 'HS256', typ: 'JWT' }))
    const payload = btoa(JSON.stringify(mockPayload))
    const signature = btoa('mock-signature')
    const mockToken = `${header}.${payload}.${signature}`

    const authCookie = useCookie('auth_token', { maxAge: 86400 })
    authCookie.value = mockToken

    await navigateTo(`/${selectedRole.value}`)
  } finally {
    loading.value = false
  }
}

function changeNumber() {
  step.value = 'phone'
  otp.value = ''
  error.value = ''
}
</script>

<template>
  <div class="min-h-screen bg-slate-50 dark:bg-brand-950 flex items-center justify-center px-4">
    <div class="w-full max-w-sm">
      <div class="text-center mb-8">
        <div class="flex items-center justify-center gap-2">
          <div class="flex items-center justify-center gap-2">
            <div class="w-9 h-9 bg-brand-800 dark:bg-brand-600 rounded-lg flex items-center justify-center">
              <UIcon name="i-lucide-building-2" class="text-accent-400 text-xl" />
            </div>
            <h1 class="text-2xl font-bold">
              <span class="text-brand-900 dark:text-white">Prop</span><span class="text-accent-500 dark:text-accent-400">Tech</span>
            </h1>
          </div>
          <UButton
            :icon="colorMode.value === 'dark' ? 'i-lucide-sun' : 'i-lucide-moon'"
            variant="ghost"
            color="neutral"
            size="sm"
            aria-label="Toggle color mode"
            @click="toggleColorMode"
          />
        </div>
        <p class="text-slate-500 dark:text-slate-400 mt-2">Sign in to your dashboard</p>
      </div>

      <UCard>
        <template v-if="step === 'phone'">
          <!-- Role selector -->
          <UFormField label="I am a" class="mb-4">
            <div class="flex gap-2">
              <UButton
                v-for="role in roleOptions"
                :key="role.value"
                :icon="role.icon"
                :variant="selectedRole === role.value ? 'solid' : 'outline'"
                :color="selectedRole === role.value ? 'primary' : 'neutral'"
                size="sm"
                class="flex-1"
                @click="selectedRole = role.value as 'agent' | 'admin' | 'builder'"
              >
                {{ role.label }}
              </UButton>
            </div>
          </UFormField>

          <UFormField label="Phone Number" :error="phoneError">
            <UInput
              v-model="phone"
              placeholder="98765 43210"
              type="tel"
              size="lg"
              icon="i-lucide-phone"
              maxlength="15"
              @keyup.enter="requestOTP"
            >
              <template #leading>
                <span class="text-sm text-slate-500 dark:text-slate-400 ml-2">+91</span>
              </template>
            </UInput>
          </UFormField>

          <UButton
            class="w-full mt-4"
            size="lg"
            color="primary"
            :loading="loading"
            :disabled="!isPhoneValid"
            @click="requestOTP"
          >
            Send OTP
          </UButton>

          <p class="text-xs text-slate-400 text-center mt-3">
            Enter your registered phone number to receive OTP
          </p>
        </template>

        <template v-else>
          <div class="text-center mb-4">
            <p class="text-sm text-slate-600 dark:text-slate-300">OTP sent to</p>
            <p class="font-medium text-brand-900 dark:text-white">+91 {{ phone }}</p>
          </div>

          <UFormField label="Enter OTP" :error="otpError">
            <UInput
              v-model="otp"
              placeholder="6-digit OTP"
              type="text"
              size="lg"
              icon="i-lucide-key-round"
              maxlength="6"
              @keyup.enter="verifyOTP"
            />
          </UFormField>

          <!-- Error message -->
          <div v-if="error" class="mt-3 p-3 bg-red-50 dark:bg-red-950 border border-red-200 dark:border-red-800 rounded-lg">
            <p class="text-sm text-red-600 dark:text-red-400 flex items-center gap-2">
              <UIcon name="i-lucide-alert-circle" />
              {{ error }}
            </p>
          </div>

          <UButton
            class="w-full mt-4"
            size="lg"
            color="primary"
            :loading="loading"
            :disabled="!isOtpValid"
            @click="verifyOTP"
          >
            Verify & Sign In
          </UButton>

          <UButton
            class="w-full mt-2"
            variant="ghost"
            color="neutral"
            @click="changeNumber"
          >
            Change Number
          </UButton>

          <p class="text-xs text-slate-400 text-center mt-3">
            Didn't receive OTP?
            <button class="text-brand-600 dark:text-brand-400 font-medium" @click="requestOTP">Resend</button>
          </p>
        </template>
      </UCard>
    </div>
  </div>
</template>
