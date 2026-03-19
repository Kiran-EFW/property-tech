<script setup lang="ts">
definePageMeta({ layout: 'builder', middleware: 'auth' })

const { units, projects } = useMockData()
const { formatINR } = useFormatting()

// Builder's projects
const builderProjectIds = ['prj-001', 'prj-002']
const builderUnits = computed(() => units.filter(u => builderProjectIds.includes(u.projectId)))
const builderProjects = computed(() => projects.filter(p => builderProjectIds.includes(p.id)))

const activeTab = ref('all')
const projectFilter = ref('')
const searchQuery = ref('')

const tabItems = computed(() => [
  { label: `All Units (${builderUnits.value.length})`, value: 'all' },
  { label: `Available (${builderUnits.value.filter(u => u.status === 'available').length})`, value: 'available' },
  { label: `Blocked (${builderUnits.value.filter(u => u.status === 'blocked').length})`, value: 'blocked' },
  { label: `Booked (${builderUnits.value.filter(u => u.status === 'booked').length})`, value: 'booked' },
  { label: `Sold (${builderUnits.value.filter(u => u.status === 'sold').length})`, value: 'sold' },
])

const projectOptions = computed(() => [
  { label: 'All Projects', value: '' },
  ...builderProjects.value.map(p => ({ label: p.name, value: p.id })),
])

const filteredUnits = computed(() => {
  let result = [...builderUnits.value]
  if (activeTab.value !== 'all') {
    result = result.filter(u => u.status === activeTab.value)
  }
  if (projectFilter.value) {
    result = result.filter(u => u.projectId === projectFilter.value)
  }
  if (searchQuery.value) {
    const q = searchQuery.value.toLowerCase()
    result = result.filter(u =>
      u.unitNumber.toLowerCase().includes(q) ||
      u.projectName.toLowerCase().includes(q) ||
      u.tower.toLowerCase().includes(q)
    )
  }
  return result
})

function getStatusColor(status: string) {
  const map: Record<string, string> = {
    available: 'success', blocked: 'warning', booked: 'info', sold: 'error',
  }
  return map[status] || 'neutral'
}

// Bulk status update
const selectedUnits = ref<string[]>([])
const bulkStatus = ref('')

function toggleUnit(unitId: string) {
  const idx = selectedUnits.value.indexOf(unitId)
  if (idx >= 0) {
    selectedUnits.value.splice(idx, 1)
  } else {
    selectedUnits.value.push(unitId)
  }
}

function toggleSelectAll() {
  if (selectedUnits.value.length === filteredUnits.value.length) {
    selectedUnits.value = []
  } else {
    selectedUnits.value = filteredUnits.value.map(u => u.id)
  }
}

function applyBulkStatus() {
  console.log('[Inventory] Bulk status update:', selectedUnits.value, '->', bulkStatus.value)
  selectedUnits.value = []
  bulkStatus.value = ''
}

function updateUnitStatus(unitId: string, status: string) {
  console.log('[Inventory] Unit status update:', unitId, '->', status)
}

function uploadInventory() {
  console.log('[Inventory] Upload spreadsheet')
}
</script>

<template>
  <div>
    <div class="flex items-center justify-between mb-6">
      <div>
        <h1 class="text-2xl font-bold text-brand-900 dark:text-white">Inventory</h1>
        <p class="text-slate-500 dark:text-slate-400 text-sm mt-1">Manage unit availability and pricing</p>
      </div>
      <UButton icon="i-lucide-upload" @click="uploadInventory">Upload Inventory</UButton>
    </div>

    <!-- Bulk actions -->
    <div v-if="selectedUnits.length > 0" class="flex items-center gap-3 mb-4 p-3 bg-blue-50 rounded-lg">
      <span class="text-sm text-blue-700 font-medium">{{ selectedUnits.length }} units selected</span>
      <USelect
        v-model="bulkStatus"
        :items="[
          { label: 'Available', value: 'available' },
          { label: 'Blocked', value: 'blocked' },
          { label: 'Booked', value: 'booked' },
          { label: 'Sold', value: 'sold' },
        ]"
        placeholder="Change status to..."
        size="sm"
      />
      <UButton size="sm" :disabled="!bulkStatus" @click="applyBulkStatus">
        Apply
      </UButton>
      <UButton size="sm" variant="ghost" color="neutral" @click="selectedUnits = []">
        Clear
      </UButton>
    </div>

    <UCard>
      <div class="flex items-center gap-3 mb-4">
        <UInput v-model="searchQuery" placeholder="Search units..." icon="i-lucide-search" class="flex-1" />
        <USelect v-model="projectFilter" :items="projectOptions" placeholder="Project" />
      </div>

      <UTabs :items="tabItems" v-model="activeTab" class="mb-4" />

      <div class="overflow-x-auto">
        <table class="w-full text-sm">
          <thead>
            <tr class="border-b border-slate-200 dark:border-slate-700">
              <th class="py-3 px-3 w-8">
                <input
                  type="checkbox"
                  :checked="selectedUnits.length === filteredUnits.length && filteredUnits.length > 0"
                  class="rounded"
                  @change="toggleSelectAll"
                />
              </th>
              <th class="text-left py-3 px-3 text-xs font-semibold text-slate-500 dark:text-slate-400 uppercase">Unit</th>
              <th class="text-left py-3 px-3 text-xs font-semibold text-slate-500 dark:text-slate-400 uppercase">Project</th>
              <th class="text-left py-3 px-3 text-xs font-semibold text-slate-500 dark:text-slate-400 uppercase">Tower</th>
              <th class="text-left py-3 px-3 text-xs font-semibold text-slate-500 dark:text-slate-400 uppercase">Floor</th>
              <th class="text-left py-3 px-3 text-xs font-semibold text-slate-500 dark:text-slate-400 uppercase">Type</th>
              <th class="text-left py-3 px-3 text-xs font-semibold text-slate-500 dark:text-slate-400 uppercase">Area</th>
              <th class="text-left py-3 px-3 text-xs font-semibold text-slate-500 dark:text-slate-400 uppercase">Facing</th>
              <th class="text-right py-3 px-3 text-xs font-semibold text-slate-500 dark:text-slate-400 uppercase">Rate/sqft</th>
              <th class="text-right py-3 px-3 text-xs font-semibold text-slate-500 dark:text-slate-400 uppercase">Price</th>
              <th class="text-left py-3 px-3 text-xs font-semibold text-slate-500 dark:text-slate-400 uppercase">Status</th>
              <th class="text-left py-3 px-3 text-xs font-semibold text-slate-500 dark:text-slate-400 uppercase">Action</th>
            </tr>
          </thead>
          <tbody>
            <tr
              v-for="unit in filteredUnits"
              :key="unit.id"
              class="border-b border-slate-100 dark:border-slate-800 hover:bg-slate-50 dark:hover:bg-brand-800"
            >
              <td class="py-3 px-3">
                <input
                  type="checkbox"
                  :checked="selectedUnits.includes(unit.id)"
                  class="rounded"
                  @change="toggleUnit(unit.id)"
                />
              </td>
              <td class="py-3 px-3 font-medium text-brand-900 dark:text-white">{{ unit.unitNumber }}</td>
              <td class="py-3 px-3 text-slate-600 dark:text-slate-300 text-xs">{{ unit.projectName }}</td>
              <td class="py-3 px-3 text-slate-600 dark:text-slate-300">{{ unit.tower }}</td>
              <td class="py-3 px-3 text-slate-600 dark:text-slate-300">{{ unit.floor }}</td>
              <td class="py-3 px-3 text-slate-600 dark:text-slate-300">{{ unit.unitType }}</td>
              <td class="py-3 px-3 text-slate-600 dark:text-slate-300">{{ unit.carpetArea }} sq.ft</td>
              <td class="py-3 px-3 text-slate-600 dark:text-slate-300">{{ unit.facing }}</td>
              <td class="py-3 px-3 text-right text-slate-600 dark:text-slate-300">{{ formatINR(unit.pricePerSqft) }}</td>
              <td class="py-3 px-3 text-right font-medium text-brand-900 dark:text-white">{{ formatINR(unit.price) }}</td>
              <td class="py-3 px-3">
                <UBadge :color="getStatusColor(unit.status)" variant="subtle" size="xs">
                  {{ unit.status }}
                </UBadge>
              </td>
              <td class="py-3 px-3">
                <USelect
                  :model-value="unit.status"
                  :items="[
                    { label: 'Available', value: 'available' },
                    { label: 'Blocked', value: 'blocked' },
                    { label: 'Booked', value: 'booked' },
                    { label: 'Sold', value: 'sold' },
                  ]"
                  size="xs"
                  class="w-28"
                  @update:model-value="(val: string) => updateUnitStatus(unit.id, val)"
                />
              </td>
            </tr>
          </tbody>
        </table>
      </div>

      <div v-if="filteredUnits.length === 0" class="text-center py-12">
        <UIcon name="i-lucide-building" class="text-4xl text-slate-300 mb-2" />
        <p class="text-slate-400 dark:text-slate-500">No units match your filters</p>
      </div>
    </UCard>
  </div>
</template>
