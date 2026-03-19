// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  devtools: { enabled: true },

  modules: [
    '@nuxt/ui',
    '@nuxt/eslint',
    '@nuxtjs/i18n'
  ],

  i18n: {
    locales: [
      { code: 'en', name: 'English', file: 'en.json' },
      { code: 'hi', name: 'हिन्दी', file: 'hi.json' },
      { code: 'mr', name: 'मराठी', file: 'mr.json' }
    ],
    defaultLocale: 'en',
    strategy: 'prefix_except_default',
    lazy: true,
    langDir: '../i18n',
    detectBrowserLanguage: false,
    bundle: {
      optimizeTranslationDirective: true
    }
  },

  css: ['~/assets/css/main.css'],

  // SSR enabled for SEO (investor-facing public site)
  ssr: true,

  // Runtime config for API URL
  runtimeConfig: {
    public: {
      apiBaseUrl: process.env.NUXT_PUBLIC_API_BASE_URL || 'http://localhost:8080/api/v1'
    }
  },

  // SEO defaults
  app: {
    head: {
      htmlAttrs: { lang: 'en' },
      charset: 'utf-8',
      viewport: 'width=device-width, initial-scale=1',
      title: 'PropTech — Curated Real Estate Investment in Mumbai',
      meta: [
        { name: 'description', content: 'RERA-verified, curated real estate investment opportunities in Mumbai micro-markets. On-ground due diligence, transparent pricing, trust-first approach.' }
      ],
      link: [
        { rel: 'preconnect', href: 'https://fonts.googleapis.com' },
        { rel: 'preconnect', href: 'https://fonts.gstatic.com', crossorigin: '' },
        { rel: 'stylesheet', href: 'https://fonts.googleapis.com/css2?family=Inter:wght@400;500;600;700&family=Noto+Sans+Devanagari:wght@400;500;600;700&display=swap' }
      ]
    }
  },

  compatibilityDate: '2025-07-16'
})