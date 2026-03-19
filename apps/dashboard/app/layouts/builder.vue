<script setup lang="ts">
const route = useRoute()

const colorMode = useColorMode()
function toggleColorMode() {
  colorMode.preference = colorMode.value === 'dark' ? 'light' : 'dark'
}

const navItems = [
  { label: 'Overview', icon: 'i-lucide-layout-dashboard', to: '/builder' },
  { label: 'Inventory', icon: 'i-lucide-building-2', to: '/builder/inventory' },
  { label: 'Leads', icon: 'i-lucide-users', to: '/builder/leads' },
  { label: 'Commissions', icon: 'i-lucide-indian-rupee', to: '/builder/commissions' },
]

// Mobile overlay state
const isMobileMenuOpen = ref(false)

function isActive(item: typeof navItems[0]) {
  if (item.to === '/builder') return route.path === '/builder'
  return route.path.startsWith(item.to)
}

// Close mobile menu on route change
watch(() => route.path, () => {
  isMobileMenuOpen.value = false
})
</script>

<template>
  <div class="min-h-screen bg-slate-50 dark:bg-brand-950 flex">
    <!-- Mobile overlay backdrop -->
    <Transition name="fade">
      <div
        v-if="isMobileMenuOpen"
        class="fixed inset-0 bg-black/50 z-40 lg:hidden"
        @click="isMobileMenuOpen = false"
      />
    </Transition>

    <!-- Sidebar -->
    <aside
      class="fixed lg:sticky top-0 h-screen w-60 bg-white dark:bg-brand-900 border-r border-slate-200 dark:border-slate-700 flex flex-col z-50 transition-transform duration-200"
      :class="isMobileMenuOpen ? 'translate-x-0' : '-translate-x-full lg:translate-x-0'"
    >
      <!-- Logo -->
      <div class="p-4 border-b border-slate-200 dark:border-slate-700 flex items-center justify-between">
        <div class="w-8 h-8 bg-brand-800 dark:bg-brand-600 rounded-lg flex items-center justify-center">
          <UIcon name="i-lucide-building-2" class="text-accent-400 text-lg" />
        </div>
        <span class="text-lg font-semibold">
          <span class="text-brand-900 dark:text-white">Prop</span><span class="text-accent-500 dark:text-accent-400">Tech</span>
        </span>
        <!-- Mobile close button -->
        <UButton
          icon="i-lucide-x"
          variant="ghost"
          color="neutral"
          size="sm"
          class="lg:hidden"
          @click="isMobileMenuOpen = false"
        />
      </div>

      <!-- Navigation -->
      <nav class="flex-1 p-2 space-y-1">
        <NuxtLink
          v-for="item in navItems"
          :key="item.to"
          :to="item.to"
          class="flex items-center gap-3 px-3 py-2 rounded-lg text-sm transition-colors"
          :class="isActive(item)
            ? 'bg-brand-800 dark:bg-brand-700 text-white'
            : 'text-slate-600 dark:text-slate-300 hover:bg-slate-100 dark:hover:bg-brand-800'"
        >
          <UIcon :name="item.icon" class="text-lg" />
          <span>{{ item.label }}</span>
        </NuxtLink>
      </nav>

      <!-- User section -->
      <div class="p-4 border-t border-slate-200 dark:border-slate-700">
        <UBadge color="warning" variant="subtle" size="sm">Builder</UBadge>
      </div>
    </aside>

    <!-- Main content area -->
    <div class="flex-1 flex flex-col min-w-0">
      <!-- Top header -->
      <header class="sticky top-0 z-30 bg-white dark:bg-brand-900 border-b border-slate-200 dark:border-slate-700 px-4 lg:px-6 py-3">
        <div class="flex items-center gap-4">
          <!-- Hamburger menu for mobile -->
          <UButton
            icon="i-lucide-menu"
            variant="ghost"
            color="neutral"
            size="sm"
            class="lg:hidden"
            aria-label="Open menu"
            @click="isMobileMenuOpen = true"
          />

          <!-- Search bar -->
          <div class="flex-1 max-w-md">
            <UInput
              icon="i-lucide-search"
              placeholder="Search inventory, leads..."
              size="sm"
              class="w-full"
            />
          </div>

          <div class="flex items-center gap-3">
            <!-- Notifications bell -->
            <UButton
              icon="i-lucide-bell"
              variant="ghost"
              color="neutral"
              size="sm"
              aria-label="Notifications"
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
            <UAvatar
              text="BL"
              size="sm"
              class="ring-2 ring-brand-600 dark:ring-brand-400 cursor-pointer"
            />
          </div>
        </div>
      </header>

      <!-- Page content -->
      <main class="flex-1 p-4 lg:p-6">
        <slot />
      </main>
    </div>
  </div>
</template>

<style scoped>
.fade-enter-active,
.fade-leave-active {
  transition: opacity 0.2s ease;
}
.fade-enter-from,
.fade-leave-to {
  opacity: 0;
}
</style>
