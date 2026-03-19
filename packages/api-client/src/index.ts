// @proptech/api-client — Typed HTTP client for the PropTech Go backend API

// Core client
export { ApiClient, ApiClientError } from './client';
export type { ApiClientConfig } from './client';

// API modules
export { AuthAPI } from './auth';
export { ProjectsAPI } from './projects';
export { LeadsAPI } from './leads';
export { AgentsAPI } from './agents';

// Types — re-export everything for consuming apps
export type {
  // Common
  Timestamp,
  UUID,
  PaginationParams,
  PaginatedResponse,
  ApiResponse,
  ApiError,

  // Auth
  UserRole,
  User,
  AuthTokens,
  RegisterData,
  LoginResponse,
  ProfileUpdateData,

  // Builder
  Builder,

  // Project
  ProjectLocation,
  ProjectMedia,
  ProjectPricingRange,
  PaymentPlanMilestone,
  PaymentPlan,
  Project,
  ProjectFilters,
  ProjectCreateData,
  ProjectUpdateData,

  // Inventory
  UnitStatus,
  UnitType,
  Unit,

  // Due Diligence
  DueDiligenceReport,

  // Investor
  Investor,

  // Agent
  Agent,
  AgentRegisterData,
  AgentPerformance,

  // Lead
  LeadNote,
  Lead,
  LeadCreateData,
  LeadFilters,

  // Site Visit
  SiteVisit,

  // Booking
  Booking,
  BookingPaymentMilestone,

  // Commission
  Commission,

  // Area
  AreaPriceTrend,
  AreaInfrastructure,
  Area,

  // Event
  EventAction,
  Event,
} from './types';

// ---------------------------------------------------------------------------
// Factory — convenience function to create a fully configured client
// ---------------------------------------------------------------------------

import { ApiClient } from './client';
import type { ApiClientConfig } from './client';
import { AuthAPI } from './auth';
import { ProjectsAPI } from './projects';
import { LeadsAPI } from './leads';
import { AgentsAPI } from './agents';

export interface PropTechAPI {
  client: ApiClient;
  auth: AuthAPI;
  projects: ProjectsAPI;
  leads: LeadsAPI;
  agents: AgentsAPI;
}

/**
 * Creates a fully configured PropTech API client with all modules.
 *
 * @example
 * ```ts
 * const api = createPropTechAPI({
 *   baseURL: 'https://api.example.com/v1',
 *   onUnauthorized: () => navigateTo('/login'),
 * });
 *
 * // Auth
 * await api.auth.requestOTP('+919876543210');
 * const { user, tokens } = await api.auth.login('+919876543210', '123456');
 *
 * // Projects
 * const projects = await api.projects.listProjects({ status: 'under_construction' });
 *
 * // Leads
 * const lead = await api.leads.createLead({ name: 'John', phone: '9876543210', source: 'web' });
 * ```
 */
export function createPropTechAPI(config: ApiClientConfig): PropTechAPI {
  const client = new ApiClient(config);

  return {
    client,
    auth: new AuthAPI(client),
    projects: new ProjectsAPI(client),
    leads: new LeadsAPI(client),
    agents: new AgentsAPI(client),
  };
}
