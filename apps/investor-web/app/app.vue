<script setup lang="ts">
const colorMode = useColorMode()
const isMobileMenuOpen = ref(false)
const route = useRoute()

const navLinks = [
  { label: 'Projects', to: '/projects' },
  { label: 'Area Guides', to: '/areas' },
  { label: 'Calculator', to: '/calculator' },
  { label: 'About', to: '/about' },
]

function toggleColorMode() {
  colorMode.preference = colorMode.value === 'dark' ? 'light' : 'dark'
}

watch(() => route.path, () => {
  isMobileMenuOpen.value = false
})
</script>

<template>
  <UApp>
    <div class="min-h-screen bg-white dark:bg-brand-950 text-slate-900 dark:text-slate-100">
      <!-- Navbar -->
      <header class="sticky top-0 z-50 bg-white/90 dark:bg-brand-950/90 backdrop-blur border-b border-slate-200 dark:border-slate-800">
        <div class="max-w-6xl mx-auto px-4 h-16 flex items-center justify-between">
          <!-- Logo -->
          <NuxtLink to="/" class="flex items-center gap-2">
            <div class="w-8 h-8 bg-brand-800 dark:bg-brand-600 rounded-lg flex items-center justify-center">
              <UIcon name="i-lucide-building-2" class="text-accent-400 text-lg" />
            </div>
            <span class="text-xl font-bold">
              <span class="text-brand-900 dark:text-white">Prop</span><span class="text-accent-500 dark:text-accent-400">Tech</span>
            </span>
          </NuxtLink>

          <!-- Desktop Nav -->
          <nav class="hidden md:flex items-center gap-6">
            <NuxtLink
              v-for="link in navLinks"
              :key="link.to"
              :to="link.to"
              class="text-sm font-medium text-slate-600 dark:text-slate-300 hover:text-brand-600 dark:hover:text-brand-400 transition-colors"
              active-class="text-brand-600 dark:text-brand-400"
            >
              {{ link.label }}
            </NuxtLink>
          </nav>

          <!-- Right section -->
          <div class="flex items-center gap-3">
            <!-- Dark mode toggle -->
            <UButton
              :icon="colorMode.value === 'dark' ? 'i-lucide-sun' : 'i-lucide-moon'"
              variant="ghost"
              color="neutral"
              size="sm"
              aria-label="Toggle color mode"
              @click="toggleColorMode"
            />

            <!-- WhatsApp CTA (desktop) -->
            <UButton
              href="https://wa.me/919876543210?text=Hi%2C%20I%20want%20to%20know%20about%20investment%20opportunities"
              target="_blank"
              size="sm"
              class="hidden md:flex bg-accent-500 hover:bg-accent-600 dark:bg-accent-400 dark:hover:bg-accent-500 text-white"
              icon="i-lucide-message-circle"
            >
              WhatsApp Us
            </UButton>

            <!-- Mobile menu toggle -->
            <UButton
              :icon="isMobileMenuOpen ? 'i-lucide-x' : 'i-lucide-menu'"
              variant="ghost"
              color="neutral"
              size="sm"
              class="md:hidden"
              aria-label="Toggle menu"
              @click="isMobileMenuOpen = !isMobileMenuOpen"
            />
          </div>
        </div>

        <!-- Mobile Nav Dropdown -->
        <Transition name="slide">
          <div v-if="isMobileMenuOpen" class="md:hidden border-t border-slate-200 dark:border-slate-800 bg-white dark:bg-brand-950">
            <nav class="px-4 py-3 space-y-1">
              <NuxtLink
                v-for="link in navLinks"
                :key="link.to"
                :to="link.to"
                class="block px-3 py-2 rounded-lg text-sm font-medium text-slate-600 dark:text-slate-300 hover:bg-slate-100 dark:hover:bg-brand-800 transition-colors"
                active-class="bg-slate-100 dark:bg-brand-800 text-brand-600 dark:text-brand-400"
              >
                {{ link.label }}
              </NuxtLink>
              <a
                href="https://wa.me/919876543210?text=Hi%2C%20I%20want%20to%20know%20about%20investment%20opportunities"
                target="_blank"
                class="block px-3 py-2 rounded-lg text-sm font-medium text-brand-600 dark:text-brand-400"
              >
                WhatsApp Us
              </a>
            </nav>
          </div>
        </Transition>
      </header>

      <!-- Page content -->
      <NuxtPage />

      <!-- Footer -->
      <footer class="bg-brand-800 dark:bg-brand-950 text-white border-t border-slate-800">
        <div class="max-w-6xl mx-auto px-4 py-12">
          <div class="grid sm:grid-cols-2 lg:grid-cols-4 gap-8">
            <!-- Brand -->
            <div class="sm:col-span-2 lg:col-span-1">
              <div class="flex items-center gap-2 mb-1">
                <div class="w-7 h-7 bg-accent-500 rounded-lg flex items-center justify-center">
                  <UIcon name="i-lucide-building-2" class="text-white text-sm" />
                </div>
                <span class="text-lg font-bold">
                  <span class="text-white">Prop</span><span class="text-accent-400">Tech</span>
                </span>
              </div>
              <p class="text-sm text-slate-400 mt-2 leading-relaxed">
                MahaRERA-licensed real estate investment platform for Mumbai micro-markets.
                Trust-first, data-driven, transparent.
              </p>
            </div>

            <!-- Quick Links -->
            <div>
              <h4 class="font-semibold text-sm mb-3">Explore</h4>
              <ul class="space-y-2 text-sm text-slate-400">
                <li><NuxtLink to="/projects" class="hover:text-brand-400 transition-colors">Projects</NuxtLink></li>
                <li><NuxtLink to="/areas" class="hover:text-brand-400 transition-colors">Area Guides</NuxtLink></li>
                <li><NuxtLink to="/calculator" class="hover:text-brand-400 transition-colors">Cost Calculator</NuxtLink></li>
                <li><NuxtLink to="/about" class="hover:text-brand-400 transition-colors">About Us</NuxtLink></li>
              </ul>
            </div>

            <!-- Areas -->
            <div>
              <h4 class="font-semibold text-sm mb-3">Areas</h4>
              <ul class="space-y-2 text-sm text-slate-400">
                <li><NuxtLink to="/areas/panvel" class="hover:text-brand-400 transition-colors">Panvel</NuxtLink></li>
                <li><NuxtLink to="/areas/dombivli-east" class="hover:text-brand-400 transition-colors">Dombivli</NuxtLink></li>
                <li><NuxtLink to="/areas/kalyan-shilphata" class="hover:text-brand-400 transition-colors">Kalyan</NuxtLink></li>
                <li><NuxtLink to="/areas/ulwe" class="hover:text-brand-400 transition-colors">Ulwe</NuxtLink></li>
                <li><NuxtLink to="/areas/kharghar" class="hover:text-brand-400 transition-colors">Kharghar</NuxtLink></li>
                <li><NuxtLink to="/areas/taloja" class="hover:text-brand-400 transition-colors">Taloja</NuxtLink></li>
              </ul>
            </div>

            <!-- Contact -->
            <div>
              <h4 class="font-semibold text-sm mb-3">Contact</h4>
              <ul class="space-y-2 text-sm text-slate-400">
                <li>
                  <a href="https://wa.me/919876543210" target="_blank" class="hover:text-brand-400 transition-colors flex items-center gap-1.5">
                    <UIcon name="i-lucide-message-circle" class="text-xs" />
                    WhatsApp
                  </a>
                </li>
                <li>
                  <a href="mailto:hello@proptech.in" class="hover:text-brand-400 transition-colors flex items-center gap-1.5">
                    <UIcon name="i-lucide-mail" class="text-xs" />
                    hello@proptech.in
                  </a>
                </li>
              </ul>
              <div class="mt-4 flex gap-3">
                <a href="https://www.instagram.com/proptech.in" target="_blank" class="text-slate-400 hover:text-brand-400 transition-colors">
                  <UIcon name="i-lucide-instagram" class="text-lg" />
                </a>
                <a href="https://www.linkedin.com/company/proptech-in" target="_blank" class="text-slate-400 hover:text-brand-400 transition-colors">
                  <UIcon name="i-lucide-linkedin" class="text-lg" />
                </a>
              </div>
            </div>
          </div>

          <div class="mt-10 pt-6 border-t border-slate-800 flex flex-col sm:flex-row justify-between items-center gap-4 text-xs text-slate-500">
            <p>PropTech. MahaRERA Licensed Platform.</p>
            <p>All project details sourced from MahaRERA. Verify independently before investing.</p>
          </div>
        </div>
      </footer>
    </div>
  </UApp>
</template>

<style scoped>
.slide-enter-active,
.slide-leave-active {
  transition: all 0.2s ease;
}
.slide-enter-from,
.slide-leave-to {
  opacity: 0;
  transform: translateY(-4px);
}
</style>
