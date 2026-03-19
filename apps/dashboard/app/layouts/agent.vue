<script setup lang="ts">
const route = useRoute()

const colorMode = useColorMode()
function toggleColorMode() {
  colorMode.preference = colorMode.value === 'dark' ? 'light' : 'dark'
}

const navItems = [
  { label: 'Home', icon: 'i-lucide-home', to: '/agent' },
  { label: 'Leads', icon: 'i-lucide-users', to: '/agent/leads' },
  { label: 'Visits', icon: 'i-lucide-map-pin', to: '/agent/visits' },
  { label: 'Projects', icon: 'i-lucide-building-2', to: '/agent/projects' },
  { label: 'Earnings', icon: 'i-lucide-indian-rupee', to: '/agent/commissions' },
]

function isActive(item: typeof navItems[0]) {
  if (item.to === '/agent') return route.path === '/agent'
  return route.path.startsWith(item.to)
}
</script>

<template>
  <div class="min-h-screen bg-slate-50 dark:bg-brand-950 pb-20">
    <!-- Top header bar -->
    <header class="sticky top-0 z-40 bg-white dark:bg-brand-900 border-b border-slate-200 dark:border-slate-700 px-4 py-3">
      <div class="flex items-center justify-between">
        <!-- Logo -->
        <div class="flex items-center gap-2">
          <div class="w-8 h-8 bg-brand-800 dark:bg-brand-600 rounded-lg flex items-center justify-center">
            <UIcon name="i-lucide-building-2" class="text-accent-400 text-lg" />
          </div>
          <h1 class="text-lg font-semibold">
            <span class="text-brand-900 dark:text-white">Prop</span><span class="text-accent-500 dark:text-accent-400">Tech</span>
          </h1>
        </div>

        <div class="flex items-center gap-3">
          <!-- Language switcher placeholder -->
          <UButton
            icon="i-lucide-languages"
            variant="ghost"
            color="neutral"
            size="sm"
            aria-label="Switch language"
          />

          <!-- Dark mode toggle -->
          <UButton
            :icon="colorMode.value === 'dark' ? 'i-lucide-sun' : 'i-lucide-moon'"
            variant="ghost"
            color="neutral"
            size="sm"
            aria-label="Toggle color mode"
            @click="toggleColorMode"
          />

          <!-- Profile avatar -->
          <NuxtLink to="/agent/profile">
            <UAvatar
              text="RS"
              size="sm"
              class="ring-2 ring-brand-600 dark:ring-brand-400 cursor-pointer"
            />
          </NuxtLink>
        </div>
      </div>
    </header>

    <!-- Page content -->
    <main class="px-4 py-4">
      <slot />
    </main>

    <!-- Bottom navigation bar (mobile-first) -->
    <nav class="fixed bottom-0 left-0 right-0 z-50 bg-white dark:bg-brand-900 border-t border-slate-200 dark:border-slate-700">
      <div class="flex justify-around py-2">
        <NuxtLink
          v-for="item in navItems"
          :key="item.to"
          :to="item.to"
          class="flex flex-col items-center gap-1 px-3 py-1 text-xs transition-colors"
          :class="isActive(item) ? 'text-brand-600 dark:text-brand-400 font-medium' : 'text-slate-400'"
        >
          <UIcon :name="item.icon" class="text-xl" />
          <span>{{ item.label }}</span>
        </NuxtLink>
      </div>
    </nav>
  </div>
</template>
