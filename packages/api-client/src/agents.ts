import type { ApiClient } from './client';
import type {
  Agent,
  AgentRegisterData,
  AgentPerformance,
  PaginatedResponse,
  PaginationParams,
} from './types';
import type { AgentTier } from '@proptech/shared';

/**
 * API client for agent management endpoints.
 *
 * Used by:
 *   - Agent dashboard: registration, profile, performance
 *   - Admin dashboard: agent management, tier assignment, verification
 */
export class AgentsAPI {
  constructor(private client: ApiClient) {}

  /**
   * Registers a new agent on the platform.
   * The agent must provide their RERA number for verification.
   * Agent starts at the 'new' tier.
   *
   * POST /agents/register
   */
  async registerAgent(data: AgentRegisterData): Promise<Agent> {
    return this.client.post<Agent>('/agents/register', data);
  }

  /**
   * Lists all agents with pagination. Admin only.
   *
   * GET /agents
   */
  async listAgents(
    params?: PaginationParams & {
      tier?: AgentTier;
      isActive?: boolean;
      search?: string;
    },
  ): Promise<PaginatedResponse<Agent>> {
    const queryParams: Record<string, string | number | boolean | undefined> =
      {};

    if (params) {
      if (params.page) queryParams.page = params.page;
      if (params.limit) queryParams.limit = params.limit;
      if (params.sortBy) queryParams.sort_by = params.sortBy;
      if (params.sortOrder) queryParams.sort_order = params.sortOrder;
      if (params.tier) queryParams.tier = params.tier;
      if (params.isActive !== undefined) queryParams.is_active = params.isActive;
      if (params.search) queryParams.search = params.search;
    }

    return this.client.get<PaginatedResponse<Agent>>('/agents', queryParams);
  }

  /**
   * Gets a single agent by ID.
   *
   * GET /agents/:id
   */
  async getAgent(id: string): Promise<Agent> {
    return this.client.get<Agent>(`/agents/${id}`);
  }

  /**
   * Gets performance metrics for an agent.
   * Agents can view their own performance.
   * Admins can view any agent's performance.
   *
   * GET /agents/:id/performance
   */
  async getPerformance(
    id: string,
    period?: string,
  ): Promise<AgentPerformance> {
    const params: Record<string, string | number | boolean | undefined> = {};
    if (period) params.period = period;

    return this.client.get<AgentPerformance>(
      `/agents/${id}/performance`,
      params,
    );
  }

  /**
   * Updates an agent's tier. Admin only.
   * Tier determines lead allocation priority.
   *
   * PUT /agents/:id/tier
   */
  async updateTier(id: string, tier: AgentTier): Promise<Agent> {
    return this.client.put<Agent>(`/agents/${id}/tier`, { tier });
  }

  /**
   * Marks an agent's RERA number as verified. Admin only.
   *
   * PUT /agents/:id/verify-rera
   */
  async verifyRERA(id: string): Promise<Agent> {
    return this.client.put<Agent>(`/agents/${id}/verify-rera`, {
      rera_verified: true,
    });
  }

  /**
   * Deactivates an agent. Admin only.
   * Deactivated agents cannot receive new leads.
   *
   * PUT /agents/:id/deactivate
   */
  async deactivateAgent(id: string): Promise<Agent> {
    return this.client.put<Agent>(`/agents/${id}/deactivate`, {
      is_active: false,
    });
  }

  /**
   * Reactivates a previously deactivated agent. Admin only.
   *
   * PUT /agents/:id/activate
   */
  async activateAgent(id: string): Promise<Agent> {
    return this.client.put<Agent>(`/agents/${id}/activate`, {
      is_active: true,
    });
  }
}
