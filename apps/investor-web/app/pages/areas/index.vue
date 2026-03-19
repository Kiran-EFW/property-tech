<script setup lang="ts">
useHead({
  title: 'Mumbai Micro-Market Guides | PropTech',
  meta: [
    { name: 'description', content: 'Explore Mumbai micro-market investment guides. Data-driven insights into Panvel, Dombivli, Kalyan, Taloja, Kharghar & Ulwe - infrastructure updates, price trends, and growth projections.' },
    { name: 'keywords', content: 'Mumbai micro-markets, Navi Mumbai investment, Panvel real estate guide, Dombivli property market, Kalyan area guide, property investment areas Mumbai' },
  ],
  link: [
    { rel: 'canonical', href: 'https://www.proptech.in/areas' },
  ],
})

useSeoMeta({
  title: 'Mumbai Micro-Market Guides | PropTech',
  description: 'Explore Mumbai micro-market investment guides. Data-driven insights into Panvel, Dombivli, Kalyan, Taloja, Kharghar & Ulwe - infrastructure updates, price trends, and growth projections.',
  ogTitle: 'Mumbai Micro-Market Guides | PropTech',
  ogDescription: 'Data-driven area guides for Mumbai\'s fastest-growing investment corridors. Infrastructure timelines, price trends, and growth projections.',
  ogType: 'website',
  ogUrl: 'https://www.proptech.in/areas',
  ogImage: 'https://www.proptech.in/og-areas.jpg',
  twitterCard: 'summary_large_image',
  twitterTitle: 'Mumbai Micro-Market Guides | PropTech',
  twitterDescription: 'Explore Mumbai micro-market investment guides with data-driven insights.',
})

useHead({
  script: [
    {
      type: 'application/ld+json',
      innerHTML: JSON.stringify({
        '@context': 'https://schema.org',
        '@type': 'ItemList',
        'name': 'Mumbai Micro-Market Investment Guides',
        'description': 'Area-wise investment guides for Mumbai metropolitan region',
        'numberOfItems': 6,
        'itemListElement': [
          { '@type': 'ListItem', position: 1, name: 'Panvel Investment Guide', url: 'https://www.proptech.in/areas/panvel' },
          { '@type': 'ListItem', position: 2, name: 'Dombivli East Investment Guide', url: 'https://www.proptech.in/areas/dombivli-east' },
          { '@type': 'ListItem', position: 3, name: 'Kalyan-Shilphata Investment Guide', url: 'https://www.proptech.in/areas/kalyan-shilphata' },
          { '@type': 'ListItem', position: 4, name: 'Taloja Investment Guide', url: 'https://www.proptech.in/areas/taloja' },
          { '@type': 'ListItem', position: 5, name: 'Kharghar Investment Guide', url: 'https://www.proptech.in/areas/kharghar' },
          { '@type': 'ListItem', position: 6, name: 'Ulwe Investment Guide', url: 'https://www.proptech.in/areas/ulwe' },
        ],
      }),
    },
  ],
})

interface Area {
  name: string
  slug: string
  tagline: string
  avgPriceSqFt: string
  yoyGrowth: string
  growthPotential: string
  growthLevel: 'High' | 'Medium' | 'Stable'
  keyInfra: string[]
  projectCount: number
  description: string
  imageColor: string
}

const areas: Area[] = [
  {
    name: 'Panvel',
    slug: 'panvel',
    tagline: 'Navi Mumbai Airport Zone',
    avgPriceSqFt: '₹5,200/sq ft',
    yoyGrowth: '+12.4%',
    growthPotential: '18-22%',
    growthLevel: 'High',
    keyInfra: ['Navi Mumbai Airport', 'MTHL Connectivity', 'Metro Line 12'],
    projectCount: 2,
    description: 'Panvel is the epicenter of Navi Mumbai\'s growth story, driven by the upcoming international airport, Mumbai Trans Harbour Link, and Metro Line 12.',
    imageColor: 'bg-teal-50',
  },
  {
    name: 'Dombivli East',
    slug: 'dombivli-east',
    tagline: 'Affordable Growth Corridor',
    avgPriceSqFt: '₹4,800/sq ft',
    yoyGrowth: '+9.8%',
    growthPotential: '14-18%',
    growthLevel: 'Medium',
    keyInfra: ['Metro Line 12', 'Kalyan-Dombivli Elevated Road', 'MIDC Expansion'],
    projectCount: 1,
    description: 'Dombivli East is a well-established residential hub with the best rail connectivity to Mumbai CST and upcoming metro and road infrastructure upgrades.',
    imageColor: 'bg-blue-50',
  },
  {
    name: 'Kalyan-Shilphata',
    slug: 'kalyan-shilphata',
    tagline: 'Metro & Expressway Corridor',
    avgPriceSqFt: '₹5,500/sq ft',
    yoyGrowth: '+11.2%',
    growthPotential: '16-20%',
    growthLevel: 'High',
    keyInfra: ['Mumbai-Nagpur Expressway', 'Kalyan Ring Road', 'Metro Extension'],
    projectCount: 2,
    description: 'The Kalyan-Shilphata corridor benefits from the Mumbai-Nagpur Expressway access point and improved road infrastructure connecting Thane, Kalyan, and Navi Mumbai.',
    imageColor: 'bg-green-50',
  },
  {
    name: 'Taloja',
    slug: 'taloja',
    tagline: 'Affordable Entry Point',
    avgPriceSqFt: '₹4,200/sq ft',
    yoyGrowth: '+10.5%',
    growthPotential: '15-19%',
    growthLevel: 'Medium',
    keyInfra: ['CIDCO Mass Housing', 'Panvel-Karjat Rail', 'Industrial Zone Growth'],
    projectCount: 1,
    description: 'Taloja offers the most affordable entry point in Navi Mumbai with CIDCO-planned infrastructure and proximity to the Panvel growth corridor.',
    imageColor: 'bg-lime-50',
  },
  {
    name: 'Kharghar',
    slug: 'kharghar',
    tagline: 'Established Premium Hub',
    avgPriceSqFt: '₹8,500/sq ft',
    yoyGrowth: '+8.3%',
    growthPotential: '12-15%',
    growthLevel: 'Stable',
    keyInfra: ['CIDCO Exhibition Centre', 'Corporate Hub Expansion', 'Golf Course'],
    projectCount: 1,
    description: 'Kharghar is Navi Mumbai\'s most established and well-planned node, offering premium lifestyle with Central Park, Golf Course, and excellent social infrastructure.',
    imageColor: 'bg-purple-50',
  },
  {
    name: 'Ulwe',
    slug: 'ulwe',
    tagline: 'Highest Growth Potential',
    avgPriceSqFt: '₹6,100/sq ft',
    yoyGrowth: '+15.2%',
    growthPotential: '20-25%',
    growthLevel: 'High',
    keyInfra: ['Mumbai Trans Harbour Link', 'Navi Mumbai Airport', 'Navi Mumbai Metro'],
    projectCount: 1,
    description: 'Ulwe is positioned for the highest growth potential in the MMR, with MTHL reducing South Mumbai travel to 20 minutes and the upcoming airport within 10 km.',
    imageColor: 'bg-orange-50',
  },
]

function getGrowthColor(level: string) {
  switch (level) {
    case 'High': return 'success' as const
    case 'Medium': return 'warning' as const
    case 'Stable': return 'info' as const
    default: return 'neutral' as const
  }
}
</script>

<template>
  <div>
    <!-- Page Header -->
    <section class="bg-brand-800 text-white py-12 px-4">
      <div class="max-w-6xl mx-auto">
        <h1 class="text-3xl md:text-4xl font-bold">Mumbai Micro-Market Guides</h1>
        <p class="text-slate-300 mt-2 max-w-2xl">
          Data-driven investment guides for Mumbai's fastest-growing corridors.
          Understand infrastructure timelines, price trends, growth projections, and available projects in each area.
        </p>
        <div class="flex items-center gap-3 mt-4">
          <UBadge color="neutral" variant="subtle">
            {{ areas.length }} Areas Covered
          </UBadge>
          <UBadge color="success" variant="subtle">
            <UIcon name="i-lucide-trending-up" class="mr-1" />
            Updated Monthly
          </UBadge>
        </div>
      </div>
    </section>

    <!-- Areas Grid -->
    <section class="py-10 px-4 bg-slate-50 dark:bg-brand-900/50 min-h-screen">
      <div class="max-w-6xl mx-auto">
        <div class="grid sm:grid-cols-2 lg:grid-cols-3 gap-6">
          <NuxtLink
            v-for="area in areas"
            :key="area.slug"
            :to="`/areas/${area.slug}`"
            class="block group"
          >
            <UCard class="h-full hover:shadow-lg transition-all duration-200 hover:-translate-y-0.5">
              <!-- Image Placeholder -->
              <div :class="[area.imageColor, 'h-36 rounded-lg flex items-center justify-center mb-4 relative']">
                <UIcon name="i-lucide-map" class="text-5xl text-slate-300" />
                <div class="absolute top-2 right-2">
                  <UBadge :color="getGrowthColor(area.growthLevel)" variant="solid" size="xs">
                    {{ area.growthLevel }} Growth
                  </UBadge>
                </div>
              </div>

              <!-- Area Name & Tagline -->
              <h3 class="text-xl font-bold text-brand-900 dark:text-white group-hover:text-brand-600 dark:group-hover:text-brand-400 transition-colors">
                {{ area.name }}
              </h3>
              <p class="text-sm text-brand-600 dark:text-brand-400 font-medium mt-0.5">{{ area.tagline }}</p>

              <!-- Description -->
              <p class="text-sm text-slate-500 dark:text-slate-400 mt-2 line-clamp-2">{{ area.description }}</p>

              <!-- Stats Grid -->
              <div class="grid grid-cols-3 gap-3 mt-4 pt-4 border-t border-slate-100 dark:border-slate-800">
                <div>
                  <p class="text-xs text-slate-400">Avg Price</p>
                  <p class="font-semibold text-slate-900 dark:text-white text-sm">{{ area.avgPriceSqFt }}</p>
                </div>
                <div>
                  <p class="text-xs text-slate-400">YoY Growth</p>
                  <p class="font-semibold text-accent-600 dark:text-accent-400 text-sm">{{ area.yoyGrowth }}</p>
                </div>
                <div>
                  <p class="text-xs text-slate-400">Projects</p>
                  <p class="font-semibold text-slate-900 dark:text-white text-sm">{{ area.projectCount }}</p>
                </div>
              </div>

              <!-- Key Infrastructure -->
              <div class="mt-4">
                <p class="text-xs text-slate-400 mb-2">Key Infrastructure</p>
                <div class="flex flex-wrap gap-1.5">
                  <UBadge
                    v-for="infra in area.keyInfra"
                    :key="infra"
                    color="neutral"
                    variant="subtle"
                    size="xs"
                  >
                    {{ infra }}
                  </UBadge>
                </div>
              </div>

              <!-- CTA -->
              <div class="mt-4 pt-3 border-t border-slate-100 dark:border-slate-800 flex items-center text-sm text-brand-600 dark:text-brand-400 font-medium">
                Read Full Area Guide
                <UIcon name="i-lucide-arrow-right" class="ml-1" />
              </div>
            </UCard>
          </NuxtLink>
        </div>
      </div>
    </section>

    <!-- Investment Comparison -->
    <section class="py-12 px-4 bg-white dark:bg-brand-900">
      <div class="max-w-6xl mx-auto">
        <h2 class="text-2xl font-bold text-brand-900 dark:text-white text-center mb-8">Area Comparison at a Glance</h2>
        <div class="overflow-x-auto">
          <table class="w-full text-sm">
            <thead>
              <tr class="border-b-2 border-slate-200 dark:border-slate-700">
                <th class="text-left py-3 px-4 text-slate-500 dark:text-slate-400 font-medium">Area</th>
                <th class="text-left py-3 px-4 text-slate-500 dark:text-slate-400 font-medium">Avg Price/sq ft</th>
                <th class="text-left py-3 px-4 text-slate-500 dark:text-slate-400 font-medium">YoY Growth</th>
                <th class="text-left py-3 px-4 text-slate-500 dark:text-slate-400 font-medium">Growth Potential</th>
                <th class="text-left py-3 px-4 text-slate-500 dark:text-slate-400 font-medium">Key Driver</th>
                <th class="text-left py-3 px-4 text-slate-500 dark:text-slate-400 font-medium">Projects</th>
              </tr>
            </thead>
            <tbody>
              <tr
                v-for="area in areas"
                :key="area.slug"
                class="border-b border-slate-100 dark:border-slate-800 hover:bg-slate-50 dark:hover:bg-brand-800"
              >
                <td class="py-3 px-4">
                  <NuxtLink :to="`/areas/${area.slug}`" class="font-semibold text-brand-900 dark:text-white hover:text-brand-600 dark:hover:text-brand-400">
                    {{ area.name }}
                  </NuxtLink>
                </td>
                <td class="py-3 px-4 font-medium">{{ area.avgPriceSqFt }}</td>
                <td class="py-3 px-4 font-semibold text-accent-600 dark:text-accent-400">{{ area.yoyGrowth }}</td>
                <td class="py-3 px-4">
                  <UBadge :color="getGrowthColor(area.growthLevel)" variant="subtle" size="xs">
                    {{ area.growthPotential }}
                  </UBadge>
                </td>
                <td class="py-3 px-4 text-slate-600 dark:text-slate-300">{{ area.keyInfra[0] }}</td>
                <td class="py-3 px-4 font-medium">{{ area.projectCount }}</td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>
    </section>

    <!-- WhatsApp CTA -->
    <section class="py-12 px-4 bg-brand-800 text-white text-center">
      <div class="max-w-2xl mx-auto">
        <h2 class="text-xl md:text-2xl font-bold">Need help choosing the right area?</h2>
        <p class="text-slate-300 mt-2 mb-6">
          Tell us your budget and investment goals, and we will recommend the best micro-market for you.
        </p>
        <UButton
          href="https://wa.me/919876543210?text=Hi%2C%20I%20need%20help%20choosing%20the%20right%20area%20for%20property%20investment%20in%20Mumbai.%20Can%20you%20guide%20me%3F"
          target="_blank"
          size="xl"
          class="bg-green-600 hover:bg-green-700 text-white"
          icon="i-lucide-message-circle"
        >
          Get Area Recommendations
        </UButton>
      </div>
    </section>
  </div>
</template>
