// @proptech/shared — Pure TypeScript utilities for the PropTech platform
// No framework dependencies. Used by both Nuxt apps and potentially the backend.

export {
  // GST
  GST_UNDER_CONSTRUCTION,
  GST_AFFORDABLE,
  GST_ON_BROKERAGE,
  GST_READY_TO_MOVE,
  GST_RATES,

  // Stamp Duty
  STAMP_DUTY_MAHARASHTRA,

  // Registration
  REGISTRATION_RATE,
  REGISTRATION_MAX,
  calculateRegistration,

  // TDS
  TDS_PROPERTY_PURCHASE,
  TDS_PROPERTY_THRESHOLD,
  TDS_BROKERAGE,
  TDS_NRI_LTCG,
  TDS_NRI_STCG,
  TDS_RATES,

  // RERA
  RERA,

  // Enums / Status arrays
  LEAD_STATUSES,
  PROJECT_STATUSES,
  AGENT_TIERS,
  BOOKING_STATUSES,
  LEAD_SOURCES,
  COMMISSION_STATUSES,
  PAYMENT_PLAN_TYPES,

  // Design tokens
  COLORS,

  // Miscellaneous
  LEAD_ESCALATION_TIMEOUT_MINUTES,
  LEAD_OWNERSHIP_EXPIRY_DAYS,
  DEFAULT_BROKERAGE_RATE,
  AGENT_COMMISSION_SHARE,
  AFFORDABLE_HOUSING_THRESHOLD,
  AFFORDABLE_HOUSING_AREA_SQMETER,
} from './constants';

export type {
  StampDutyRate,
  MaharashtraRegion,
  LeadStatus,
  ProjectStatus,
  AgentTier,
  BookingStatus,
  LeadSource,
  CommissionStatus,
  PaymentPlanType,
} from './constants';

export {
  calculateAllInclusiveCost,
  calculateEMI,
} from './calculator';

export type {
  CostCalculatorInput,
  CostBreakdownItem,
  CostBreakdown,
} from './calculator';

export {
  formatINR,
  formatCrores,
  formatLakhs,
  formatCompactINR,
  formatArea,
  sqftToSqm,
  sqmToSqft,
  formatPhone,
  formatPhoneDisplay,
  formatPricePerSqft,
} from './formatting';

export {
  validatePhone,
  validateRERA,
  validatePAN,
  validateGST,
  validateEmail,
  validateAadhaar,
} from './validation';

export type { ValidationResult } from './validation';
