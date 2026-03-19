<script setup lang="ts">
useHead({
  title: 'PropTech - Mumbai Real Estate Investment Platform',
  meta: [
    { name: 'description', content: 'Invest in Mumbai\'s fastest-growing micro-markets. RERA-verified projects in Panvel, Dombivli, Kalyan, Taloja, Kharghar & Ulwe with transparent all-inclusive pricing and on-ground due diligence.' },
    { name: 'keywords', content: 'Mumbai real estate investment, RERA verified projects, Navi Mumbai property, Panvel investment, micro-market investing, PropTech India' },
  ],
  link: [
    { rel: 'canonical', href: 'https://www.proptech.in/' },
  ],
})

useSeoMeta({
  title: 'PropTech - Mumbai Real Estate Investment Platform',
  description: 'Invest in Mumbai\'s fastest-growing micro-markets. RERA-verified projects in Panvel, Dombivli, Kalyan, Taloja, Kharghar & Ulwe with transparent all-inclusive pricing and on-ground due diligence.',
  ogTitle: 'PropTech - Mumbai Real Estate Investment Platform',
  ogDescription: 'RERA-verified, curated real estate investment opportunities in Mumbai micro-markets. On-ground due diligence, transparent pricing, trust-first approach.',
  ogType: 'website',
  ogUrl: 'https://www.proptech.in/',
  ogImage: 'https://www.proptech.in/og-home.jpg',
  twitterCard: 'summary_large_image',
  twitterTitle: 'PropTech - Mumbai Real Estate Investment Platform',
  twitterDescription: 'Invest in Mumbai\'s fastest-growing micro-markets with RERA-verified projects and transparent pricing.',
})

useHead({
  script: [
    {
      type: 'application/ld+json',
      innerHTML: JSON.stringify({
        '@context': 'https://schema.org',
        '@type': 'RealEstateAgent',
        'name': 'PropTech',
        'description': 'MahaRERA-licensed real estate investment platform for Mumbai micro-markets',
        'url': 'https://www.proptech.in',
        'logo': 'https://www.proptech.in/logo.png',
        'address': {
          '@type': 'PostalAddress',
          'addressLocality': 'Mumbai',
          'addressRegion': 'Maharashtra',
          'addressCountry': 'IN',
        },
        'areaServed': {
          '@type': 'City',
          'name': 'Mumbai',
        },
        'sameAs': [
          'https://www.instagram.com/proptech.in',
          'https://www.linkedin.com/company/proptech-in',
        ],
      }),
    },
  ],
})

const featuredProjects = [
  {
    slug: 'royal-greens-panvel',
    name: 'Royal Greens',
    location: 'Panvel',
    builder: 'Arihant Superstructures',
    priceRange: '₹45 L - ₹1.1 Cr',
    configs: ['1 BHK', '2 BHK', '3 BHK'],
    reraNo: 'P52100048001',
    growthPotential: '18-22%',
    possession: 'Dec 2026',
    imageColor: 'bg-brand-100',
  },
  {
    slug: 'metro-heights-dombivli',
    name: 'Metro Heights',
    location: 'Dombivli East',
    builder: 'Runwal Group',
    priceRange: '₹38 L - ₹85 L',
    configs: ['1 BHK', '2 BHK'],
    reraNo: 'P51700032145',
    growthPotential: '14-18%',
    possession: 'Mar 2027',
    imageColor: 'bg-blue-100',
  },
  {
    slug: 'green-valley-kalyan',
    name: 'Green Valley Residences',
    location: 'Kalyan-Shilphata',
    builder: 'Lodha Group',
    priceRange: '₹52 L - ₹1.4 Cr',
    configs: ['1 BHK', '2 BHK', '3 BHK'],
    reraNo: 'P51700041287',
    growthPotential: '16-20%',
    possession: 'Jun 2027',
    imageColor: 'bg-green-100',
  },
  {
    slug: 'skyline-towers-kharghar',
    name: 'Skyline Towers',
    location: 'Kharghar',
    builder: 'Adhiraj Constructions',
    priceRange: '₹78 L - ₹1.8 Cr',
    configs: ['2 BHK', '3 BHK'],
    reraNo: 'P52100051439',
    growthPotential: '12-15%',
    possession: 'Sep 2026',
    imageColor: 'bg-purple-100',
  },
]

const microMarkets = [
  { name: 'Panvel', slug: 'panvel', avgPrice: '₹5,200/sq ft', growth: '18-22%', driver: 'Navi Mumbai Airport' },
  { name: 'Dombivli East', slug: 'dombivli-east', avgPrice: '₹4,800/sq ft', growth: '14-18%', driver: 'Metro Line 12' },
  { name: 'Kalyan-Shilphata', slug: 'kalyan-shilphata', avgPrice: '₹5,500/sq ft', growth: '16-20%', driver: 'Mumbai-Nagpur Expressway' },
  { name: 'Taloja', slug: 'taloja', avgPrice: '₹4,200/sq ft', growth: '15-19%', driver: 'CIDCO Mass Housing' },
  { name: 'Kharghar', slug: 'kharghar', avgPrice: '₹8,500/sq ft', growth: '12-15%', driver: 'Corporate Hub Growth' },
  { name: 'Ulwe', slug: 'ulwe', avgPrice: '₹6,100/sq ft', growth: '20-25%', driver: 'Trans Harbour Link' },
]

const howItWorks = [
  {
    step: 1,
    title: 'We Curate & Verify',
    description: 'Our team shortlists projects through MahaRERA verification, builder background checks, title clearance, and on-site inspections. Only 1 in 10 projects make the cut.',
    icon: 'i-lucide-shield-check',
  },
  {
    step: 2,
    title: 'You Browse & Compare',
    description: 'Explore curated projects with all-inclusive pricing, area guides, infrastructure timelines, and growth projections. No hidden costs, no fine print.',
    icon: 'i-lucide-search',
  },
  {
    step: 3,
    title: 'We Guide Your Visit',
    description: 'Schedule a site visit at your convenience. Our trained agents accompany you, explain everything clearly, and never push a sale.',
    icon: 'i-lucide-map-pin',
  },
  {
    step: 4,
    title: 'Invest with Confidence',
    description: 'Complete your purchase with full documentation support. We handle RERA compliance, agreement drafting, and post-sale registration assistance.',
    icon: 'i-lucide-check-circle',
  },
]

const trustSignals = [
  { label: 'RERA Verified', icon: 'i-lucide-shield-check', description: 'Every project is MahaRERA registered and verified' },
  { label: 'All-Inclusive Pricing', icon: 'i-lucide-indian-rupee', description: 'GST, stamp duty, registration — everything included' },
  { label: 'On-Ground DD Completed', icon: 'i-lucide-clipboard-check', description: 'Physical site inspections and builder reputation checks' },
  { label: 'Builder Reputation Checked', icon: 'i-lucide-building-2', description: 'Track record, delivery history, and financial health verified' },
]
</script>

<template>
  <div>
    <!-- Hero Section -->
    <section class="bg-brand-800 text-white py-24 md:py-32 px-4">
      <div class="max-w-5xl mx-auto text-center">
        <UBadge color="neutral" variant="subtle" class="mb-6 text-sm px-4 py-1.5">
          <UIcon name="i-lucide-shield-check" class="mr-1" />
          MahaRERA Licensed Platform
        </UBadge>
        <h1 class="text-3xl sm:text-4xl md:text-5xl lg:text-6xl font-bold leading-tight tracking-tight">
          Invest in Mumbai's Fastest-Growing
          <br class="hidden sm:block" />
          <span class="text-accent-400">Micro-Markets</span>
        </h1>
        <p class="mt-6 text-base sm:text-lg text-brand-200 max-w-2xl mx-auto leading-relaxed">
          RERA-verified projects in Panvel, Dombivli, Kalyan, Taloja, Kharghar &amp; Ulwe.
          On-ground due diligence, transparent all-inclusive pricing, and zero spam calls.
        </p>

        <!-- Stats strip -->
        <div class="mt-10 flex flex-wrap justify-center gap-8 text-sm">
          <div>
            <span class="block text-2xl font-bold text-accent-400">200+</span>
            <span class="text-brand-300">Verified Projects</span>
          </div>
          <div class="hidden sm:block w-px bg-brand-600" />
          <div>
            <span class="block text-2xl font-bold text-accent-400">6</span>
            <span class="text-brand-300">Micro-Markets</span>
          </div>
          <div class="hidden sm:block w-px bg-brand-600" />
          <div>
            <span class="block text-2xl font-bold text-accent-400">12-25%</span>
            <span class="text-brand-300">Growth Potential</span>
          </div>
        </div>

        <div class="mt-10 flex flex-col sm:flex-row gap-4 justify-center">
          <UButton
            to="/projects"
            size="xl"
            class="bg-accent-500 hover:bg-accent-600 text-white"
            icon="i-lucide-building-2"
          >
            Explore Projects
          </UButton>
          <UButton
            href="https://wa.me/919876543210?text=Hi%2C%20I%20want%20to%20know%20about%20investment%20opportunities%20in%20Mumbai%20micro-markets"
            target="_blank"
            size="xl"
            variant="outline"
            class="border-white text-white hover:bg-white/10"
            icon="i-lucide-message-circle"
          >
            WhatsApp Us
          </UButton>
        </div>
      </div>
    </section>

    <!-- Trust Signals Bar -->
    <section class="py-10 px-4 bg-white dark:bg-brand-900 border-b border-slate-100 dark:border-slate-800">
      <div class="max-w-5xl mx-auto grid grid-cols-2 md:grid-cols-4 gap-6 md:gap-8">
        <div v-for="signal in trustSignals" :key="signal.label" class="text-center">
          <div class="w-12 h-12 bg-brand-600/10 dark:bg-brand-400/15 rounded-xl flex items-center justify-center mx-auto mb-3">
            <UIcon :name="signal.icon" class="text-2xl text-brand-600 dark:text-brand-400" />
          </div>
          <p class="font-semibold text-brand-900 dark:text-white text-sm">{{ signal.label }}</p>
          <p class="text-xs text-slate-500 dark:text-slate-400 mt-1 hidden md:block">{{ signal.description }}</p>
        </div>
      </div>
    </section>

    <!-- Featured Projects -->
    <section class="py-16 px-4 bg-slate-50 dark:bg-brand-900/50">
      <div class="max-w-6xl mx-auto">
        <div class="text-center mb-10">
          <h2 class="text-2xl md:text-3xl font-bold text-brand-900 dark:text-white">Curated Investment Projects</h2>
          <p class="text-slate-500 dark:text-slate-400 mt-2 max-w-xl mx-auto">
            Every project is RERA-verified, field-inspected, and priced with full transparency. No hidden costs.
          </p>
        </div>
        <div class="grid sm:grid-cols-2 lg:grid-cols-4 gap-6">
          <NuxtLink
            v-for="project in featuredProjects"
            :key="project.slug"
            :to="`/projects/${project.slug}`"
            class="block group"
          >
            <UCard class="h-full hover:shadow-lg transition-shadow duration-200">
              <div :class="[project.imageColor, 'h-40 rounded-lg flex items-center justify-center mb-4 group-hover:opacity-90 transition-opacity']">
                <UIcon name="i-lucide-building-2" class="text-4xl text-slate-400" />
              </div>
              <div class="flex items-center gap-2 mb-2">
                <UBadge color="success" variant="subtle" size="xs">
                  <UIcon name="i-lucide-shield-check" class="mr-1" />
                  RERA
                </UBadge>
                <UBadge color="info" variant="subtle" size="xs">
                  {{ project.growthPotential }} Growth
                </UBadge>
              </div>
              <h3 class="font-semibold text-brand-900 dark:text-white group-hover:text-brand-600 dark:group-hover:text-brand-400 transition-colors">
                {{ project.name }}
              </h3>
              <p class="text-sm text-slate-500 dark:text-slate-400 flex items-center gap-1 mt-1">
                <UIcon name="i-lucide-map-pin" class="text-xs" />
                {{ project.location }}
              </p>
              <p class="text-lg font-bold text-brand-900 dark:text-white mt-2">{{ project.priceRange }}</p>
              <div class="flex flex-wrap gap-1.5 mt-2">
                <UBadge v-for="config in project.configs" :key="config" color="neutral" variant="subtle" size="xs">
                  {{ config }}
                </UBadge>
              </div>
              <p class="text-xs text-slate-400 mt-3">
                {{ project.builder }} &middot; Possession {{ project.possession }}
              </p>
            </UCard>
          </NuxtLink>
        </div>
        <div class="text-center mt-8">
          <UButton to="/projects" variant="outline" size="lg" icon="i-lucide-arrow-right" trailing>
            View All Projects
          </UButton>
        </div>
      </div>
    </section>

    <!-- Mumbai Micro-Market Highlights -->
    <section class="py-16 px-4 bg-white dark:bg-brand-900">
      <div class="max-w-6xl mx-auto">
        <div class="text-center mb-10">
          <h2 class="text-2xl md:text-3xl font-bold text-brand-900 dark:text-white">Mumbai Micro-Market Highlights</h2>
          <p class="text-slate-500 dark:text-slate-400 mt-2 max-w-xl mx-auto">
            Data-driven insights into Mumbai's fastest-growing corridors, powered by infrastructure development and connectivity upgrades.
          </p>
        </div>
        <div class="grid sm:grid-cols-2 lg:grid-cols-3 gap-6">
          <NuxtLink
            v-for="market in microMarkets"
            :key="market.slug"
            :to="`/areas/${market.slug}`"
            class="block group"
          >
            <UCard class="h-full hover:shadow-lg transition-shadow duration-200 hover:border-brand-500/30">
              <div class="flex items-start justify-between mb-3">
                <h3 class="font-semibold text-brand-900 dark:text-white text-lg group-hover:text-brand-600 dark:group-hover:text-brand-400 transition-colors">
                  {{ market.name }}
                </h3>
                <UBadge color="success" variant="subtle" size="xs">
                  {{ market.growth }}
                </UBadge>
              </div>
              <div class="space-y-2 text-sm">
                <div class="flex justify-between">
                  <span class="text-slate-500 dark:text-slate-400">Avg. Price</span>
                  <span class="font-medium text-slate-900 dark:text-white">{{ market.avgPrice }}</span>
                </div>
                <div class="flex justify-between">
                  <span class="text-slate-500 dark:text-slate-400">Growth Potential</span>
                  <span class="font-medium text-accent-600 dark:text-accent-400">{{ market.growth }}</span>
                </div>
                <div class="flex justify-between">
                  <span class="text-slate-500 dark:text-slate-400">Key Driver</span>
                  <span class="font-medium text-slate-900 dark:text-white">{{ market.driver }}</span>
                </div>
              </div>
              <div class="mt-4 pt-3 border-t border-slate-100 dark:border-slate-800 flex items-center text-sm text-brand-600 dark:text-brand-400 font-medium">
                View Area Guide
                <UIcon name="i-lucide-arrow-right" class="ml-1" />
              </div>
            </UCard>
          </NuxtLink>
        </div>
      </div>
    </section>

    <!-- How It Works -->
    <section class="py-16 px-4 bg-slate-50 dark:bg-brand-900/50">
      <div class="max-w-5xl mx-auto">
        <div class="text-center mb-12">
          <h2 class="text-2xl md:text-3xl font-bold text-brand-900 dark:text-white">How It Works</h2>
          <p class="text-slate-500 dark:text-slate-400 mt-2">A trust-first, no-pressure approach to real estate investment</p>
        </div>
        <div class="grid sm:grid-cols-2 lg:grid-cols-4 gap-8">
          <div v-for="step in howItWorks" :key="step.step" class="text-center">
            <div class="w-14 h-14 bg-brand-800 dark:bg-brand-700 text-white rounded-2xl flex items-center justify-center mx-auto mb-4">
              <UIcon :name="step.icon" class="text-2xl" />
            </div>
            <div class="text-xs font-bold text-brand-600 dark:text-brand-400 mb-1">STEP {{ step.step }}</div>
            <h3 class="font-semibold text-brand-900 dark:text-white mb-2">{{ step.title }}</h3>
            <p class="text-sm text-slate-500 dark:text-slate-400 leading-relaxed">{{ step.description }}</p>
          </div>
        </div>
      </div>
    </section>

    <!-- Cost Transparency Section -->
    <section class="py-16 px-4 bg-white dark:bg-brand-900">
      <div class="max-w-4xl mx-auto">
        <div class="grid md:grid-cols-2 gap-10 items-center">
          <div>
            <h2 class="text-2xl md:text-3xl font-bold text-brand-900 dark:text-white">Know the Real Cost<br>Before You Invest</h2>
            <p class="text-slate-500 dark:text-slate-400 mt-4 leading-relaxed">
              Other platforms show base price only. We show you everything &mdash;
              GST, stamp duty, registration, maintenance deposit, legal charges, and parking.
              Our all-inclusive pricing ensures zero surprises at the time of booking.
            </p>
            <div class="mt-6 space-y-3">
              <div class="flex items-center gap-3">
                <UIcon name="i-lucide-check-circle" class="text-brand-600 dark:text-brand-400 text-lg" />
                <span class="text-sm text-slate-700 dark:text-slate-200">Base price + GST included</span>
              </div>
              <div class="flex items-center gap-3">
                <UIcon name="i-lucide-check-circle" class="text-brand-600 dark:text-brand-400 text-lg" />
                <span class="text-sm text-slate-700 dark:text-slate-200">Stamp duty &amp; registration charges calculated</span>
              </div>
              <div class="flex items-center gap-3">
                <UIcon name="i-lucide-check-circle" class="text-brand-600 dark:text-brand-400 text-lg" />
                <span class="text-sm text-slate-700 dark:text-slate-200">Maintenance deposit &amp; legal fees factored in</span>
              </div>
              <div class="flex items-center gap-3">
                <UIcon name="i-lucide-check-circle" class="text-brand-600 dark:text-brand-400 text-lg" />
                <span class="text-sm text-slate-700 dark:text-slate-200">Parking charges and amenity costs included</span>
              </div>
            </div>
          </div>
          <div class="bg-slate-50 dark:bg-brand-800 rounded-2xl p-6 border border-slate-200 dark:border-slate-700">
            <h4 class="font-semibold text-brand-900 dark:text-white mb-4">Sample: 2 BHK in Panvel</h4>
            <div class="space-y-3 text-sm">
              <div class="flex justify-between">
                <span class="text-slate-500 dark:text-slate-400">Base Price (650 sq ft)</span>
                <span class="font-medium">₹33.80 L</span>
              </div>
              <div class="flex justify-between">
                <span class="text-slate-500 dark:text-slate-400">GST (5%)</span>
                <span class="font-medium">₹1.69 L</span>
              </div>
              <div class="flex justify-between">
                <span class="text-slate-500 dark:text-slate-400">Stamp Duty (6%)</span>
                <span class="font-medium">₹2.03 L</span>
              </div>
              <div class="flex justify-between">
                <span class="text-slate-500 dark:text-slate-400">Registration</span>
                <span class="font-medium">₹0.30 L</span>
              </div>
              <div class="flex justify-between">
                <span class="text-slate-500 dark:text-slate-400">Legal &amp; Documentation</span>
                <span class="font-medium">₹0.25 L</span>
              </div>
              <div class="flex justify-between">
                <span class="text-slate-500 dark:text-slate-400">Maintenance Deposit</span>
                <span class="font-medium">₹0.50 L</span>
              </div>
              <hr class="border-slate-200 dark:border-slate-700" />
              <div class="flex justify-between font-bold text-base text-slate-900 dark:text-white">
                <span>All-Inclusive Cost</span>
                <span>₹38.57 L</span>
              </div>
            </div>
          </div>
        </div>
      </div>
    </section>

    <!-- WhatsApp CTA Banner -->
    <section class="py-16 px-4 bg-brand-800 text-white text-center">
      <div class="max-w-2xl mx-auto">
        <UIcon name="i-lucide-message-circle" class="text-4xl text-accent-400 mb-4" />
        <h2 class="text-2xl md:text-3xl font-bold">Ready to Explore Investment Options?</h2>
        <p class="text-brand-200 mt-3 mb-8 leading-relaxed">
          No spam calls. No pressure. No registration required.
          Just a WhatsApp conversation about your investment goals.
        </p>
        <div class="flex flex-col sm:flex-row gap-4 justify-center">
          <UButton
            href="https://wa.me/919876543210?text=Hi%2C%20I%20want%20to%20explore%20investment%20options%20in%20Mumbai%20micro-markets"
            target="_blank"
            size="xl"
            class="bg-green-600 hover:bg-green-700 text-white"
            icon="i-lucide-message-circle"
          >
            Start on WhatsApp
          </UButton>
          <UButton
            to="/projects"
            size="xl"
            variant="outline"
            class="border-white text-white hover:bg-white/10"
          >
            Browse Projects First
          </UButton>
        </div>
      </div>
    </section>
  </div>
</template>
