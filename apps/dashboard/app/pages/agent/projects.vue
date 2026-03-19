<script setup lang="ts">
definePageMeta({ layout: 'agent', middleware: 'auth' })

const { projects, units } = useMockData()
const { formatINRShort, formatINR, formatDate } = useFormatting()

const locationFilter = ref('')
const bhkFilter = ref('')

const locations = computed(() => {
  const locs = [...new Set(projects.map(p => p.location))]
  return [{ label: 'All Locations', value: '' }, ...locs.map(l => ({ label: l, value: l }))]
})

const bhkOptions = [
  { label: 'All BHK', value: '' },
  { label: '1 BHK', value: '1 BHK' },
  { label: '2 BHK', value: '2 BHK' },
  { label: '3 BHK', value: '3 BHK' },
  { label: '4 BHK', value: '4 BHK' },
]

const filteredProjects = computed(() => {
  return projects
    .filter(p => p.status === 'active')
    .filter(p => !locationFilter.value || p.location === locationFilter.value)
    .filter(p => !bhkFilter.value || p.configurations.includes(bhkFilter.value))
})

// Inventory modal
const showInventoryModal = ref(false)
const selectedProject = ref<typeof projects[0] | null>(null)

function openInventory(project: typeof projects[0]) {
  selectedProject.value = project
  showInventoryModal.value = true
}

const projectUnits = computed(() => {
  if (!selectedProject.value) return []
  return units.filter(u => u.projectId === selectedProject.value!.id)
})

function getUnitStatusColor(status: string) {
  const map: Record<string, string> = {
    available: 'success',
    blocked: 'warning',
    booked: 'info',
    sold: 'error',
  }
  return map[status] || 'neutral'
}
</script>

<template>
  <div>
    <h1 class="text-xl font-bold text-brand-900 dark:text-white">Projects</h1>
    <p class="text-slate-500 dark:text-slate-400 text-sm mt-1 mb-4">Browse available projects for your leads</p>

    <div class="flex gap-3 mb-6">
      <USelect v-model="locationFilter" :items="locations" placeholder="Location" class="flex-1" />
      <USelect v-model="bhkFilter" :items="bhkOptions" placeholder="BHK" class="flex-1" />
    </div>

    <div class="space-y-4">
      <UCard
        v-for="project in filteredProjects"
        :key="project.id"
        class="hover:shadow-md transition-shadow cursor-pointer"
        @click="openInventory(project)"
      >
        <div class="flex items-start justify-between mb-2">
          <div>
            <h3 class="font-semibold text-brand-900 dark:text-white">{{ project.name }}</h3>
            <p class="text-xs text-slate-500 dark:text-slate-400">by {{ project.builder }}</p>
          </div>
          <UBadge color="success" variant="subtle" size="xs">Active</UBadge>
        </div>

        <div class="grid grid-cols-2 gap-2 text-xs text-slate-600 dark:text-slate-300 mb-3">
          <div class="flex items-center gap-1">
            <UIcon name="i-lucide-map-pin" class="text-slate-400" />
            <span>{{ project.location }}</span>
          </div>
          <div class="flex items-center gap-1">
            <UIcon name="i-lucide-home" class="text-slate-400" />
            <span>{{ project.configurations.join(', ') }}</span>
          </div>
          <div class="flex items-center gap-1">
            <UIcon name="i-lucide-indian-rupee" class="text-slate-400" />
            <span>{{ formatINRShort(project.priceRange.min) }} - {{ formatINRShort(project.priceRange.max) }}</span>
          </div>
          <div class="flex items-center gap-1">
            <UIcon name="i-lucide-building" class="text-slate-400" />
            <span>{{ project.availableUnits }} / {{ project.totalUnits }} units</span>
          </div>
        </div>

        <div class="flex items-center justify-between">
          <span class="text-xs text-slate-400">RERA: {{ project.reraNumber }}</span>
          <UButton variant="link" size="xs" icon="i-lucide-eye" @click.stop="openInventory(project)">
            View Units
          </UButton>
        </div>

        <!-- Construction progress bar -->
        <div class="mt-3">
          <div class="flex justify-between text-xs text-slate-500 dark:text-slate-400 mb-1">
            <span>Construction</span>
            <span>{{ project.constructionProgress }}%</span>
          </div>
          <div class="w-full bg-slate-100 dark:bg-brand-800 rounded-full h-1.5">
            <div
              class="bg-brand-600 dark:bg-brand-500 h-1.5 rounded-full transition-all"
              :style="{ width: `${project.constructionProgress}%` }"
            />
          </div>
        </div>
      </UCard>

      <div v-if="filteredProjects.length === 0" class="text-center py-12">
        <UIcon name="i-lucide-building-2" class="text-4xl text-slate-300 mb-2" />
        <p class="text-slate-400">No projects match your filters</p>
      </div>
    </div>

    <!-- Inventory Modal -->
    <UModal v-model:open="showInventoryModal">
      <template #content>
        <div class="p-6 max-h-[80vh] overflow-y-auto">
          <div class="flex items-center justify-between mb-4">
            <div>
              <h3 class="text-lg font-semibold text-brand-900 dark:text-white">{{ selectedProject?.name }}</h3>
              <p class="text-sm text-slate-500 dark:text-slate-400">Unit Inventory</p>
            </div>
            <UButton icon="i-lucide-x" variant="ghost" color="neutral" size="sm" @click="showInventoryModal = false" />
          </div>

          <div v-if="projectUnits.length > 0" class="space-y-2">
            <div
              v-for="unit in projectUnits"
              :key="unit.id"
              class="flex items-center justify-between p-3 bg-slate-50 dark:bg-brand-900/50 rounded-lg"
            >
              <div>
                <p class="text-sm font-medium text-brand-900 dark:text-white">{{ unit.unitNumber }}</p>
                <p class="text-xs text-slate-500 dark:text-slate-400">
                  Tower {{ unit.tower }} | Floor {{ unit.floor }} | {{ unit.unitType }} | {{ unit.carpetArea }} sq.ft | {{ unit.facing }}
                </p>
              </div>
              <div class="text-right">
                <p class="text-sm font-semibold text-brand-900 dark:text-white">{{ formatINR(unit.price) }}</p>
                <UBadge :color="getUnitStatusColor(unit.status)" variant="subtle" size="xs">
                  {{ unit.status }}
                </UBadge>
              </div>
            </div>
          </div>

          <div v-else class="text-center py-8">
            <p class="text-slate-400 text-sm">No unit data available for this project</p>
          </div>
        </div>
      </template>
    </UModal>
  </div>
</template>
