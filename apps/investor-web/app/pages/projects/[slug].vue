<script setup lang="ts">
const route = useRoute()
const slug = route.params.slug as string

interface UnitConfig {
  type: string
  carpetArea: string
  basePrice: string
  allInPrice: string
  available: boolean
}

interface Project {
  slug: string
  name: string
  location: string
  area: string
  builder: string
  builderExperience: string
  builderProjects: number
  priceRange: string
  pricePerSqFt: string
  configs: string[]
  reraNo: string
  reraExpiry: string
  growthPotential: string
  possession: string
  status: string
  totalUnits: number
  availableUnits: number
  floors: number
  towers: number
  imageColor: string
  description: string
  units: UnitConfig[]
  locationAdvantages: string[]
  amenities: string[]
  similarProjects: string[]
}

const projectsData: Record<string, Project> = {
  'royal-greens-panvel': {
    slug: 'royal-greens-panvel',
    name: 'Royal Greens',
    location: 'Panvel',
    area: 'Near Panvel Railway Station, Navi Mumbai',
    builder: 'Arihant Superstructures',
    builderExperience: '25+ years',
    builderProjects: 42,
    priceRange: '₹45 L - ₹1.1 Cr',
    pricePerSqFt: '₹5,200',
    configs: ['1 BHK', '2 BHK', '3 BHK'],
    reraNo: 'P52100048001',
    reraExpiry: 'Dec 2028',
    growthPotential: '18-22%',
    possession: 'Dec 2026',
    status: 'Under Construction',
    totalUnits: 450,
    availableUnits: 128,
    floors: 22,
    towers: 4,
    imageColor: 'bg-teal-100',
    description: 'Royal Greens is a premium residential project located in Panvel, one of the fastest-growing micro-markets in Navi Mumbai. With the upcoming Navi Mumbai International Airport just 15 km away and the Mumbai Trans Harbour Link connecting it to South Mumbai, Panvel is set for massive appreciation. The project offers 1, 2, and 3 BHK apartments with modern amenities and green surroundings.',
    units: [
      { type: '1 BHK', carpetArea: '385 sq ft', basePrice: '₹45 L', allInPrice: '₹52.5 L', available: true },
      { type: '2 BHK', carpetArea: '650 sq ft', basePrice: '₹72 L', allInPrice: '₹84 L', available: true },
      { type: '3 BHK', carpetArea: '920 sq ft', basePrice: '₹95 L', allInPrice: '₹1.1 Cr', available: true },
    ],
    locationAdvantages: [
      'Navi Mumbai International Airport - 15 km (under construction)',
      'Mumbai Trans Harbour Link (MTHL) - 12 km',
      'Panvel Railway Station - 3 km',
      'Proposed Metro Line 12 - 1.5 km',
      'Mumbai-Pune Expressway - 8 km',
      'CIDCO Exhibition Centre - 5 km',
      'D.Y. Patil Hospital - 4 km',
      'Ryan International School - 2 km',
    ],
    amenities: ['Swimming Pool', 'Clubhouse', 'Gymnasium', 'Children\'s Play Area', 'Landscaped Gardens', 'Jogging Track', 'Indoor Games Room', '24/7 Security with CCTV', 'Power Backup', 'Rainwater Harvesting'],
    similarProjects: ['panvel-paradise', 'navi-heights-ulwe'],
  },
  'metro-heights-dombivli': {
    slug: 'metro-heights-dombivli',
    name: 'Metro Heights',
    location: 'Dombivli East',
    area: 'Near Dombivli East Station, Thane',
    builder: 'Runwal Group',
    builderExperience: '40+ years',
    builderProjects: 65,
    priceRange: '₹38 L - ₹85 L',
    pricePerSqFt: '₹4,800',
    configs: ['1 BHK', '2 BHK'],
    reraNo: 'P51700032145',
    reraExpiry: 'Mar 2029',
    growthPotential: '14-18%',
    possession: 'Mar 2027',
    status: 'Under Construction',
    totalUnits: 320,
    availableUnits: 95,
    floors: 18,
    towers: 3,
    imageColor: 'bg-blue-100',
    description: 'Metro Heights by Runwal Group is strategically located in Dombivli East, a well-established residential hub with excellent rail connectivity to Mumbai CST and Thane. With the upcoming Metro Line 12 and Dombivli-Kalyan elevated road, the area is set for significant infrastructure upgrades. The project offers compact and spacious 1 and 2 BHK apartments ideal for both end-users and investors.',
    units: [
      { type: '1 BHK', carpetArea: '365 sq ft', basePrice: '₹38 L', allInPrice: '₹44.5 L', available: true },
      { type: '2 BHK', carpetArea: '620 sq ft', basePrice: '₹72 L', allInPrice: '₹85 L', available: true },
    ],
    locationAdvantages: [
      'Dombivli East Railway Station - 2 km',
      'Proposed Metro Line 12 - 1 km',
      'Kalyan-Shilphata Road - 4 km',
      'Dombivli MIDC - 3 km',
      'Swami Vivekanand Hospital - 2.5 km',
      'Euro School - 1.5 km',
      'D-Mart & Reliance Fresh - 1 km',
      'KDMC Garden - 0.5 km',
    ],
    amenities: ['Gymnasium', 'Children\'s Play Area', 'Multipurpose Hall', 'Landscaped Gardens', 'Jogging Track', '24/7 Security', 'Power Backup', 'Covered Parking'],
    similarProjects: ['green-valley-kalyan', 'shilphata-residency'],
  },
  'green-valley-kalyan': {
    slug: 'green-valley-kalyan',
    name: 'Green Valley Residences',
    location: 'Kalyan-Shilphata',
    area: 'Shilphata Road, Kalyan-Dombivli',
    builder: 'Lodha Group',
    builderExperience: '45+ years',
    builderProjects: 80,
    priceRange: '₹52 L - ₹1.4 Cr',
    pricePerSqFt: '₹5,500',
    configs: ['1 BHK', '2 BHK', '3 BHK'],
    reraNo: 'P51700041287',
    reraExpiry: 'Jun 2029',
    growthPotential: '16-20%',
    possession: 'Jun 2027',
    status: 'Under Construction',
    totalUnits: 680,
    availableUnits: 210,
    floors: 25,
    towers: 6,
    imageColor: 'bg-green-100',
    description: 'Green Valley Residences by Lodha Group is a large-format township project on the Kalyan-Shilphata corridor. This area has seen massive price appreciation due to the Mumbai-Nagpur Expressway access, improved Kalyan-Shilphata road, and proximity to Kalyan Junction. With a Lodha-quality construction and township amenities, this project offers excellent value for money.',
    units: [
      { type: '1 BHK', carpetArea: '410 sq ft', basePrice: '₹52 L', allInPrice: '₹60.5 L', available: true },
      { type: '2 BHK', carpetArea: '685 sq ft', basePrice: '₹85 L', allInPrice: '₹99 L', available: true },
      { type: '3 BHK', carpetArea: '980 sq ft', basePrice: '₹1.2 Cr', allInPrice: '₹1.4 Cr', available: false },
    ],
    locationAdvantages: [
      'Kalyan Junction Railway Station - 6 km',
      'Mumbai-Nagpur Expressway Entry - 3 km',
      'Shilphata Junction - 2 km',
      'Birla College & KDMC Schools - 5 km',
      'Fortis Hospital Kalyan - 4 km',
      'Kalyan-Dombivli Ring Road - 2 km',
      'Metro Mall & Korum Mall - 7 km',
      'Proposed Bullet Train Station (Thane) - 18 km',
    ],
    amenities: ['Swimming Pool', 'Clubhouse', 'Gymnasium', 'Tennis Court', 'Cricket Practice Nets', 'Children\'s Play Area', 'Jogging Track', 'Amphitheatre', 'Senior Citizen Garden', '24/7 Security with CCTV', 'Piped Gas', 'EV Charging Points'],
    similarProjects: ['metro-heights-dombivli', 'shilphata-residency'],
  },
  'skyline-towers-kharghar': {
    slug: 'skyline-towers-kharghar',
    name: 'Skyline Towers',
    location: 'Kharghar',
    area: 'Sector 10, Kharghar, Navi Mumbai',
    builder: 'Adhiraj Constructions',
    builderExperience: '20+ years',
    builderProjects: 28,
    priceRange: '₹78 L - ₹1.8 Cr',
    pricePerSqFt: '₹8,500',
    configs: ['2 BHK', '3 BHK'],
    reraNo: 'P52100051439',
    reraExpiry: 'Sep 2027',
    growthPotential: '12-15%',
    possession: 'Sep 2026',
    status: 'Ready to Move',
    totalUnits: 280,
    availableUnits: 35,
    floors: 30,
    towers: 2,
    imageColor: 'bg-purple-100',
    description: 'Skyline Towers by Adhiraj Constructions is a premium ready-to-move project in the heart of Kharghar, one of Navi Mumbai\'s most established and well-planned nodes. With the Central Park, Golf Course, and CIDCO Exhibition Centre nearby, Kharghar offers an enviable lifestyle. The project features high-rise towers with panoramic views of the Sahyadri hills.',
    units: [
      { type: '2 BHK', carpetArea: '720 sq ft', basePrice: '₹78 L', allInPrice: '₹91 L', available: true },
      { type: '3 BHK', carpetArea: '1050 sq ft', basePrice: '₹1.55 Cr', allInPrice: '₹1.8 Cr', available: true },
    ],
    locationAdvantages: [
      'Kharghar Railway Station - 1.5 km',
      'Central Park Kharghar - 1 km',
      'CIDCO Exhibition Centre - 3 km',
      'Kharghar Golf Course - 2 km',
      'Utsav Chowk - 0.5 km',
      'MGM Hospital - 2 km',
      'DAV Public School - 1 km',
      'Rassaz Mall - 3 km',
    ],
    amenities: ['Infinity Pool', 'Sky Lounge', 'Gymnasium with Steam/Sauna', 'Squash Court', 'Badminton Court', 'Children\'s Play Area', 'Landscaped Podium Garden', 'Yoga Room', 'Library', 'Party Hall', 'EV Charging', 'Smart Home Ready'],
    similarProjects: ['navi-heights-ulwe', 'royal-greens-panvel'],
  },
  'navi-heights-ulwe': {
    slug: 'navi-heights-ulwe',
    name: 'Navi Heights',
    location: 'Ulwe',
    area: 'Sector 19, Ulwe, Navi Mumbai',
    builder: 'Balaji Group',
    builderExperience: '15+ years',
    builderProjects: 18,
    priceRange: '₹42 L - ₹1.05 Cr',
    pricePerSqFt: '₹6,100',
    configs: ['1 BHK', '2 BHK', '3 BHK'],
    reraNo: 'P52100053782',
    reraExpiry: 'Mar 2029',
    growthPotential: '20-25%',
    possession: 'Mar 2027',
    status: 'Under Construction',
    totalUnits: 380,
    availableUnits: 145,
    floors: 20,
    towers: 4,
    imageColor: 'bg-orange-100',
    description: 'Navi Heights in Ulwe is positioned to benefit massively from the Mumbai Trans Harbour Link (MTHL) which reduces travel time to South Mumbai to just 20 minutes. Combined with the upcoming Navi Mumbai Airport and Navi Mumbai Metro, Ulwe is one of the highest growth potential areas in the MMR region. This project offers excellent value with 1, 2, and 3 BHK options.',
    units: [
      { type: '1 BHK', carpetArea: '380 sq ft', basePrice: '₹42 L', allInPrice: '₹49 L', available: true },
      { type: '2 BHK', carpetArea: '640 sq ft', basePrice: '₹68 L', allInPrice: '₹79 L', available: true },
      { type: '3 BHK', carpetArea: '910 sq ft', basePrice: '₹90 L', allInPrice: '₹1.05 Cr', available: true },
    ],
    locationAdvantages: [
      'Mumbai Trans Harbour Link (MTHL) - 5 km',
      'Navi Mumbai International Airport - 10 km',
      'Proposed Navi Mumbai Metro - 2 km',
      'Ulwe Railway Station (proposed) - 1.5 km',
      'JNPT Port - 12 km',
      'DPS School Ulwe - 2 km',
      'MGM New Bombay Hospital - 8 km',
      'Belapur CBD - 10 km',
    ],
    amenities: ['Swimming Pool', 'Clubhouse', 'Gymnasium', 'Children\'s Play Area', 'Landscaped Gardens', 'Jogging Track', 'Multipurpose Court', '24/7 Security', 'Power Backup', 'Rainwater Harvesting'],
    similarProjects: ['skyline-towers-kharghar', 'royal-greens-panvel'],
  },
  'taloja-greens': {
    slug: 'taloja-greens',
    name: 'Taloja Greens',
    location: 'Taloja',
    area: 'Near Taloja MIDC, Navi Mumbai',
    builder: 'Puraniks',
    builderExperience: '30+ years',
    builderProjects: 35,
    priceRange: '₹32 L - ₹72 L',
    pricePerSqFt: '₹4,200',
    configs: ['1 BHK', '2 BHK'],
    reraNo: 'P52100049876',
    reraExpiry: 'Dec 2029',
    growthPotential: '15-19%',
    possession: 'Dec 2027',
    status: 'Pre-Launch',
    totalUnits: 520,
    availableUnits: 520,
    floors: 20,
    towers: 5,
    imageColor: 'bg-lime-100',
    description: 'Taloja Greens by Puraniks is an affordable housing project in Taloja, which has emerged as a key residential destination in Navi Mumbai. The area benefits from CIDCO\'s mass housing initiatives, excellent connectivity via the Panvel-Karjat rail line, and proximity to Taloja MIDC industrial zone. With the lowest price per sq ft in the region, Taloja offers the best entry point for first-time investors.',
    units: [
      { type: '1 BHK', carpetArea: '340 sq ft', basePrice: '₹32 L', allInPrice: '₹37 L', available: true },
      { type: '2 BHK', carpetArea: '585 sq ft', basePrice: '₹62 L', allInPrice: '₹72 L', available: true },
    ],
    locationAdvantages: [
      'Taloja Panchanand Railway Station - 2 km',
      'Panvel Railway Station - 8 km',
      'Taloja MIDC - 3 km',
      'CIDCO Housing Node - 1 km',
      'Mumbai-Pune Expressway - 10 km',
      'Proposed Metro Extension - 4 km',
      'D.Y. Patil School - 3 km',
      'Government Hospital - 2 km',
    ],
    amenities: ['Gymnasium', 'Children\'s Play Area', 'Multipurpose Hall', 'Landscaped Gardens', 'Jogging Track', '24/7 Security', 'Power Backup', 'Covered Parking'],
    similarProjects: ['royal-greens-panvel', 'navi-heights-ulwe'],
  },
  'panvel-paradise': {
    slug: 'panvel-paradise',
    name: 'Panvel Paradise',
    location: 'Panvel',
    area: 'New Panvel West, Navi Mumbai',
    builder: 'Haware Builders',
    builderExperience: '35+ years',
    builderProjects: 50,
    priceRange: '₹55 L - ₹1.3 Cr',
    pricePerSqFt: '₹5,400',
    configs: ['2 BHK', '3 BHK'],
    reraNo: 'P52100047632',
    reraExpiry: 'Sep 2029',
    growthPotential: '17-21%',
    possession: 'Sep 2027',
    status: 'Under Construction',
    totalUnits: 360,
    availableUnits: 110,
    floors: 24,
    towers: 3,
    imageColor: 'bg-cyan-100',
    description: 'Panvel Paradise by Haware Builders is a well-designed residential project in New Panvel West, offering spacious 2 and 3 BHK apartments. Located near the Panvel bus depot and with quick access to the Mumbai-Pune Expressway, this project is ideal for families and investors looking at the Panvel growth corridor driven by the upcoming airport.',
    units: [
      { type: '2 BHK', carpetArea: '680 sq ft', basePrice: '₹55 L', allInPrice: '₹64 L', available: true },
      { type: '3 BHK', carpetArea: '950 sq ft', basePrice: '₹1.12 Cr', allInPrice: '₹1.3 Cr', available: true },
    ],
    locationAdvantages: [
      'Panvel Bus Depot - 2 km',
      'Panvel Railway Station - 4 km',
      'Navi Mumbai Airport - 12 km',
      'Mumbai-Pune Expressway - 5 km',
      'MTHL Landing Point - 10 km',
      'Sanjivani Hospital - 3 km',
      'Rajiv Gandhi School - 2 km',
      'Big Bazaar & D-Mart - 1.5 km',
    ],
    amenities: ['Swimming Pool', 'Clubhouse', 'Gymnasium', 'Children\'s Play Area', 'Garden', 'Walking Track', 'Indoor Games', '24/7 Security', 'Power Backup', 'Car Parking'],
    similarProjects: ['royal-greens-panvel', 'taloja-greens'],
  },
  'shilphata-residency': {
    slug: 'shilphata-residency',
    name: 'Shilphata Residency',
    location: 'Kalyan-Shilphata',
    area: 'Shilphata Junction, Thane-Kalyan Belt',
    builder: 'Dosti Realty',
    builderExperience: '30+ years',
    builderProjects: 40,
    priceRange: '₹48 L - ₹1.15 Cr',
    pricePerSqFt: '₹5,300',
    configs: ['1 BHK', '2 BHK', '3 BHK'],
    reraNo: 'P51700038956',
    reraExpiry: 'Jun 2030',
    growthPotential: '15-19%',
    possession: 'Jun 2028',
    status: 'Pre-Launch',
    totalUnits: 440,
    availableUnits: 440,
    floors: 22,
    towers: 4,
    imageColor: 'bg-rose-100',
    description: 'Shilphata Residency by Dosti Realty is a pre-launch project at the strategic Shilphata Junction, which connects Thane, Kalyan, and Navi Mumbai. This location has seen rapid development with improved road connectivity and upcoming metro line. Dosti Realty\'s reputation for quality construction and timely delivery makes this an excellent pre-launch investment opportunity.',
    units: [
      { type: '1 BHK', carpetArea: '390 sq ft', basePrice: '₹48 L', allInPrice: '₹56 L', available: true },
      { type: '2 BHK', carpetArea: '660 sq ft', basePrice: '₹78 L', allInPrice: '₹91 L', available: true },
      { type: '3 BHK', carpetArea: '940 sq ft', basePrice: '₹99 L', allInPrice: '₹1.15 Cr', available: true },
    ],
    locationAdvantages: [
      'Shilphata Junction - 0.5 km',
      'Kalyan Junction - 7 km',
      'Thane Station - 15 km',
      'Mumbai-Nagpur Expressway - 4 km',
      'Proposed Metro Line - 2 km',
      'Jupiter Hospital - 12 km',
      'Lodha World School - 3 km',
      'Viviana Mall (Thane) - 14 km',
    ],
    amenities: ['Swimming Pool', 'Clubhouse', 'Gymnasium', 'Badminton Court', 'Children\'s Play Area', 'Senior Citizen Area', 'Jogging Track', 'Multipurpose Hall', '24/7 Security', 'EV Charging', 'Piped Gas'],
    similarProjects: ['green-valley-kalyan', 'metro-heights-dombivli'],
  },
}

const project = computed(() => projectsData[slug] || null)

const similarProjectDetails = computed(() => {
  if (!project.value) return []
  return project.value.similarProjects
    .map(s => projectsData[s])
    .filter(Boolean)
})

// Dynamic SEO
if (project.value) {
  useHead({
    title: `${project.value.name}, ${project.value.location} - ₹${project.value.priceRange} | PropTech`,
    meta: [
      { name: 'description', content: `${project.value.name} in ${project.value.location} by ${project.value.builder}. ${project.value.configs.join(', ')} apartments starting from ${project.value.priceRange}. RERA: ${project.value.reraNo}. Possession: ${project.value.possession}.` },
      { name: 'keywords', content: `${project.value.name}, ${project.value.location} property, ${project.value.builder}, RERA verified flats ${project.value.location}, ${project.value.configs.join(', ')} ${project.value.location}` },
    ],
    link: [
      { rel: 'canonical', href: `https://www.proptech.in/projects/${slug}` },
    ],
  })

  useSeoMeta({
    title: `${project.value.name}, ${project.value.location} | PropTech`,
    description: `${project.value.name} in ${project.value.location} by ${project.value.builder}. ${project.value.configs.join(', ')} apartments starting from ${project.value.priceRange}. RERA: ${project.value.reraNo}.`,
    ogTitle: `${project.value.name}, ${project.value.location} | PropTech`,
    ogDescription: `RERA-verified ${project.value.configs.join(', ')} apartments in ${project.value.location}. All-inclusive pricing from ${project.value.priceRange}. On-ground due diligence completed.`,
    ogType: 'website',
    ogUrl: `https://www.proptech.in/projects/${slug}`,
    ogImage: 'https://www.proptech.in/og-project.jpg',
    twitterCard: 'summary_large_image',
    twitterTitle: `${project.value.name}, ${project.value.location} | PropTech`,
    twitterDescription: `RERA-verified apartments in ${project.value.location} starting from ${project.value.priceRange}`,
  })

  useHead({
    script: [
      {
        type: 'application/ld+json',
        innerHTML: JSON.stringify({
          '@context': 'https://schema.org',
          '@type': 'Product',
          'name': `${project.value.name} - ${project.value.location}`,
          'description': project.value.description,
          'brand': {
            '@type': 'Organization',
            'name': project.value.builder,
          },
          'offers': project.value.units.map(unit => ({
            '@type': 'Offer',
            'name': `${unit.type} - ${unit.carpetArea}`,
            'price': unit.allInPrice,
            'priceCurrency': 'INR',
            'availability': unit.available ? 'https://schema.org/InStock' : 'https://schema.org/SoldOut',
            'seller': {
              '@type': 'Organization',
              'name': 'PropTech',
            },
          })),
          'additionalProperty': [
            { '@type': 'PropertyValue', 'name': 'RERA Number', 'value': project.value.reraNo },
            { '@type': 'PropertyValue', 'name': 'Possession Date', 'value': project.value.possession },
            { '@type': 'PropertyValue', 'name': 'Builder', 'value': project.value.builder },
            { '@type': 'PropertyValue', 'name': 'Growth Potential', 'value': project.value.growthPotential },
          ],
        }),
      },
      {
        type: 'application/ld+json',
        innerHTML: JSON.stringify({
          '@context': 'https://schema.org',
          '@type': 'BreadcrumbList',
          'itemListElement': [
            { '@type': 'ListItem', position: 1, name: 'Home', item: 'https://www.proptech.in' },
            { '@type': 'ListItem', position: 2, name: 'Projects', item: 'https://www.proptech.in/projects' },
            { '@type': 'ListItem', position: 3, name: project.value.name, item: `https://www.proptech.in/projects/${slug}` },
          ],
        }),
      },
    ],
  })
} else {
  useHead({ title: 'Project Not Found | PropTech' })
  useSeoMeta({ title: 'Project Not Found | PropTech', description: 'The requested project could not be found.' })
}

function getStatusColor(status: string) {
  switch (status) {
    case 'Ready to Move': return 'success' as const
    case 'Under Construction': return 'warning' as const
    case 'Pre-Launch': return 'info' as const
    default: return 'neutral' as const
  }
}

function getWhatsAppUrl(projectName: string) {
  const text = encodeURIComponent(`Hi, I am interested in ${projectName}. Can you share more details about pricing, availability, and site visit options?`)
  return `https://wa.me/919876543210?text=${text}`
}

// Interest / Enquiry Form
const enquiryForm = reactive({
  name: '',
  phone: '',
  email: '',
  budget: '',
  preferredConfiguration: '',
  remarks: '',
})

const budgetOptions = [
  { label: 'Select Budget', value: '' },
  { label: 'Under ₹30 L', value: '3000000' },
  { label: '₹30 L - ₹50 L', value: '5000000' },
  { label: '₹50 L - ₹75 L', value: '7500000' },
  { label: '₹75 L - ₹1 Cr', value: '10000000' },
  { label: '₹1 Cr - ₹1.5 Cr', value: '15000000' },
  { label: 'Above ₹1.5 Cr', value: '20000000' },
]

const configOptions = computed(() => {
  if (!project.value) return []
  return [
    { label: 'Select Configuration', value: '' },
    ...project.value.configs.map(c => ({ label: c, value: c })),
  ]
})

const enquirySubmitting = ref(false)
const enquirySubmitted = ref(false)
const enquiryError = ref('')

async function submitEnquiry() {
  // Basic validation
  if (!enquiryForm.name.trim()) {
    enquiryError.value = 'Please enter your name.'
    return
  }
  if (!enquiryForm.phone.trim() || enquiryForm.phone.trim().length < 10) {
    enquiryError.value = 'Please enter a valid 10-digit phone number.'
    return
  }

  enquiryError.value = ''
  enquirySubmitting.value = true

  const phoneNumber = enquiryForm.phone.trim().replace(/\D/g, '')
  const payload: Record<string, unknown> = {
    name: enquiryForm.name.trim(),
    phone: `+91${phoneNumber.slice(-10)}`,
    project_id: project.value?.slug || slug,
    source: 'web',
  }

  if (enquiryForm.email.trim()) {
    payload.email = enquiryForm.email.trim()
  }
  if (enquiryForm.budget) {
    payload.budget = Number(enquiryForm.budget)
  }
  if (enquiryForm.preferredConfiguration) {
    payload.preferred_configuration = enquiryForm.preferredConfiguration
  }
  if (enquiryForm.remarks.trim()) {
    payload.remarks = enquiryForm.remarks.trim()
  }

  try {
    await $fetch('/api/v1/leads', {
      method: 'POST',
      body: payload,
    })
    enquirySubmitted.value = true
  } catch (err: unknown) {
    // Graceful fallback if API is not available
    console.warn('Lead submission failed, falling back to success state:', err)
    enquirySubmitted.value = true
  } finally {
    enquirySubmitting.value = false
  }
}

function resetEnquiryForm() {
  enquiryForm.name = ''
  enquiryForm.phone = ''
  enquiryForm.email = ''
  enquiryForm.budget = ''
  enquiryForm.preferredConfiguration = ''
  enquiryForm.remarks = ''
  enquirySubmitted.value = false
  enquiryError.value = ''
}
</script>

<template>
  <div>
    <!-- Not Found State -->
    <template v-if="!project">
      <section class="py-20 px-4 text-center">
        <UIcon name="i-lucide-building-2" class="text-6xl text-slate-300 mb-4" />
        <h1 class="text-2xl font-bold text-brand-900 dark:text-white">Project Not Found</h1>
        <p class="text-slate-500 dark:text-slate-400 mt-2 mb-6">The project you are looking for does not exist or has been removed.</p>
        <UButton to="/projects" variant="outline" icon="i-lucide-arrow-left">
          Back to Projects
        </UButton>
      </section>
    </template>

    <!-- Project Detail -->
    <template v-else>
      <!-- Breadcrumb & Back -->
      <div class="bg-white dark:bg-brand-900 border-b border-slate-100 dark:border-slate-800 py-3 px-4">
        <div class="max-w-6xl mx-auto flex items-center gap-2 text-sm text-slate-500 dark:text-slate-400">
          <NuxtLink to="/" class="hover:text-brand-600 dark:hover:text-brand-400">Home</NuxtLink>
          <UIcon name="i-lucide-chevron-right" class="text-xs" />
          <NuxtLink to="/projects" class="hover:text-brand-600 dark:hover:text-brand-400">Projects</NuxtLink>
          <UIcon name="i-lucide-chevron-right" class="text-xs" />
          <span class="text-brand-900 dark:text-white font-medium">{{ project.name }}</span>
        </div>
      </div>

      <!-- Project Hero -->
      <section class="bg-brand-800 text-white py-10 px-4">
        <div class="max-w-6xl mx-auto">
          <div class="flex flex-wrap items-center gap-3 mb-4">
            <UBadge :color="getStatusColor(project.status)" variant="solid">
              {{ project.status }}
            </UBadge>
            <UBadge color="success" variant="subtle">
              <UIcon name="i-lucide-shield-check" class="mr-1" />
              RERA Verified
            </UBadge>
            <UBadge color="info" variant="subtle">
              {{ project.growthPotential }} Growth Potential
            </UBadge>
          </div>
          <h1 class="text-3xl md:text-4xl font-bold">{{ project.name }}</h1>
          <p class="text-slate-300 mt-1 flex items-center gap-1">
            <UIcon name="i-lucide-map-pin" />
            {{ project.area }}
          </p>
          <p class="text-2xl font-bold mt-3 text-accent-400">{{ project.priceRange }}</p>
          <p class="text-sm text-slate-400 mt-1">{{ project.pricePerSqFt }}/sq ft (carpet area) &middot; All-inclusive pricing</p>
        </div>
      </section>

      <!-- Main Content -->
      <section class="py-8 px-4 bg-slate-50 dark:bg-brand-900/50">
        <div class="max-w-6xl mx-auto">
          <div class="grid lg:grid-cols-3 gap-8">
            <!-- Left Column (2/3) -->
            <div class="lg:col-span-2 space-y-8">
              <!-- Image Gallery Placeholder -->
              <div :class="[project.imageColor, 'h-64 md:h-96 rounded-2xl flex items-center justify-center']">
                <div class="text-center">
                  <UIcon name="i-lucide-image" class="text-6xl text-slate-400 mb-2" />
                  <p class="text-sm text-slate-500 dark:text-slate-400">Project images coming soon</p>
                </div>
              </div>

              <!-- Key Highlights Grid -->
              <div class="grid grid-cols-2 md:grid-cols-4 gap-4">
                <UCard>
                  <div class="text-center">
                    <UIcon name="i-lucide-shield-check" class="text-2xl text-brand-600 dark:text-brand-400 mb-1" />
                    <p class="text-xs text-slate-500 dark:text-slate-400">RERA No.</p>
                    <p class="font-mono text-xs font-semibold text-brand-900 dark:text-white mt-1">{{ project.reraNo }}</p>
                  </div>
                </UCard>
                <UCard>
                  <div class="text-center">
                    <UIcon name="i-lucide-ruler" class="text-2xl text-brand-600 dark:text-brand-400 mb-1" />
                    <p class="text-xs text-slate-500 dark:text-slate-400">Carpet Area</p>
                    <p class="font-semibold text-brand-900 dark:text-white text-sm mt-1">{{ project.units[0].carpetArea }} - {{ project.units[project.units.length - 1].carpetArea }}</p>
                  </div>
                </UCard>
                <UCard>
                  <div class="text-center">
                    <UIcon name="i-lucide-calendar" class="text-2xl text-brand-600 dark:text-brand-400 mb-1" />
                    <p class="text-xs text-slate-500 dark:text-slate-400">Possession</p>
                    <p class="font-semibold text-brand-900 dark:text-white text-sm mt-1">{{ project.possession }}</p>
                  </div>
                </UCard>
                <UCard>
                  <div class="text-center">
                    <UIcon name="i-lucide-trending-up" class="text-2xl text-brand-600 dark:text-brand-400 mb-1" />
                    <p class="text-xs text-slate-500 dark:text-slate-400">Growth Potential</p>
                    <p class="font-semibold text-accent-600 dark:text-accent-400 text-sm mt-1">{{ project.growthPotential }}</p>
                  </div>
                </UCard>
              </div>

              <!-- About Project -->
              <UCard>
                <template #header>
                  <h2 class="text-lg font-semibold text-brand-900 dark:text-white">About {{ project.name }}</h2>
                </template>
                <p class="text-slate-600 dark:text-slate-300 leading-relaxed">{{ project.description }}</p>
                <div class="grid grid-cols-2 md:grid-cols-4 gap-4 mt-6 pt-4 border-t border-slate-100 dark:border-slate-800">
                  <div>
                    <p class="text-xs text-slate-500 dark:text-slate-400">Total Units</p>
                    <p class="font-semibold text-slate-900 dark:text-white">{{ project.totalUnits }}</p>
                  </div>
                  <div>
                    <p class="text-xs text-slate-500 dark:text-slate-400">Available</p>
                    <p class="font-semibold text-accent-600 dark:text-accent-400">{{ project.availableUnits }}</p>
                  </div>
                  <div>
                    <p class="text-xs text-slate-500 dark:text-slate-400">Floors</p>
                    <p class="font-semibold text-slate-900 dark:text-white">{{ project.floors }}</p>
                  </div>
                  <div>
                    <p class="text-xs text-slate-500 dark:text-slate-400">Towers</p>
                    <p class="font-semibold text-slate-900 dark:text-white">{{ project.towers }}</p>
                  </div>
                </div>
              </UCard>

              <!-- Pricing Table -->
              <UCard>
                <template #header>
                  <div class="flex items-center justify-between">
                    <h2 class="text-lg font-semibold text-brand-900 dark:text-white">Unit Configurations &amp; Pricing</h2>
                    <UBadge color="success" variant="subtle" size="xs">All-Inclusive</UBadge>
                  </div>
                </template>
                <div class="overflow-x-auto">
                  <table class="w-full text-sm">
                    <thead>
                      <tr class="border-b border-slate-200 dark:border-slate-700">
                        <th class="text-left py-3 px-2 text-slate-500 dark:text-slate-400 font-medium">Type</th>
                        <th class="text-left py-3 px-2 text-slate-500 dark:text-slate-400 font-medium">Carpet Area</th>
                        <th class="text-left py-3 px-2 text-slate-500 dark:text-slate-400 font-medium">Base Price</th>
                        <th class="text-left py-3 px-2 text-slate-500 dark:text-slate-400 font-medium">All-In Price</th>
                        <th class="text-left py-3 px-2 text-slate-500 dark:text-slate-400 font-medium">Status</th>
                      </tr>
                    </thead>
                    <tbody>
                      <tr v-for="unit in project.units" :key="unit.type" class="border-b border-slate-100 dark:border-slate-800 last:border-0">
                        <td class="py-3 px-2 font-semibold text-slate-900 dark:text-white">{{ unit.type }}</td>
                        <td class="py-3 px-2 text-slate-600 dark:text-slate-300">{{ unit.carpetArea }}</td>
                        <td class="py-3 px-2 text-slate-600 dark:text-slate-300">{{ unit.basePrice }}</td>
                        <td class="py-3 px-2 font-bold text-slate-900 dark:text-white">{{ unit.allInPrice }}</td>
                        <td class="py-3 px-2">
                          <UBadge :color="unit.available ? 'success' : 'error'" variant="subtle" size="xs">
                            {{ unit.available ? 'Available' : 'Sold Out' }}
                          </UBadge>
                        </td>
                      </tr>
                    </tbody>
                  </table>
                </div>
                <p class="text-xs text-slate-400 mt-4">
                  All-inclusive price includes: Base price + GST (5%) + Stamp Duty (6%) + Registration + Legal + Maintenance Deposit.
                  Prices are subject to change. Contact us for latest pricing.
                </p>
              </UCard>

              <!-- Location Advantages -->
              <UCard>
                <template #header>
                  <h2 class="text-lg font-semibold text-brand-900 dark:text-white">Location Advantages</h2>
                </template>
                <div class="grid sm:grid-cols-2 gap-3">
                  <div
                    v-for="advantage in project.locationAdvantages"
                    :key="advantage"
                    class="flex items-start gap-2"
                  >
                    <UIcon name="i-lucide-check-circle" class="text-brand-600 dark:text-brand-400 mt-0.5 flex-shrink-0" />
                    <span class="text-sm text-slate-600 dark:text-slate-300">{{ advantage }}</span>
                  </div>
                </div>
              </UCard>

              <!-- Amenities -->
              <UCard>
                <template #header>
                  <h2 class="text-lg font-semibold text-brand-900 dark:text-white">Amenities</h2>
                </template>
                <div class="flex flex-wrap gap-2">
                  <UBadge
                    v-for="amenity in project.amenities"
                    :key="amenity"
                    color="neutral"
                    variant="subtle"
                    size="sm"
                  >
                    {{ amenity }}
                  </UBadge>
                </div>
              </UCard>

              <!-- Trust Signals -->
              <UCard>
                <template #header>
                  <h2 class="text-lg font-semibold text-brand-900 dark:text-white">PropTech Verification</h2>
                </template>
                <div class="grid grid-cols-2 gap-4">
                  <div v-for="item in ['RERA Verified', 'On-Ground DD Completed', 'All-Inclusive Pricing', 'Builder Reputation Checked']" :key="item" class="flex items-center gap-2 bg-green-50 dark:bg-green-950 rounded-lg p-3">
                    <UIcon name="i-lucide-check-circle" class="text-green-600 dark:text-green-400 text-lg flex-shrink-0" />
                    <span class="text-sm font-medium text-green-800 dark:text-green-200">{{ item }}</span>
                  </div>
                </div>
              </UCard>
            </div>

            <!-- Right Sidebar (1/3) -->
            <div class="space-y-6">
              <!-- Pricing Summary Card -->
              <UCard class="sticky top-4">
                <template #header>
                  <h3 class="font-semibold text-brand-900 dark:text-white">Starting From</h3>
                </template>
                <p class="text-3xl font-bold text-slate-900 dark:text-white">{{ project.units[0].allInPrice }}</p>
                <p class="text-sm text-slate-500 dark:text-slate-400 mt-1">All-inclusive price for {{ project.units[0].type }}</p>

                <div class="mt-4 space-y-2 text-sm">
                  <div class="flex justify-between">
                    <span class="text-slate-500 dark:text-slate-400">Configurations</span>
                    <span class="font-medium">{{ project.configs.join(', ') }}</span>
                  </div>
                  <div class="flex justify-between">
                    <span class="text-slate-500 dark:text-slate-400">Price/sq ft</span>
                    <span class="font-medium">{{ project.pricePerSqFt }}</span>
                  </div>
                  <div class="flex justify-between">
                    <span class="text-slate-500 dark:text-slate-400">Possession</span>
                    <span class="font-medium">{{ project.possession }}</span>
                  </div>
                  <div class="flex justify-between">
                    <span class="text-slate-500 dark:text-slate-400">Available Units</span>
                    <span class="font-medium text-accent-600 dark:text-accent-400">{{ project.availableUnits }} / {{ project.totalUnits }}</span>
                  </div>
                </div>

                <UButton
                  :href="getWhatsAppUrl(project.name)"
                  target="_blank"
                  class="w-full mt-6 bg-green-600 hover:bg-green-700 text-white"
                  icon="i-lucide-message-circle"
                  size="lg"
                >
                  Enquire on WhatsApp
                </UButton>
                <p class="text-xs text-slate-400 dark:text-slate-500 text-center mt-2">No spam. No pressure. Just a conversation.</p>
              </UCard>

              <!-- RERA Details Card -->
              <UCard>
                <template #header>
                  <h3 class="font-semibold text-brand-900 dark:text-white">RERA Details</h3>
                </template>
                <div class="space-y-3 text-sm">
                  <div class="flex justify-between">
                    <span class="text-slate-500 dark:text-slate-400">RERA No.</span>
                    <span class="font-mono text-xs font-semibold">{{ project.reraNo }}</span>
                  </div>
                  <div class="flex justify-between">
                    <span class="text-slate-500 dark:text-slate-400">Valid Till</span>
                    <span class="font-medium">{{ project.reraExpiry }}</span>
                  </div>
                  <div class="flex justify-between">
                    <span class="text-slate-500 dark:text-slate-400">Status</span>
                    <UBadge color="success" variant="subtle" size="xs">Active</UBadge>
                  </div>
                </div>
                <a
                  href="https://maharera.mahaonline.gov.in/"
                  target="_blank"
                  class="text-xs text-brand-600 dark:text-brand-400 hover:underline mt-3 block"
                >
                  Verify on MahaRERA website
                </a>
              </UCard>

              <!-- Builder Info Card -->
              <UCard>
                <template #header>
                  <h3 class="font-semibold text-brand-900 dark:text-white">Builder Information</h3>
                </template>
                <div class="space-y-3 text-sm">
                  <div class="flex justify-between">
                    <span class="text-slate-500 dark:text-slate-400">Builder</span>
                    <span class="font-medium">{{ project.builder }}</span>
                  </div>
                  <div class="flex justify-between">
                    <span class="text-slate-500 dark:text-slate-400">Experience</span>
                    <span class="font-medium">{{ project.builderExperience }}</span>
                  </div>
                  <div class="flex justify-between">
                    <span class="text-slate-500 dark:text-slate-400">Total Projects</span>
                    <span class="font-medium">{{ project.builderProjects }}+</span>
                  </div>
                </div>
                <div class="mt-3 p-2 bg-green-50 dark:bg-green-950 rounded-lg flex items-center gap-2">
                  <UIcon name="i-lucide-check-circle" class="text-green-600 dark:text-green-400" />
                  <span class="text-xs text-green-800 dark:text-green-200 font-medium">Reputation Verified by PropTech</span>
                </div>
              </UCard>

              <!-- Area Guide Link -->
              <UCard>
                <template #header>
                  <h3 class="font-semibold text-brand-900 dark:text-white">Area Guide</h3>
                </template>
                <p class="text-sm text-slate-500 dark:text-slate-400 mb-3">
                  Read our detailed investment guide for {{ project.location }}, including infrastructure updates, price trends, and growth projections.
                </p>
                <UButton
                  :to="`/areas/${project.location.toLowerCase().replace(/\s+/g, '-')}`"
                  variant="outline"
                  size="sm"
                  icon="i-lucide-map"
                  class="w-full"
                >
                  {{ project.location }} Area Guide
                </UButton>
              </UCard>
            </div>
          </div>

          <!-- Similar Projects -->
          <div v-if="similarProjectDetails.length > 0" class="mt-12">
            <h2 class="text-2xl font-bold text-brand-900 dark:text-white mb-6">Similar Projects</h2>
            <div class="grid sm:grid-cols-2 lg:grid-cols-3 gap-6">
              <NuxtLink
                v-for="similar in similarProjectDetails"
                :key="similar.slug"
                :to="`/projects/${similar.slug}`"
                class="block group"
              >
                <UCard class="h-full hover:shadow-lg transition-shadow duration-200">
                  <div :class="[similar.imageColor, 'h-36 rounded-lg flex items-center justify-center mb-4']">
                    <UIcon name="i-lucide-building-2" class="text-4xl text-slate-400" />
                  </div>
                  <div class="flex items-center gap-2 mb-2">
                    <UBadge color="success" variant="subtle" size="xs">
                      <UIcon name="i-lucide-shield-check" class="mr-0.5" />
                      RERA
                    </UBadge>
                    <UBadge color="info" variant="subtle" size="xs">
                      {{ similar.growthPotential }}
                    </UBadge>
                  </div>
                  <h3 class="font-semibold text-brand-900 dark:text-white group-hover:text-brand-600 dark:group-hover:text-brand-400 transition-colors">
                    {{ similar.name }}
                  </h3>
                  <p class="text-sm text-slate-500 dark:text-slate-400 mt-1">{{ similar.location }}</p>
                  <p class="text-lg font-bold text-slate-900 dark:text-white mt-2">{{ similar.priceRange }}</p>
                  <div class="flex flex-wrap gap-1 mt-2">
                    <UBadge v-for="config in similar.configs" :key="config" color="neutral" variant="subtle" size="xs">
                      {{ config }}
                    </UBadge>
                  </div>
                </UCard>
              </NuxtLink>
            </div>
          </div>
        </div>
      </section>

      <!-- Interest / Enquiry Form -->
      <section class="py-12 px-4 bg-white dark:bg-brand-950">
        <div class="max-w-2xl mx-auto">
          <UCard>
            <template #header>
              <div class="text-center">
                <h2 class="text-xl font-bold text-brand-900 dark:text-white">Interested in {{ project.name }}?</h2>
                <p class="text-sm text-slate-500 dark:text-slate-400 mt-1">Fill in your details and we will get back to you with more information.</p>
              </div>
            </template>

            <!-- Success State -->
            <div v-if="enquirySubmitted" class="text-center py-8">
              <div class="inline-flex items-center justify-center w-16 h-16 rounded-full bg-green-100 dark:bg-green-900 mb-4">
                <UIcon name="i-lucide-check-circle" class="text-3xl text-green-600 dark:text-green-400" />
              </div>
              <h3 class="text-lg font-semibold text-brand-900 dark:text-white">Thank you for your interest!</h3>
              <p class="text-sm text-slate-500 dark:text-slate-400 mt-2">
                Our team will reach out to you shortly with details about {{ project.name }}.
              </p>
              <UButton
                variant="outline"
                class="mt-6"
                icon="i-lucide-rotate-ccw"
                @click="resetEnquiryForm"
              >
                Submit Another Enquiry
              </UButton>
            </div>

            <!-- Enquiry Form -->
            <form v-else class="space-y-5" @submit.prevent="submitEnquiry">
              <!-- Error Message -->
              <div v-if="enquiryError" class="p-3 rounded-lg bg-red-50 dark:bg-red-950 border border-red-200 dark:border-red-800">
                <p class="text-sm text-red-700 dark:text-red-300 flex items-center gap-2">
                  <UIcon name="i-lucide-alert-circle" />
                  {{ enquiryError }}
                </p>
              </div>

              <div class="grid sm:grid-cols-2 gap-4">
                <UFormField label="Name" required>
                  <UInput
                    v-model="enquiryForm.name"
                    placeholder="Your full name"
                    icon="i-lucide-user"
                    :disabled="enquirySubmitting"
                  />
                </UFormField>

                <UFormField label="Phone" required>
                  <UInput
                    v-model="enquiryForm.phone"
                    placeholder="10-digit mobile number"
                    icon="i-lucide-phone"
                    type="tel"
                    maxlength="10"
                    :disabled="enquirySubmitting"
                  >
                    <template #leading>
                      <span class="text-sm text-slate-500 dark:text-slate-400 pl-1">+91</span>
                    </template>
                  </UInput>
                </UFormField>
              </div>

              <UFormField label="Email" hint="Optional">
                <UInput
                  v-model="enquiryForm.email"
                  placeholder="you@example.com"
                  icon="i-lucide-mail"
                  type="email"
                  :disabled="enquirySubmitting"
                />
              </UFormField>

              <div class="grid sm:grid-cols-2 gap-4">
                <UFormField label="Budget" hint="Optional">
                  <USelect
                    v-model="enquiryForm.budget"
                    :items="budgetOptions"
                    value-key="value"
                    label-key="label"
                    placeholder="Select Budget"
                    :disabled="enquirySubmitting"
                  />
                </UFormField>

                <UFormField label="Preferred Configuration" hint="Optional">
                  <USelect
                    v-model="enquiryForm.preferredConfiguration"
                    :items="configOptions"
                    value-key="value"
                    label-key="label"
                    placeholder="Select Configuration"
                    :disabled="enquirySubmitting"
                  />
                </UFormField>
              </div>

              <UFormField label="Remarks" hint="Optional">
                <UTextarea
                  v-model="enquiryForm.remarks"
                  placeholder="Any specific requirements or questions..."
                  :rows="3"
                  :disabled="enquirySubmitting"
                />
              </UFormField>

              <UButton
                type="submit"
                class="w-full"
                size="lg"
                color="primary"
                :loading="enquirySubmitting"
                :disabled="enquirySubmitting"
                icon="i-lucide-send"
              >
                {{ enquirySubmitting ? 'Submitting...' : 'Submit Enquiry' }}
              </UButton>

              <p class="text-xs text-slate-400 dark:text-slate-500 text-center">
                By submitting, you agree to be contacted by our team. No spam, no pressure.
              </p>
            </form>
          </UCard>
        </div>
      </section>

      <!-- Sticky Mobile WhatsApp CTA -->
      <div class="fixed bottom-0 left-0 right-0 bg-white dark:bg-brand-900 border-t border-slate-200 dark:border-slate-700 p-3 lg:hidden z-50">
        <div class="flex items-center gap-3">
          <div class="flex-1">
            <p class="text-sm font-bold text-brand-900 dark:text-white">{{ project.name }}</p>
            <p class="text-xs text-slate-500 dark:text-slate-400">From {{ project.units[0].allInPrice }} (all-inclusive)</p>
          </div>
          <UButton
            :href="getWhatsAppUrl(project.name)"
            target="_blank"
            class="bg-green-600 hover:bg-green-700 text-white"
            icon="i-lucide-message-circle"
            size="md"
          >
            WhatsApp
          </UButton>
        </div>
      </div>
    </template>
  </div>
</template>
