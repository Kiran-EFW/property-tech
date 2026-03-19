<script setup lang="ts">
const route = useRoute()
const slug = route.params.slug as string

interface InfraProject {
  name: string
  status: 'Completed' | 'Under Construction' | 'Planned'
  expectedCompletion: string
  impact: string
}

interface PriceTrend {
  year: string
  price: string
  change: string
}

interface AreaProject {
  slug: string
  name: string
  builder: string
  priceRange: string
  configs: string[]
  status: string
  imageColor: string
}

interface AreaData {
  name: string
  slug: string
  tagline: string
  description: string
  investmentThesis: string
  avgPriceSqFt: string
  yoyGrowth: string
  growthPotential: string
  growthLevel: 'High' | 'Medium' | 'Stable'
  population: string
  pincode: string
  district: string
  nearestStation: string
  distanceFromMumbai: string
  infraProjects: InfraProject[]
  priceTrends: PriceTrend[]
  projects: AreaProject[]
  pros: string[]
  cons: string[]
  idealFor: string[]
}

const areasData: Record<string, AreaData> = {
  'panvel': {
    name: 'Panvel',
    slug: 'panvel',
    tagline: 'Navi Mumbai Airport Zone',
    description: 'Panvel is the growth epicenter of Navi Mumbai, positioned to become one of the most important nodes in the Mumbai Metropolitan Region. The upcoming Navi Mumbai International Airport (NMIA), just 15 km away, is the single biggest catalyst for real estate appreciation in this area. Combined with the Mumbai Trans Harbour Link (MTHL) providing direct connectivity to South Mumbai, Metro Line 12 connecting to Kalyan, and the Mumbai-Pune Expressway access, Panvel is witnessing unprecedented infrastructure development that is driving sustained price appreciation.',
    investmentThesis: 'Panvel represents the best risk-adjusted return opportunity in the MMR. The convergence of airport proximity, MTHL connectivity, metro access, and expressway links creates a multi-catalyst growth story. Current prices at ₹5,200/sq ft are significantly undervalued compared to established Navi Mumbai nodes like Kharghar (₹8,500/sq ft) and Vashi (₹15,000+/sq ft). With the airport expected to be operational by 2027, we project 18-22% appreciation over the next 3-5 years, making this an ideal time to enter.',
    avgPriceSqFt: '₹5,200/sq ft',
    yoyGrowth: '+12.4%',
    growthPotential: '18-22%',
    growthLevel: 'High',
    population: '4.5 Lakh+',
    pincode: '410206',
    district: 'Raigad',
    nearestStation: 'Panvel Junction',
    distanceFromMumbai: '45 km from CST',
    infraProjects: [
      { name: 'Navi Mumbai International Airport (NMIA)', status: 'Under Construction', expectedCompletion: '2027', impact: 'Game-changer for Panvel real estate. Will put Panvel on the global map and drive commercial/residential demand.' },
      { name: 'Mumbai Trans Harbour Link (MTHL)', status: 'Completed', expectedCompletion: '2024', impact: 'Reduces travel time to South Mumbai from 2+ hours to 20 minutes. Already boosting demand.' },
      { name: 'Metro Line 12 (Kalyan-Taloja-Panvel)', status: 'Planned', expectedCompletion: '2028', impact: 'Will connect Panvel to Kalyan and Dombivli, opening up commuter corridors.' },
      { name: 'Mumbai-Pune Expressway Widening', status: 'Under Construction', expectedCompletion: '2026', impact: 'Improved connectivity to Pune and reduced congestion on existing expressway.' },
      { name: 'CIDCO Exhibition Centre', status: 'Completed', expectedCompletion: '2023', impact: 'Major commercial anchor driving foot traffic and employment.' },
    ],
    priceTrends: [
      { year: '2020', price: '₹3,400', change: '-' },
      { year: '2021', price: '₹3,700', change: '+8.8%' },
      { year: '2022', price: '₹4,100', change: '+10.8%' },
      { year: '2023', price: '₹4,500', change: '+9.8%' },
      { year: '2024', price: '₹4,800', change: '+6.7%' },
      { year: '2025', price: '₹5,200', change: '+12.4%' },
    ],
    projects: [
      { slug: 'royal-greens-panvel', name: 'Royal Greens', builder: 'Arihant Superstructures', priceRange: '₹45 L - ₹1.1 Cr', configs: ['1 BHK', '2 BHK', '3 BHK'], status: 'Under Construction', imageColor: 'bg-teal-100' },
      { slug: 'panvel-paradise', name: 'Panvel Paradise', builder: 'Haware Builders', priceRange: '₹55 L - ₹1.3 Cr', configs: ['2 BHK', '3 BHK'], status: 'Under Construction', imageColor: 'bg-cyan-100' },
    ],
    pros: ['Closest area to upcoming NMIA airport', 'MTHL connectivity to South Mumbai (20 min)', 'Established railway junction with Central & Harbour line', 'Competitive pricing compared to established nodes', 'Multiple infrastructure catalysts converging'],
    cons: ['Airport completion timelines have been delayed before', 'Some areas still lack social infrastructure', 'Monsoon waterlogging in low-lying areas'],
    idealFor: ['Long-term investors (3-5 year horizon)', 'NRIs looking for airport-adjacent properties', 'First-time investors with ₹40-60 L budget'],
  },
  'dombivli-east': {
    name: 'Dombivli East',
    slug: 'dombivli-east',
    tagline: 'Affordable Growth Corridor',
    description: 'Dombivli East is one of the most well-connected residential hubs in the Mumbai Metropolitan Region, with the Central Railway line providing direct connectivity to Mumbai CST, Thane, and Kalyan. The area has a mature social infrastructure with established schools, hospitals, and shopping destinations. With the upcoming Metro Line 12 and Dombivli-Kalyan elevated road, the area is poised for the next phase of growth while maintaining its affordability advantage.',
    investmentThesis: 'Dombivli East offers the best combination of established infrastructure and affordability in the Thane-Kalyan belt. Current prices at ₹4,800/sq ft make it accessible for first-time investors while the upcoming Metro Line 12 and road upgrades will drive appreciation. The area already has 50+ years of established residential history, reducing development risk. We project 14-18% growth over 3-5 years as metro and road infrastructure catalysts materialize.',
    avgPriceSqFt: '₹4,800/sq ft',
    yoyGrowth: '+9.8%',
    growthPotential: '14-18%',
    growthLevel: 'Medium',
    population: '12 Lakh+',
    pincode: '421201',
    district: 'Thane',
    nearestStation: 'Dombivli',
    distanceFromMumbai: '55 km from CST',
    infraProjects: [
      { name: 'Metro Line 12 (Kalyan-Dombivli-Taloja)', status: 'Planned', expectedCompletion: '2028', impact: 'Will connect Dombivli to Taloja and Panvel, creating a new commuter corridor.' },
      { name: 'Dombivli-Kalyan Elevated Road', status: 'Under Construction', expectedCompletion: '2026', impact: 'Will reduce Kalyan travel time to 10 minutes and ease road congestion.' },
      { name: 'Dombivli MIDC Expansion', status: 'Under Construction', expectedCompletion: '2026', impact: 'New industrial zones driving employment and housing demand.' },
      { name: 'KDMC Smart City Projects', status: 'Under Construction', expectedCompletion: '2027', impact: 'Improved civic infrastructure, smart water management, and green spaces.' },
    ],
    priceTrends: [
      { year: '2020', price: '₹3,200', change: '-' },
      { year: '2021', price: '₹3,500', change: '+9.4%' },
      { year: '2022', price: '₹3,800', change: '+8.6%' },
      { year: '2023', price: '₹4,100', change: '+7.9%' },
      { year: '2024', price: '₹4,400', change: '+7.3%' },
      { year: '2025', price: '₹4,800', change: '+9.8%' },
    ],
    projects: [
      { slug: 'metro-heights-dombivli', name: 'Metro Heights', builder: 'Runwal Group', priceRange: '₹38 L - ₹85 L', configs: ['1 BHK', '2 BHK'], status: 'Under Construction', imageColor: 'bg-blue-100' },
    ],
    pros: ['Excellent Central Railway connectivity to CST', 'Mature social infrastructure (schools, hospitals, markets)', 'Most affordable pricing in established areas', 'Large residential population ensuring strong rental demand', 'KDMC Smart City upgrades underway'],
    cons: ['Railway overcrowding during peak hours', 'Limited premium/luxury housing options', 'Some older areas have narrow roads'],
    idealFor: ['Budget-conscious investors (₹35-50 L)', 'Rental yield focused investors', 'End-users seeking affordable homes with good connectivity'],
  },
  'kalyan-shilphata': {
    name: 'Kalyan-Shilphata',
    slug: 'kalyan-shilphata',
    tagline: 'Metro & Expressway Corridor',
    description: 'The Kalyan-Shilphata corridor has emerged as one of the most dynamic growth corridors in the MMR. Located at the junction of Thane, Kalyan, and Navi Mumbai, this area benefits from the Mumbai-Nagpur Expressway (Samruddhi Mahamarg) access, improved Kalyan-Shilphata road infrastructure, and proximity to Kalyan Junction - one of the busiest railway stations in India. Large township projects by premium developers like Lodha and Dosti are transforming this corridor.',
    investmentThesis: 'Kalyan-Shilphata represents a strategic convergence point where the Mumbai-Nagpur Expressway meets the Thane-Kalyan-Navi Mumbai triangle. This positioning creates multi-directional connectivity that is rare in the MMR. Premium developers entering the market signal confidence in the area\'s trajectory. At ₹5,500/sq ft, the area offers value compared to Thane (₹12,000+) while benefiting from similar connectivity. We project 16-20% appreciation as expressway and metro catalysts drive demand.',
    avgPriceSqFt: '₹5,500/sq ft',
    yoyGrowth: '+11.2%',
    growthPotential: '16-20%',
    growthLevel: 'High',
    population: '6 Lakh+',
    pincode: '421306',
    district: 'Thane',
    nearestStation: 'Kalyan Junction',
    distanceFromMumbai: '60 km from CST',
    infraProjects: [
      { name: 'Mumbai-Nagpur Expressway (Samruddhi Mahamarg)', status: 'Completed', expectedCompletion: '2024', impact: 'Direct access to Nagpur in 8 hours. Shilphata is a key access point.' },
      { name: 'Kalyan-Shilphata Road Widening', status: 'Completed', expectedCompletion: '2023', impact: 'Reduced travel time between Kalyan and Shilphata to 15 minutes.' },
      { name: 'Metro Line Extension (Thane-Kalyan)', status: 'Planned', expectedCompletion: '2029', impact: 'Will connect Shilphata to Thane and Mumbai metro network.' },
      { name: 'Kalyan Ring Road', status: 'Under Construction', expectedCompletion: '2027', impact: 'Will provide seamless connectivity around Kalyan, reducing congestion.' },
    ],
    priceTrends: [
      { year: '2020', price: '₹3,600', change: '-' },
      { year: '2021', price: '₹3,900', change: '+8.3%' },
      { year: '2022', price: '₹4,300', change: '+10.3%' },
      { year: '2023', price: '₹4,700', change: '+9.3%' },
      { year: '2024', price: '₹5,100', change: '+8.5%' },
      { year: '2025', price: '₹5,500', change: '+11.2%' },
    ],
    projects: [
      { slug: 'green-valley-kalyan', name: 'Green Valley Residences', builder: 'Lodha Group', priceRange: '₹52 L - ₹1.4 Cr', configs: ['1 BHK', '2 BHK', '3 BHK'], status: 'Under Construction', imageColor: 'bg-green-100' },
      { slug: 'shilphata-residency', name: 'Shilphata Residency', builder: 'Dosti Realty', priceRange: '₹48 L - ₹1.15 Cr', configs: ['1 BHK', '2 BHK', '3 BHK'], status: 'Pre-Launch', imageColor: 'bg-rose-100' },
    ],
    pros: ['Mumbai-Nagpur Expressway access point', 'Premium developers (Lodha, Dosti) entering market', 'Strategic tri-junction of Thane-Kalyan-Navi Mumbai', 'Kalyan Junction provides excellent rail connectivity', 'Significant price advantage over Thane'],
    cons: ['Some stretches still have traffic congestion', 'Area is still developing social infrastructure', 'Metro connectivity is still years away'],
    idealFor: ['Medium-term investors (3-5 year horizon)', 'Township buyers who want community living', 'Investors comparing Thane alternatives at lower price points'],
  },
  'taloja': {
    name: 'Taloja',
    slug: 'taloja',
    tagline: 'Affordable Entry Point',
    description: 'Taloja is Navi Mumbai\'s most affordable residential destination, planned and developed by CIDCO. The area benefits from its proximity to the Panvel growth corridor and the Taloja MIDC industrial zone, which provides a steady base of employment. With the Panvel-Karjat railway line providing connectivity and CIDCO\'s mass housing initiatives driving planned development, Taloja is emerging as the go-to destination for first-time homebuyers and budget-conscious investors.',
    investmentThesis: 'Taloja offers the lowest entry point in the entire Navi Mumbai region at ₹4,200/sq ft. For investors with limited capital, this is the most accessible way to participate in the Navi Mumbai growth story. CIDCO\'s planned development reduces execution risk, and the area benefits from spillover demand from Panvel and Kharghar. The proposed metro extension and proximity to the upcoming airport corridor will drive sustained appreciation. We project 15-19% growth as infrastructure catalysts mature.',
    avgPriceSqFt: '₹4,200/sq ft',
    yoyGrowth: '+10.5%',
    growthPotential: '15-19%',
    growthLevel: 'Medium',
    population: '3 Lakh+',
    pincode: '410208',
    district: 'Raigad',
    nearestStation: 'Taloja Panchanand',
    distanceFromMumbai: '50 km from CST',
    infraProjects: [
      { name: 'CIDCO Mass Housing Phase 3', status: 'Under Construction', expectedCompletion: '2027', impact: 'Planned residential development with social infrastructure, attracting families.' },
      { name: 'Panvel-Karjat Rail Upgrades', status: 'Under Construction', expectedCompletion: '2026', impact: 'Improved frequency and new stations, better connectivity to Panvel and beyond.' },
      { name: 'Proposed Metro Extension to Taloja', status: 'Planned', expectedCompletion: '2029', impact: 'Will connect Taloja to Navi Mumbai metro network and CBD Belapur.' },
      { name: 'Taloja MIDC Phase 2 Expansion', status: 'Under Construction', expectedCompletion: '2026', impact: 'New industrial units driving employment and housing demand.' },
    ],
    priceTrends: [
      { year: '2020', price: '₹2,800', change: '-' },
      { year: '2021', price: '₹3,000', change: '+7.1%' },
      { year: '2022', price: '₹3,300', change: '+10.0%' },
      { year: '2023', price: '₹3,600', change: '+9.1%' },
      { year: '2024', price: '₹3,900', change: '+8.3%' },
      { year: '2025', price: '₹4,200', change: '+10.5%' },
    ],
    projects: [
      { slug: 'taloja-greens', name: 'Taloja Greens', builder: 'Puraniks', priceRange: '₹32 L - ₹72 L', configs: ['1 BHK', '2 BHK'], status: 'Pre-Launch', imageColor: 'bg-lime-100' },
    ],
    pros: ['Lowest price per sq ft in Navi Mumbai', 'CIDCO-planned development reduces risk', 'Proximity to Panvel and Kharghar growth corridors', 'Taloja MIDC provides employment base', 'Ideal first-time investor entry point'],
    cons: ['Social infrastructure still developing', 'Limited premium retail and entertainment options', 'Public transport frequency needs improvement'],
    idealFor: ['First-time investors with ₹30-45 L budget', 'Affordable housing end-users', 'Long-term investors betting on Navi Mumbai growth'],
  },
  'kharghar': {
    name: 'Kharghar',
    slug: 'kharghar',
    tagline: 'Established Premium Hub',
    description: 'Kharghar is Navi Mumbai\'s most prestigious and well-planned residential node, developed by CIDCO. Known for its Central Park (one of Asia\'s largest), Golf Course, and excellent social infrastructure, Kharghar offers a premium lifestyle that few areas in the MMR can match. The area has a mature ecosystem of schools (DAV, Ryan, Vibgyor), hospitals (MGM), malls (Rassaz), and corporate offices. It is ideal for investors seeking stable appreciation with lower risk.',
    investmentThesis: 'Kharghar is the "blue-chip" of Navi Mumbai real estate. While it offers lower percentage returns compared to emerging micro-markets, it provides the most stable and predictable appreciation with minimal risk. Current pricing at ₹8,500/sq ft is still significantly below Mumbai mainland (₹20,000+/sq ft) and offers excellent value for the lifestyle and infrastructure it delivers. The upcoming CIDCO Exhibition Centre expansion and corporate hub growth will drive steady demand. We project 12-15% stable appreciation, suitable for conservative investors.',
    avgPriceSqFt: '₹8,500/sq ft',
    yoyGrowth: '+8.3%',
    growthPotential: '12-15%',
    growthLevel: 'Stable',
    population: '5 Lakh+',
    pincode: '410210',
    district: 'Raigad',
    nearestStation: 'Kharghar',
    distanceFromMumbai: '40 km from CST',
    infraProjects: [
      { name: 'CIDCO Exhibition Centre Expansion', status: 'Under Construction', expectedCompletion: '2026', impact: 'India\'s largest exhibition centre expansion, driving commercial activity and employment.' },
      { name: 'Kharghar-Uran Railway Line', status: 'Under Construction', expectedCompletion: '2027', impact: 'New rail connectivity to Uran and JNPT area, opening southern corridor.' },
      { name: 'Corporate Hub Phase 2', status: 'Planned', expectedCompletion: '2028', impact: 'New office spaces attracting IT/ITES companies, driving housing demand.' },
      { name: 'Kharghar Hills Conservation', status: 'Completed', expectedCompletion: '2024', impact: 'Protected green belt ensuring long-term environmental quality.' },
    ],
    priceTrends: [
      { year: '2020', price: '₹6,200', change: '-' },
      { year: '2021', price: '₹6,600', change: '+6.5%' },
      { year: '2022', price: '₹7,100', change: '+7.6%' },
      { year: '2023', price: '₹7,600', change: '+7.0%' },
      { year: '2024', price: '₹8,000', change: '+5.3%' },
      { year: '2025', price: '₹8,500', change: '+8.3%' },
    ],
    projects: [
      { slug: 'skyline-towers-kharghar', name: 'Skyline Towers', builder: 'Adhiraj Constructions', priceRange: '₹78 L - ₹1.8 Cr', configs: ['2 BHK', '3 BHK'], status: 'Ready to Move', imageColor: 'bg-purple-100' },
    ],
    pros: ['Most established and well-planned node in Navi Mumbai', 'Central Park, Golf Course, premium lifestyle', 'Excellent schools, hospitals, and social infrastructure', 'Strong and predictable appreciation history', 'High rental demand from corporate professionals'],
    cons: ['Highest price per sq ft among covered areas', 'Limited new inventory due to mature development', 'Peak hour congestion on internal roads'],
    idealFor: ['Conservative investors seeking stable returns', 'End-users wanting premium lifestyle', 'NRIs looking for ready-to-move-in options', 'Rental yield investors targeting corporate tenants'],
  },
  'ulwe': {
    name: 'Ulwe',
    slug: 'ulwe',
    tagline: 'Highest Growth Potential',
    description: 'Ulwe has emerged as the highest growth potential area in the entire Mumbai Metropolitan Region. The game-changing Mumbai Trans Harbour Link (MTHL), now operational, has reduced travel time to South Mumbai from over 2 hours to just 20 minutes. Combined with the upcoming Navi Mumbai International Airport within 10 km and the proposed Navi Mumbai Metro, Ulwe is positioned for a transformation from a developing node to a premium residential destination. The area is also close to JNPT (India\'s largest container port) and the upcoming Navi Mumbai SEZ.',
    investmentThesis: 'Ulwe represents the highest upside potential in our coverage universe. The MTHL is already operational and has fundamentally changed Ulwe\'s accessibility equation. Current prices at ₹6,100/sq ft represent a significant discount to the connectivity premium this area now commands. With the airport and metro still to come, we see multiple catalysts that are yet to be priced in. We project 20-25% appreciation over 3-5 years, the highest among all covered areas. Early investors stand to benefit the most as each infrastructure milestone triggers re-pricing.',
    avgPriceSqFt: '₹6,100/sq ft',
    yoyGrowth: '+15.2%',
    growthPotential: '20-25%',
    growthLevel: 'High',
    population: '2.5 Lakh+',
    pincode: '410206',
    district: 'Raigad',
    nearestStation: 'Proposed Ulwe Station',
    distanceFromMumbai: '35 km from CST (via MTHL)',
    infraProjects: [
      { name: 'Mumbai Trans Harbour Link (MTHL)', status: 'Completed', expectedCompletion: '2024', impact: 'Game-changer. South Mumbai in 20 minutes. Already driving massive demand.' },
      { name: 'Navi Mumbai International Airport', status: 'Under Construction', expectedCompletion: '2027', impact: 'Within 10 km. Will make Ulwe a premium airport zone residential area.' },
      { name: 'Navi Mumbai Metro Line 1', status: 'Under Construction', expectedCompletion: '2026', impact: 'Will connect Ulwe to Belapur CBD and the wider Navi Mumbai metro network.' },
      { name: 'Proposed Ulwe Railway Station', status: 'Planned', expectedCompletion: '2028', impact: 'Will add rail connectivity to Harbour line, completing the transport grid.' },
      { name: 'Navi Mumbai SEZ', status: 'Planned', expectedCompletion: '2029', impact: 'Large-scale commercial development driving employment and housing demand.' },
    ],
    priceTrends: [
      { year: '2020', price: '₹3,500', change: '-' },
      { year: '2021', price: '₹3,800', change: '+8.6%' },
      { year: '2022', price: '₹4,200', change: '+10.5%' },
      { year: '2023', price: '₹4,700', change: '+11.9%' },
      { year: '2024', price: '₹5,300', change: '+12.8%' },
      { year: '2025', price: '₹6,100', change: '+15.2%' },
    ],
    projects: [
      { slug: 'navi-heights-ulwe', name: 'Navi Heights', builder: 'Balaji Group', priceRange: '₹42 L - ₹1.05 Cr', configs: ['1 BHK', '2 BHK', '3 BHK'], status: 'Under Construction', imageColor: 'bg-orange-100' },
    ],
    pros: ['Highest growth potential in MMR (20-25%)', 'MTHL operational - South Mumbai in 20 min', 'Closest area to Navi Mumbai Airport', 'Multiple infrastructure catalysts yet to be priced in', 'CIDCO-planned development with wide roads'],
    cons: ['Social infrastructure still developing rapidly', 'Railway station is still proposed/planned', 'Some sectors are still under active development'],
    idealFor: ['Aggressive growth investors', 'Investors comfortable with 3-5 year holding period', 'NRIs looking for high-appreciation assets', 'First-time investors with ₹40-55 L budget'],
  },
}

const area = computed(() => areasData[slug] || null)

// Dynamic SEO
if (area.value) {
  useHead({
    title: `${area.value.name} Real Estate Investment Guide | PropTech`,
    meta: [
      { name: 'description', content: `${area.value.name} real estate investment guide. Avg price: ${area.value.avgPriceSqFt}. Growth potential: ${area.value.growthPotential}. Infrastructure updates, price trends, and RERA-verified projects in ${area.value.name}.` },
      { name: 'keywords', content: `${area.value.name} real estate, ${area.value.name} property investment, ${area.value.name} price trends, ${area.value.name} infrastructure, buy flat in ${area.value.name}` },
    ],
    link: [
      { rel: 'canonical', href: `https://www.proptech.in/areas/${slug}` },
    ],
  })

  useSeoMeta({
    title: `${area.value.name} Real Estate Investment Guide | PropTech`,
    description: `${area.value.name} real estate investment guide. Avg price: ${area.value.avgPriceSqFt}. Growth potential: ${area.value.growthPotential}. Infrastructure updates, price trends, and RERA-verified projects.`,
    ogTitle: `${area.value.name} Real Estate Investment Guide | PropTech`,
    ogDescription: `Data-driven investment guide for ${area.value.name}. Avg price: ${area.value.avgPriceSqFt}. Growth: ${area.value.growthPotential}. Infrastructure timelines and RERA-verified projects.`,
    ogType: 'article',
    ogUrl: `https://www.proptech.in/areas/${slug}`,
    ogImage: 'https://www.proptech.in/og-area.jpg',
    twitterCard: 'summary_large_image',
    twitterTitle: `${area.value.name} Investment Guide | PropTech`,
    twitterDescription: `Real estate investment guide for ${area.value.name}. Growth potential: ${area.value.growthPotential}`,
  })

  useHead({
    script: [
      {
        type: 'application/ld+json',
        innerHTML: JSON.stringify({
          '@context': 'https://schema.org',
          '@type': 'Place',
          'name': area.value.name,
          'description': area.value.description,
          'address': {
            '@type': 'PostalAddress',
            'addressLocality': area.value.name,
            'addressRegion': 'Maharashtra',
            'postalCode': area.value.pincode,
            'addressCountry': 'IN',
          },
          'geo': {
            '@type': 'GeoCoordinates',
            'addressCountry': 'IN',
          },
          'additionalProperty': [
            { '@type': 'PropertyValue', 'name': 'Average Price per sq ft', 'value': area.value.avgPriceSqFt },
            { '@type': 'PropertyValue', 'name': 'Year-over-Year Growth', 'value': area.value.yoyGrowth },
            { '@type': 'PropertyValue', 'name': 'Growth Potential', 'value': area.value.growthPotential },
            { '@type': 'PropertyValue', 'name': 'District', 'value': area.value.district },
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
            { '@type': 'ListItem', position: 2, name: 'Areas', item: 'https://www.proptech.in/areas' },
            { '@type': 'ListItem', position: 3, name: area.value.name, item: `https://www.proptech.in/areas/${slug}` },
          ],
        }),
      },
    ],
  })
} else {
  useHead({ title: 'Area Not Found | PropTech' })
  useSeoMeta({ title: 'Area Not Found | PropTech', description: 'The requested area guide could not be found.' })
}

function getStatusColor(status: string) {
  switch (status) {
    case 'Completed': return 'success' as const
    case 'Under Construction': return 'warning' as const
    case 'Planned': return 'info' as const
    default: return 'neutral' as const
  }
}

function getGrowthColor(level: string) {
  switch (level) {
    case 'High': return 'success' as const
    case 'Medium': return 'warning' as const
    case 'Stable': return 'info' as const
    default: return 'neutral' as const
  }
}

function getWhatsAppUrl(areaName: string) {
  const text = encodeURIComponent(`Hi, I am interested in investing in ${areaName}. Can you share more details about available projects and growth potential?`)
  return `https://wa.me/919876543210?text=${text}`
}
</script>

<template>
  <div>
    <!-- Not Found State -->
    <template v-if="!area">
      <section class="py-20 px-4 text-center">
        <UIcon name="i-lucide-map" class="text-6xl text-slate-300 mb-4" />
        <h1 class="text-2xl font-bold text-brand-900 dark:text-white">Area Not Found</h1>
        <p class="text-slate-500 dark:text-slate-400 mt-2 mb-6">The area guide you are looking for does not exist.</p>
        <UButton to="/areas" variant="outline" icon="i-lucide-arrow-left">
          Back to Areas
        </UButton>
      </section>
    </template>

    <!-- Area Detail -->
    <template v-else>
      <!-- Breadcrumb -->
      <div class="bg-white dark:bg-brand-900 border-b border-slate-100 dark:border-slate-800 py-3 px-4">
        <div class="max-w-6xl mx-auto flex items-center gap-2 text-sm text-slate-500 dark:text-slate-400">
          <NuxtLink to="/" class="hover:text-brand-600 dark:hover:text-brand-400">Home</NuxtLink>
          <UIcon name="i-lucide-chevron-right" class="text-xs" />
          <NuxtLink to="/areas" class="hover:text-brand-600 dark:hover:text-brand-400">Areas</NuxtLink>
          <UIcon name="i-lucide-chevron-right" class="text-xs" />
          <span class="text-brand-900 dark:text-white font-medium">{{ area.name }}</span>
        </div>
      </div>

      <!-- Area Hero -->
      <section class="bg-brand-800 text-white py-10 px-4">
        <div class="max-w-6xl mx-auto">
          <div class="flex flex-wrap items-center gap-3 mb-4">
            <UBadge :color="getGrowthColor(area.growthLevel)" variant="solid">
              {{ area.growthLevel }} Growth Potential
            </UBadge>
            <UBadge color="neutral" variant="subtle">
              {{ area.district }} District
            </UBadge>
          </div>
          <h1 class="text-3xl md:text-4xl font-bold">{{ area.name }}</h1>
          <p class="text-accent-400 font-medium text-lg mt-1">{{ area.tagline }}</p>

          <!-- Quick Stats -->
          <div class="grid grid-cols-2 sm:grid-cols-4 gap-4 mt-6">
            <div class="bg-white/10 rounded-xl p-3">
              <p class="text-xs text-slate-400">Avg Price/sq ft</p>
              <p class="text-xl font-bold mt-1">{{ area.avgPriceSqFt }}</p>
            </div>
            <div class="bg-white/10 rounded-xl p-3">
              <p class="text-xs text-slate-400">YoY Growth</p>
              <p class="text-xl font-bold mt-1 text-accent-400">{{ area.yoyGrowth }}</p>
            </div>
            <div class="bg-white/10 rounded-xl p-3">
              <p class="text-xs text-slate-400">Growth Potential</p>
              <p class="text-xl font-bold mt-1">{{ area.growthPotential }}</p>
            </div>
            <div class="bg-white/10 rounded-xl p-3">
              <p class="text-xs text-slate-400">Projects</p>
              <p class="text-xl font-bold mt-1">{{ area.projects.length }}</p>
            </div>
          </div>
        </div>
      </section>

      <!-- Main Content -->
      <section class="py-8 px-4 bg-slate-50 dark:bg-brand-900/50">
        <div class="max-w-6xl mx-auto">
          <div class="grid lg:grid-cols-3 gap-8">
            <!-- Left Column (2/3) -->
            <div class="lg:col-span-2 space-y-8">
              <!-- Area Overview -->
              <UCard>
                <template #header>
                  <h2 class="text-lg font-semibold text-brand-900 dark:text-white">Area Overview</h2>
                </template>
                <p class="text-slate-600 dark:text-slate-300 leading-relaxed">{{ area.description }}</p>
                <div class="grid grid-cols-2 md:grid-cols-3 gap-4 mt-6 pt-4 border-t border-slate-100 dark:border-slate-800">
                  <div>
                    <p class="text-xs text-slate-400">Population</p>
                    <p class="font-semibold text-slate-900 dark:text-white">{{ area.population }}</p>
                  </div>
                  <div>
                    <p class="text-xs text-slate-400">Nearest Station</p>
                    <p class="font-semibold text-slate-900 dark:text-white">{{ area.nearestStation }}</p>
                  </div>
                  <div>
                    <p class="text-xs text-slate-400">Distance from Mumbai</p>
                    <p class="font-semibold text-slate-900 dark:text-white">{{ area.distanceFromMumbai }}</p>
                  </div>
                </div>
              </UCard>

              <!-- Infrastructure Development -->
              <UCard>
                <template #header>
                  <h2 class="text-lg font-semibold text-brand-900 dark:text-white">Infrastructure Development</h2>
                </template>
                <div class="space-y-4">
                  <div
                    v-for="infra in area.infraProjects"
                    :key="infra.name"
                    class="border border-slate-200 dark:border-slate-700 rounded-lg p-4"
                  >
                    <div class="flex items-start justify-between gap-3">
                      <div class="flex-1">
                        <div class="flex items-center gap-2 mb-1">
                          <h4 class="font-semibold text-slate-900 dark:text-white text-sm">{{ infra.name }}</h4>
                          <UBadge :color="getStatusColor(infra.status)" variant="subtle" size="xs">
                            {{ infra.status }}
                          </UBadge>
                        </div>
                        <p class="text-xs text-slate-500 dark:text-slate-400 mb-2">Expected: {{ infra.expectedCompletion }}</p>
                        <p class="text-sm text-slate-600 dark:text-slate-300">{{ infra.impact }}</p>
                      </div>
                    </div>
                  </div>
                </div>
              </UCard>

              <!-- Price Trends -->
              <UCard>
                <template #header>
                  <h2 class="text-lg font-semibold text-brand-900 dark:text-white">Price Trends (per sq ft)</h2>
                </template>
                <!-- Price Trend Bar Chart Placeholder -->
                <div class="h-48 bg-slate-100 dark:bg-brand-800 rounded-lg flex items-center justify-center mb-6">
                  <div class="text-center">
                    <UIcon name="i-lucide-bar-chart-3" class="text-4xl text-slate-400 mb-2" />
                    <p class="text-sm text-slate-500 dark:text-slate-400">Interactive chart coming soon</p>
                  </div>
                </div>
                <!-- Price Table -->
                <div class="overflow-x-auto">
                  <table class="w-full text-sm">
                    <thead>
                      <tr class="border-b-2 border-slate-200 dark:border-slate-700">
                        <th class="text-left py-2 px-3 text-slate-500 dark:text-slate-400 font-medium">Year</th>
                        <th class="text-left py-2 px-3 text-slate-500 dark:text-slate-400 font-medium">Avg Price/sq ft</th>
                        <th class="text-left py-2 px-3 text-slate-500 dark:text-slate-400 font-medium">YoY Change</th>
                      </tr>
                    </thead>
                    <tbody>
                      <tr
                        v-for="trend in area.priceTrends"
                        :key="trend.year"
                        class="border-b border-slate-100 dark:border-slate-800"
                      >
                        <td class="py-2 px-3 font-medium text-slate-900 dark:text-white">{{ trend.year }}</td>
                        <td class="py-2 px-3">{{ trend.price }}</td>
                        <td class="py-2 px-3">
                          <span v-if="trend.change !== '-'" class="font-semibold text-accent-600 dark:text-accent-400">{{ trend.change }}</span>
                          <span v-else class="text-slate-400">-</span>
                        </td>
                      </tr>
                    </tbody>
                  </table>
                </div>
              </UCard>

              <!-- Investment Thesis -->
              <UCard>
                <template #header>
                  <h2 class="text-lg font-semibold text-brand-900 dark:text-white">Investment Thesis</h2>
                </template>
                <p class="text-slate-600 dark:text-slate-300 leading-relaxed">{{ area.investmentThesis }}</p>

                <!-- Pros & Cons -->
                <div class="grid md:grid-cols-2 gap-6 mt-6 pt-4 border-t border-slate-100 dark:border-slate-800">
                  <div>
                    <h4 class="font-semibold text-green-700 dark:text-green-300 mb-3 flex items-center gap-1">
                      <UIcon name="i-lucide-thumbs-up" />
                      Strengths
                    </h4>
                    <div class="space-y-2">
                      <div v-for="pro in area.pros" :key="pro" class="flex items-start gap-2">
                        <UIcon name="i-lucide-check" class="text-green-600 dark:text-green-400 mt-0.5 flex-shrink-0" />
                        <span class="text-sm text-slate-600 dark:text-slate-300">{{ pro }}</span>
                      </div>
                    </div>
                  </div>
                  <div>
                    <h4 class="font-semibold text-amber-700 dark:text-amber-300 mb-3 flex items-center gap-1">
                      <UIcon name="i-lucide-alert-triangle" />
                      Risks to Consider
                    </h4>
                    <div class="space-y-2">
                      <div v-for="con in area.cons" :key="con" class="flex items-start gap-2">
                        <UIcon name="i-lucide-minus" class="text-amber-600 dark:text-amber-400 mt-0.5 flex-shrink-0" />
                        <span class="text-sm text-slate-600 dark:text-slate-300">{{ con }}</span>
                      </div>
                    </div>
                  </div>
                </div>
              </UCard>

              <!-- Ideal Investor Profile -->
              <UCard>
                <template #header>
                  <h2 class="text-lg font-semibold text-brand-900 dark:text-white">Ideal For</h2>
                </template>
                <div class="flex flex-wrap gap-2">
                  <UBadge
                    v-for="ideal in area.idealFor"
                    :key="ideal"
                    color="info"
                    variant="subtle"
                    size="sm"
                  >
                    {{ ideal }}
                  </UBadge>
                </div>
              </UCard>
            </div>

            <!-- Right Sidebar (1/3) -->
            <div class="space-y-6">
              <!-- Quick Stats Card -->
              <UCard class="sticky top-4">
                <template #header>
                  <h3 class="font-semibold text-brand-900 dark:text-white">Quick Stats</h3>
                </template>
                <div class="space-y-3 text-sm">
                  <div class="flex justify-between">
                    <span class="text-slate-500 dark:text-slate-400">Avg Price/sq ft</span>
                    <span class="font-bold text-slate-900 dark:text-white">{{ area.avgPriceSqFt }}</span>
                  </div>
                  <div class="flex justify-between">
                    <span class="text-slate-500 dark:text-slate-400">YoY Growth</span>
                    <span class="font-bold text-accent-600 dark:text-accent-400">{{ area.yoyGrowth }}</span>
                  </div>
                  <div class="flex justify-between">
                    <span class="text-slate-500 dark:text-slate-400">Growth Potential</span>
                    <span class="font-bold text-accent-600 dark:text-accent-400">{{ area.growthPotential }}</span>
                  </div>
                  <div class="flex justify-between">
                    <span class="text-slate-500 dark:text-slate-400">Projects Listed</span>
                    <span class="font-medium">{{ area.projects.length }}</span>
                  </div>
                  <div class="flex justify-between">
                    <span class="text-slate-500 dark:text-slate-400">Pincode</span>
                    <span class="font-mono text-xs">{{ area.pincode }}</span>
                  </div>
                  <div class="flex justify-between">
                    <span class="text-slate-500 dark:text-slate-400">District</span>
                    <span class="font-medium">{{ area.district }}</span>
                  </div>
                </div>
                <UButton
                  :href="getWhatsAppUrl(area.name)"
                  target="_blank"
                  class="w-full mt-6 bg-green-600 hover:bg-green-700 text-white"
                  icon="i-lucide-message-circle"
                  size="lg"
                >
                  Discuss {{ area.name }} Investment
                </UButton>
                <p class="text-xs text-slate-400 dark:text-slate-500 text-center mt-2">No spam. Just guidance.</p>
              </UCard>

              <!-- Projects in Area -->
              <UCard>
                <template #header>
                  <h3 class="font-semibold text-brand-900 dark:text-white">Projects in {{ area.name }}</h3>
                </template>
                <div class="space-y-4">
                  <NuxtLink
                    v-for="project in area.projects"
                    :key="project.slug"
                    :to="`/projects/${project.slug}`"
                    class="block group"
                  >
                    <div class="border border-slate-200 dark:border-slate-700 rounded-lg p-3 hover:border-brand-500/50 transition-colors">
                      <div :class="[project.imageColor, 'h-20 rounded-md flex items-center justify-center mb-2']">
                        <UIcon name="i-lucide-building-2" class="text-2xl text-slate-400" />
                      </div>
                      <h4 class="font-semibold text-sm text-brand-900 dark:text-white group-hover:text-brand-600 dark:group-hover:text-brand-400">{{ project.name }}</h4>
                      <p class="text-xs text-slate-500 dark:text-slate-400">{{ project.builder }}</p>
                      <p class="text-sm font-bold text-slate-900 dark:text-white mt-1">{{ project.priceRange }}</p>
                      <div class="flex flex-wrap gap-1 mt-1">
                        <UBadge v-for="config in project.configs" :key="config" color="neutral" variant="subtle" size="xs">
                          {{ config }}
                        </UBadge>
                      </div>
                    </div>
                  </NuxtLink>
                </div>
                <UButton to="/projects" variant="outline" size="sm" class="w-full mt-4" icon="i-lucide-arrow-right" trailing>
                  View All Projects
                </UButton>
              </UCard>

              <!-- Other Areas -->
              <UCard>
                <template #header>
                  <h3 class="font-semibold text-brand-900 dark:text-white">Explore Other Areas</h3>
                </template>
                <div class="space-y-2">
                  <NuxtLink
                    v-for="otherSlug in Object.keys(areasData).filter(s => s !== slug)"
                    :key="otherSlug"
                    :to="`/areas/${otherSlug}`"
                    class="flex items-center justify-between text-sm py-2 px-1 rounded hover:bg-slate-50 dark:hover:bg-brand-800 transition-colors"
                  >
                    <span class="font-medium text-brand-900 dark:text-white hover:text-brand-600 dark:hover:text-brand-400">{{ areasData[otherSlug].name }}</span>
                    <UBadge :color="getGrowthColor(areasData[otherSlug].growthLevel)" variant="subtle" size="xs">
                      {{ areasData[otherSlug].growthPotential }}
                    </UBadge>
                  </NuxtLink>
                </div>
              </UCard>
            </div>
          </div>
        </div>
      </section>

      <!-- WhatsApp CTA -->
      <section class="py-12 px-4 bg-brand-800 text-white text-center">
        <div class="max-w-2xl mx-auto">
          <h2 class="text-xl md:text-2xl font-bold">Interested in {{ area.name }}?</h2>
          <p class="text-slate-300 mt-2 mb-6">
            Our team can share detailed project reports, site visit schedules, and personalized investment recommendations for {{ area.name }}.
          </p>
          <UButton
            :href="getWhatsAppUrl(area.name)"
            target="_blank"
            size="xl"
            class="bg-green-600 hover:bg-green-700 text-white"
            icon="i-lucide-message-circle"
          >
            Chat About {{ area.name }}
          </UButton>
        </div>
      </section>
    </template>
  </div>
</template>
