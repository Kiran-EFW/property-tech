<script setup lang="ts">
import { computed } from 'vue';
import { formatCompactINR, formatArea } from '@proptech/shared';
import type { ProjectStatus } from '@proptech/shared';

export interface ProjectCardProps {
  /** Project name */
  name: string;
  /** URL slug for linking */
  slug: string;
  /** Locality / area name */
  locality: string;
  /** City name */
  city?: string;
  /** Thumbnail image URL */
  thumbnailUrl: string | null;
  /** Minimum total price */
  minPrice: number;
  /** Maximum total price (optional — if same as min, shows single price) */
  maxPrice?: number;
  /** Minimum carpet area in sqft */
  minCarpetArea: number;
  /** Maximum carpet area in sqft (optional) */
  maxCarpetArea?: number;
  /** RERA registration number */
  reraNumber: string;
  /** Whether RERA is verified by the platform */
  reraVerified?: boolean;
  /** Project status */
  status: ProjectStatus;
  /** Builder name */
  builderName?: string;
  /** Available configurations (e.g., ["2BHK", "3BHK"]) */
  configurations?: string[];
  /** WhatsApp number for CTA (without +) */
  whatsappNumber?: string;
  /** Pre-filled WhatsApp message */
  whatsappMessage?: string;
}

const props = withDefaults(defineProps<ProjectCardProps>(), {
  city: '',
  maxPrice: 0,
  maxCarpetArea: 0,
  reraVerified: false,
  builderName: '',
  configurations: () => [],
  whatsappNumber: '',
  whatsappMessage: '',
});

const emit = defineEmits<{
  (e: 'click', slug: string): void;
  (e: 'whatsapp', slug: string): void;
}>();

const priceDisplay = computed(() => {
  if (props.maxPrice && props.maxPrice !== props.minPrice) {
    return `${formatCompactINR(props.minPrice)} - ${formatCompactINR(props.maxPrice)}`;
  }
  return formatCompactINR(props.minPrice);
});

const areaDisplay = computed(() => {
  if (props.maxCarpetArea && props.maxCarpetArea !== props.minCarpetArea) {
    return `${formatArea(props.minCarpetArea)} - ${formatArea(props.maxCarpetArea)}`;
  }
  return formatArea(props.minCarpetArea);
});

const statusLabel = computed(() => {
  const map: Record<ProjectStatus, string> = {
    pre_launch: 'Pre-Launch',
    under_construction: 'Under Construction',
    ready_to_move: 'Ready to Move',
  };
  return map[props.status] || props.status;
});

const statusClasses = computed(() => {
  const map: Record<ProjectStatus, string> = {
    pre_launch:
      'bg-amber-50 text-amber-700 border border-amber-200',
    under_construction:
      'bg-blue-50 text-blue-700 border border-blue-200',
    ready_to_move:
      'bg-emerald-50 text-emerald-700 border border-emerald-200',
  };
  return map[props.status] || 'bg-slate-50 text-slate-700 border border-slate-200';
});

const whatsappUrl = computed(() => {
  if (!props.whatsappNumber) return '';
  const message = props.whatsappMessage || `Hi, I'm interested in ${props.name} in ${props.locality}. Please share more details.`;
  return `https://wa.me/${props.whatsappNumber}?text=${encodeURIComponent(message)}`;
});

function handleCardClick() {
  emit('click', props.slug);
}

function handleWhatsAppClick(event: Event) {
  event.stopPropagation();
  emit('whatsapp', props.slug);
}
</script>

<template>
  <article
    class="group bg-white rounded-xl border border-slate-200 shadow-sm overflow-hidden transition-shadow duration-150 hover:shadow-md cursor-pointer"
    role="button"
    :tabindex="0"
    @click="handleCardClick"
    @keydown.enter="handleCardClick"
  >
    <!-- Thumbnail -->
    <div class="relative aspect-[16/10] bg-slate-100 overflow-hidden">
      <img
        v-if="thumbnailUrl"
        :src="thumbnailUrl"
        :alt="`${name} - ${locality}`"
        class="w-full h-full object-cover transition-transform duration-150 group-hover:scale-[1.02]"
        loading="lazy"
      />
      <div
        v-else
        class="w-full h-full flex items-center justify-center"
        style="background-color: #1a2b4a;"
      >
        <span class="text-white/50 text-sm font-medium">No Image</span>
      </div>

      <!-- Status badge (top-left) -->
      <span
        :class="statusClasses"
        class="absolute top-3 left-3 px-2.5 py-1 text-xs font-semibold rounded-md"
      >
        {{ statusLabel }}
      </span>

      <!-- RERA badge (top-right) -->
      <span
        v-if="reraVerified"
        class="absolute top-3 right-3 px-2.5 py-1 text-xs font-semibold rounded-md bg-emerald-50 text-emerald-700 border border-emerald-200"
      >
        RERA Verified
      </span>
    </div>

    <!-- Content -->
    <div class="p-4 space-y-3">
      <!-- Builder name -->
      <p v-if="builderName" class="text-xs font-medium text-slate-400 uppercase tracking-wide">
        {{ builderName }}
      </p>

      <!-- Project name -->
      <h3 class="text-lg font-semibold leading-tight" style="color: #1a2b4a;">
        {{ name }}
      </h3>

      <!-- Location -->
      <p class="text-sm text-slate-500">
        {{ locality }}<span v-if="city">, {{ city }}</span>
      </p>

      <!-- Price & Area -->
      <div class="flex items-baseline justify-between gap-4">
        <div>
          <p class="text-lg font-bold" style="color: #1a2b4a;">
            {{ priceDisplay }}
          </p>
          <p class="text-xs text-slate-400 mt-0.5">onwards</p>
        </div>
        <div class="text-right">
          <p class="text-sm font-medium text-slate-600">
            {{ areaDisplay }}
          </p>
          <p class="text-xs text-slate-400 mt-0.5">carpet area</p>
        </div>
      </div>

      <!-- Configurations -->
      <div v-if="configurations.length" class="flex flex-wrap gap-1.5">
        <span
          v-for="config in configurations"
          :key="config"
          class="px-2 py-0.5 text-xs font-medium rounded bg-slate-100 text-slate-600"
        >
          {{ config }}
        </span>
      </div>

      <!-- RERA number -->
      <p class="text-xs text-slate-400 font-mono truncate" :title="reraNumber">
        RERA: {{ reraNumber }}
      </p>

      <!-- WhatsApp CTA -->
      <a
        v-if="whatsappNumber"
        :href="whatsappUrl"
        target="_blank"
        rel="noopener noreferrer"
        class="flex items-center justify-center gap-2 w-full mt-2 px-4 py-2.5 rounded-lg text-sm font-semibold text-white transition-colors duration-150"
        style="background-color: #2a9d8f;"
        @mouseenter="($event.target as HTMLElement).style.backgroundColor = '#238b7f'"
        @mouseleave="($event.target as HTMLElement).style.backgroundColor = '#2a9d8f'"
        @click="handleWhatsAppClick"
      >
        <svg class="w-4 h-4" viewBox="0 0 24 24" fill="currentColor">
          <path d="M17.472 14.382c-.297-.149-1.758-.867-2.03-.967-.273-.099-.471-.148-.67.15-.197.297-.767.966-.94 1.164-.173.199-.347.223-.644.075-.297-.15-1.255-.463-2.39-1.475-.883-.788-1.48-1.761-1.653-2.059-.173-.297-.018-.458.13-.606.134-.133.298-.347.446-.52.149-.174.198-.298.298-.497.099-.198.05-.371-.025-.52-.075-.149-.669-1.612-.916-2.207-.242-.579-.487-.5-.669-.51-.173-.008-.371-.01-.57-.01-.198 0-.52.074-.792.372-.272.297-1.04 1.016-1.04 2.479 0 1.462 1.065 2.875 1.213 3.074.149.198 2.096 3.2 5.077 4.487.709.306 1.262.489 1.694.625.712.227 1.36.195 1.871.118.571-.085 1.758-.719 2.006-1.413.248-.694.248-1.289.173-1.413-.074-.124-.272-.198-.57-.347m-5.421 7.403h-.004a9.87 9.87 0 01-5.031-1.378l-.361-.214-3.741.982.998-3.648-.235-.374a9.86 9.86 0 01-1.51-5.26c.001-5.45 4.436-9.884 9.888-9.884 2.64 0 5.122 1.03 6.988 2.898a9.825 9.825 0 012.893 6.994c-.003 5.45-4.437 9.884-9.885 9.884m8.413-18.297A11.815 11.815 0 0012.05 0C5.495 0 .16 5.335.157 11.892c0 2.096.547 4.142 1.588 5.945L.057 24l6.305-1.654a11.882 11.882 0 005.683 1.448h.005c6.554 0 11.89-5.335 11.893-11.893a11.821 11.821 0 00-3.48-8.413z"/>
        </svg>
        Chat on WhatsApp
      </a>
    </div>
  </article>
</template>
