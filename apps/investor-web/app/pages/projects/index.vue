<script setup lang="ts">
useHead({
  title: 'Investment Properties in Mumbai | PropTech',
  meta: [
    { name: 'description', content: 'Browse RERA-verified investment properties in Mumbai micro-markets. Compare projects in Panvel, Dombivli, Kalyan, Taloja, Kharghar & Ulwe with transparent all-inclusive pricing.' },
    { name: 'keywords', content: 'Mumbai property investment, RERA verified flats, Panvel projects, Dombivli flats, Kalyan real estate, Navi Mumbai apartments' },
  ],
  link: [
    { rel: 'canonical', href: 'https://www.proptech.in/projects' },
  ],
})

useSeoMeta({
  title: 'Investment Properties in Mumbai | PropTech',
  description: 'Browse RERA-verified investment properties in Mumbai micro-markets. Compare projects in Panvel, Dombivli, Kalyan, Taloja, Kharghar & Ulwe with transparent all-inclusive pricing.',
  ogTitle: 'Investment Properties in Mumbai | PropTech',
  ogDescription: 'RERA-verified projects with on-ground due diligence and all-inclusive pricing. Find your next investment property in Mumbai.',
  ogType: 'website',
  ogUrl: 'https://www.proptech.in/projects',
  ogImage: 'https://www.proptech.in/og-projects.jpg',
  twitterCard: 'summary_large_image',
  twitterTitle: 'Investment Properties in Mumbai | PropTech',
  twitterDescription: 'Browse RERA-verified investment properties in Mumbai micro-markets with transparent pricing.',
})

useHead({
  script: [
    {
      type: 'application/ld+json',
      innerHTML: JSON.stringify({
        '@context': 'https://schema.org',
        '@type': 'ItemList',
        'name': 'RERA Verified Investment Properties in Mumbai',
        'description': 'Curated list of RERA-verified real estate investment projects in Mumbai micro-markets',
        'numberOfItems': 8,
        'itemListElement': [
          { '@type': 'ListItem', position: 1, name: 'Royal Greens, Panvel', url: 'https://www.proptech.in/projects/royal-greens-panvel' },
          { '@type': 'ListItem', position: 2, name: 'Metro Heights, Dombivli East', url: 'https://www.proptech.in/projects/metro-heights-dombivli' },
          { '@type': 'ListItem', position: 3, name: 'Green Valley Residences, Kalyan-Shilphata', url: 'https://www.proptech.in/projects/green-valley-kalyan' },
          { '@type': 'ListItem', position: 4, name: 'Skyline Towers, Kharghar', url: 'https://www.proptech.in/projects/skyline-towers-kharghar' },
          { '@type': 'ListItem', position: 5, name: 'Navi Heights, Ulwe', url: 'https://www.proptech.in/projects/navi-heights-ulwe' },
          { '@type': 'ListItem', position: 6, name: 'Taloja Greens, Taloja', url: 'https://www.proptech.in/projects/taloja-greens' },
          { '@type': 'ListItem', position: 7, name: 'Panvel Paradise, Panvel', url: 'https://www.proptech.in/projects/panvel-paradise' },
          { '@type': 'ListItem', position: 8, name: 'Shilphata Residency, Kalyan-Shilphata', url: 'https://www.proptech.in/projects/shilphata-residency' },
        ],
      }),
    },
  ],
})

interface Project {
  slug: string
  name: string
  location: string
  builder: string
  priceRange: string
  pricePerSqFt: string
  configs: string[]
  reraNo: string
  growthPotential: string
  possession: string
  status: string
  imageColor: string
  featured: boolean
}

const allProjects: Project[] = [
  {
    slug: 'royal-greens-panvel',
    name: 'Royal Greens',
    location: 'Panvel',
    builder: 'Arihant Superstructures',
    priceRange: '₹45 L - ₹1.1 Cr',
    pricePerSqFt: '₹5,200',
    configs: ['1 BHK', '2 BHK', '3 BHK'],
    reraNo: 'P52100048001',
    growthPotential: '18-22%',
    possession: 'Dec 2026',
    status: 'Under Construction',
    imageColor: 'bg-teal-100',
    featured: true,
  },
  {
    slug: 'metro-heights-dombivli',
    name: 'Metro Heights',
    location: 'Dombivli East',
    builder: 'Runwal Group',
    priceRange: '₹38 L - ₹85 L',
    pricePerSqFt: '₹4,800',
    configs: ['1 BHK', '2 BHK'],
    reraNo: 'P51700032145',
    growthPotential: '14-18%',
    possession: 'Mar 2027',
    status: 'Under Construction',
    imageColor: 'bg-blue-100',
    featured: true,
  },
  {
    slug: 'green-valley-kalyan',
    name: 'Green Valley Residences',
    location: 'Kalyan-Shilphata',
    builder: 'Lodha Group',
    priceRange: '₹52 L - ₹1.4 Cr',
    pricePerSqFt: '₹5,500',
    configs: ['1 BHK', '2 BHK', '3 BHK'],
    reraNo: 'P51700041287',
    growthPotential: '16-20%',
    possession: 'Jun 2027',
    status: 'Under Construction',
    imageColor: 'bg-green-100',
    featured: true,
  },
  {
    slug: 'skyline-towers-kharghar',
    name: 'Skyline Towers',
    location: 'Kharghar',
    builder: 'Adhiraj Constructions',
    priceRange: '₹78 L - ₹1.8 Cr',
    pricePerSqFt: '₹8,500',
    configs: ['2 BHK', '3 BHK'],
    reraNo: 'P52100051439',
    growthPotential: '12-15%',
    possession: 'Sep 2026',
    status: 'Ready to Move',
    imageColor: 'bg-purple-100',
    featured: true,
  },
  {
    slug: 'navi-heights-ulwe',
    name: 'Navi Heights',
    location: 'Ulwe',
    builder: 'Balaji Group',
    priceRange: '₹42 L - ₹1.05 Cr',
    pricePerSqFt: '₹6,100',
    configs: ['1 BHK', '2 BHK', '3 BHK'],
    reraNo: 'P52100053782',
    growthPotential: '20-25%',
    possession: 'Mar 2027',
    status: 'Under Construction',
    imageColor: 'bg-orange-100',
    featured: false,
  },
  {
    slug: 'taloja-greens',
    name: 'Taloja Greens',
    location: 'Taloja',
    builder: 'Puraniks',
    priceRange: '₹32 L - ₹72 L',
    pricePerSqFt: '₹4,200',
    configs: ['1 BHK', '2 BHK'],
    reraNo: 'P52100049876',
    growthPotential: '15-19%',
    possession: 'Dec 2027',
    status: 'Pre-Launch',
    imageColor: 'bg-lime-100',
    featured: false,
  },
  {
    slug: 'panvel-paradise',
    name: 'Panvel Paradise',
    location: 'Panvel',
    builder: 'Haware Builders',
    priceRange: '₹55 L - ₹1.3 Cr',
    pricePerSqFt: '₹5,400',
    configs: ['2 BHK', '3 BHK'],
    reraNo: 'P52100047632',
    growthPotential: '17-21%',
    possession: 'Sep 2027',
    status: 'Under Construction',
    imageColor: 'bg-cyan-100',
    featured: false,
  },
  {
    slug: 'shilphata-residency',
    name: 'Shilphata Residency',
    location: 'Kalyan-Shilphata',
    builder: 'Dosti Realty',
    priceRange: '₹48 L - ₹1.15 Cr',
    pricePerSqFt: '₹5,300',
    configs: ['1 BHK', '2 BHK', '3 BHK'],
    reraNo: 'P51700038956',
    growthPotential: '15-19%',
    possession: 'Jun 2028',
    status: 'Pre-Launch',
    imageColor: 'bg-rose-100',
    featured: false,
  },
]

const selectedLocation = ref('')
const selectedBhk = ref('')
const selectedBudget = ref('')
const selectedSort = ref('')

const locations = ['All Locations', 'Panvel', 'Dombivli East', 'Kalyan-Shilphata', 'Taloja', 'Kharghar', 'Ulwe']
const bhkOptions = ['All BHK', '1 BHK', '2 BHK', '3 BHK']
const budgetOptions = ['All Budgets', 'Under ₹50 L', '₹50 L - ₹1 Cr', '₹1 Cr - ₹2 Cr']
const sortOptions = ['Newest First', 'Price: Low to High', 'Price: High to Low', 'Growth Potential']

const filteredProjects = computed(() => {
  let result = [...allProjects]

  if (selectedLocation.value && selectedLocation.value !== 'All Locations') {
    result = result.filter(p => p.location === selectedLocation.value)
  }

  if (selectedBhk.value && selectedBhk.value !== 'All BHK') {
    result = result.filter(p => p.configs.includes(selectedBhk.value))
  }

  return result
})

function getStatusColor(status: string) {
  switch (status) {
    case 'Ready to Move': return 'success' as const
    case 'Under Construction': return 'warning' as const
    case 'Pre-Launch': return 'info' as const
    default: return 'neutral' as const
  }
}
</script>

<template>
  <div>
    <!-- Page Header -->
    <section class="bg-brand-800 text-white py-12 px-4">
      <div class="max-w-6xl mx-auto">
        <h1 class="text-3xl md:text-4xl font-bold">Investment Properties</h1>
        <p class="text-slate-300 mt-2 max-w-xl">
          Browse RERA-verified projects across Mumbai's fastest-growing micro-markets.
          Every project includes on-ground due diligence and all-inclusive pricing.
        </p>
        <div class="flex items-center gap-3 mt-4">
          <UBadge color="success" variant="subtle">
            <UIcon name="i-lucide-shield-check" class="mr-1" />
            All RERA Verified
          </UBadge>
          <UBadge color="neutral" variant="subtle">
            {{ allProjects.length }} Projects Listed
          </UBadge>
        </div>
      </div>
    </section>

    <!-- Filters -->
    <section class="bg-white dark:bg-brand-950 border-b border-slate-200 dark:border-slate-700 py-4 px-4 sticky top-0 z-10">
      <div class="max-w-6xl mx-auto">
        <div class="flex flex-wrap gap-3 items-center">
          <USelect
            v-model="selectedLocation"
            :items="locations"
            placeholder="Location"
            class="w-40"
          />
          <USelect
            v-model="selectedBhk"
            :items="bhkOptions"
            placeholder="BHK"
            class="w-32"
          />
          <USelect
            v-model="selectedBudget"
            :items="budgetOptions"
            placeholder="Budget"
            class="w-44"
          />
          <USelect
            v-model="selectedSort"
            :items="sortOptions"
            placeholder="Sort By"
            class="w-44"
          />
          <div class="ml-auto text-sm text-slate-500 dark:text-slate-400 hidden md:block">
            Showing {{ filteredProjects.length }} of {{ allProjects.length }} projects
          </div>
        </div>
      </div>
    </section>

    <!-- Projects Grid -->
    <section class="py-8 px-4 bg-slate-50 dark:bg-brand-900/50 min-h-screen">
      <div class="max-w-6xl mx-auto">
        <div class="grid sm:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-6">
          <NuxtLink
            v-for="project in filteredProjects"
            :key="project.slug"
            :to="`/projects/${project.slug}`"
            class="block group"
          >
            <UCard class="h-full hover:shadow-lg transition-all duration-200 hover:-translate-y-0.5">
              <!-- Image Placeholder -->
              <div :class="[project.imageColor, 'h-44 rounded-lg flex items-center justify-center mb-4 relative overflow-hidden']">
                <UIcon name="i-lucide-building-2" class="text-5xl text-slate-400" />
                <!-- Status Badge -->
                <div class="absolute top-2 left-2">
                  <UBadge :color="getStatusColor(project.status)" variant="solid" size="xs">
                    {{ project.status }}
                  </UBadge>
                </div>
                <!-- Featured Badge -->
                <div v-if="project.featured" class="absolute top-2 right-2">
                  <UBadge color="warning" variant="solid" size="xs">
                    <UIcon name="i-lucide-star" class="mr-0.5" />
                    Featured
                  </UBadge>
                </div>
              </div>

              <!-- RERA & Growth Badges -->
              <div class="flex items-center gap-2 mb-2">
                <UBadge color="success" variant="subtle" size="xs">
                  <UIcon name="i-lucide-shield-check" class="mr-0.5" />
                  RERA
                </UBadge>
                <UBadge color="info" variant="subtle" size="xs">
                  {{ project.growthPotential }} Growth
                </UBadge>
              </div>

              <!-- Project Name & Location -->
              <h3 class="font-semibold text-brand-900 dark:text-white group-hover:text-brand-600 dark:group-hover:text-brand-400 transition-colors text-base">
                {{ project.name }}
              </h3>
              <p class="text-sm text-slate-500 dark:text-slate-400 flex items-center gap-1 mt-1">
                <UIcon name="i-lucide-map-pin" class="text-xs flex-shrink-0" />
                {{ project.location }}
              </p>

              <!-- Price -->
              <p class="text-lg font-bold text-slate-900 dark:text-white mt-2">{{ project.priceRange }}</p>
              <p class="text-xs text-slate-400">{{ project.pricePerSqFt }}/sq ft (carpet)</p>

              <!-- BHK Configs -->
              <div class="flex flex-wrap gap-1.5 mt-3">
                <UBadge v-for="config in project.configs" :key="config" color="neutral" variant="subtle" size="xs">
                  {{ config }}
                </UBadge>
              </div>

              <!-- Footer Info -->
              <div class="mt-3 pt-3 border-t border-slate-100 dark:border-slate-800 flex justify-between items-center text-xs text-slate-400">
                <span>{{ project.builder }}</span>
                <span>{{ project.possession }}</span>
              </div>

              <!-- RERA Number -->
              <p class="text-xs text-slate-400 mt-1 font-mono">
                RERA: {{ project.reraNo }}
              </p>
            </UCard>
          </NuxtLink>
        </div>

        <!-- No Results -->
        <div v-if="filteredProjects.length === 0" class="text-center py-16">
          <UIcon name="i-lucide-search-x" class="text-5xl text-slate-300 mb-4" />
          <h3 class="text-lg font-semibold text-brand-900 dark:text-white">No projects match your filters</h3>
          <p class="text-slate-500 dark:text-slate-400 mt-1 mb-4">Try adjusting your search criteria or browse all projects.</p>
          <UButton variant="outline" @click="selectedLocation = ''; selectedBhk = ''; selectedBudget = ''">
            Clear All Filters
          </UButton>
        </div>
      </div>
    </section>

    <!-- WhatsApp CTA -->
    <section class="py-12 px-4 bg-brand-800 text-white text-center">
      <div class="max-w-2xl mx-auto">
        <h2 class="text-xl md:text-2xl font-bold">Can't find what you're looking for?</h2>
        <p class="text-slate-300 mt-2 mb-6">
          Tell us your requirements on WhatsApp and we will curate options for you. No spam, no pressure.
        </p>
        <UButton
          href="https://wa.me/919876543210?text=Hi%2C%20I%20am%20looking%20for%20an%20investment%20property%20in%20Mumbai.%20Can%20you%20help%20me%20find%20the%20right%20one%3F"
          target="_blank"
          size="xl"
          class="bg-green-600 hover:bg-green-700 text-white"
          icon="i-lucide-message-circle"
        >
          Tell Us Your Requirements
        </UButton>
      </div>
    </section>
  </div>
</template>
