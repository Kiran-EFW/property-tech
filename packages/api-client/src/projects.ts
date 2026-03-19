import type { ApiClient } from './client';
import type {
  Project,
  ProjectFilters,
  ProjectCreateData,
  ProjectUpdateData,
  Unit,
  DueDiligenceReport,
  PaginatedResponse,
} from './types';

/**
 * API client for project-related endpoints.
 *
 * Handles both public project browsing (investor web) and
 * admin CRUD operations (dashboard).
 */
export class ProjectsAPI {
  constructor(private client: ApiClient) {}

  /**
   * Lists projects with optional filters and pagination.
   * Public endpoint — used by investor web for browsing.
   *
   * GET /projects
   */
  async listProjects(
    filters?: ProjectFilters,
  ): Promise<PaginatedResponse<Project>> {
    const params: Record<string, string | number | boolean | undefined> = {};

    if (filters) {
      if (filters.page) params.page = filters.page;
      if (filters.limit) params.limit = filters.limit;
      if (filters.sortBy) params.sort_by = filters.sortBy;
      if (filters.sortOrder) params.sort_order = filters.sortOrder;
      if (filters.status) params.status = filters.status;
      if (filters.city) params.city = filters.city;
      if (filters.locality) params.locality = filters.locality;
      if (filters.microMarket) params.micro_market = filters.microMarket;
      if (filters.minPrice !== undefined) params.min_price = filters.minPrice;
      if (filters.maxPrice !== undefined) params.max_price = filters.maxPrice;
      if (filters.minArea !== undefined) params.min_area = filters.minArea;
      if (filters.maxArea !== undefined) params.max_area = filters.maxArea;
      if (filters.builderId) params.builder_id = filters.builderId;
      if (filters.search) params.search = filters.search;
      if (filters.configurations?.length) {
        params.configurations = filters.configurations.join(',');
      }
    }

    return this.client.get<PaginatedResponse<Project>>('/projects', params);
  }

  /**
   * Gets a single project by its URL slug.
   * Public endpoint — used for project detail pages.
   *
   * GET /projects/:slug
   */
  async getProject(slug: string): Promise<Project> {
    return this.client.get<Project>(`/projects/${encodeURIComponent(slug)}`);
  }

  /**
   * Creates a new project. Admin only.
   *
   * POST /projects
   */
  async createProject(data: ProjectCreateData): Promise<Project> {
    return this.client.post<Project>('/projects', data);
  }

  /**
   * Updates an existing project. Admin only.
   *
   * PUT /projects/:id
   */
  async updateProject(id: string, data: ProjectUpdateData): Promise<Project> {
    return this.client.put<Project>(`/projects/${id}`, data);
  }

  /**
   * Gets the unit inventory for a project.
   * Shows available, booked, and sold units with pricing.
   *
   * GET /projects/:id/inventory
   */
  async getInventory(projectId: string): Promise<Unit[]> {
    return this.client.get<Unit[]>(`/projects/${projectId}/inventory`);
  }

  /**
   * Gets the due diligence report for a project.
   * Public endpoint — shows verification status to investors.
   *
   * GET /projects/:id/due-diligence
   */
  async getDueDiligence(projectId: string): Promise<DueDiligenceReport> {
    return this.client.get<DueDiligenceReport>(
      `/projects/${projectId}/due-diligence`,
    );
  }
}
