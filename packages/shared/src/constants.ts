// ---------------------------------------------------------------------------
// GST Rates on Real Estate
// ---------------------------------------------------------------------------

/** GST rate for under-construction properties (without ITC) */
export const GST_UNDER_CONSTRUCTION = 0.05;

/** GST rate for affordable housing (carpet area <= 60 sqm in metros, price <= 45L) */
export const GST_AFFORDABLE = 0.01;

/** GST rate on brokerage/commission income */
export const GST_ON_BROKERAGE = 0.18;

/** GST rate for ready-to-move-in properties with OC — no GST applicable */
export const GST_READY_TO_MOVE = 0;

export const GST_RATES = {
  underConstruction: GST_UNDER_CONSTRUCTION,
  affordable: GST_AFFORDABLE,
  readyToMove: GST_READY_TO_MOVE,
  brokerage: GST_ON_BROKERAGE,
} as const;

// ---------------------------------------------------------------------------
// Stamp Duty Rates by State — Maharashtra
// ---------------------------------------------------------------------------

export interface StampDutyRate {
  /** Stamp duty rate as a decimal (e.g., 0.05 = 5%) */
  male: number;
  /** Stamp duty rate for female buyers */
  female: number;
  /** Stamp duty rate for joint (male + female) */
  joint: number;
  /** Additional metro cess (applicable in Mumbai, Pune, etc.) */
  metroCess: number;
  /** Local body tax / surcharge */
  lbt: number;
}

/**
 * Maharashtra stamp duty rates.
 * Rates as of 2024-25 — verify against latest government notifications.
 *
 * Within Mumbai Municipal Corporation limits:
 *   Male: 6% (5% stamp duty + 1% metro cess)
 *   Female: 5% (5% stamp duty, metro cess waived for women in some schemes)
 *
 * Outside Mumbai (rest of Maharashtra):
 *   Male: 5%
 *   Female: 4% (1% concession for women)
 */
export const STAMP_DUTY_MAHARASHTRA = {
  mumbai: {
    male: 0.06,
    female: 0.05,
    joint: 0.06,
    metroCess: 0.01,
    lbt: 0.01,
  } satisfies StampDutyRate,

  mumbaiSuburban: {
    male: 0.06,
    female: 0.05,
    joint: 0.06,
    metroCess: 0.01,
    lbt: 0.01,
  } satisfies StampDutyRate,

  thane: {
    male: 0.06,
    female: 0.05,
    joint: 0.06,
    metroCess: 0.01,
    lbt: 0.01,
  } satisfies StampDutyRate,

  naviMumbai: {
    male: 0.06,
    female: 0.05,
    joint: 0.06,
    metroCess: 0.01,
    lbt: 0.01,
  } satisfies StampDutyRate,

  pune: {
    male: 0.06,
    female: 0.05,
    joint: 0.06,
    metroCess: 0.01,
    lbt: 0.01,
  } satisfies StampDutyRate,

  restOfMaharashtra: {
    male: 0.05,
    female: 0.04,
    joint: 0.05,
    metroCess: 0,
    lbt: 0.01,
  } satisfies StampDutyRate,
} as const;

export type MaharashtraRegion = keyof typeof STAMP_DUTY_MAHARASHTRA;

// ---------------------------------------------------------------------------
// Registration Charges
// ---------------------------------------------------------------------------

/** Registration charge rate */
export const REGISTRATION_RATE = 0.01;

/** Maximum registration charge cap in INR (Maharashtra) */
export const REGISTRATION_MAX = 30_000;

/** Registration charge calculation */
export function calculateRegistration(propertyValue: number): number {
  const computed = Math.round(propertyValue * REGISTRATION_RATE);
  return Math.min(computed, REGISTRATION_MAX);
}

// ---------------------------------------------------------------------------
// TDS Rates
// ---------------------------------------------------------------------------

/** TDS rate on property purchase above INR 50L — Section 194-IA */
export const TDS_PROPERTY_PURCHASE = 0.01;

/** Threshold above which TDS on property purchase applies */
export const TDS_PROPERTY_THRESHOLD = 50_00_000;

/** TDS rate on brokerage/commission income — Section 194H */
export const TDS_BROKERAGE = 0.05;

/** TDS rate for NRI property sellers — Section 195 (depends on holding period) */
export const TDS_NRI_LTCG = 0.2;
export const TDS_NRI_STCG = 0.30;

export const TDS_RATES = {
  propertyPurchase: TDS_PROPERTY_PURCHASE,
  propertyThreshold: TDS_PROPERTY_THRESHOLD,
  brokerage: TDS_BROKERAGE,
  nriLongTermCapitalGains: TDS_NRI_LTCG,
  nriShortTermCapitalGains: TDS_NRI_STCG,
} as const;

// ---------------------------------------------------------------------------
// RERA Rules & Compliance Constants
// ---------------------------------------------------------------------------

export const RERA = {
  /** RERA uses carpet area, not super built-up area */
  areaStandard: 'carpet' as const,

  /** Maximum defect liability period post-possession (years) */
  defectLiabilityYears: 5,

  /** Builder must deposit 70% of collections in escrow account */
  escrowPercentage: 0.7,

  /** Projects > 500 sqm or > 8 units must register with RERA */
  registrationThresholdSqm: 500,
  registrationThresholdUnits: 8,

  /** Agent registration fee for individual (INR) */
  agentFeeIndividual: 10_000,

  /** Agent registration fee for entity/company (INR) */
  agentFeeEntity: 50_000,

  /** RERA registration validity period (years) */
  registrationValidityYears: 5,

  /** Quarterly Progress Report — mandatory filing by builders */
  qprFrequency: 'quarterly' as const,

  /** MahaRERA website for verification */
  maharashtraPortalUrl: 'https://maharera.maharashtra.gov.in',

  /** MahaRERA RERA number prefix pattern */
  maharashtraPrefix: 'P52000',
} as const;

// ---------------------------------------------------------------------------
// Lead Pipeline Stages
// ---------------------------------------------------------------------------

export const LEAD_STATUSES = [
  'new',
  'contacted',
  'site_visit',
  'negotiation',
  'booking',
  'won',
  'lost',
] as const;

export type LeadStatus = (typeof LEAD_STATUSES)[number];

// ---------------------------------------------------------------------------
// Project Statuses
// ---------------------------------------------------------------------------

export const PROJECT_STATUSES = [
  'pre_launch',
  'under_construction',
  'ready_to_move',
] as const;

export type ProjectStatus = (typeof PROJECT_STATUSES)[number];

// ---------------------------------------------------------------------------
// Agent Tiers
// ---------------------------------------------------------------------------

export const AGENT_TIERS = ['platinum', 'gold', 'silver', 'new'] as const;

export type AgentTier = (typeof AGENT_TIERS)[number];

// ---------------------------------------------------------------------------
// Booking Statuses
// ---------------------------------------------------------------------------

export const BOOKING_STATUSES = [
  'booked',
  'agreement',
  'construction',
  'possession',
  'cancelled',
] as const;

export type BookingStatus = (typeof BOOKING_STATUSES)[number];

// ---------------------------------------------------------------------------
// Lead Sources
// ---------------------------------------------------------------------------

export const LEAD_SOURCES = [
  'web',
  'whatsapp',
  'referral',
  'social',
  'walk_in',
  'builder',
] as const;

export type LeadSource = (typeof LEAD_SOURCES)[number];

// ---------------------------------------------------------------------------
// Commission Statuses
// ---------------------------------------------------------------------------

export const COMMISSION_STATUSES = [
  'pending',
  'approved',
  'paid',
  'disputed',
] as const;

export type CommissionStatus = (typeof COMMISSION_STATUSES)[number];

// ---------------------------------------------------------------------------
// Payment Plan Types
// ---------------------------------------------------------------------------

export const PAYMENT_PLAN_TYPES = ['clp', 'dpp', 'flexi'] as const;

export type PaymentPlanType = (typeof PAYMENT_PLAN_TYPES)[number];

// ---------------------------------------------------------------------------
// Design Tokens
// ---------------------------------------------------------------------------

export const COLORS = {
  primary: '#1e3a5f', // brand-800  (Deep Blue)
  accent: '#d4932e', // accent-500 (Gold)
  success: '#16a34a',
  warning: '#d97706',
  error: '#dc2626',
  background: '#ffffff',
  surface: '#f8fafc', // slate-50
  border: '#e2e8f0', // slate-200
} as const;

// ---------------------------------------------------------------------------
// Miscellaneous Constants
// ---------------------------------------------------------------------------

/** Lead auto-escalation timeout in minutes */
export const LEAD_ESCALATION_TIMEOUT_MINUTES = 5;

/** Lead ownership expiry in days */
export const LEAD_OWNERSHIP_EXPIRY_DAYS = 90;

/** Default brokerage rate (platform receives from builder) */
export const DEFAULT_BROKERAGE_RATE = 0.03;

/** Agent commission share of platform brokerage */
export const AGENT_COMMISSION_SHARE = 0.5;

/** Affordable housing price threshold (INR) */
export const AFFORDABLE_HOUSING_THRESHOLD = 45_00_000;

/** Affordable housing carpet area threshold in metros (sqm) */
export const AFFORDABLE_HOUSING_AREA_SQMETER = 60;
