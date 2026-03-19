<script setup lang="ts">
import { computed } from 'vue';
import { formatINR, formatCompactINR, calculateEMI } from '@proptech/shared';

export interface PriceDisplayProps {
  /** Primary price to display (in INR) */
  price: number;
  /** Optional label above the price (e.g., "Starting from") */
  label?: string;
  /** Show compact format (Cr/L) instead of full format */
  compact?: boolean;
  /** Optional all-inclusive cost (shown below primary price) */
  allInclusiveCost?: number;
  /** Show EMI estimate */
  showEMI?: boolean;
  /** Loan percentage of price (default 80%) */
  loanPercentage?: number;
  /** Annual interest rate (default 8.5%) */
  interestRate?: number;
  /** Loan tenure in years (default 20) */
  tenureYears?: number;
  /** Size variant */
  size?: 'sm' | 'md' | 'lg';
}

const props = withDefaults(defineProps<PriceDisplayProps>(), {
  label: '',
  compact: false,
  allInclusiveCost: 0,
  showEMI: false,
  loanPercentage: 80,
  interestRate: 8.5,
  tenureYears: 20,
  size: 'md',
});

const formattedPrice = computed(() => {
  return props.compact ? formatCompactINR(props.price) : formatINR(props.price);
});

const formattedAllInclusive = computed(() => {
  if (!props.allInclusiveCost) return '';
  return props.compact
    ? formatCompactINR(props.allInclusiveCost)
    : formatINR(props.allInclusiveCost);
});

const emiDisplay = computed(() => {
  if (!props.showEMI) return '';

  const loanAmount = props.price * (props.loanPercentage / 100);
  const emi = calculateEMI(loanAmount, props.interestRate / 100, props.tenureYears);

  if (emi <= 0) return '';
  return formatINR(emi);
});

const priceSizeClasses = computed(() => {
  const map: Record<string, string> = {
    sm: 'text-base font-bold',
    md: 'text-xl font-bold',
    lg: 'text-2xl font-bold',
  };
  return map[props.size] || map.md;
});
</script>

<template>
  <div class="space-y-1">
    <!-- Label -->
    <p
      v-if="label"
      class="text-xs font-medium text-slate-400 uppercase tracking-wide"
    >
      {{ label }}
    </p>

    <!-- Primary price -->
    <p :class="priceSizeClasses" style="color: #1a2b4a;">
      {{ formattedPrice }}
    </p>

    <!-- All-inclusive cost -->
    <div
      v-if="allInclusiveCost"
      class="flex items-center gap-1.5"
    >
      <span class="text-xs text-slate-400">All-inclusive:</span>
      <span class="text-sm font-semibold" style="color: #2a9d8f;">
        {{ formattedAllInclusive }}
      </span>
    </div>

    <!-- EMI estimate -->
    <div
      v-if="showEMI && emiDisplay"
      class="flex items-center gap-1.5 mt-1 px-2.5 py-1.5 rounded-md bg-slate-50 border border-slate-100"
    >
      <span class="text-xs text-slate-400">EMI from</span>
      <span class="text-sm font-semibold text-slate-700">
        {{ emiDisplay }}/mo
      </span>
      <span class="text-[10px] text-slate-400">
        ({{ loanPercentage }}% loan, {{ interestRate }}%, {{ tenureYears }}yr)
      </span>
    </div>
  </div>
</template>
