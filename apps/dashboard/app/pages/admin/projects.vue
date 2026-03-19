<script setup lang="ts">
definePageMeta({ layout: 'admin', middleware: 'auth' })

const { projects } = useMockData()
const { formatINRShort, formatDate } = useFormatting()

const searchQuery = ref('')
const statusFilter = ref('')

const statusOptions = [
  { label: 'All Status', value: '' },
  { label: 'Active', value: 'active' },
  { label: 'Draft', value: 'draft' },
  { label: 'Sold Out', value: 'sold_out' },
  { label: 'Suspended', value: 'suspended' },
]

const filteredProjects = computed(() => {
  let result = [...projects]
  if (statusFilter.value) {
    result = result.filter(p => p.status === statusFilter.value)
  }
  if (searchQuery.value) {
    const q = searchQuery.value.toLowerCase()
    result = result.filter(p =>
      p.name.toLowerCase().includes(q) ||
      p.builder.toLowerCase().includes(q) ||
      p.reraNumber.toLowerCase().includes(q) ||
      p.location.toLowerCase().includes(q)
    )
  }
  return result
})

function getStatusColor(status: string) {
  const map: Record<string, string> = {
    active: 'success',
    draft: 'neutral',
    sold_out: 'warning',
    suspended: 'error',
  }
  return map[status] || 'neutral'
}

function getStatusLabel(status: string) {
  const map: Record<string, string> = {
    active: 'Active',
    draft: 'Draft',
    sold_out: 'Sold Out',
    suspended: 'Suspended',
  }
  return map[status] || status
}

// Add/Edit modal
const showModal = ref(false)
const editingProject = ref<typeof projects[0] | null>(null)

const formData = reactive({
  name: '',
  builder: '',
  reraNumber: '',
  location: '',
  status: 'draft' as string,
  totalUnits: 0,
  totalFloors: 0,
  totalTowers: 1,
  shortDescription: '',
  configurations: [] as string[],
  amenities: '',
  possessionDate: '',
})

function openAddModal() {
  editingProject.value = null
  Object.assign(formData, {
    name: '', builder: '', reraNumber: '', location: '', status: 'draft',
    totalUnits: 0, totalFloors: 0, totalTowers: 1, shortDescription: '',
    configurations: [], amenities: '', possessionDate: '',
  })
  showModal.value = true
}

function openEditModal(project: typeof projects[0]) {
  editingProject.value = project
  Object.assign(formData, {
    name: project.name,
    builder: project.builder,
    reraNumber: project.reraNumber,
    location: project.location,
    status: project.status,
    totalUnits: project.totalUnits,
    totalFloors: 0,
    totalTowers: 1,
    shortDescription: '',
    configurations: project.configurations,
    amenities: project.amenities.join(', '),
    possessionDate: project.possessionDate,
  })
  showModal.value = true
}

function submitForm() {
  if (editingProject.value) {
    console.log('[Project] Updated:', editingProject.value.id, { ...formData })
  } else {
    console.log('[Project] Created:', { ...formData })
  }
  showModal.value = false
}

function archiveProject(projectId: string) {
  console.log('[Project] Archived:', projectId)
}

const configOptions = [
  { label: '1 BHK', value: '1 BHK' },
  { label: '2 BHK', value: '2 BHK' },
  { label: '3 BHK', value: '3 BHK' },
  { label: '4 BHK', value: '4 BHK' },
  { label: '5 BHK', value: '5 BHK' },
  { label: 'Penthouse', value: 'Penthouse' },
]

const locationOptions = [
  { label: 'Panvel', value: 'Panvel' },
  { label: 'Ulwe', value: 'Ulwe' },
  { label: 'Kharghar', value: 'Kharghar' },
  { label: 'Dombivli', value: 'Dombivli' },
  { label: 'Kalyan', value: 'Kalyan' },
  { label: 'Thane', value: 'Thane' },
  { label: 'Vashi', value: 'Vashi' },
  { label: 'Belapur', value: 'Belapur' },
]
</script>

<template>
  <div>
    <div class="flex items-center justify-between mb-6">
      <div>
        <h1 class="text-2xl font-bold text-brand-900 dark:text-white">Projects</h1>
        <p class="text-slate-500 dark:text-slate-400 text-sm mt-1">Manage listed projects</p>
      </div>
      <UButton icon="i-lucide-plus" @click="openAddModal">Add Project</UButton>
    </div>

    <UCard>
      <div class="flex items-center gap-3 mb-4">
        <UInput v-model="searchQuery" placeholder="Search projects..." icon="i-lucide-search" class="flex-1" />
        <USelect v-model="statusFilter" :items="statusOptions" placeholder="Status" />
      </div>

      <div class="overflow-x-auto">
        <table class="w-full text-sm">
          <thead>
            <tr class="border-b border-slate-200 dark:border-slate-700">
              <th class="text-left py-3 px-4 text-xs font-semibold text-slate-500 dark:text-slate-400 uppercase">Name</th>
              <th class="text-left py-3 px-4 text-xs font-semibold text-slate-500 dark:text-slate-400 uppercase">Builder</th>
              <th class="text-left py-3 px-4 text-xs font-semibold text-slate-500 dark:text-slate-400 uppercase">RERA</th>
              <th class="text-left py-3 px-4 text-xs font-semibold text-slate-500 dark:text-slate-400 uppercase">Location</th>
              <th class="text-left py-3 px-4 text-xs font-semibold text-slate-500 dark:text-slate-400 uppercase">Status</th>
              <th class="text-left py-3 px-4 text-xs font-semibold text-slate-500 dark:text-slate-400 uppercase">Units</th>
              <th class="text-left py-3 px-4 text-xs font-semibold text-slate-500 dark:text-slate-400 uppercase">Price Range</th>
              <th class="text-left py-3 px-4 text-xs font-semibold text-slate-500 dark:text-slate-400 uppercase">Created</th>
              <th class="text-left py-3 px-4 text-xs font-semibold text-slate-500 dark:text-slate-400 uppercase">Actions</th>
            </tr>
          </thead>
          <tbody>
            <tr
              v-for="project in filteredProjects"
              :key="project.id"
              class="border-b border-slate-100 dark:border-slate-800 hover:bg-slate-50 dark:hover:bg-brand-800"
            >
              <td class="py-3 px-4">
                <p class="font-medium text-brand-900 dark:text-white">{{ project.name }}</p>
                <p class="text-xs text-slate-400">{{ project.configurations.join(', ') }}</p>
              </td>
              <td class="py-3 px-4 text-slate-600 dark:text-slate-300">{{ project.builder }}</td>
              <td class="py-3 px-4">
                <span class="text-xs font-mono text-slate-500 dark:text-slate-400">{{ project.reraNumber }}</span>
              </td>
              <td class="py-3 px-4 text-slate-600 dark:text-slate-300">{{ project.location }}</td>
              <td class="py-3 px-4">
                <UBadge :color="getStatusColor(project.status)" variant="subtle" size="xs">
                  {{ getStatusLabel(project.status) }}
                </UBadge>
              </td>
              <td class="py-3 px-4">
                <span class="text-slate-600 dark:text-slate-300">{{ project.availableUnits }}</span>
                <span class="text-slate-400 dark:text-slate-500"> / {{ project.totalUnits }}</span>
              </td>
              <td class="py-3 px-4 text-slate-600 dark:text-slate-300">
                {{ formatINRShort(project.priceRange.min) }} - {{ formatINRShort(project.priceRange.max) }}
              </td>
              <td class="py-3 px-4 text-slate-500 dark:text-slate-400 text-xs">{{ formatDate(project.createdAt) }}</td>
              <td class="py-3 px-4">
                <div class="flex items-center gap-1">
                  <UButton
                    icon="i-lucide-pencil"
                    variant="ghost"
                    color="neutral"
                    size="xs"
                    @click="openEditModal(project)"
                  />
                  <UButton
                    icon="i-lucide-archive"
                    variant="ghost"
                    color="error"
                    size="xs"
                    @click="archiveProject(project.id)"
                  />
                </div>
              </td>
            </tr>
          </tbody>
        </table>
      </div>

      <div v-if="filteredProjects.length === 0" class="text-center py-12">
        <UIcon name="i-lucide-building-2" class="text-4xl text-slate-300 mb-2" />
        <p class="text-slate-400">No projects found</p>
      </div>
    </UCard>

    <!-- Add/Edit Project Modal -->
    <UModal v-model:open="showModal">
      <template #content>
        <div class="p-6 max-h-[85vh] overflow-y-auto">
          <h3 class="text-lg font-semibold text-brand-900 dark:text-white mb-4">
            {{ editingProject ? 'Edit Project' : 'Add New Project' }}
          </h3>

          <div class="space-y-4">
            <UFormField label="Project Name" required>
              <UInput v-model="formData.name" placeholder="Enter project name" class="w-full" />
            </UFormField>

            <UFormField label="Builder Name" required>
              <UInput v-model="formData.builder" placeholder="Builder/Developer name" class="w-full" />
            </UFormField>

            <UFormField label="RERA Number" required>
              <UInput v-model="formData.reraNumber" placeholder="P52000XXXXXX" class="w-full" />
            </UFormField>

            <UFormField label="Location" required>
              <USelect v-model="formData.location" :items="locationOptions" placeholder="Select location" class="w-full" />
            </UFormField>

            <UFormField label="Status">
              <USelect v-model="formData.status" :items="statusOptions.slice(1)" class="w-full" />
            </UFormField>

            <div class="grid grid-cols-3 gap-3">
              <UFormField label="Total Units">
                <UInput v-model.number="formData.totalUnits" type="number" class="w-full" />
              </UFormField>
              <UFormField label="Total Floors">
                <UInput v-model.number="formData.totalFloors" type="number" class="w-full" />
              </UFormField>
              <UFormField label="Towers">
                <UInput v-model.number="formData.totalTowers" type="number" class="w-full" />
              </UFormField>
            </div>

            <UFormField label="Description">
              <UInput v-model="formData.shortDescription" placeholder="Brief project description" class="w-full" />
            </UFormField>

            <UFormField label="Possession Date">
              <UInput v-model="formData.possessionDate" type="date" class="w-full" />
            </UFormField>

            <UFormField label="Amenities">
              <UInput v-model="formData.amenities" placeholder="Swimming Pool, Gym, Clubhouse..." class="w-full" />
            </UFormField>

            <div class="flex gap-3 pt-4">
              <UButton variant="outline" class="flex-1" @click="showModal = false">Cancel</UButton>
              <UButton class="flex-1" @click="submitForm">
                {{ editingProject ? 'Update Project' : 'Create Project' }}
              </UButton>
            </div>
          </div>
        </div>
      </template>
    </UModal>
  </div>
</template>
