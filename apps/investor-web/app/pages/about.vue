<script setup lang="ts">
useHead({
  title: 'About PropTech - RERA Licensed Real Estate Investment Platform',
  meta: [
    { name: 'description', content: 'PropTech is a MahaRERA-licensed real estate investment platform for Mumbai micro-markets. We offer RERA-verified projects, on-ground due diligence, transparent all-inclusive pricing, and a no-pressure investment experience.' },
    { name: 'keywords', content: 'PropTech about, MahaRERA licensed platform, Mumbai real estate investment, RERA verified, transparent pricing, on-ground due diligence' },
  ],
  link: [
    { rel: 'canonical', href: 'https://www.proptech.in/about' },
  ],
})

useSeoMeta({
  title: 'About PropTech - RERA Licensed Real Estate Investment Platform',
  description: 'PropTech is a MahaRERA-licensed real estate investment platform for Mumbai micro-markets. RERA-verified projects, on-ground due diligence, transparent all-inclusive pricing.',
  ogTitle: 'About PropTech - RERA Licensed Real Estate Investment',
  ogDescription: 'MahaRERA-licensed platform offering curated, RERA-verified real estate investment opportunities in Mumbai micro-markets.',
  ogType: 'website',
  ogUrl: 'https://www.proptech.in/about',
  ogImage: 'https://www.proptech.in/og-about.jpg',
  twitterCard: 'summary_large_image',
  twitterTitle: 'About PropTech - RERA Licensed Real Estate Investment',
  twitterDescription: 'MahaRERA-licensed platform for curated real estate investment in Mumbai micro-markets.',
})

useHead({
  script: [
    {
      type: 'application/ld+json',
      innerHTML: JSON.stringify({
        '@context': 'https://schema.org',
        '@type': 'Organization',
        'name': 'PropTech',
        'alternateName': 'PropTech India',
        'description': 'MahaRERA-licensed real estate investment platform for Mumbai micro-markets. We offer RERA-verified projects with on-ground due diligence and transparent all-inclusive pricing.',
        'url': 'https://www.proptech.in',
        'logo': 'https://www.proptech.in/logo.png',
        'foundingDate': '2024',
        'founders': [
          {
            '@type': 'Person',
            'name': 'PropTech Founding Team',
          },
        ],
        'address': {
          '@type': 'PostalAddress',
          'streetAddress': 'Navi Mumbai',
          'addressLocality': 'Mumbai',
          'addressRegion': 'Maharashtra',
          'postalCode': '400001',
          'addressCountry': 'IN',
        },
        'contactPoint': {
          '@type': 'ContactPoint',
          'contactType': 'customer service',
          'telephone': '+91-9876543210',
          'availableLanguage': ['English', 'Hindi', 'Marathi'],
        },
        'areaServed': {
          '@type': 'City',
          'name': 'Mumbai',
        },
        'sameAs': [
          'https://www.instagram.com/proptech.in',
          'https://www.linkedin.com/company/proptech-in',
        ],
        'knowsAbout': [
          'Real Estate Investment',
          'RERA Compliance',
          'Mumbai Property Market',
          'Navi Mumbai Development',
          'Micro-Market Analysis',
        ],
      }),
    },
    {
      type: 'application/ld+json',
      innerHTML: JSON.stringify({
        '@context': 'https://schema.org',
        '@type': 'RealEstateAgent',
        'name': 'PropTech',
        'description': 'MahaRERA-licensed real estate investment platform',
        'url': 'https://www.proptech.in',
        'areaServed': [
          { '@type': 'City', 'name': 'Panvel' },
          { '@type': 'City', 'name': 'Dombivli' },
          { '@type': 'City', 'name': 'Kalyan' },
          { '@type': 'City', 'name': 'Taloja' },
          { '@type': 'City', 'name': 'Kharghar' },
          { '@type': 'City', 'name': 'Ulwe' },
        ],
      }),
    },
  ],
})

const differentiators = [
  {
    icon: 'i-lucide-shield-check',
    title: 'RERA Verified Only',
    description: 'Every project on our platform is MahaRERA registered and verified. We reject projects that do not meet our compliance standards. We verify RERA registration, builder credentials, title clearance, and environmental approvals before listing any project.',
  },
  {
    icon: 'i-lucide-eye',
    title: 'Full Cost Transparency',
    description: 'We show the all-inclusive cost of every property, including base price, GST, stamp duty, registration, legal fees, maintenance deposit, and parking charges. What you see on our platform is what you pay. No hidden charges, no last-minute surprises.',
  },
  {
    icon: 'i-lucide-map-pin',
    title: 'On-Ground Due Diligence',
    description: 'Our team physically visits every project before listing it. We inspect construction progress, verify builder claims, assess the neighbourhood, check connectivity, and evaluate the surrounding infrastructure. Our DD reports give you ground-level truth.',
  },
]

const howItWorks = [
  {
    step: 1,
    title: 'Curate & Verify',
    description: 'We shortlist projects through MahaRERA verification, builder background checks, title clearance, financial health assessment, and on-site inspections. Only 1 in 10 projects make the cut. We reject projects with unclear titles, delayed builders, or non-compliant documentation.',
    icon: 'i-lucide-shield-check',
  },
  {
    step: 2,
    title: 'Browse & Compare',
    description: 'Explore our curated projects with all-inclusive pricing, detailed area guides, infrastructure timelines, and growth projections. Compare projects across locations, configurations, and price ranges. Every data point is verified and updated monthly.',
    icon: 'i-lucide-search',
  },
  {
    step: 3,
    title: 'Guide Your Visit',
    description: 'Schedule a site visit at your convenience through WhatsApp. Our trained agents accompany you, explain everything clearly, show you the actual site (not just the model flat), and answer your questions honestly. No pressure, no hard-selling.',
    icon: 'i-lucide-map-pin',
  },
  {
    step: 4,
    title: 'Invest with Confidence',
    description: 'Complete your purchase with full documentation support. We assist with RERA agreement review, stamp duty payment, registration process, loan assistance coordination, and post-sale registration. Our relationship does not end at the sale.',
    icon: 'i-lucide-check-circle',
  },
]

const stats = [
  { value: 'MahaRERA', label: 'Licensed Platform' },
  { value: '8+', label: 'Curated Projects' },
  { value: '6', label: 'Micro-Markets Covered' },
  { value: '100%', label: 'RERA Verified' },
]

const values = [
  { icon: 'i-lucide-heart', title: 'Trust First', description: 'We believe real estate transactions should be built on trust, not pressure. Every interaction with us is designed to earn your confidence.' },
  { icon: 'i-lucide-bar-chart-3', title: 'Data Driven', description: 'Our area guides, price trends, and growth projections are based on real data from MahaRERA, infrastructure authorities, and on-ground research.' },
  { icon: 'i-lucide-users', title: 'Investor Centric', description: 'We are not a builder platform. We work for the investor. Our recommendations are based on what is best for your investment goals, not builder commissions.' },
  { icon: 'i-lucide-lock', title: 'No Spam Promise', description: 'We will never make unsolicited calls. All communication happens on WhatsApp, at your pace. You control the conversation, not us.' },
]
</script>

<template>
  <div>
    <!-- Hero -->
    <section class="bg-brand-800 text-white py-16 md:py-20 px-4">
      <div class="max-w-4xl mx-auto text-center">
        <UBadge color="neutral" variant="subtle" class="mb-6 text-sm px-4 py-1.5">
          MahaRERA Licensed Platform
        </UBadge>
        <h1 class="text-3xl md:text-4xl lg:text-5xl font-bold leading-tight">
          Built on Trust,
          <span class="text-accent-400">Backed by Data</span>
        </h1>
        <p class="mt-6 text-base sm:text-lg text-slate-300 max-w-2xl mx-auto leading-relaxed">
          PropTech is a MahaRERA-licensed investment platform that brings transparency,
          curation, and trust to Mumbai's real estate market. We help investors make informed
          decisions through RERA-verified projects, on-ground due diligence, and all-inclusive pricing.
        </p>
      </div>
    </section>

    <!-- Stats Bar -->
    <section class="py-10 px-4 bg-white dark:bg-brand-900 border-b border-slate-100 dark:border-slate-800">
      <div class="max-w-4xl mx-auto grid grid-cols-2 md:grid-cols-4 gap-8 text-center">
        <div v-for="stat in stats" :key="stat.label">
          <p class="text-3xl font-bold text-brand-900 dark:text-white">{{ stat.value }}</p>
          <p class="text-sm text-slate-500 dark:text-slate-400 mt-1">{{ stat.label }}</p>
        </div>
      </div>
    </section>

    <!-- Our Story -->
    <section class="py-16 px-4 bg-white dark:bg-brand-900">
      <div class="max-w-4xl mx-auto">
        <div class="grid md:grid-cols-2 gap-10 items-center">
          <div>
            <h2 class="text-2xl md:text-3xl font-bold text-brand-900 dark:text-white">Why We Built PropTech</h2>
            <div class="mt-4 space-y-4 text-slate-600 dark:text-slate-300 leading-relaxed">
              <p>
                The Mumbai real estate market is massive, but it is broken. Buyers face aggressive
                sales calls, hidden charges, inflated promises, and projects that look nothing like the
                brochure. RERA was meant to fix this, but most platforms still operate the old way.
              </p>
              <p>
                We built PropTech because we believe investors deserve better. Better information,
                better transparency, and a better experience. Every project on our platform is RERA-verified,
                physically inspected, and priced with complete cost transparency.
              </p>
              <p>
                We focus exclusively on Mumbai's fastest-growing micro-markets &mdash; Panvel, Dombivli,
                Kalyan, Taloja, Kharghar, and Ulwe. These areas are backed by massive infrastructure
                investments and offer the best risk-adjusted returns in the MMR.
              </p>
            </div>
          </div>
          <div class="bg-slate-50 dark:bg-brand-800 rounded-2xl p-8 border border-slate-200 dark:border-slate-700">
            <h3 class="font-semibold text-brand-900 dark:text-white text-lg mb-4">Our Mission</h3>
            <p class="text-slate-600 dark:text-slate-300 leading-relaxed">
              To make real estate investment in Mumbai accessible, transparent, and trust-based.
              We want every investor &mdash; whether first-time buyer or experienced NRI &mdash; to
              have the information and confidence they need to make the right decision.
            </p>
            <div class="mt-6 pt-4 border-t border-slate-200 dark:border-slate-700">
              <h4 class="font-semibold text-brand-900 dark:text-white text-sm mb-3">We Serve</h4>
              <div class="flex flex-wrap gap-2">
                <UBadge color="neutral" variant="subtle" size="sm">First-Time Investors</UBadge>
                <UBadge color="neutral" variant="subtle" size="sm">NRI Investors</UBadge>
                <UBadge color="neutral" variant="subtle" size="sm">Portfolio Diversifiers</UBadge>
                <UBadge color="neutral" variant="subtle" size="sm">End-Use Buyers</UBadge>
                <UBadge color="neutral" variant="subtle" size="sm">Salaried Professionals</UBadge>
              </div>
            </div>
          </div>
        </div>
      </div>
    </section>

    <!-- 3 Differentiators -->
    <section class="py-16 px-4 bg-slate-50 dark:bg-brand-900/50">
      <div class="max-w-5xl mx-auto">
        <div class="text-center mb-12">
          <h2 class="text-2xl md:text-3xl font-bold text-brand-900 dark:text-white">How We Are Different</h2>
          <p class="text-slate-500 dark:text-slate-400 mt-2">Three pillars that define the PropTech experience</p>
        </div>
        <div class="grid md:grid-cols-3 gap-8">
          <UCard v-for="diff in differentiators" :key="diff.title" class="text-center">
            <div class="w-14 h-14 bg-brand-600/10 dark:bg-brand-400/15 rounded-2xl flex items-center justify-center mx-auto mb-4">
              <UIcon :name="diff.icon" class="text-3xl text-brand-600 dark:text-brand-400" />
            </div>
            <h3 class="font-bold text-brand-900 dark:text-white text-lg mb-3">{{ diff.title }}</h3>
            <p class="text-sm text-slate-500 dark:text-slate-400 leading-relaxed">{{ diff.description }}</p>
          </UCard>
        </div>
      </div>
    </section>

    <!-- How It Works -->
    <section class="py-16 px-4 bg-white dark:bg-brand-900">
      <div class="max-w-4xl mx-auto">
        <div class="text-center mb-12">
          <h2 class="text-2xl md:text-3xl font-bold text-brand-900 dark:text-white">How It Works</h2>
          <p class="text-slate-500 dark:text-slate-400 mt-2">A simple, trust-first process from discovery to investment</p>
        </div>
        <div class="space-y-8">
          <div v-for="step in howItWorks" :key="step.step" class="flex gap-6">
            <div class="flex flex-col items-center">
              <div class="w-12 h-12 bg-brand-800 dark:bg-brand-700 text-white rounded-2xl flex items-center justify-center shrink-0">
                <UIcon :name="step.icon" class="text-xl" />
              </div>
              <div v-if="step.step < 4" class="w-0.5 h-full bg-slate-200 dark:bg-slate-700 mt-2" />
            </div>
            <div class="pb-8">
              <div class="text-xs font-bold text-brand-600 dark:text-brand-400 mb-1">STEP {{ step.step }}</div>
              <h3 class="text-lg font-bold text-brand-900 dark:text-white mb-2">{{ step.title }}</h3>
              <p class="text-sm text-slate-500 dark:text-slate-400 leading-relaxed">{{ step.description }}</p>
            </div>
          </div>
        </div>
      </div>
    </section>

    <!-- Our Values -->
    <section class="py-16 px-4 bg-slate-50 dark:bg-brand-900/50">
      <div class="max-w-5xl mx-auto">
        <div class="text-center mb-12">
          <h2 class="text-2xl md:text-3xl font-bold text-brand-900 dark:text-white">What We Stand For</h2>
          <p class="text-slate-500 dark:text-slate-400 mt-2">The principles that guide every decision we make</p>
        </div>
        <div class="grid sm:grid-cols-2 gap-6">
          <UCard v-for="value in values" :key="value.title">
            <div class="flex gap-4">
              <div class="w-10 h-10 bg-brand-800 dark:bg-brand-700 text-white rounded-xl flex items-center justify-center shrink-0">
                <UIcon :name="value.icon" class="text-lg" />
              </div>
              <div>
                <h3 class="font-bold text-brand-900 dark:text-white mb-1">{{ value.title }}</h3>
                <p class="text-sm text-slate-500 dark:text-slate-400 leading-relaxed">{{ value.description }}</p>
              </div>
            </div>
          </UCard>
        </div>
      </div>
    </section>

    <!-- Trust Section -->
    <section class="py-16 px-4 bg-white dark:bg-brand-900">
      <div class="max-w-4xl mx-auto">
        <div class="text-center mb-10">
          <h2 class="text-2xl md:text-3xl font-bold text-brand-900 dark:text-white">Trust Signals</h2>
          <p class="text-slate-500 dark:text-slate-400 mt-2">Verifiable commitments we make to every investor</p>
        </div>
        <div class="grid sm:grid-cols-2 gap-4">
          <div
            v-for="item in [
              { icon: 'i-lucide-shield-check', text: 'MahaRERA Licensed Platform', detail: 'Registered and compliant with MahaRERA regulations' },
              { icon: 'i-lucide-check-circle', text: '100% RERA Verified Projects', detail: 'Every project is verified on the MahaRERA portal' },
              { icon: 'i-lucide-clipboard-check', text: 'On-Ground DD Reports', detail: 'Physical site visits and independent assessment' },
              { icon: 'i-lucide-indian-rupee', text: 'All-Inclusive Pricing', detail: 'GST, stamp duty, registration, legal — everything included' },
              { icon: 'i-lucide-building-2', text: 'Builder Reputation Verified', detail: 'Track record, delivery history, and financial health checked' },
              { icon: 'i-lucide-phone-off', text: 'Zero Spam Calls', detail: 'All communication on WhatsApp, at your pace' },
            ]"
            :key="item.text"
            class="flex items-start gap-3 bg-green-50 dark:bg-green-950 rounded-xl p-4"
          >
            <UIcon :name="item.icon" class="text-green-600 dark:text-green-400 text-xl mt-0.5 flex-shrink-0" />
            <div>
              <p class="font-semibold text-green-800 dark:text-green-200 text-sm">{{ item.text }}</p>
              <p class="text-xs text-green-600 dark:text-green-400 mt-0.5">{{ item.detail }}</p>
            </div>
          </div>
        </div>
      </div>
    </section>

    <!-- Coverage Areas -->
    <section class="py-16 px-4 bg-slate-50 dark:bg-brand-900/50">
      <div class="max-w-5xl mx-auto">
        <div class="text-center mb-10">
          <h2 class="text-2xl md:text-3xl font-bold text-brand-900 dark:text-white">Areas We Cover</h2>
          <p class="text-slate-500 dark:text-slate-400 mt-2">Six high-growth micro-markets in the Mumbai Metropolitan Region</p>
        </div>
        <div class="grid sm:grid-cols-2 lg:grid-cols-3 gap-4">
          <NuxtLink
            v-for="area in [
              { name: 'Panvel', slug: 'panvel', growth: '18-22%', driver: 'Navi Mumbai Airport' },
              { name: 'Dombivli East', slug: 'dombivli-east', growth: '14-18%', driver: 'Metro Line 12' },
              { name: 'Kalyan-Shilphata', slug: 'kalyan-shilphata', growth: '16-20%', driver: 'Mumbai-Nagpur Expressway' },
              { name: 'Taloja', slug: 'taloja', growth: '15-19%', driver: 'CIDCO Mass Housing' },
              { name: 'Kharghar', slug: 'kharghar', growth: '12-15%', driver: 'Corporate Hub Growth' },
              { name: 'Ulwe', slug: 'ulwe', growth: '20-25%', driver: 'Trans Harbour Link' },
            ]"
            :key="area.slug"
            :to="`/areas/${area.slug}`"
            class="block group"
          >
            <div class="bg-white dark:bg-brand-900 rounded-xl p-4 border border-slate-200 dark:border-slate-700 hover:border-brand-500/30 hover:shadow-md transition-all">
              <div class="flex items-center justify-between">
                <h3 class="font-semibold text-brand-900 dark:text-white group-hover:text-brand-600 dark:group-hover:text-brand-400 transition-colors">{{ area.name }}</h3>
                <UBadge color="success" variant="subtle" size="xs">{{ area.growth }}</UBadge>
              </div>
              <p class="text-xs text-slate-500 dark:text-slate-400 mt-1">{{ area.driver }}</p>
            </div>
          </NuxtLink>
        </div>
        <div class="text-center mt-6">
          <UButton to="/areas" variant="outline" size="lg" icon="i-lucide-map" trailing>
            Explore All Area Guides
          </UButton>
        </div>
      </div>
    </section>

    <!-- CTA -->
    <section class="py-16 px-4 bg-brand-800 text-white text-center">
      <div class="max-w-2xl mx-auto">
        <UIcon name="i-lucide-message-circle" class="text-4xl text-accent-400 mb-4" />
        <h2 class="text-2xl md:text-3xl font-bold">Start Your Investment Journey</h2>
        <p class="text-slate-300 mt-3 mb-8 leading-relaxed">
          Talk to us on WhatsApp. No registration needed. No spam calls. No pressure.
          Just an honest conversation about your investment goals and how we can help.
        </p>
        <div class="flex flex-col sm:flex-row gap-4 justify-center">
          <UButton
            href="https://wa.me/919876543210?text=Hi%2C%20I%20want%20to%20explore%20real%20estate%20investment%20options%20in%20Mumbai%20micro-markets"
            target="_blank"
            size="xl"
            class="bg-green-600 hover:bg-green-700 text-white"
            icon="i-lucide-message-circle"
          >
            Chat with Us on WhatsApp
          </UButton>
          <UButton
            to="/projects"
            size="xl"
            variant="outline"
            class="border-white text-white hover:bg-white/10"
          >
            Browse Projects
          </UButton>
        </div>
      </div>
    </section>
  </div>
</template>
