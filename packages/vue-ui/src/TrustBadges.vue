<script setup lang="ts">
export interface TrustBadge {
  /** Unique identifier */
  id: string;
  /** Badge label */
  label: string;
  /** Badge type determines the icon and color */
  type: 'rera_verified' | 'title_clear' | 'builder_track_record' | 'construction_on_track' | 'escrow_verified' | 'legal_clear';
  /** Whether this badge is active/verified */
  verified: boolean;
  /** Optional tooltip/description */
  description?: string;
}

export interface TrustBadgesProps {
  /** Array of badge objects to display */
  badges: TrustBadge[];
  /** Layout direction */
  direction?: 'horizontal' | 'vertical';
  /** Size variant */
  size?: 'sm' | 'md';
}

const props = withDefaults(defineProps<TrustBadgesProps>(), {
  direction: 'horizontal',
  size: 'md',
});

function getBadgeIcon(type: TrustBadge['type']): string {
  const icons: Record<TrustBadge['type'], string> = {
    rera_verified: 'M9 12l2 2 4-4m5.618-4.016A11.955 11.955 0 0112 2.944a11.955 11.955 0 01-8.618 3.04A12.02 12.02 0 003 9c0 5.591 3.824 10.29 9 11.622 5.176-1.332 9-6.03 9-11.622 0-1.042-.133-2.052-.382-3.016z',
    title_clear: 'M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z',
    builder_track_record: 'M19 21V5a2 2 0 00-2-2H7a2 2 0 00-2 2v16m14 0h2m-2 0h-5m-9 0H3m2 0h5M9 7h1m-1 4h1m4-4h1m-1 4h1m-5 10v-5a1 1 0 011-1h2a1 1 0 011 1v5m-4 0h4',
    construction_on_track: 'M13 10V3L4 14h7v7l9-11h-7z',
    escrow_verified: 'M12 15v2m-6 4h12a2 2 0 002-2v-6a2 2 0 00-2-2H6a2 2 0 00-2 2v6a2 2 0 002 2zm10-10V7a4 4 0 00-8 0v4h8z',
    legal_clear: 'M9 12h6m-6 4h6m2 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z',
  };
  return icons[type] || icons.rera_verified;
}

function getBadgeColors(type: TrustBadge['type'], verified: boolean): { bg: string; text: string; icon: string; border: string } {
  if (!verified) {
    return {
      bg: 'bg-slate-50',
      text: 'text-slate-400',
      icon: 'text-slate-300',
      border: 'border-slate-200',
    };
  }

  const colorMap: Record<TrustBadge['type'], { bg: string; text: string; icon: string; border: string }> = {
    rera_verified: {
      bg: 'bg-emerald-50',
      text: 'text-emerald-700',
      icon: 'text-emerald-500',
      border: 'border-emerald-200',
    },
    title_clear: {
      bg: 'bg-emerald-50',
      text: 'text-emerald-700',
      icon: 'text-emerald-500',
      border: 'border-emerald-200',
    },
    builder_track_record: {
      bg: 'bg-blue-50',
      text: 'text-blue-700',
      icon: 'text-blue-500',
      border: 'border-blue-200',
    },
    construction_on_track: {
      bg: 'bg-teal-50',
      text: 'text-teal-700',
      icon: 'text-teal-500',
      border: 'border-teal-200',
    },
    escrow_verified: {
      bg: 'bg-indigo-50',
      text: 'text-indigo-700',
      icon: 'text-indigo-500',
      border: 'border-indigo-200',
    },
    legal_clear: {
      bg: 'bg-emerald-50',
      text: 'text-emerald-700',
      icon: 'text-emerald-500',
      border: 'border-emerald-200',
    },
  };

  return colorMap[type] || colorMap.rera_verified;
}
</script>

<template>
  <div
    :class="[
      'flex gap-2',
      direction === 'horizontal' ? 'flex-row flex-wrap' : 'flex-col',
    ]"
  >
    <div
      v-for="badge in badges"
      :key="badge.id"
      :class="[
        'flex items-center gap-2 rounded-lg border transition-colors duration-150',
        getBadgeColors(badge.type, badge.verified).bg,
        getBadgeColors(badge.type, badge.verified).border,
        size === 'sm' ? 'px-2.5 py-1.5' : 'px-3 py-2',
      ]"
      :title="badge.description || badge.label"
    >
      <!-- Icon -->
      <svg
        :class="[
          getBadgeColors(badge.type, badge.verified).icon,
          size === 'sm' ? 'w-3.5 h-3.5' : 'w-4 h-4',
        ]"
        fill="none"
        viewBox="0 0 24 24"
        stroke="currentColor"
        stroke-width="2"
        stroke-linecap="round"
        stroke-linejoin="round"
      >
        <path :d="getBadgeIcon(badge.type)" />
      </svg>

      <!-- Label -->
      <span
        :class="[
          'font-medium whitespace-nowrap',
          getBadgeColors(badge.type, badge.verified).text,
          size === 'sm' ? 'text-xs' : 'text-sm',
        ]"
      >
        {{ badge.label }}
      </span>

      <!-- Unverified indicator -->
      <span
        v-if="!badge.verified"
        :class="[
          'text-slate-400 italic',
          size === 'sm' ? 'text-[10px]' : 'text-xs',
        ]"
      >
        Pending
      </span>
    </div>
  </div>
</template>
