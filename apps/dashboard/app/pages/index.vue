<script setup lang="ts">
definePageMeta({ middleware: 'auth' })

const authCookie = useCookie('auth_token')

onMounted(() => {
  if (!authCookie.value) {
    navigateTo('/login')
    return
  }

  try {
    const parts = authCookie.value.split('.')
    if (parts.length === 3) {
      const payload = JSON.parse(atob(parts[1]))
      const role = payload.role as string
      if (role === 'admin') navigateTo('/admin')
      else if (role === 'builder') navigateTo('/builder')
      else navigateTo('/agent')
    } else {
      navigateTo('/login')
    }
  } catch {
    navigateTo('/login')
  }
})
</script>

<template>
  <div class="flex items-center justify-center h-screen">
    <UIcon name="i-lucide-loader-2" class="text-3xl text-slate-400 animate-spin" />
  </div>
</template>
