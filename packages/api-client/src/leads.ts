import type { ApiClient } from './client';
import type {
  Lead,
  LeadCreateData,
  LeadFilters,
  LeadNote,
  PaginatedResponse,
} from './types';
import type { LeadStatus } from '@proptech/shared';

/**
 * API client for lead management endpoints.
 *
 * Used by:
 *   - Investor web: creating leads (express interest)
 *   - Agent dashboard: viewing/updating assigned leads
 *   - Admin dashboard: full pipeline management
 */
export class LeadsAPI {
  constructor(private client: ApiClient) {}

  /**
   * Creates a new lead (investor expresses interest).
   * Public endpoint — can be called without authentication.
   * Triggers auto-assignment to an agent via background job.
   *
   * POST /leads
   */
  async createLead(data: LeadCreateData): Promise<Lead> {
    return this.client.post<Lead>('/leads', data);
  }

  /**
   * Lists leads with filters and pagination.
   *
   * Agents see only their assigned leads.
   * Admins see all leads across agents.
   * Role-based filtering is handled server-side based on JWT.
   *
   * GET /leads
   */
  async listLeads(filters?: LeadFilters): Promise<PaginatedResponse<Lead>> {
    const params: Record<string, string | number | boolean | undefined> = {};

    if (filters) {
      if (filters.page) params.page = filters.page;
      if (filters.limit) params.limit = filters.limit;
      if (filters.sortBy) params.sort_by = filters.sortBy;
      if (filters.sortOrder) params.sort_order = filters.sortOrder;
      if (filters.status) params.status = filters.status;
      if (filters.agentId) params.agent_id = filters.agentId;
      if (filters.projectId) params.project_id = filters.projectId;
      if (filters.source) params.source = filters.source;
      if (filters.isNRI !== undefined) params.is_nri = filters.isNRI;
      if (filters.isHot !== undefined) params.is_hot = filters.isHot;
      if (filters.search) params.search = filters.search;
      if (filters.dateFrom) params.date_from = filters.dateFrom;
      if (filters.dateTo) params.date_to = filters.dateTo;
    }

    return this.client.get<PaginatedResponse<Lead>>('/leads', params);
  }

  /**
   * Gets a single lead by ID.
   *
   * GET /leads/:id
   */
  async getLead(id: string): Promise<Lead> {
    return this.client.get<Lead>(`/leads/${id}`);
  }

  /**
   * Updates the pipeline status of a lead.
   * Agents can update their own leads' status.
   * Admins can update any lead's status.
   *
   * PUT /leads/:id/status
   */
  async updateLeadStatus(
    id: string,
    status: LeadStatus,
    remarks?: string,
  ): Promise<Lead> {
    return this.client.put<Lead>(`/leads/${id}/status`, { status, remarks });
  }

  /**
   * Adds a note to a lead's communication history.
   *
   * POST /leads/:id/notes
   */
  async addNote(id: string, content: string): Promise<LeadNote> {
    return this.client.post<LeadNote>(`/leads/${id}/notes`, { content });
  }

  /**
   * Assigns a lead to a specific agent. Admin only.
   * Overrides any existing assignment.
   *
   * PUT /leads/:id/assign
   */
  async assignLead(id: string, agentId: string): Promise<Lead> {
    return this.client.put<Lead>(`/leads/${id}/assign`, {
      agent_id: agentId,
    });
  }

  /**
   * Updates the next follow-up date for a lead.
   *
   * PUT /leads/:id/follow-up
   */
  async setFollowUp(id: string, followUpAt: string): Promise<Lead> {
    return this.client.put<Lead>(`/leads/${id}/follow-up`, {
      next_follow_up_at: followUpAt,
    });
  }
}
