import type {
  LeadStatus,
  ProjectStatus,
  AgentTier,
  BookingStatus,
  LeadSource,
  CommissionStatus,
  PaymentPlanType,
} from '@proptech/shared';

// ---------------------------------------------------------------------------
// Common / Shared Types
// ---------------------------------------------------------------------------

/** Standard API timestamp string (ISO 8601) */
export type Timestamp = string;

/** UUID string */
export type UUID = string;

/** Pagination parameters for list endpoints */
export interface PaginationParams {
  page?: number;
  limit?: number;
  sortBy?: string;
  sortOrder?: 'asc' | 'desc';
}

/** Paginated API response wrapper */
export interface PaginatedResponse<T> {
  data: T[];
  meta: {
    page: number;
    limit: number;
    total: number;
    totalPages: number;
  };
}

/** Standard API response wrapper */
export interface ApiResponse<T> {
  data: T;
  message?: string;
}

/** Standard API error */
export interface ApiError {
  statusCode: number;
  message: string;
  errors?: Record<string, string[]>;
}

// ---------------------------------------------------------------------------
// User / Auth
// ---------------------------------------------------------------------------

export type UserRole = 'investor' | 'agent' | 'admin' | 'builder';

export interface User {
  id: UUID;
  phone: string;
  email: string | null;
  name: string;
  role: UserRole;
  avatarUrl: string | null;
  isActive: boolean;
  createdAt: Timestamp;
  updatedAt: Timestamp;
}

export interface AuthTokens {
  accessToken: string;
  refreshToken: string;
  expiresAt: Timestamp;
}

export interface RegisterData {
  phone: string;
  name: string;
  email?: string;
  role: UserRole;
}

export interface LoginResponse {
  user: User;
  tokens: AuthTokens;
}

export interface ProfileUpdateData {
  name?: string;
  email?: string;
  avatarUrl?: string;
}

// ---------------------------------------------------------------------------
// Builder
// ---------------------------------------------------------------------------

export interface Builder {
  id: UUID;
  userId: UUID;
  companyName: string;
  legalName: string;
  reraNumber: string;
  pan: string;
  gstin: string | null;
  phone: string;
  email: string;
  website: string | null;
  logoUrl: string | null;

  /** Address */
  addressLine1: string;
  addressLine2: string | null;
  city: string;
  state: string;
  pincode: string;

  /** Track record */
  establishedYear: number | null;
  totalProjectsDelivered: number;
  ongoingProjects: number;
  averageDeliveryDelayMonths: number | null;

  /** Commission structure */
  defaultBrokerageRate: number;

  isVerified: boolean;
  createdAt: Timestamp;
  updatedAt: Timestamp;
}

// ---------------------------------------------------------------------------
// Project
// ---------------------------------------------------------------------------

export interface ProjectLocation {
  addressLine1: string;
  addressLine2: string | null;
  locality: string;
  city: string;
  state: string;
  pincode: string;
  latitude: number;
  longitude: number;
  microMarket: string;
}

export interface ProjectMedia {
  id: UUID;
  url: string;
  type: 'photo' | 'video' | 'floor_plan' | 'brochure';
  caption: string | null;
  sortOrder: number;
}

export interface ProjectPricingRange {
  minPricePerSqft: number;
  maxPricePerSqft: number;
  minCarpetArea: number;
  maxCarpetArea: number;
  minTotalPrice: number;
  maxTotalPrice: number;
}

export interface PaymentPlanMilestone {
  id: UUID;
  name: string;
  percentage: number;
  description: string | null;
  sortOrder: number;
}

export interface PaymentPlan {
  id: UUID;
  type: PaymentPlanType;
  name: string;
  description: string | null;
  milestones: PaymentPlanMilestone[];
}

export interface Project {
  id: UUID;
  slug: string;
  name: string;
  builderId: UUID;
  builder: Builder | null;

  /** RERA */
  reraNumber: string;
  reraRegistrationDate: Timestamp;
  reraCompletionDate: Timestamp;
  reraVerified: boolean;

  /** Status */
  status: ProjectStatus;
  launchDate: Timestamp | null;
  possessionDate: Timestamp | null;
  constructionProgressPercent: number;

  /** Location */
  location: ProjectLocation;

  /** Description */
  shortDescription: string;
  longDescription: string;

  /** Pricing */
  pricing: ProjectPricingRange;
  floorRisePerSqftPerFloor: number;
  gstRate: number;
  stampDutyRate: number;
  maintenancePerSqft: number;
  legalCharges: number;
  parkingCost: number;

  /** Configuration */
  configurations: string[];
  totalUnits: number;
  totalFloors: number;
  totalTowers: number;

  /** Amenities */
  amenities: string[];

  /** Media */
  media: ProjectMedia[];
  thumbnailUrl: string | null;

  /** Payment plans */
  paymentPlans: PaymentPlan[];

  /** SEO */
  metaTitle: string | null;
  metaDescription: string | null;

  /** Timestamps */
  isActive: boolean;
  createdAt: Timestamp;
  updatedAt: Timestamp;
}

export interface ProjectFilters extends PaginationParams {
  status?: ProjectStatus;
  city?: string;
  locality?: string;
  microMarket?: string;
  minPrice?: number;
  maxPrice?: number;
  minArea?: number;
  maxArea?: number;
  builderId?: UUID;
  configurations?: string[];
  search?: string;
}

export interface ProjectCreateData {
  name: string;
  builderId: UUID;
  reraNumber: string;
  reraRegistrationDate: string;
  reraCompletionDate: string;
  status: ProjectStatus;
  launchDate?: string;
  possessionDate?: string;
  location: ProjectLocation;
  shortDescription: string;
  longDescription: string;
  configurations: string[];
  totalUnits: number;
  totalFloors: number;
  totalTowers: number;
  amenities: string[];
  floorRisePerSqftPerFloor: number;
  gstRate: number;
  stampDutyRate: number;
  maintenancePerSqft: number;
  legalCharges: number;
  parkingCost: number;
}

export interface ProjectUpdateData extends Partial<ProjectCreateData> {
  isActive?: boolean;
  constructionProgressPercent?: number;
  metaTitle?: string;
  metaDescription?: string;
}

// ---------------------------------------------------------------------------
// Inventory (Units)
// ---------------------------------------------------------------------------

export type UnitStatus = 'available' | 'blocked' | 'booked' | 'sold';
export type UnitType = '1BHK' | '2BHK' | '3BHK' | '4BHK' | '5BHK' | 'penthouse' | 'studio' | 'duplex';

export interface Unit {
  id: UUID;
  projectId: UUID;
  towerName: string;
  floorNumber: number;
  unitNumber: string;
  unitType: UnitType;
  carpetAreaSqft: number;
  balconyAreaSqft: number;
  pricePerSqft: number;
  totalPrice: number;
  status: UnitStatus;
  floorPlanUrl: string | null;
  facing: string | null;
  view: string | null;
  createdAt: Timestamp;
  updatedAt: Timestamp;
}

// ---------------------------------------------------------------------------
// Due Diligence
// ---------------------------------------------------------------------------

export interface DueDiligenceReport {
  id: UUID;
  projectId: UUID;

  /** Title verification */
  titleClear: boolean;
  titleRemarks: string | null;

  /** Approvals */
  hasIOD: boolean;
  hasCC: boolean;
  hasOC: boolean;
  hasEnvironmentClearance: boolean;
  approvalRemarks: string | null;

  /** Builder assessment */
  builderTrackRecord: 'excellent' | 'good' | 'average' | 'poor' | 'new';
  builderFinancialHealth: 'strong' | 'moderate' | 'weak' | 'unknown';
  builderRemarks: string | null;

  /** Escrow */
  escrowAccountVerified: boolean;
  escrowBankName: string | null;

  /** Construction */
  constructionQuality: 'premium' | 'standard' | 'basic' | 'not_assessed';
  siteVisitDate: Timestamp | null;
  siteVisitPhotos: string[];
  constructionRemarks: string | null;

  /** Legal */
  legalOpinionAvailable: boolean;
  legalOpinionUrl: string | null;

  /** Overall */
  overallRating: 'recommended' | 'conditional' | 'not_recommended';
  summary: string;

  verifiedBy: UUID;
  verifiedAt: Timestamp;
  createdAt: Timestamp;
  updatedAt: Timestamp;
}

// ---------------------------------------------------------------------------
// Investor
// ---------------------------------------------------------------------------

export interface Investor {
  id: UUID;
  userId: UUID;
  name: string;
  phone: string;
  email: string | null;
  isNRI: boolean;
  nriCountry: string | null;

  /** Investment profile */
  budgetMin: number | null;
  budgetMax: number | null;
  preferredMicroMarkets: string[];
  preferredConfigurations: string[];
  investmentTimeline: 'immediate' | '3_months' | '6_months' | '1_year' | 'exploring';
  purpose: 'investment' | 'end_use' | 'both';

  /** Saved / interested projects */
  savedProjectIds: UUID[];

  createdAt: Timestamp;
  updatedAt: Timestamp;
}

// ---------------------------------------------------------------------------
// Agent
// ---------------------------------------------------------------------------

export interface Agent {
  id: UUID;
  userId: UUID;
  name: string;
  phone: string;
  email: string | null;

  /** RERA */
  reraNumber: string;
  reraVerified: boolean;
  reraExpiryDate: Timestamp | null;

  /** Identity */
  pan: string;
  gstin: string | null;

  /** Profile */
  photoUrl: string | null;
  experienceYears: number;
  specializations: string[];
  operatingAreas: string[];

  /** Performance */
  tier: AgentTier;
  totalLeads: number;
  totalBookings: number;
  conversionRate: number;
  averageResponseTimeMinutes: number;

  /** Status */
  isActive: boolean;
  onboardedAt: Timestamp;
  createdAt: Timestamp;
  updatedAt: Timestamp;
}

export interface AgentRegisterData {
  name: string;
  phone: string;
  email?: string;
  reraNumber: string;
  pan: string;
  gstin?: string;
  experienceYears: number;
  specializations?: string[];
  operatingAreas?: string[];
}

export interface AgentPerformance {
  agentId: UUID;
  period: string;
  totalLeads: number;
  contactedLeads: number;
  siteVisits: number;
  bookings: number;
  revenue: number;
  contactRate: number;
  visitRate: number;
  conversionRate: number;
  averageResponseTimeMinutes: number;
  tier: AgentTier;
}

// ---------------------------------------------------------------------------
// Lead
// ---------------------------------------------------------------------------

export interface LeadNote {
  id: UUID;
  leadId: UUID;
  authorId: UUID;
  authorName: string;
  content: string;
  createdAt: Timestamp;
}

export interface Lead {
  id: UUID;
  investorId: UUID;
  investor: Investor | null;
  agentId: UUID | null;
  agent: Agent | null;
  projectId: UUID | null;
  project: Project | null;

  /** Source & attribution */
  source: LeadSource;
  sourceDetail: string | null;
  utmSource: string | null;
  utmMedium: string | null;
  utmCampaign: string | null;

  /** Pipeline */
  status: LeadStatus;
  statusChangedAt: Timestamp;

  /** Contact info (denormalized for quick access) */
  name: string;
  phone: string;
  email: string | null;

  /** Interest details */
  budget: number | null;
  preferredConfiguration: string | null;
  remarks: string | null;
  isNRI: boolean;

  /** Follow-up */
  nextFollowUpAt: Timestamp | null;
  lastContactedAt: Timestamp | null;
  firstContactedAt: Timestamp | null;

  /** Notes */
  notes: LeadNote[];

  /** Scoring */
  score: number;
  isHot: boolean;

  /** Ownership */
  assignedAt: Timestamp | null;
  ownershipExpiresAt: Timestamp | null;

  createdAt: Timestamp;
  updatedAt: Timestamp;
}

export interface LeadCreateData {
  name: string;
  phone: string;
  email?: string;
  projectId?: UUID;
  source: LeadSource;
  sourceDetail?: string;
  budget?: number;
  preferredConfiguration?: string;
  remarks?: string;
  isNRI?: boolean;
  utmSource?: string;
  utmMedium?: string;
  utmCampaign?: string;
}

export interface LeadFilters extends PaginationParams {
  status?: LeadStatus;
  agentId?: UUID;
  projectId?: UUID;
  source?: LeadSource;
  isNRI?: boolean;
  isHot?: boolean;
  search?: string;
  dateFrom?: string;
  dateTo?: string;
}

// ---------------------------------------------------------------------------
// Site Visit
// ---------------------------------------------------------------------------

export interface SiteVisit {
  id: UUID;
  leadId: UUID;
  lead: Lead | null;
  agentId: UUID;
  agent: Agent | null;
  projectId: UUID;
  project: Project | null;
  investorId: UUID;

  /** Scheduling */
  scheduledAt: Timestamp;
  startedAt: Timestamp | null;
  completedAt: Timestamp | null;
  status: 'scheduled' | 'in_progress' | 'completed' | 'cancelled' | 'no_show';

  /** Feedback */
  investorFeedback: string | null;
  investorInterestLevel: 'high' | 'medium' | 'low' | 'not_interested' | null;
  agentNotes: string | null;
  photos: string[];
  videoUrl: string | null;

  /** Outcome */
  outcome: 'interested' | 'follow_up' | 'not_interested' | 'booked' | null;
  nextSteps: string | null;

  createdAt: Timestamp;
  updatedAt: Timestamp;
}

// ---------------------------------------------------------------------------
// Booking
// ---------------------------------------------------------------------------

export interface Booking {
  id: UUID;
  leadId: UUID;
  lead: Lead | null;
  projectId: UUID;
  project: Project | null;
  unitId: UUID;
  unit: Unit | null;
  investorId: UUID;
  investor: Investor | null;
  agentId: UUID;
  agent: Agent | null;

  /** Financial */
  agreementValue: number;
  bookingAmount: number;
  stampDutyAmount: number;
  registrationAmount: number;
  gstAmount: number;
  totalAllInclusiveAmount: number;

  /** Payment plan */
  paymentPlanType: PaymentPlanType;
  paymentMilestones: BookingPaymentMilestone[];

  /** Status */
  status: BookingStatus;
  bookingDate: Timestamp;
  agreementDate: Timestamp | null;
  possessionDate: Timestamp | null;
  cancellationDate: Timestamp | null;
  cancellationReason: string | null;

  /** Documents */
  bookingFormUrl: string | null;
  agreementUrl: string | null;

  createdAt: Timestamp;
  updatedAt: Timestamp;
}

export interface BookingPaymentMilestone {
  id: UUID;
  bookingId: UUID;
  name: string;
  amount: number;
  percentage: number;
  dueDate: Timestamp | null;
  paidDate: Timestamp | null;
  status: 'pending' | 'due' | 'paid' | 'overdue';
  sortOrder: number;
}

// ---------------------------------------------------------------------------
// Commission
// ---------------------------------------------------------------------------

export interface Commission {
  id: UUID;
  bookingId: UUID;
  booking: Booking | null;
  agentId: UUID;
  agent: Agent | null;
  projectId: UUID;

  /** Amounts */
  agreementValue: number;
  brokerageRate: number;
  totalBrokerage: number;
  agentShareRate: number;
  agentCommission: number;
  platformCommission: number;
  tdsAmount: number;
  gstAmount: number;
  netPayableToAgent: number;

  /** Status */
  status: CommissionStatus;
  approvedBy: UUID | null;
  approvedAt: Timestamp | null;
  paidAt: Timestamp | null;
  paymentReference: string | null;

  /** Milestones */
  bookingMilestonePaid: boolean;
  agreementMilestonePaid: boolean;
  collectionMilestonePaid: boolean;

  remarks: string | null;
  createdAt: Timestamp;
  updatedAt: Timestamp;
}

// ---------------------------------------------------------------------------
// Area (Micro-market content)
// ---------------------------------------------------------------------------

export interface AreaPriceTrend {
  month: string;
  averagePricePerSqft: number;
  totalTransactions: number;
}

export interface AreaInfrastructure {
  name: string;
  type: 'metro' | 'road' | 'airport' | 'railway' | 'hospital' | 'school' | 'mall' | 'park';
  status: 'operational' | 'under_construction' | 'planned';
  distanceKm: number;
  expectedCompletionDate: string | null;
  description: string | null;
}

export interface Area {
  id: UUID;
  slug: string;
  name: string;
  city: string;
  state: string;

  /** Content */
  title: string;
  summary: string;
  description: string;

  /** Pricing data */
  averagePricePerSqft: number;
  priceGrowthYoY: number;
  priceTrends: AreaPriceTrend[];

  /** Infrastructure */
  infrastructure: AreaInfrastructure[];

  /** Related projects */
  nearbyProjectIds: UUID[];

  /** Location */
  latitude: number;
  longitude: number;

  /** Media */
  coverImageUrl: string | null;

  /** SEO */
  metaTitle: string | null;
  metaDescription: string | null;

  isActive: boolean;
  createdAt: Timestamp;
  updatedAt: Timestamp;
}

// ---------------------------------------------------------------------------
// Event (Audit Log)
// ---------------------------------------------------------------------------

export type EventAction =
  | 'lead_created'
  | 'lead_assigned'
  | 'lead_status_changed'
  | 'lead_note_added'
  | 'site_visit_scheduled'
  | 'site_visit_completed'
  | 'booking_created'
  | 'booking_status_changed'
  | 'booking_cancelled'
  | 'commission_calculated'
  | 'commission_approved'
  | 'commission_paid'
  | 'commission_disputed'
  | 'project_created'
  | 'project_updated'
  | 'agent_onboarded'
  | 'agent_tier_changed'
  | 'agent_rera_verified'
  | 'user_registered'
  | 'user_logged_in';

export interface Event {
  id: UUID;
  actorId: UUID;
  actorRole: UserRole;
  action: EventAction;
  entityType: string;
  entityId: UUID;
  payload: Record<string, unknown>;
  ipAddress: string | null;
  userAgent: string | null;
  createdAt: Timestamp;
}
