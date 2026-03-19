<script setup lang="ts">
definePageMeta({ layout: 'admin', middleware: 'auth' })

const contentSections = [
  {
    title: 'Area Guides',
    description: 'Micro-market area pages with infrastructure data, price trends, and maps.',
    icon: 'i-lucide-map-pin',
    stats: '5 areas published',
    items: [
      { name: 'Panvel', status: 'published', pages: 3 },
      { name: 'Ulwe', status: 'published', pages: 2 },
      { name: 'Kharghar', status: 'published', pages: 4 },
      { name: 'Dombivli', status: 'draft', pages: 1 },
      { name: 'Kalyan', status: 'draft', pages: 1 },
    ],
  },
  {
    title: 'Media Library',
    description: 'Upload and manage project images, floor plans, and documents.',
    icon: 'i-lucide-image',
    stats: '128 files uploaded',
    items: [
      { name: 'Project Photos', status: 'published', pages: 85 },
      { name: 'Floor Plans', status: 'published', pages: 24 },
      { name: 'Brochures', status: 'published', pages: 12 },
      { name: 'Videos', status: 'published', pages: 7 },
    ],
  },
  {
    title: 'SEO & Pages',
    description: 'Edit meta tags, landing page content, and SEO settings.',
    icon: 'i-lucide-file-text',
    stats: '8 pages configured',
    items: [
      { name: 'Homepage', status: 'published', pages: 1 },
      { name: 'About Us', status: 'published', pages: 1 },
      { name: 'Contact', status: 'published', pages: 1 },
      { name: 'Privacy Policy', status: 'published', pages: 1 },
      { name: 'Terms & Conditions', status: 'draft', pages: 1 },
    ],
  },
]

function manageSection(title: string) {
  console.log('[Content] Navigate to manage:', title)
}
</script>

<template>
  <div>
    <div class="flex items-center justify-between mb-6">
      <div>
        <h1 class="text-2xl font-bold text-brand-900 dark:text-white">Content</h1>
        <p class="text-slate-500 dark:text-slate-400 text-sm mt-1">Manage website content, images, and area guides</p>
      </div>
    </div>

    <div class="grid md:grid-cols-3 gap-6 mb-8">
      <UCard v-for="section in contentSections" :key="section.title">
        <template #header>
          <div class="flex items-center gap-2">
            <UIcon :name="section.icon" class="text-brand-600 dark:text-brand-400 text-lg" />
            <h3 class="font-semibold text-brand-900 dark:text-white">{{ section.title }}</h3>
          </div>
        </template>

        <p class="text-sm text-slate-500 dark:text-slate-400 mb-3">{{ section.description }}</p>
        <p class="text-xs text-brand-600 dark:text-brand-400 font-medium mb-3">{{ section.stats }}</p>

        <div class="space-y-2">
          <div
            v-for="item in section.items"
            :key="item.name"
            class="flex items-center justify-between p-2 bg-slate-50 dark:bg-brand-900/50 rounded text-sm"
          >
            <span class="text-slate-700 dark:text-slate-300">{{ item.name }}</span>
            <div class="flex items-center gap-2">
              <span class="text-xs text-slate-400">{{ item.pages }} {{ item.pages === 1 ? 'page' : 'pages' }}</span>
              <UBadge
                :color="item.status === 'published' ? 'success' : 'neutral'"
                variant="subtle"
                size="xs"
              >
                {{ item.status }}
              </UBadge>
            </div>
          </div>
        </div>

        <template #footer>
          <UButton variant="outline" size="sm" @click="manageSection(section.title)">
            Manage {{ section.title }}
          </UButton>
        </template>
      </UCard>
    </div>
  </div>
</template>
