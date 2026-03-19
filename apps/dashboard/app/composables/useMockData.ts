/**
 * Centralised mock data for the PropTech dashboard demo.
 * Uses realistic Indian names, Mumbai-region locations, and INR amounts.
 */
export function useMockData() {
  // ---- Agents ----------------------------------------------------------------
  const agents = [
    {
      id: 'agt-001', name: 'Rajesh Sharma', phone: '9876543210', email: 'rajesh.sharma@gmail.com',
      reraNumber: 'A51900045678', reraVerified: true, pan: 'ABCPS1234K', tier: 'gold' as const,
      totalLeads: 42, totalBookings: 8, conversionRate: 19.0, experienceYears: 6,
      operatingAreas: ['Panvel', 'Ulwe', 'Kharghar'], isActive: true, status: 'active' as const,
      onboardedAt: '2024-06-15T10:00:00Z',
    },
    {
      id: 'agt-002', name: 'Priya Patel', phone: '9823456789', email: 'priya.patel@gmail.com',
      reraNumber: 'A51900056789', reraVerified: true, pan: 'DEFPP5678L', tier: 'silver' as const,
      totalLeads: 28, totalBookings: 5, conversionRate: 17.8, experienceYears: 4,
      operatingAreas: ['Dombivli', 'Kalyan', 'Thane'], isActive: true, status: 'active' as const,
      onboardedAt: '2024-08-20T10:00:00Z',
    },
    {
      id: 'agt-003', name: 'Amit Deshmukh', phone: '9834567890', email: 'amit.deshmukh@gmail.com',
      reraNumber: 'A51900067890', reraVerified: true, pan: 'GHIAD7890M', tier: 'platinum' as const,
      totalLeads: 65, totalBookings: 15, conversionRate: 23.1, experienceYears: 9,
      operatingAreas: ['Panvel', 'Kharghar', 'Vashi'], isActive: true, status: 'active' as const,
      onboardedAt: '2024-03-10T10:00:00Z',
    },
    {
      id: 'agt-004', name: 'Sneha Kulkarni', phone: '9845678901', email: 'sneha.k@gmail.com',
      reraNumber: 'A51900078901', reraVerified: true, pan: 'JKLSK0123N', tier: 'silver' as const,
      totalLeads: 31, totalBookings: 6, conversionRate: 19.4, experienceYears: 3,
      operatingAreas: ['Ulwe', 'Panvel'], isActive: true, status: 'active' as const,
      onboardedAt: '2024-09-05T10:00:00Z',
    },
    {
      id: 'agt-005', name: 'Vikram Joshi', phone: '9856789012', email: 'vikram.joshi@gmail.com',
      reraNumber: 'A51900089012', reraVerified: false, pan: 'MNOPV3456O', tier: 'bronze' as const,
      totalLeads: 12, totalBookings: 1, conversionRate: 8.3, experienceYears: 1,
      operatingAreas: ['Kalyan', 'Dombivli'], isActive: true, status: 'active' as const,
      onboardedAt: '2025-01-10T10:00:00Z',
    },
    {
      id: 'agt-006', name: 'Meera Nair', phone: '9867890123', email: 'meera.nair@gmail.com',
      reraNumber: 'A51900090123', reraVerified: true, pan: 'QRSTM6789P', tier: 'gold' as const,
      totalLeads: 38, totalBookings: 9, conversionRate: 23.7, experienceYears: 5,
      operatingAreas: ['Kharghar', 'Vashi', 'Belapur'], isActive: true, status: 'active' as const,
      onboardedAt: '2024-05-22T10:00:00Z',
    },
    {
      id: 'agt-007', name: 'Suresh Gupta', phone: '9878901234', email: 'suresh.gupta@gmail.com',
      reraNumber: 'A51900001234', reraVerified: false, pan: 'UVWSG9012Q', tier: 'bronze' as const,
      totalLeads: 5, totalBookings: 0, conversionRate: 0, experienceYears: 0,
      operatingAreas: ['Panvel'], isActive: false, status: 'pending' as const,
      onboardedAt: '2025-03-01T10:00:00Z',
    },
    {
      id: 'agt-008', name: 'Anita Bhosale', phone: '9889012345', email: 'anita.b@gmail.com',
      reraNumber: 'A51900012345', reraVerified: false, pan: 'XYZAB2345R', tier: 'bronze' as const,
      totalLeads: 3, totalBookings: 0, conversionRate: 0, experienceYears: 1,
      operatingAreas: ['Ulwe', 'Panvel'], isActive: false, status: 'pending' as const,
      onboardedAt: '2025-03-10T10:00:00Z',
    },
  ]

  // ---- Projects ---------------------------------------------------------------
  const projects = [
    {
      id: 'prj-001', name: 'Arihant Aspire', builder: 'Arihant Superstructures', location: 'Panvel',
      reraNumber: 'P52000048573', status: 'active' as const, totalUnits: 240, availableUnits: 85,
      configurations: ['1 BHK', '2 BHK', '3 BHK'], priceRange: { min: 4500000, max: 12500000 },
      pricePerSqft: 7200, amenities: ['Swimming Pool', 'Gym', 'Clubhouse', 'Garden', 'Parking'],
      possessionDate: '2027-06-01', constructionProgress: 45, createdAt: '2024-08-15T10:00:00Z',
    },
    {
      id: 'prj-002', name: 'Balaji Symphony', builder: 'Balaji Developers', location: 'Ulwe',
      reraNumber: 'P52000052891', status: 'active' as const, totalUnits: 180, availableUnits: 62,
      configurations: ['1 BHK', '2 BHK'], priceRange: { min: 3800000, max: 8500000 },
      pricePerSqft: 6800, amenities: ['Gym', 'Children Play Area', 'Garden', 'Parking'],
      possessionDate: '2026-12-01', constructionProgress: 68, createdAt: '2024-05-20T10:00:00Z',
    },
    {
      id: 'prj-003', name: 'Paradise Heights', builder: 'Paradise Group', location: 'Kharghar',
      reraNumber: 'P52000039145', status: 'active' as const, totalUnits: 320, availableUnits: 42,
      configurations: ['2 BHK', '3 BHK', '4 BHK'], priceRange: { min: 7500000, max: 22000000 },
      pricePerSqft: 9500, amenities: ['Swimming Pool', 'Gym', 'Clubhouse', 'Tennis Court', 'Jogging Track', 'Garden'],
      possessionDate: '2026-09-01', constructionProgress: 82, createdAt: '2024-02-10T10:00:00Z',
    },
    {
      id: 'prj-004', name: 'Sai Residency', builder: 'Sai Builders', location: 'Dombivli',
      reraNumber: 'P52000061234', status: 'active' as const, totalUnits: 120, availableUnits: 95,
      configurations: ['1 BHK', '2 BHK'], priceRange: { min: 3200000, max: 7200000 },
      pricePerSqft: 5800, amenities: ['Garden', 'Parking', 'Children Play Area'],
      possessionDate: '2028-03-01', constructionProgress: 15, createdAt: '2025-01-05T10:00:00Z',
    },
    {
      id: 'prj-005', name: 'Green Valley', builder: 'Godrej Properties', location: 'Kalyan',
      reraNumber: 'P52000055678', status: 'active' as const, totalUnits: 400, availableUnits: 180,
      configurations: ['1 BHK', '2 BHK', '3 BHK'], priceRange: { min: 4000000, max: 11000000 },
      pricePerSqft: 6200, amenities: ['Swimming Pool', 'Gym', 'Clubhouse', 'Garden', 'Sports Complex'],
      possessionDate: '2027-12-01', constructionProgress: 30, createdAt: '2024-11-15T10:00:00Z',
    },
    {
      id: 'prj-006', name: 'Azure Bay', builder: 'Azure Developers', location: 'Ulwe',
      reraNumber: 'P52000043567', status: 'active' as const, totalUnits: 150, availableUnits: 28,
      configurations: ['2 BHK', '3 BHK'], priceRange: { min: 6500000, max: 15000000 },
      pricePerSqft: 8200, amenities: ['Swimming Pool', 'Gym', 'Rooftop Garden', 'Parking'],
      possessionDate: '2026-06-01', constructionProgress: 90, createdAt: '2023-12-01T10:00:00Z',
    },
    {
      id: 'prj-007', name: 'Royal Palms', builder: 'Royal Group', location: 'Panvel',
      reraNumber: 'P52000070891', status: 'draft' as const, totalUnits: 200, availableUnits: 200,
      configurations: ['2 BHK', '3 BHK', '4 BHK'], priceRange: { min: 8000000, max: 20000000 },
      pricePerSqft: 8800, amenities: ['Swimming Pool', 'Gym', 'Clubhouse', 'Tennis Court', 'Spa'],
      possessionDate: '2029-01-01', constructionProgress: 0, createdAt: '2025-02-28T10:00:00Z',
    },
    {
      id: 'prj-008', name: 'Sunrise Towers', builder: 'Sunrise Realty', location: 'Kharghar',
      reraNumber: 'P52000025678', status: 'sold_out' as const, totalUnits: 100, availableUnits: 0,
      configurations: ['2 BHK', '3 BHK'], priceRange: { min: 8500000, max: 18000000 },
      pricePerSqft: 10200, amenities: ['Swimming Pool', 'Gym', 'Garden'],
      possessionDate: '2026-03-01', constructionProgress: 95, createdAt: '2023-06-10T10:00:00Z',
    },
    {
      id: 'prj-009', name: 'Viva City', builder: 'Viva Developers', location: 'Dombivli',
      reraNumber: 'P52000033456', status: 'active' as const, totalUnits: 280, availableUnits: 110,
      configurations: ['1 BHK', '2 BHK', '3 BHK'], priceRange: { min: 3500000, max: 9500000 },
      pricePerSqft: 5500, amenities: ['Gym', 'Garden', 'Parking', 'Children Play Area'],
      possessionDate: '2027-09-01', constructionProgress: 35, createdAt: '2024-10-01T10:00:00Z',
    },
    {
      id: 'prj-010', name: 'Skyline Avenue', builder: 'Skyline Constructions', location: 'Panvel',
      reraNumber: 'P52000081234', status: 'active' as const, totalUnits: 160, availableUnits: 72,
      configurations: ['2 BHK', '3 BHK'], priceRange: { min: 6000000, max: 14000000 },
      pricePerSqft: 7800, amenities: ['Swimming Pool', 'Gym', 'Clubhouse', 'Jogging Track'],
      possessionDate: '2027-03-01', constructionProgress: 55, createdAt: '2024-07-20T10:00:00Z',
    },
    {
      id: 'prj-011', name: 'Lotus Greens', builder: 'Lotus Builders', location: 'Kalyan',
      reraNumber: 'P52000092345', status: 'active' as const, totalUnits: 350, availableUnits: 155,
      configurations: ['1 BHK', '2 BHK'], priceRange: { min: 2800000, max: 7000000 },
      pricePerSqft: 4800, amenities: ['Garden', 'Parking', 'Community Hall'],
      possessionDate: '2028-06-01', constructionProgress: 20, createdAt: '2025-01-20T10:00:00Z',
    },
    {
      id: 'prj-012', name: 'Metro Residences', builder: 'Metro Group', location: 'Kharghar',
      reraNumber: 'P52000015678', status: 'suspended' as const, totalUnits: 90, availableUnits: 45,
      configurations: ['2 BHK', '3 BHK'], priceRange: { min: 7000000, max: 16000000 },
      pricePerSqft: 9000, amenities: ['Gym', 'Parking'],
      possessionDate: '2027-06-01', constructionProgress: 40, createdAt: '2024-04-05T10:00:00Z',
    },
  ]

  // ---- Leads ------------------------------------------------------------------
  const leads = [
    {
      id: 'lead-001', name: 'Rahul Mehta', phone: '9901234567', email: 'rahul.mehta@gmail.com',
      projectId: 'prj-001', projectName: 'Arihant Aspire', agentId: 'agt-001', agentName: 'Rajesh Sharma',
      source: 'website' as const, status: 'new' as const, budget: 7500000,
      preferredConfiguration: '2 BHK', isNRI: false, score: 75, isHot: true,
      remarks: 'Looking for investment property near Panvel station',
      nextFollowUpAt: '2026-03-19T10:00:00Z', lastContactedAt: null,
      createdAt: '2026-03-17T14:30:00Z', updatedAt: '2026-03-17T14:30:00Z',
    },
    {
      id: 'lead-002', name: 'Sunita Reddy', phone: '9912345678', email: 'sunita.r@yahoo.com',
      projectId: 'prj-002', projectName: 'Balaji Symphony', agentId: 'agt-001', agentName: 'Rajesh Sharma',
      source: 'whatsapp' as const, status: 'contacted' as const, budget: 5500000,
      preferredConfiguration: '2 BHK', isNRI: false, score: 60, isHot: false,
      remarks: 'Working in Seawoods, wants Ulwe for proximity',
      nextFollowUpAt: '2026-03-20T11:00:00Z', lastContactedAt: '2026-03-16T15:00:00Z',
      createdAt: '2026-03-14T09:00:00Z', updatedAt: '2026-03-16T15:00:00Z',
    },
    {
      id: 'lead-003', name: 'Deepak Tiwari', phone: '9923456789', email: 'deepak.t@outlook.com',
      projectId: 'prj-003', projectName: 'Paradise Heights', agentId: 'agt-001', agentName: 'Rajesh Sharma',
      source: 'referral' as const, status: 'site_visit' as const, budget: 15000000,
      preferredConfiguration: '3 BHK', isNRI: false, score: 85, isHot: true,
      remarks: 'Visited site, very interested. Wife wants to visit again.',
      nextFollowUpAt: '2026-03-21T14:00:00Z', lastContactedAt: '2026-03-15T10:00:00Z',
      createdAt: '2026-03-10T11:00:00Z', updatedAt: '2026-03-15T10:00:00Z',
    },
    {
      id: 'lead-004', name: 'Kavita Shah', phone: '9934567890', email: 'kavita.shah@gmail.com',
      projectId: 'prj-001', projectName: 'Arihant Aspire', agentId: 'agt-001', agentName: 'Rajesh Sharma',
      source: 'website' as const, status: 'new' as const, budget: 5000000,
      preferredConfiguration: '1 BHK', isNRI: false, score: 55, isHot: false,
      remarks: 'First-time buyer, needs financing guidance',
      nextFollowUpAt: '2026-03-19T15:00:00Z', lastContactedAt: null,
      createdAt: '2026-03-17T16:00:00Z', updatedAt: '2026-03-17T16:00:00Z',
    },
    {
      id: 'lead-005', name: 'Manish Agarwal', phone: '9945678901', email: 'manish.a@gmail.com',
      projectId: 'prj-005', projectName: 'Green Valley', agentId: 'agt-001', agentName: 'Rajesh Sharma',
      source: 'whatsapp' as const, status: 'contacted' as const, budget: 8000000,
      preferredConfiguration: '2 BHK', isNRI: false, score: 65, isHot: false,
      remarks: 'Comparing with other projects in Kalyan',
      nextFollowUpAt: '2026-03-22T10:00:00Z', lastContactedAt: '2026-03-15T14:00:00Z',
      createdAt: '2026-03-12T10:00:00Z', updatedAt: '2026-03-15T14:00:00Z',
    },
    {
      id: 'lead-006', name: 'Pooja Iyer', phone: '9956789012', email: 'pooja.iyer@hotmail.com',
      projectId: 'prj-006', projectName: 'Azure Bay', agentId: 'agt-002', agentName: 'Priya Patel',
      source: 'referral' as const, status: 'booked' as const, budget: 12000000,
      preferredConfiguration: '3 BHK', isNRI: false, score: 95, isHot: true,
      remarks: 'Booked 3BHK in Tower B, 12th floor. Agreement pending.',
      nextFollowUpAt: null, lastContactedAt: '2026-03-16T12:00:00Z',
      createdAt: '2026-02-20T09:00:00Z', updatedAt: '2026-03-16T12:00:00Z',
    },
    {
      id: 'lead-007', name: 'Arjun Kapoor', phone: '9967890123', email: null,
      projectId: 'prj-003', projectName: 'Paradise Heights', agentId: 'agt-002', agentName: 'Priya Patel',
      source: 'whatsapp' as const, status: 'new' as const, budget: 18000000,
      preferredConfiguration: '4 BHK', isNRI: true, score: 80, isHot: true,
      remarks: 'NRI based in Dubai, visiting Mumbai next month',
      nextFollowUpAt: '2026-03-25T10:00:00Z', lastContactedAt: null,
      createdAt: '2026-03-16T20:00:00Z', updatedAt: '2026-03-16T20:00:00Z',
    },
    {
      id: 'lead-008', name: 'Nandini Rao', phone: '9978901234', email: 'nandini.rao@gmail.com',
      projectId: 'prj-004', projectName: 'Sai Residency', agentId: 'agt-002', agentName: 'Priya Patel',
      source: 'website' as const, status: 'contacted' as const, budget: 4500000,
      preferredConfiguration: '1 BHK', isNRI: false, score: 50, isHot: false,
      remarks: 'Budget conscious, wants EMI details',
      nextFollowUpAt: '2026-03-20T16:00:00Z', lastContactedAt: '2026-03-14T11:00:00Z',
      createdAt: '2026-03-13T08:00:00Z', updatedAt: '2026-03-14T11:00:00Z',
    },
    {
      id: 'lead-009', name: 'Vikash Sinha', phone: '9989012345', email: 'vikash.s@gmail.com',
      projectId: 'prj-002', projectName: 'Balaji Symphony', agentId: 'agt-003', agentName: 'Amit Deshmukh',
      source: 'referral' as const, status: 'contacted' as const, budget: 6500000,
      preferredConfiguration: '2 BHK', isNRI: false, score: 70, isHot: false,
      remarks: 'Referred by existing customer, interested in Ulwe connectivity',
      nextFollowUpAt: '2026-03-19T12:00:00Z', lastContactedAt: '2026-03-16T09:00:00Z',
      createdAt: '2026-03-11T10:00:00Z', updatedAt: '2026-03-16T09:00:00Z',
    },
    {
      id: 'lead-010', name: 'Rekha Verma', phone: '9990123456', email: 'rekha.v@gmail.com',
      projectId: 'prj-010', projectName: 'Skyline Avenue', agentId: 'agt-003', agentName: 'Amit Deshmukh',
      source: 'website' as const, status: 'site_visit' as const, budget: 11000000,
      preferredConfiguration: '3 BHK', isNRI: false, score: 82, isHot: true,
      remarks: 'Visited once, wants to compare floor plans of 8th vs 12th floor',
      nextFollowUpAt: '2026-03-20T14:00:00Z', lastContactedAt: '2026-03-15T16:00:00Z',
      createdAt: '2026-03-08T14:00:00Z', updatedAt: '2026-03-15T16:00:00Z',
    },
    {
      id: 'lead-011', name: 'Rohan Patil', phone: '9811223344', email: 'rohan.p@gmail.com',
      projectId: 'prj-001', projectName: 'Arihant Aspire', agentId: 'agt-003', agentName: 'Amit Deshmukh',
      source: 'whatsapp' as const, status: 'new' as const, budget: 6000000,
      preferredConfiguration: '2 BHK', isNRI: false, score: 58, isHot: false,
      remarks: 'Enquired via WhatsApp ad, early stage',
      nextFollowUpAt: '2026-03-19T09:00:00Z', lastContactedAt: null,
      createdAt: '2026-03-17T11:00:00Z', updatedAt: '2026-03-17T11:00:00Z',
    },
    {
      id: 'lead-012', name: 'Swati Jain', phone: '9822334455', email: 'swati.jain@outlook.com',
      projectId: 'prj-005', projectName: 'Green Valley', agentId: 'agt-004', agentName: 'Sneha Kulkarni',
      source: 'website' as const, status: 'contacted' as const, budget: 7000000,
      preferredConfiguration: '2 BHK', isNRI: false, score: 62, isHot: false,
      remarks: 'Works in Thane, exploring Kalyan for better value',
      nextFollowUpAt: '2026-03-21T10:00:00Z', lastContactedAt: '2026-03-15T13:00:00Z',
      createdAt: '2026-03-13T15:00:00Z', updatedAt: '2026-03-15T13:00:00Z',
    },
    {
      id: 'lead-013', name: 'Anil Kumar', phone: '9833445566', email: 'anil.k@gmail.com',
      projectId: 'prj-009', projectName: 'Viva City', agentId: 'agt-004', agentName: 'Sneha Kulkarni',
      source: 'referral' as const, status: 'booked' as const, budget: 5500000,
      preferredConfiguration: '2 BHK', isNRI: false, score: 92, isHot: true,
      remarks: 'Booked 2BHK in Tower A, 5th floor. Token paid.',
      nextFollowUpAt: null, lastContactedAt: '2026-03-14T10:00:00Z',
      createdAt: '2026-02-15T10:00:00Z', updatedAt: '2026-03-14T10:00:00Z',
    },
    {
      id: 'lead-014', name: 'Divya Menon', phone: '9844556677', email: 'divya.m@yahoo.com',
      projectId: 'prj-003', projectName: 'Paradise Heights', agentId: 'agt-004', agentName: 'Sneha Kulkarni',
      source: 'website' as const, status: 'contacted' as const, budget: 16000000,
      preferredConfiguration: '3 BHK', isNRI: true, score: 78, isHot: true,
      remarks: 'NRI in Singapore, plans to return in 2 years',
      nextFollowUpAt: '2026-03-22T18:00:00Z', lastContactedAt: '2026-03-16T18:00:00Z',
      createdAt: '2026-03-12T20:00:00Z', updatedAt: '2026-03-16T18:00:00Z',
    },
    {
      id: 'lead-015', name: 'Nitin Chavan', phone: '9855667788', email: null,
      projectId: 'prj-004', projectName: 'Sai Residency', agentId: 'agt-005', agentName: 'Vikram Joshi',
      source: 'whatsapp' as const, status: 'new' as const, budget: 3500000,
      preferredConfiguration: '1 BHK', isNRI: false, score: 45, isHot: false,
      remarks: 'Young professional looking for first home',
      nextFollowUpAt: '2026-03-19T11:00:00Z', lastContactedAt: null,
      createdAt: '2026-03-17T08:00:00Z', updatedAt: '2026-03-17T08:00:00Z',
    },
    {
      id: 'lead-016', name: 'Sanjay Bhatt', phone: '9866778899', email: 'sanjay.b@gmail.com',
      projectId: 'prj-006', projectName: 'Azure Bay', agentId: 'agt-006', agentName: 'Meera Nair',
      source: 'referral' as const, status: 'site_visit' as const, budget: 13000000,
      preferredConfiguration: '3 BHK', isNRI: false, score: 88, isHot: true,
      remarks: 'Second visit done, discussing with family',
      nextFollowUpAt: '2026-03-20T10:00:00Z', lastContactedAt: '2026-03-16T11:00:00Z',
      createdAt: '2026-03-05T09:00:00Z', updatedAt: '2026-03-16T11:00:00Z',
    },
    {
      id: 'lead-017', name: 'Geeta Pandey', phone: '9877889900', email: 'geeta.p@gmail.com',
      projectId: 'prj-010', projectName: 'Skyline Avenue', agentId: 'agt-006', agentName: 'Meera Nair',
      source: 'website' as const, status: 'contacted' as const, budget: 9000000,
      preferredConfiguration: '2 BHK', isNRI: false, score: 55, isHot: false,
      remarks: 'Comparing Panvel options, wants loan pre-approval first',
      nextFollowUpAt: '2026-03-23T10:00:00Z', lastContactedAt: '2026-03-15T15:00:00Z',
      createdAt: '2026-03-13T12:00:00Z', updatedAt: '2026-03-15T15:00:00Z',
    },
    {
      id: 'lead-018', name: 'Prakash Desai', phone: '9888990011', email: 'prakash.d@outlook.com',
      projectId: 'prj-011', projectName: 'Lotus Greens', agentId: 'agt-005', agentName: 'Vikram Joshi',
      source: 'website' as const, status: 'contacted' as const, budget: 4000000,
      preferredConfiguration: '1 BHK', isNRI: false, score: 48, isHot: false,
      remarks: 'Interested but financing not yet sorted',
      nextFollowUpAt: '2026-03-24T10:00:00Z', lastContactedAt: '2026-03-15T10:00:00Z',
      createdAt: '2026-03-14T10:00:00Z', updatedAt: '2026-03-15T10:00:00Z',
    },
    {
      id: 'lead-019', name: 'Vandana Thakur', phone: '9899001122', email: 'vandana.t@gmail.com',
      projectId: 'prj-002', projectName: 'Balaji Symphony', agentId: 'agt-006', agentName: 'Meera Nair',
      source: 'whatsapp' as const, status: 'new' as const, budget: 5800000,
      preferredConfiguration: '2 BHK', isNRI: false, score: 60, isHot: false,
      remarks: 'Saw project on Instagram, wants brochure',
      nextFollowUpAt: '2026-03-19T14:00:00Z', lastContactedAt: null,
      createdAt: '2026-03-17T12:00:00Z', updatedAt: '2026-03-17T12:00:00Z',
    },
    {
      id: 'lead-020', name: 'Ajay Mishra', phone: '9800112233', email: 'ajay.m@gmail.com',
      projectId: 'prj-001', projectName: 'Arihant Aspire', agentId: 'agt-003', agentName: 'Amit Deshmukh',
      source: 'referral' as const, status: 'contacted' as const, budget: 8500000,
      preferredConfiguration: '2 BHK', isNRI: false, score: 68, isHot: false,
      remarks: 'High budget, wants premium floor',
      nextFollowUpAt: '2026-03-21T11:00:00Z', lastContactedAt: '2026-03-16T14:00:00Z',
      createdAt: '2026-03-10T10:00:00Z', updatedAt: '2026-03-16T14:00:00Z',
    },
    {
      id: 'lead-021', name: 'Neha Saxena', phone: '9811223355', email: 'neha.s@gmail.com',
      projectId: 'prj-005', projectName: 'Green Valley', agentId: 'agt-002', agentName: 'Priya Patel',
      source: 'website' as const, status: 'new' as const, budget: 6000000,
      preferredConfiguration: '2 BHK', isNRI: false, score: 52, isHot: false,
      remarks: 'Browsed website, filled enquiry form',
      nextFollowUpAt: '2026-03-19T16:00:00Z', lastContactedAt: null,
      createdAt: '2026-03-17T18:00:00Z', updatedAt: '2026-03-17T18:00:00Z',
    },
    {
      id: 'lead-022', name: 'Ramesh Yadav', phone: '9822334466', email: null,
      projectId: 'prj-009', projectName: 'Viva City', agentId: 'agt-005', agentName: 'Vikram Joshi',
      source: 'whatsapp' as const, status: 'new' as const, budget: 4200000,
      preferredConfiguration: '1 BHK', isNRI: false, score: 40, isHot: false,
      remarks: 'Early enquiry via WhatsApp',
      nextFollowUpAt: '2026-03-20T09:00:00Z', lastContactedAt: null,
      createdAt: '2026-03-17T07:00:00Z', updatedAt: '2026-03-17T07:00:00Z',
    },
    {
      id: 'lead-023', name: 'Lakshmi Narayan', phone: '9833445577', email: 'lakshmi.n@gmail.com',
      projectId: 'prj-003', projectName: 'Paradise Heights', agentId: 'agt-006', agentName: 'Meera Nair',
      source: 'referral' as const, status: 'site_visit' as const, budget: 20000000,
      preferredConfiguration: '4 BHK', isNRI: true, score: 90, isHot: true,
      remarks: 'NRI from USA, visited site during India trip. Very impressed.',
      nextFollowUpAt: '2026-03-22T08:00:00Z', lastContactedAt: '2026-03-14T16:00:00Z',
      createdAt: '2026-03-01T10:00:00Z', updatedAt: '2026-03-14T16:00:00Z',
    },
    {
      id: 'lead-024', name: 'Tushar More', phone: '9844556688', email: 'tushar.m@gmail.com',
      projectId: 'prj-010', projectName: 'Skyline Avenue', agentId: 'agt-004', agentName: 'Sneha Kulkarni',
      source: 'website' as const, status: 'contacted' as const, budget: 10000000,
      preferredConfiguration: '3 BHK', isNRI: false, score: 64, isHot: false,
      remarks: 'Interested in Panvel market, comparing 3 projects',
      nextFollowUpAt: '2026-03-21T15:00:00Z', lastContactedAt: '2026-03-16T10:00:00Z',
      createdAt: '2026-03-11T14:00:00Z', updatedAt: '2026-03-16T10:00:00Z',
    },
  ]

  // ---- Site Visits ------------------------------------------------------------
  const visits = [
    {
      id: 'vis-001', leadId: 'lead-003', investorName: 'Deepak Tiwari', investorPhone: '9923456789',
      projectId: 'prj-003', projectName: 'Paradise Heights', agentId: 'agt-001', agentName: 'Rajesh Sharma',
      scheduledAt: '2026-03-20T11:00:00Z', status: 'scheduled' as const,
      investorFeedback: null, investorInterestLevel: null, agentNotes: null,
      outcome: null, nextSteps: null,
    },
    {
      id: 'vis-002', leadId: 'lead-010', investorName: 'Rekha Verma', investorPhone: '9990123456',
      projectId: 'prj-010', projectName: 'Skyline Avenue', agentId: 'agt-003', agentName: 'Amit Deshmukh',
      scheduledAt: '2026-03-21T14:00:00Z', status: 'scheduled' as const,
      investorFeedback: null, investorInterestLevel: null, agentNotes: null,
      outcome: null, nextSteps: null,
    },
    {
      id: 'vis-003', leadId: 'lead-016', investorName: 'Sanjay Bhatt', investorPhone: '9866778899',
      projectId: 'prj-006', projectName: 'Azure Bay', agentId: 'agt-006', agentName: 'Meera Nair',
      scheduledAt: '2026-03-19T10:00:00Z', status: 'scheduled' as const,
      investorFeedback: null, investorInterestLevel: null, agentNotes: null,
      outcome: null, nextSteps: null,
    },
    {
      id: 'vis-004', leadId: 'lead-003', investorName: 'Deepak Tiwari', investorPhone: '9923456789',
      projectId: 'prj-003', projectName: 'Paradise Heights', agentId: 'agt-001', agentName: 'Rajesh Sharma',
      scheduledAt: '2026-03-14T11:00:00Z', status: 'completed' as const,
      investorFeedback: 'Beautiful project, loved the amenities and view from 15th floor. Want to discuss pricing.',
      investorInterestLevel: 'high' as const, agentNotes: 'Client is very keen. Showed 3BHK on 15th floor facing garden.',
      outcome: 'interested' as const, nextSteps: 'Schedule second visit with wife',
    },
    {
      id: 'vis-005', leadId: 'lead-006', investorName: 'Pooja Iyer', investorPhone: '9956789012',
      projectId: 'prj-006', projectName: 'Azure Bay', agentId: 'agt-002', agentName: 'Priya Patel',
      scheduledAt: '2026-03-10T14:00:00Z', status: 'completed' as const,
      investorFeedback: 'Perfect location, very happy with the quality. Ready to book.',
      investorInterestLevel: 'high' as const, agentNotes: 'Client booked on the spot. 3BHK Tower B, 12th floor.',
      outcome: 'booked' as const, nextSteps: 'Process booking documentation',
    },
    {
      id: 'vis-006', leadId: 'lead-016', investorName: 'Sanjay Bhatt', investorPhone: '9866778899',
      projectId: 'prj-006', projectName: 'Azure Bay', agentId: 'agt-006', agentName: 'Meera Nair',
      scheduledAt: '2026-03-12T11:00:00Z', status: 'completed' as const,
      investorFeedback: 'Good project but need to discuss with family before deciding.',
      investorInterestLevel: 'medium' as const, agentNotes: 'Showed 2BHK and 3BHK options. Client leaning towards 3BHK.',
      outcome: 'follow_up' as const, nextSteps: 'Follow up in a week, schedule second visit',
    },
    {
      id: 'vis-007', leadId: 'lead-023', investorName: 'Lakshmi Narayan', investorPhone: '9833445577',
      projectId: 'prj-003', projectName: 'Paradise Heights', agentId: 'agt-006', agentName: 'Meera Nair',
      scheduledAt: '2026-03-13T10:00:00Z', status: 'completed' as const,
      investorFeedback: 'Excellent location and construction quality. Very impressed with the clubhouse.',
      investorInterestLevel: 'high' as const, agentNotes: 'NRI client from USA. Showed 4BHK penthouse option. Very interested.',
      outcome: 'interested' as const, nextSteps: 'Send detailed pricing and payment plan via email',
    },
    {
      id: 'vis-008', leadId: 'lead-013', investorName: 'Anil Kumar', investorPhone: '9833445566',
      projectId: 'prj-009', projectName: 'Viva City', agentId: 'agt-004', agentName: 'Sneha Kulkarni',
      scheduledAt: '2026-03-08T11:00:00Z', status: 'completed' as const,
      investorFeedback: 'Good value for money. Location suits my commute.',
      investorInterestLevel: 'high' as const, agentNotes: 'Booked 2BHK Tower A, 5th floor after visit.',
      outcome: 'booked' as const, nextSteps: 'Coordinate with builder for agreement',
    },
    {
      id: 'vis-009', leadId: 'lead-010', investorName: 'Rekha Verma', investorPhone: '9990123456',
      projectId: 'prj-010', projectName: 'Skyline Avenue', agentId: 'agt-003', agentName: 'Amit Deshmukh',
      scheduledAt: '2026-03-11T15:00:00Z', status: 'completed' as const,
      investorFeedback: 'Good project, but want to compare floor plans.',
      investorInterestLevel: 'medium' as const, agentNotes: 'Showed 3BHK on 8th floor. Client wants 12th floor comparison.',
      outcome: 'follow_up' as const, nextSteps: 'Schedule second visit with floor plan comparison',
    },
    {
      id: 'vis-010', leadId: 'lead-005', investorName: 'Manish Agarwal', investorPhone: '9945678901',
      projectId: 'prj-005', projectName: 'Green Valley', agentId: 'agt-001', agentName: 'Rajesh Sharma',
      scheduledAt: '2026-03-13T14:00:00Z', status: 'completed' as const,
      investorFeedback: 'Nice project but construction is early stage. Will wait and watch.',
      investorInterestLevel: 'low' as const, agentNotes: 'Client not convinced about early-stage investment.',
      outcome: 'not_interested' as const, nextSteps: 'Follow up after 2 months when construction progresses',
    },
    {
      id: 'vis-011', leadId: 'lead-008', investorName: 'Nandini Rao', investorPhone: '9978901234',
      projectId: 'prj-004', projectName: 'Sai Residency', agentId: 'agt-002', agentName: 'Priya Patel',
      scheduledAt: '2026-03-09T11:00:00Z', status: 'cancelled' as const,
      investorFeedback: null, investorInterestLevel: null,
      agentNotes: 'Client cancelled due to personal emergency. Reschedule later.',
      outcome: null, nextSteps: 'Reschedule when client is available',
    },
  ]

  // ---- Commissions -----------------------------------------------------------
  const commissions = [
    {
      id: 'com-001', bookingId: 'bkg-001', bookingRef: 'BKG-2026-001',
      agentId: 'agt-002', agentName: 'Priya Patel',
      projectId: 'prj-006', projectName: 'Azure Bay',
      investorName: 'Pooja Iyer',
      agreementValue: 12500000, brokerageRate: 2, totalBrokerage: 250000,
      agentCommission: 200000, tdsAmount: 10000, gstAmount: 0, netPayableToAgent: 190000,
      status: 'paid' as const, paidAt: '2026-03-15T10:00:00Z',
      createdAt: '2026-03-11T10:00:00Z',
    },
    {
      id: 'com-002', bookingId: 'bkg-002', bookingRef: 'BKG-2026-002',
      agentId: 'agt-004', agentName: 'Sneha Kulkarni',
      projectId: 'prj-009', projectName: 'Viva City',
      investorName: 'Anil Kumar',
      agreementValue: 5500000, brokerageRate: 2.5, totalBrokerage: 137500,
      agentCommission: 110000, tdsAmount: 5500, gstAmount: 0, netPayableToAgent: 104500,
      status: 'approved' as const, paidAt: null,
      createdAt: '2026-03-09T10:00:00Z',
    },
    {
      id: 'com-003', bookingId: 'bkg-003', bookingRef: 'BKG-2025-018',
      agentId: 'agt-003', agentName: 'Amit Deshmukh',
      projectId: 'prj-003', projectName: 'Paradise Heights',
      investorName: 'Mahesh Gupta',
      agreementValue: 18000000, brokerageRate: 2, totalBrokerage: 360000,
      agentCommission: 288000, tdsAmount: 14400, gstAmount: 0, netPayableToAgent: 273600,
      status: 'paid' as const, paidAt: '2026-02-20T10:00:00Z',
      createdAt: '2026-02-15T10:00:00Z',
    },
    {
      id: 'com-004', bookingId: 'bkg-004', bookingRef: 'BKG-2025-019',
      agentId: 'agt-001', agentName: 'Rajesh Sharma',
      projectId: 'prj-001', projectName: 'Arihant Aspire',
      investorName: 'Sunil Khanna',
      agreementValue: 7800000, brokerageRate: 2, totalBrokerage: 156000,
      agentCommission: 125000, tdsAmount: 6250, gstAmount: 0, netPayableToAgent: 118750,
      status: 'paid' as const, paidAt: '2026-01-25T10:00:00Z',
      createdAt: '2026-01-20T10:00:00Z',
    },
    {
      id: 'com-005', bookingId: 'bkg-005', bookingRef: 'BKG-2025-020',
      agentId: 'agt-003', agentName: 'Amit Deshmukh',
      projectId: 'prj-006', projectName: 'Azure Bay',
      investorName: 'Pradeep Nair',
      agreementValue: 14000000, brokerageRate: 2, totalBrokerage: 280000,
      agentCommission: 224000, tdsAmount: 11200, gstAmount: 0, netPayableToAgent: 212800,
      status: 'paid' as const, paidAt: '2026-02-05T10:00:00Z',
      createdAt: '2026-01-28T10:00:00Z',
    },
    {
      id: 'com-006', bookingId: 'bkg-006', bookingRef: 'BKG-2025-021',
      agentId: 'agt-006', agentName: 'Meera Nair',
      projectId: 'prj-002', projectName: 'Balaji Symphony',
      investorName: 'Kiran Rao',
      agreementValue: 6200000, brokerageRate: 2.5, totalBrokerage: 155000,
      agentCommission: 124000, tdsAmount: 6200, gstAmount: 0, netPayableToAgent: 117800,
      status: 'paid' as const, paidAt: '2026-02-28T10:00:00Z',
      createdAt: '2026-02-22T10:00:00Z',
    },
    {
      id: 'com-007', bookingId: 'bkg-007', bookingRef: 'BKG-2026-003',
      agentId: 'agt-001', agentName: 'Rajesh Sharma',
      projectId: 'prj-002', projectName: 'Balaji Symphony',
      investorName: 'Harsh Vardhan',
      agreementValue: 5200000, brokerageRate: 2.5, totalBrokerage: 130000,
      agentCommission: 104000, tdsAmount: 5200, gstAmount: 0, netPayableToAgent: 98800,
      status: 'pending' as const, paidAt: null,
      createdAt: '2026-03-16T10:00:00Z',
    },
    {
      id: 'com-008', bookingId: 'bkg-008', bookingRef: 'BKG-2026-004',
      agentId: 'agt-006', agentName: 'Meera Nair',
      projectId: 'prj-010', projectName: 'Skyline Avenue',
      investorName: 'Dinesh Patel',
      agreementValue: 9800000, brokerageRate: 2, totalBrokerage: 196000,
      agentCommission: 156800, tdsAmount: 7840, gstAmount: 0, netPayableToAgent: 148960,
      status: 'pending' as const, paidAt: null,
      createdAt: '2026-03-14T10:00:00Z',
    },
    {
      id: 'com-009', bookingId: 'bkg-009', bookingRef: 'BKG-2025-015',
      agentId: 'agt-003', agentName: 'Amit Deshmukh',
      projectId: 'prj-001', projectName: 'Arihant Aspire',
      investorName: 'Vijay Malhotra',
      agreementValue: 8500000, brokerageRate: 2, totalBrokerage: 170000,
      agentCommission: 136000, tdsAmount: 6800, gstAmount: 0, netPayableToAgent: 129200,
      status: 'paid' as const, paidAt: '2025-12-15T10:00:00Z',
      createdAt: '2025-12-10T10:00:00Z',
    },
    {
      id: 'com-010', bookingId: 'bkg-010', bookingRef: 'BKG-2025-016',
      agentId: 'agt-001', agentName: 'Rajesh Sharma',
      projectId: 'prj-003', projectName: 'Paradise Heights',
      investorName: 'Ashok Shetty',
      agreementValue: 16500000, brokerageRate: 2, totalBrokerage: 330000,
      agentCommission: 264000, tdsAmount: 13200, gstAmount: 0, netPayableToAgent: 250800,
      status: 'paid' as const, paidAt: '2025-11-20T10:00:00Z',
      createdAt: '2025-11-15T10:00:00Z',
    },
  ]

  // ---- Inventory (Units) for builder view ------------------------------------
  const units = [
    { id: 'u-001', projectId: 'prj-001', projectName: 'Arihant Aspire', tower: 'A', floor: 2, unitNumber: 'A-201', unitType: '1 BHK' as const, carpetArea: 450, price: 4500000, pricePerSqft: 7200, status: 'available' as const, facing: 'East' },
    { id: 'u-002', projectId: 'prj-001', projectName: 'Arihant Aspire', tower: 'A', floor: 2, unitNumber: 'A-202', unitType: '2 BHK' as const, carpetArea: 680, price: 6800000, pricePerSqft: 7200, status: 'available' as const, facing: 'West' },
    { id: 'u-003', projectId: 'prj-001', projectName: 'Arihant Aspire', tower: 'A', floor: 3, unitNumber: 'A-301', unitType: '1 BHK' as const, carpetArea: 450, price: 4600000, pricePerSqft: 7400, status: 'booked' as const, facing: 'East' },
    { id: 'u-004', projectId: 'prj-001', projectName: 'Arihant Aspire', tower: 'A', floor: 3, unitNumber: 'A-302', unitType: '2 BHK' as const, carpetArea: 680, price: 6900000, pricePerSqft: 7400, status: 'available' as const, facing: 'West' },
    { id: 'u-005', projectId: 'prj-001', projectName: 'Arihant Aspire', tower: 'B', floor: 5, unitNumber: 'B-501', unitType: '3 BHK' as const, carpetArea: 980, price: 9800000, pricePerSqft: 7500, status: 'sold' as const, facing: 'North' },
    { id: 'u-006', projectId: 'prj-001', projectName: 'Arihant Aspire', tower: 'B', floor: 5, unitNumber: 'B-502', unitType: '3 BHK' as const, carpetArea: 980, price: 9800000, pricePerSqft: 7500, status: 'available' as const, facing: 'South' },
    { id: 'u-007', projectId: 'prj-001', projectName: 'Arihant Aspire', tower: 'B', floor: 8, unitNumber: 'B-801', unitType: '2 BHK' as const, carpetArea: 720, price: 7500000, pricePerSqft: 7600, status: 'blocked' as const, facing: 'East' },
    { id: 'u-008', projectId: 'prj-001', projectName: 'Arihant Aspire', tower: 'B', floor: 12, unitNumber: 'B-1201', unitType: '3 BHK' as const, carpetArea: 1020, price: 12500000, pricePerSqft: 7800, status: 'available' as const, facing: 'North' },
    { id: 'u-009', projectId: 'prj-002', projectName: 'Balaji Symphony', tower: 'A', floor: 1, unitNumber: 'A-101', unitType: '1 BHK' as const, carpetArea: 420, price: 3800000, pricePerSqft: 6800, status: 'sold' as const, facing: 'East' },
    { id: 'u-010', projectId: 'prj-002', projectName: 'Balaji Symphony', tower: 'A', floor: 1, unitNumber: 'A-102', unitType: '2 BHK' as const, carpetArea: 650, price: 5850000, pricePerSqft: 6800, status: 'available' as const, facing: 'West' },
    { id: 'u-011', projectId: 'prj-002', projectName: 'Balaji Symphony', tower: 'A', floor: 4, unitNumber: 'A-401', unitType: '1 BHK' as const, carpetArea: 420, price: 4000000, pricePerSqft: 7000, status: 'available' as const, facing: 'East' },
    { id: 'u-012', projectId: 'prj-002', projectName: 'Balaji Symphony', tower: 'B', floor: 6, unitNumber: 'B-601', unitType: '2 BHK' as const, carpetArea: 680, price: 6300000, pricePerSqft: 7100, status: 'blocked' as const, facing: 'South' },
    { id: 'u-013', projectId: 'prj-002', projectName: 'Balaji Symphony', tower: 'B', floor: 10, unitNumber: 'B-1001', unitType: '2 BHK' as const, carpetArea: 700, price: 7000000, pricePerSqft: 7200, status: 'available' as const, facing: 'North' },
    { id: 'u-014', projectId: 'prj-002', projectName: 'Balaji Symphony', tower: 'B', floor: 10, unitNumber: 'B-1002', unitType: '2 BHK' as const, carpetArea: 700, price: 7000000, pricePerSqft: 7200, status: 'booked' as const, facing: 'South' },
    { id: 'u-015', projectId: 'prj-003', projectName: 'Paradise Heights', tower: 'A', floor: 15, unitNumber: 'A-1501', unitType: '3 BHK' as const, carpetArea: 1100, price: 14500000, pricePerSqft: 9500, status: 'available' as const, facing: 'East' },
    { id: 'u-016', projectId: 'prj-003', projectName: 'Paradise Heights', tower: 'A', floor: 15, unitNumber: 'A-1502', unitType: '4 BHK' as const, carpetArea: 1600, price: 22000000, pricePerSqft: 9800, status: 'available' as const, facing: 'West' },
    { id: 'u-017', projectId: 'prj-003', projectName: 'Paradise Heights', tower: 'B', floor: 8, unitNumber: 'B-801', unitType: '2 BHK' as const, carpetArea: 850, price: 7500000, pricePerSqft: 9200, status: 'sold' as const, facing: 'North' },
    { id: 'u-018', projectId: 'prj-003', projectName: 'Paradise Heights', tower: 'B', floor: 12, unitNumber: 'B-1201', unitType: '3 BHK' as const, carpetArea: 1100, price: 15000000, pricePerSqft: 9500, status: 'booked' as const, facing: 'South' },
  ]

  // ---- Activity Feed (for admin dashboard) ------------------------------------
  const recentActivity = [
    { id: 'act-01', action: 'lead_created', description: 'New lead Rahul Mehta enquired about Arihant Aspire', actor: 'System', time: '2026-03-17T14:30:00Z' },
    { id: 'act-02', action: 'site_visit_completed', description: 'Deepak Tiwari completed site visit at Paradise Heights', actor: 'Rajesh Sharma', time: '2026-03-14T12:00:00Z' },
    { id: 'act-03', action: 'booking_created', description: 'Pooja Iyer booked 3BHK at Azure Bay', actor: 'Priya Patel', time: '2026-03-10T15:00:00Z' },
    { id: 'act-04', action: 'commission_paid', description: 'Commission of ₹1,90,000 paid to Priya Patel', actor: 'Admin', time: '2026-03-15T10:00:00Z' },
    { id: 'act-05', action: 'agent_onboarded', description: 'New agent Suresh Gupta registered, pending approval', actor: 'System', time: '2026-03-01T10:00:00Z' },
    { id: 'act-06', action: 'lead_status_changed', description: 'Anil Kumar lead status changed to Booked', actor: 'Sneha Kulkarni', time: '2026-03-08T12:00:00Z' },
    { id: 'act-07', action: 'site_visit_scheduled', description: 'Visit scheduled for Sanjay Bhatt at Azure Bay', actor: 'Meera Nair', time: '2026-03-17T09:00:00Z' },
    { id: 'act-08', action: 'lead_created', description: 'New lead Kavita Shah enquired about Arihant Aspire', actor: 'System', time: '2026-03-17T16:00:00Z' },
  ]

  // ---- Builders ---------------------------------------------------------------
  const builders = [
    {
      id: 'bld-001', name: 'Arihant Superstructures', contactPerson: 'Ashok Chhajer', phone: '9820012345', email: 'info@arihantgroup.co.in',
      reraNumber: 'P52100012345', activeProjects: 3, totalUnits: 600, rating: 4.2,
    },
    {
      id: 'bld-002', name: 'Balaji Developers', contactPerson: 'Rajendra Jain', phone: '9821123456', email: 'contact@balajidevelopers.com',
      reraNumber: 'P52100023456', activeProjects: 2, totalUnits: 330, rating: 3.8,
    },
    {
      id: 'bld-003', name: 'Paradise Group', contactPerson: 'Manish Narang', phone: '9822234567', email: 'sales@paradisegroup.co.in',
      reraNumber: 'P52100034567', activeProjects: 4, totalUnits: 920, rating: 4.5,
    },
    {
      id: 'bld-004', name: 'Godrej Properties', contactPerson: 'Pirojsha Godrej', phone: '9823345678', email: 'homes@godrejproperties.com',
      reraNumber: 'P52100045678', activeProjects: 5, totalUnits: 1200, rating: 4.7,
    },
  ]

  // ---- Areas (Micro-markets) ---------------------------------------------------
  const areas = [
    {
      id: 'area-001', name: 'Panvel', slug: 'panvel', avgPricePerSqft: 7500, growthPotential: 22,
      keyInfra: ['Navi Mumbai Airport', 'Mumbai Trans Harbour Link', 'Virar-Alibaug Multimodal Corridor'],
      totalProjects: 18,
      description: 'Panvel is a rapidly developing node in Navi Mumbai, benefiting from the upcoming Navi Mumbai International Airport and excellent road connectivity.',
    },
    {
      id: 'area-002', name: 'Dombivli East', slug: 'dombivli-east', avgPricePerSqft: 5800, growthPotential: 15,
      keyInfra: ['Dombivli-Kalyan Metro Line', 'Dombivli MIDC Expansion', 'Thane-Dombivli Flyover'],
      totalProjects: 12,
      description: 'Dombivli East offers affordable housing with strong rental demand thanks to proximity to MIDC industrial zone and good railway connectivity.',
    },
    {
      id: 'area-003', name: 'Kalyan-Shilphata', slug: 'kalyan-shilphata', avgPricePerSqft: 6200, growthPotential: 18,
      keyInfra: ['Kalyan-Dombivli Smart City', 'Bhiwandi-Kalyan Corridor', 'Metro Line 12 Extension'],
      totalProjects: 14,
      description: 'Kalyan-Shilphata corridor is emerging as a prime residential hub with improving infrastructure and competitive pricing compared to Thane and Mumbai.',
    },
    {
      id: 'area-004', name: 'Taloja', slug: 'taloja', avgPricePerSqft: 5200, growthPotential: 20,
      keyInfra: ['Taloja MIDC', 'Navi Mumbai Airport Influence Zone', 'Proposed Metro Line 4'],
      totalProjects: 10,
      description: 'Taloja is gaining traction as a value-for-money investment destination with upcoming airport proximity and industrial growth driving residential demand.',
    },
    {
      id: 'area-005', name: 'Kharghar', slug: 'kharghar', avgPricePerSqft: 9500, growthPotential: 12,
      keyInfra: ['Central Park', 'Golf Course', 'Proposed Kharghar Railway Station Upgrade'],
      totalProjects: 22,
      description: 'Kharghar is one of the most developed nodes of Navi Mumbai offering premium lifestyle, green spaces, and strong social infrastructure.',
    },
    {
      id: 'area-006', name: 'Ulwe', slug: 'ulwe', avgPricePerSqft: 7200, growthPotential: 25,
      keyInfra: ['Navi Mumbai Airport (nearest node)', 'Mumbai Trans Harbour Link', 'Proposed Coastal Road Extension'],
      totalProjects: 16,
      description: 'Ulwe is the closest residential node to the upcoming Navi Mumbai International Airport, making it a top investment hotspot with highest growth potential.',
    },
  ]

  // ---- Events ------------------------------------------------------------------
  const events = [
    {
      id: 'evt-001', title: 'Investment Opportunities in Navi Mumbai 2026',
      type: 'webinar' as const,
      scheduledAt: '2026-04-05T15:00:00Z', location: 'Online (Zoom)',
      description: 'Join industry experts as they discuss the latest trends and opportunities in Navi Mumbai real estate post-airport announcement.',
      maxAttendees: 200, registeredCount: 142,
    },
    {
      id: 'evt-002', title: 'Panvel & Ulwe Site Visit Day',
      type: 'site_visit_day' as const,
      scheduledAt: '2026-04-12T09:00:00Z', location: 'Panvel Railway Station (Meeting Point)',
      description: 'Exclusive guided site visit covering top projects in Panvel and Ulwe. Transportation and lunch included.',
      maxAttendees: 50, registeredCount: 38,
    },
    {
      id: 'evt-003', title: 'Q1 2026 Broker Meet & Greet',
      type: 'broker_meet' as const,
      scheduledAt: '2026-04-20T18:00:00Z', location: 'Hotel Sahara Star, Vile Parle East, Mumbai',
      description: 'Quarterly networking event for channel partners. New project launches, commission updates, and top performer awards.',
      maxAttendees: 150, registeredCount: 98,
    },
    {
      id: 'evt-004', title: 'Home Loan & Tax Benefits Workshop',
      type: 'webinar' as const,
      scheduledAt: '2026-05-03T11:00:00Z', location: 'Online (Google Meet)',
      description: 'A workshop for agents and investors covering home loan processes, tax benefits under Section 80C/24, and PMAY subsidy eligibility.',
      maxAttendees: 300, registeredCount: 67,
    },
    {
      id: 'evt-005', title: 'Kharghar Premium Properties Showcase',
      type: 'site_visit_day' as const,
      scheduledAt: '2026-05-10T10:00:00Z', location: 'Kharghar Central Park Entrance',
      description: 'Premium project showcase in Kharghar featuring 3BHK and 4BHK luxury apartments. Exclusive discounts for attendees.',
      maxAttendees: 40, registeredCount: 29,
    },
  ]

  // ---- Helpers ----------------------------------------------------------------

  /** Get leads for a specific agent (by agentId) */
  function getAgentLeads(agentId: string) {
    return leads.filter(l => l.agentId === agentId)
  }

  /** Get visits for a specific agent */
  function getAgentVisits(agentId: string) {
    return visits.filter(v => v.agentId === agentId)
  }

  /** Get commissions for a specific agent */
  function getAgentCommissions(agentId: string) {
    return commissions.filter(c => c.agentId === agentId)
  }

  /** Get units for a specific project */
  function getProjectUnits(projectId: string) {
    return units.filter(u => u.projectId === projectId)
  }

  return {
    agents,
    projects,
    leads,
    visits,
    commissions,
    units,
    recentActivity,
    builders,
    areas,
    events,
    getAgentLeads,
    getAgentVisits,
    getAgentCommissions,
    getProjectUnits,
  }
}
