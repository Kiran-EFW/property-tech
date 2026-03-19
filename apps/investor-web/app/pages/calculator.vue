<script setup lang="ts">
useSeoMeta({
  title: 'All-Inclusive Property Cost Calculator — PropTech',
  description: 'Calculate the real total cost of buying property in Mumbai. Includes GST, stamp duty, registration, legal charges, and maintenance — no hidden costs.',
})

const basePrice = ref(8000)
const carpetArea = ref(650)
const floorNumber = ref(5)
const floorRise = ref(50)
const gstRate = ref(5)
const stampDutyRate = ref(6)
const registrationRate = ref(1)
const registrationCap = ref(30000)
const legalCharges = ref(25000)
const maintenancePerSqft = ref(50)
const maintenanceMonths = ref(24)
const parkingCost = ref(500000)

const breakdown = computed(() => {
  const baseCost = basePrice.value * carpetArea.value
  const floorRiseCost = floorRise.value * (floorNumber.value - 1) * carpetArea.value
  const agreementValue = baseCost + floorRiseCost
  const gst = Math.round(agreementValue * (gstRate.value / 100))
  const stampDuty = Math.round(agreementValue * (stampDutyRate.value / 100))
  const registration = Math.min(
    Math.round(agreementValue * (registrationRate.value / 100)),
    registrationCap.value
  )
  const maintenance = maintenancePerSqft.value * carpetArea.value * maintenanceMonths.value
  const total = agreementValue + gst + stampDuty + registration + legalCharges.value + maintenance + parkingCost.value

  return {
    baseCost,
    floorRiseCost,
    agreementValue,
    gst,
    stampDuty,
    registration,
    legal: legalCharges.value,
    maintenance,
    parking: parkingCost.value,
    total,
  }
})

function formatINR(amount: number): string {
  const formatter = new Intl.NumberFormat('en-IN', {
    style: 'currency',
    currency: 'INR',
    maximumFractionDigits: 0,
  })
  return formatter.format(amount)
}
</script>

<template>
  <div class="max-w-4xl mx-auto px-4 py-8">
    <h1 class="text-3xl font-bold text-brand-900 dark:text-white">All-Inclusive Cost Calculator</h1>
    <p class="text-slate-500 dark:text-slate-400 mt-2 mb-8">
      Know the real total cost. No hidden charges. Every rupee accounted for.
    </p>

    <div class="grid md:grid-cols-2 gap-8">
      <!-- Inputs -->
      <UCard>
        <template #header>
          <h2 class="text-lg font-semibold text-brand-900 dark:text-white">Property Details</h2>
        </template>
        <div class="space-y-4">
          <UFormField label="Base Price (₹/sq ft)">
            <UInput v-model.number="basePrice" type="number" />
          </UFormField>
          <UFormField label="Carpet Area (sq ft)">
            <UInput v-model.number="carpetArea" type="number" />
          </UFormField>
          <UFormField label="Floor Number">
            <UInput v-model.number="floorNumber" type="number" />
          </UFormField>
          <UFormField label="Floor Rise (₹/sq ft per floor)">
            <UInput v-model.number="floorRise" type="number" />
          </UFormField>
          <UFormField label="Parking Cost (₹)">
            <UInput v-model.number="parkingCost" type="number" />
          </UFormField>
        </div>
      </UCard>

      <!-- Breakdown -->
      <UCard>
        <template #header>
          <h2 class="text-lg font-semibold text-brand-900 dark:text-white">Cost Breakdown</h2>
        </template>
        <div class="space-y-3">
          <div class="flex justify-between text-sm">
            <span class="text-slate-600 dark:text-slate-300">Base cost ({{ carpetArea }} sq ft × {{ formatINR(basePrice) }})</span>
            <span class="font-medium">{{ formatINR(breakdown.baseCost) }}</span>
          </div>
          <div v-if="breakdown.floorRiseCost > 0" class="flex justify-between text-sm">
            <span class="text-slate-600 dark:text-slate-300">Floor rise (Floor {{ floorNumber }})</span>
            <span class="font-medium">{{ formatINR(breakdown.floorRiseCost) }}</span>
          </div>
          <USeparator />
          <div class="flex justify-between text-sm">
            <span class="text-slate-600 dark:text-slate-300">Agreement Value</span>
            <span class="font-semibold">{{ formatINR(breakdown.agreementValue) }}</span>
          </div>
          <USeparator />
          <div class="flex justify-between text-sm">
            <span class="text-slate-600 dark:text-slate-300">GST ({{ gstRate }}%)</span>
            <span class="font-medium">{{ formatINR(breakdown.gst) }}</span>
          </div>
          <div class="flex justify-between text-sm">
            <span class="text-slate-600 dark:text-slate-300">Stamp Duty ({{ stampDutyRate }}%)</span>
            <span class="font-medium">{{ formatINR(breakdown.stampDuty) }}</span>
          </div>
          <div class="flex justify-between text-sm">
            <span class="text-slate-600 dark:text-slate-300">Registration ({{ registrationRate }}%, max ₹30,000)</span>
            <span class="font-medium">{{ formatINR(breakdown.registration) }}</span>
          </div>
          <div class="flex justify-between text-sm">
            <span class="text-slate-600 dark:text-slate-300">Legal Charges</span>
            <span class="font-medium">{{ formatINR(breakdown.legal) }}</span>
          </div>
          <div class="flex justify-between text-sm">
            <span class="text-slate-600 dark:text-slate-300">Maintenance ({{ maintenanceMonths }} months advance)</span>
            <span class="font-medium">{{ formatINR(breakdown.maintenance) }}</span>
          </div>
          <div class="flex justify-between text-sm">
            <span class="text-slate-600 dark:text-slate-300">Parking</span>
            <span class="font-medium">{{ formatINR(breakdown.parking) }}</span>
          </div>
          <USeparator />
          <div class="flex justify-between pt-2">
            <span class="text-lg font-bold text-brand-900 dark:text-white">Total All-In Cost</span>
            <span class="text-lg font-bold text-brand-900 dark:text-white">{{ formatINR(breakdown.total) }}</span>
          </div>
          <p class="text-xs text-slate-400 mt-2">
            * Excludes brokerage (paid by builder, not investor). Actual amounts may vary by project.
          </p>
        </div>
      </UCard>
    </div>
  </div>
</template>
