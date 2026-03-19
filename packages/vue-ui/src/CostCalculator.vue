<script setup lang="ts">
import { ref, computed, watch } from 'vue';
import {
  calculateAllInclusiveCost,
  calculateEMI,
  formatINR,
  formatCompactINR,
  GST_UNDER_CONSTRUCTION,
  GST_AFFORDABLE,
  STAMP_DUTY_MAHARASHTRA,
} from '@proptech/shared';
import type { CostBreakdown, MaharashtraRegion } from '@proptech/shared';

export interface CostCalculatorProps {
  /** Default base price per sqft */
  defaultBasePricePerSqft?: number;
  /** Default carpet area in sqft */
  defaultCarpetArea?: number;
  /** Default floor number */
  defaultFloorNumber?: number;
  /** Floor rise per sqft per floor (project-specific) */
  floorRisePerSqftPerFloor?: number;
  /** Default GST rate */
  defaultGstRate?: number;
  /** Default region for stamp duty */
  defaultRegion?: MaharashtraRegion;
  /** Default buyer gender for stamp duty */
  defaultGender?: 'male' | 'female' | 'joint';
  /** Maintenance charge per sqft per month */
  maintenancePerSqft?: number;
  /** Number of months advance maintenance */
  maintenanceMonths?: number;
  /** Legal charges */
  legalCharges?: number;
  /** Parking cost */
  parkingCost?: number;
  /** Project name (shown in header) */
  projectName?: string;
}

const props = withDefaults(defineProps<CostCalculatorProps>(), {
  defaultBasePricePerSqft: 10000,
  defaultCarpetArea: 650,
  defaultFloorNumber: 5,
  floorRisePerSqftPerFloor: 50,
  defaultGstRate: GST_UNDER_CONSTRUCTION,
  defaultRegion: 'mumbai',
  defaultGender: 'male',
  maintenancePerSqft: 4,
  maintenanceMonths: 24,
  legalCharges: 25000,
  parkingCost: 300000,
  projectName: '',
});

const emit = defineEmits<{
  (e: 'calculated', breakdown: CostBreakdown): void;
}>();

// ---- Reactive inputs ----
const basePricePerSqft = ref(props.defaultBasePricePerSqft);
const carpetArea = ref(props.defaultCarpetArea);
const floorNumber = ref(props.defaultFloorNumber);
const region = ref<MaharashtraRegion>(props.defaultRegion);
const gender = ref<'male' | 'female' | 'joint'>(props.defaultGender);
const gstType = ref<'standard' | 'affordable'>(
  props.defaultGstRate === GST_AFFORDABLE ? 'affordable' : 'standard',
);

// ---- EMI inputs ----
const showEMI = ref(false);
const loanPercentage = ref(80);
const interestRate = ref(8.5);
const tenureYears = ref(20);

// ---- Region options ----
const regionOptions: { value: MaharashtraRegion; label: string }[] = [
  { value: 'mumbai', label: 'Mumbai (BMC)' },
  { value: 'mumbaiSuburban', label: 'Mumbai Suburban' },
  { value: 'thane', label: 'Thane' },
  { value: 'naviMumbai', label: 'Navi Mumbai' },
  { value: 'pune', label: 'Pune' },
  { value: 'restOfMaharashtra', label: 'Rest of Maharashtra' },
];

const genderOptions: { value: 'male' | 'female' | 'joint'; label: string }[] = [
  { value: 'male', label: 'Male' },
  { value: 'female', label: 'Female' },
  { value: 'joint', label: 'Joint' },
];

// ---- Computed stamp duty & GST rate ----
const stampDutyRate = computed(() => {
  const regionRates = STAMP_DUTY_MAHARASHTRA[region.value];
  return regionRates[gender.value];
});

const gstRate = computed(() => {
  return gstType.value === 'affordable' ? GST_AFFORDABLE : GST_UNDER_CONSTRUCTION;
});

// ---- Cost breakdown ----
const breakdown = computed<CostBreakdown>(() => {
  return calculateAllInclusiveCost({
    basePricePerSqft: basePricePerSqft.value,
    carpetAreaSqft: carpetArea.value,
    floorNumber: floorNumber.value,
    floorRisePerSqftPerFloor: props.floorRisePerSqftPerFloor,
    gstRate: gstRate.value,
    stampDutyRate: stampDutyRate.value,
    maintenancePerSqft: props.maintenancePerSqft,
    maintenanceMonths: props.maintenanceMonths,
    legalCharges: props.legalCharges,
    parkingCost: props.parkingCost,
  });
});

// ---- EMI calculation ----
const monthlyEMI = computed(() => {
  if (!showEMI.value) return 0;
  const loanAmount = breakdown.value.total.amount * (loanPercentage.value / 100);
  return calculateEMI(loanAmount, interestRate.value / 100, tenureYears.value);
});

// ---- Emit on change ----
watch(breakdown, (val) => {
  emit('calculated', val);
}, { immediate: true });

// ---- Breakdown line items for display ----
const lineItems = computed(() => {
  const items = [
    breakdown.value.baseCost,
    ...(breakdown.value.floorRise.amount > 0 ? [breakdown.value.floorRise] : []),
    breakdown.value.gst,
    breakdown.value.stampDuty,
    breakdown.value.registration,
  ];

  if (breakdown.value.legalCharges.amount > 0) {
    items.push(breakdown.value.legalCharges);
  }
  if (breakdown.value.maintenance.amount > 0) {
    items.push(breakdown.value.maintenance);
  }
  if (breakdown.value.parking.amount > 0) {
    items.push(breakdown.value.parking);
  }

  return items;
});
</script>

<template>
  <div class="bg-white rounded-xl border border-slate-200 shadow-sm overflow-hidden">
    <!-- Header -->
    <div
      class="px-5 py-4 border-b border-slate-100"
      style="background-color: #1a2b4a;"
    >
      <h3 class="text-lg font-semibold text-white">
        All-Inclusive Cost Calculator
      </h3>
      <p v-if="projectName" class="text-sm text-slate-300 mt-0.5">
        {{ projectName }}
      </p>
    </div>

    <div class="p-5 space-y-6">
      <!-- Input section -->
      <div class="grid grid-cols-1 sm:grid-cols-3 gap-4">
        <!-- Base price per sqft -->
        <div>
          <label class="block text-xs font-medium text-slate-500 mb-1.5">
            Base Price (per sq ft)
          </label>
          <div class="relative">
            <span class="absolute left-3 top-1/2 -translate-y-1/2 text-slate-400 text-sm">
              ₹
            </span>
            <input
              v-model.number="basePricePerSqft"
              type="number"
              min="1000"
              step="100"
              class="w-full pl-7 pr-3 py-2.5 rounded-lg border border-slate-200 text-sm font-medium focus:outline-none focus:ring-2 focus:border-transparent"
              style="color: #1a2b4a;"
              :style="{ '--tw-ring-color': '#2a9d8f' }"
            />
          </div>
        </div>

        <!-- Carpet area -->
        <div>
          <label class="block text-xs font-medium text-slate-500 mb-1.5">
            Carpet Area (sq ft)
          </label>
          <input
            v-model.number="carpetArea"
            type="number"
            min="100"
            step="10"
            class="w-full px-3 py-2.5 rounded-lg border border-slate-200 text-sm font-medium focus:outline-none focus:ring-2 focus:border-transparent"
            style="color: #1a2b4a;"
            :style="{ '--tw-ring-color': '#2a9d8f' }"
          />
        </div>

        <!-- Floor number -->
        <div>
          <label class="block text-xs font-medium text-slate-500 mb-1.5">
            Floor Number
          </label>
          <input
            v-model.number="floorNumber"
            type="number"
            min="0"
            max="80"
            class="w-full px-3 py-2.5 rounded-lg border border-slate-200 text-sm font-medium focus:outline-none focus:ring-2 focus:border-transparent"
            style="color: #1a2b4a;"
            :style="{ '--tw-ring-color': '#2a9d8f' }"
          />
        </div>
      </div>

      <!-- Tax configuration -->
      <div class="grid grid-cols-1 sm:grid-cols-3 gap-4">
        <!-- Region -->
        <div>
          <label class="block text-xs font-medium text-slate-500 mb-1.5">
            Region
          </label>
          <select
            v-model="region"
            class="w-full px-3 py-2.5 rounded-lg border border-slate-200 text-sm font-medium bg-white focus:outline-none focus:ring-2 focus:border-transparent"
            style="color: #1a2b4a;"
            :style="{ '--tw-ring-color': '#2a9d8f' }"
          >
            <option
              v-for="opt in regionOptions"
              :key="opt.value"
              :value="opt.value"
            >
              {{ opt.label }}
            </option>
          </select>
        </div>

        <!-- Gender (stamp duty) -->
        <div>
          <label class="block text-xs font-medium text-slate-500 mb-1.5">
            Buyer Gender (Stamp Duty)
          </label>
          <select
            v-model="gender"
            class="w-full px-3 py-2.5 rounded-lg border border-slate-200 text-sm font-medium bg-white focus:outline-none focus:ring-2 focus:border-transparent"
            style="color: #1a2b4a;"
            :style="{ '--tw-ring-color': '#2a9d8f' }"
          >
            <option
              v-for="opt in genderOptions"
              :key="opt.value"
              :value="opt.value"
            >
              {{ opt.label }}
            </option>
          </select>
        </div>

        <!-- GST type -->
        <div>
          <label class="block text-xs font-medium text-slate-500 mb-1.5">
            GST Rate
          </label>
          <select
            v-model="gstType"
            class="w-full px-3 py-2.5 rounded-lg border border-slate-200 text-sm font-medium bg-white focus:outline-none focus:ring-2 focus:border-transparent"
            style="color: #1a2b4a;"
            :style="{ '--tw-ring-color': '#2a9d8f' }"
          >
            <option value="standard">5% (Under Construction)</option>
            <option value="affordable">1% (Affordable Housing)</option>
          </select>
        </div>
      </div>

      <!-- Divider -->
      <hr class="border-slate-100" />

      <!-- Cost breakdown -->
      <div class="space-y-2">
        <h4 class="text-sm font-semibold text-slate-500 uppercase tracking-wide">
          Cost Breakdown
        </h4>

        <!-- Agreement value highlight -->
        <div class="px-4 py-3 rounded-lg bg-slate-50 border border-slate-100">
          <div class="flex justify-between items-center">
            <span class="text-sm text-slate-600">{{ breakdown.agreementValue.label }}</span>
            <span class="text-sm font-bold" style="color: #1a2b4a;">
              {{ breakdown.agreementValue.formatted }}
            </span>
          </div>
          <p class="text-xs text-slate-400 mt-1">
            Effective rate: {{ formatINR(breakdown.effectivePricePerSqft) }}/sq ft
          </p>
        </div>

        <!-- Line items -->
        <div class="space-y-1.5 pl-1">
          <div
            v-for="item in lineItems"
            :key="item.label"
            class="flex justify-between items-center py-1.5"
          >
            <span class="text-sm text-slate-500">{{ item.label }}</span>
            <span class="text-sm font-medium text-slate-700 font-mono tabular-nums">
              {{ item.formatted }}
            </span>
          </div>
        </div>

        <!-- Total -->
        <div
          class="flex justify-between items-center px-4 py-3.5 rounded-lg mt-3"
          style="background-color: #1a2b4a;"
        >
          <span class="text-sm font-semibold text-white">
            {{ breakdown.total.label }}
          </span>
          <span class="text-lg font-bold text-white">
            {{ breakdown.total.formatted }}
          </span>
        </div>

        <!-- Compact total -->
        <p class="text-center text-xs text-slate-400 mt-1">
          {{ formatCompactINR(breakdown.total.amount) }}
        </p>
      </div>

      <!-- EMI section -->
      <div class="border-t border-slate-100 pt-4">
        <button
          type="button"
          class="flex items-center gap-2 text-sm font-medium transition-colors duration-150"
          style="color: #2a9d8f;"
          @click="showEMI = !showEMI"
        >
          <svg
            class="w-4 h-4 transition-transform duration-150"
            :class="{ 'rotate-90': showEMI }"
            fill="none"
            viewBox="0 0 24 24"
            stroke="currentColor"
            stroke-width="2"
          >
            <path stroke-linecap="round" stroke-linejoin="round" d="M9 5l7 7-7 7" />
          </svg>
          {{ showEMI ? 'Hide' : 'Show' }} EMI Estimate
        </button>

        <div v-if="showEMI" class="mt-4 space-y-4">
          <div class="grid grid-cols-3 gap-3">
            <div>
              <label class="block text-xs font-medium text-slate-500 mb-1">
                Loan %
              </label>
              <input
                v-model.number="loanPercentage"
                type="number"
                min="10"
                max="90"
                step="5"
                class="w-full px-2.5 py-2 rounded-lg border border-slate-200 text-sm font-medium focus:outline-none focus:ring-2 focus:border-transparent"
                style="color: #1a2b4a;"
                :style="{ '--tw-ring-color': '#2a9d8f' }"
              />
            </div>
            <div>
              <label class="block text-xs font-medium text-slate-500 mb-1">
                Interest %
              </label>
              <input
                v-model.number="interestRate"
                type="number"
                min="5"
                max="15"
                step="0.25"
                class="w-full px-2.5 py-2 rounded-lg border border-slate-200 text-sm font-medium focus:outline-none focus:ring-2 focus:border-transparent"
                style="color: #1a2b4a;"
                :style="{ '--tw-ring-color': '#2a9d8f' }"
              />
            </div>
            <div>
              <label class="block text-xs font-medium text-slate-500 mb-1">
                Tenure (yr)
              </label>
              <input
                v-model.number="tenureYears"
                type="number"
                min="5"
                max="30"
                step="1"
                class="w-full px-2.5 py-2 rounded-lg border border-slate-200 text-sm font-medium focus:outline-none focus:ring-2 focus:border-transparent"
                style="color: #1a2b4a;"
                :style="{ '--tw-ring-color': '#2a9d8f' }"
              />
            </div>
          </div>

          <div
            class="flex justify-between items-center px-4 py-3 rounded-lg border"
            style="background-color: rgba(42, 157, 143, 0.06); border-color: rgba(42, 157, 143, 0.2);"
          >
            <div>
              <p class="text-xs text-slate-500">Estimated Monthly EMI</p>
              <p class="text-xs text-slate-400 mt-0.5">
                Loan: {{ formatCompactINR(breakdown.total.amount * (loanPercentage / 100)) }}
              </p>
            </div>
            <p class="text-lg font-bold" style="color: #2a9d8f;">
              {{ formatINR(monthlyEMI) }}<span class="text-sm font-normal text-slate-400">/mo</span>
            </p>
          </div>
        </div>
      </div>

      <!-- Disclaimer -->
      <p class="text-[11px] text-slate-400 leading-relaxed">
        * All calculations are indicative. Actual costs may vary based on builder pricing,
        government charges, and applicable taxes at the time of transaction. Stamp duty and
        registration charges are as per current Maharashtra government rates. GST is applicable
        only on under-construction properties.
      </p>
    </div>
  </div>
</template>
