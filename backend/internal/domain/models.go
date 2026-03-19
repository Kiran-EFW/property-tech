package domain

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

// ---------------------------------------------------------------------------
// Enums
// ---------------------------------------------------------------------------

// ProjectStatus represents the lifecycle state of a real-estate project.
type ProjectStatus string

const (
	ProjectStatusDraft      ProjectStatus = "draft"
	ProjectStatusActive     ProjectStatus = "active"
	ProjectStatusSoldOut    ProjectStatus = "sold_out"
	ProjectStatusCompleted  ProjectStatus = "completed"
	ProjectStatusSuspended  ProjectStatus = "suspended"
)

// LeadStatus represents the pipeline stage of an investor lead.
type LeadStatus string

const (
	LeadStatusNew          LeadStatus = "new"
	LeadStatusContacted    LeadStatus = "contacted"
	LeadStatusQualified    LeadStatus = "qualified"
	LeadStatusSiteVisit    LeadStatus = "site_visit"
	LeadStatusNegotiation  LeadStatus = "negotiation"
	LeadStatusConverted    LeadStatus = "converted"
	LeadStatusLost         LeadStatus = "lost"
)

// LeadSource indicates where a lead originated.
type LeadSource string

const (
	LeadSourceWalkIn    LeadSource = "walk_in"
	LeadSourceReferral  LeadSource = "referral"
	LeadSourceWebsite   LeadSource = "website"
	LeadSourceApp       LeadSource = "app"
	LeadSourceWhatsApp  LeadSource = "whatsapp"
	LeadSourceCampaign  LeadSource = "campaign"
	LeadSourceOther     LeadSource = "other"
)

// BookingStatus tracks the state of a unit booking.
type BookingStatus string

const (
	BookingStatusPending    BookingStatus = "pending"
	BookingStatusConfirmed  BookingStatus = "confirmed"
	BookingStatusCancelled  BookingStatus = "cancelled"
	BookingStatusCompleted  BookingStatus = "completed"
)

// CommissionStatus tracks agent commission payout state.
type CommissionStatus string

const (
	CommissionStatusPending   CommissionStatus = "pending"
	CommissionStatusApproved  CommissionStatus = "approved"
	CommissionStatusPaid      CommissionStatus = "paid"
	CommissionStatusRejected  CommissionStatus = "rejected"
)

// AgentTier represents the channel-partner tier.
type AgentTier string

const (
	AgentTierBronze   AgentTier = "bronze"
	AgentTierSilver   AgentTier = "silver"
	AgentTierGold     AgentTier = "gold"
	AgentTierPlatinum AgentTier = "platinum"
)

// ---------------------------------------------------------------------------
// Domain Models
// ---------------------------------------------------------------------------

// UserRole represents the role of a user in the system.
type UserRole string

const (
	UserRoleInvestor   UserRole = "investor"
	UserRoleAgent      UserRole = "agent"
	UserRoleBuilder    UserRole = "builder"
	UserRoleAdmin      UserRole = "admin"
	UserRoleSuperAdmin UserRole = "super_admin"
)

// User represents a registered user on the platform.
type User struct {
	ID        uuid.UUID `json:"id"         db:"id"`
	Phone     string    `json:"phone"      db:"phone"`
	Name      string    `json:"name"       db:"name"`
	Email     *string   `json:"email"      db:"email"`
	Role      UserRole  `json:"role"       db:"role"`
	IsActive  bool      `json:"is_active"  db:"is_active"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

// ProjectUnit represents an individual unit (flat/villa/plot) within a project.
type ProjectUnit struct {
	ID          uuid.UUID       `json:"id"           db:"id"`
	ProjectID   uuid.UUID       `json:"project_id"   db:"project_id"`
	Name        string          `json:"name"         db:"name"`
	Type        string          `json:"type"         db:"type"`         // 1BHK, 2BHK, 3BHK, villa, plot
	Floor       *int            `json:"floor"        db:"floor"`
	CarpetArea  float64         `json:"carpet_area"  db:"carpet_area"`
	Price       float64         `json:"price"        db:"price"`
	Status      string          `json:"status"       db:"status"`       // available, booked, sold
	Amenities   json.RawMessage `json:"amenities"    db:"amenities"`
	CreatedAt   time.Time       `json:"created_at"   db:"created_at"`
	UpdatedAt   time.Time       `json:"updated_at"   db:"updated_at"`
}

// LeadNote represents a note attached to a lead by an agent or admin.
type LeadNote struct {
	ID        uuid.UUID `json:"id"         db:"id"`
	LeadID    uuid.UUID `json:"lead_id"    db:"lead_id"`
	AuthorID  uuid.UUID `json:"author_id"  db:"author_id"`
	Content   string    `json:"content"    db:"content"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

// DueDiligenceReport holds the due-diligence data for a project.
type DueDiligenceReport struct {
	ProjectID     uuid.UUID       `json:"project_id"     db:"project_id"`
	RERAStatus    string          `json:"rera_status"    db:"rera_status"`
	TitleClear    bool            `json:"title_clear"    db:"title_clear"`
	Encumbrances  json.RawMessage `json:"encumbrances"   db:"encumbrances"`
	LegalOpinion  *string         `json:"legal_opinion"  db:"legal_opinion"`
	LastCheckedAt time.Time       `json:"last_checked_at" db:"last_checked_at"`
}

// Project represents a real-estate project listed on the platform.
type Project struct {
	ID            uuid.UUID       `json:"id"             db:"id"`
	Name          string          `json:"name"           db:"name"`
	Slug          string          `json:"slug"           db:"slug"`
	RERANumber    string          `json:"rera_number"    db:"rera_number"`
	BuilderID     uuid.UUID       `json:"builder_id"     db:"builder_id"`
	Description   *string         `json:"description"    db:"description"`
	CarpetAreaMin *float64        `json:"carpet_area_min" db:"carpet_area_min"`
	CarpetAreaMax *float64        `json:"carpet_area_max" db:"carpet_area_max"`
	PriceMin      *float64        `json:"price_min"      db:"price_min"`
	PriceMax      *float64        `json:"price_max"      db:"price_max"`
	Location      string          `json:"location"       db:"location"`       // PostGIS geometry (POINT)
	Address       *string         `json:"address"        db:"address"`
	City          *string         `json:"city"           db:"city"`
	State         *string         `json:"state"          db:"state"`
	Pincode       *string         `json:"pincode"        db:"pincode"`
	Status        ProjectStatus   `json:"status"         db:"status"`
	Amenities     json.RawMessage `json:"amenities"      db:"amenities"`
	Media         json.RawMessage `json:"media"          db:"media"`
	CreatedAt     time.Time       `json:"created_at"     db:"created_at"`
	UpdatedAt     time.Time       `json:"updated_at"     db:"updated_at"`
}

// Builder represents a real-estate developer / builder entity.
type Builder struct {
	ID           uuid.UUID `json:"id"            db:"id"`
	Name         string    `json:"name"          db:"name"`
	Slug         string    `json:"slug"          db:"slug"`
	RERANumber   string    `json:"rera_number"   db:"rera_number"`
	PAN          *string   `json:"pan"           db:"pan"`
	GST          *string   `json:"gst"           db:"gst"`
	TrackRecord  *string   `json:"track_record"  db:"track_record"`
	ContactPhone *string   `json:"contact_phone" db:"contact_phone"`
	ContactEmail *string   `json:"contact_email" db:"contact_email"`
	LogoURL      *string   `json:"logo_url"      db:"logo_url"`
	IsVerified   bool      `json:"is_verified"   db:"is_verified"`
	CreatedAt    time.Time `json:"created_at"    db:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"    db:"updated_at"`
}

// Lead represents a prospective investor lead in the sales pipeline.
type Lead struct {
	ID         uuid.UUID  `json:"id"          db:"id"`
	InvestorID *uuid.UUID `json:"investor_id" db:"investor_id"` // nullable until user registers
	ProjectID  uuid.UUID  `json:"project_id"  db:"project_id"`
	AgentID    *uuid.UUID `json:"agent_id"    db:"agent_id"`    // assigned channel partner
	Source     LeadSource `json:"source"      db:"source"`
	Status     LeadStatus `json:"status"      db:"status"`
	Phone      string     `json:"phone"       db:"phone"`
	Name       string     `json:"name"        db:"name"`
	Email      *string    `json:"email"       db:"email"`
	Budget     *float64   `json:"budget"      db:"budget"`
	Notes      *string    `json:"notes"       db:"notes"`
	CreatedAt  time.Time  `json:"created_at"  db:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"  db:"updated_at"`
}

// Agent represents a channel partner (real-estate agent) on the platform.
type Agent struct {
	ID         uuid.UUID `json:"id"          db:"id"`
	UserID     uuid.UUID `json:"user_id"     db:"user_id"`
	RERANumber *string   `json:"rera_number" db:"rera_number"`
	PAN        *string   `json:"pan"         db:"pan"`
	GST        *string   `json:"gst"         db:"gst"`
	Tier       AgentTier `json:"tier"        db:"tier"`
	IsActive   bool      `json:"is_active"   db:"is_active"`
	CreatedAt  time.Time `json:"created_at"  db:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"  db:"updated_at"`
}

// Booking represents a unit booking made by an investor.
type Booking struct {
	ID             uuid.UUID     `json:"id"              db:"id"`
	ProjectID      uuid.UUID     `json:"project_id"      db:"project_id"`
	UnitID         *uuid.UUID    `json:"unit_id"         db:"unit_id"`
	InvestorID     uuid.UUID     `json:"investor_id"     db:"investor_id"`
	AgentID        *uuid.UUID    `json:"agent_id"        db:"agent_id"`
	BookingAmount  float64       `json:"booking_amount"  db:"booking_amount"`
	AgreementValue float64       `json:"agreement_value" db:"agreement_value"`
	Status         BookingStatus `json:"status"          db:"status"`
	BookedAt       time.Time     `json:"booked_at"       db:"booked_at"`
	CreatedAt      time.Time     `json:"created_at"      db:"created_at"`
	UpdatedAt      time.Time     `json:"updated_at"      db:"updated_at"`
}

// Commission represents the commission earned by an agent on a booking.
type Commission struct {
	ID        uuid.UUID        `json:"id"         db:"id"`
	BookingID uuid.UUID        `json:"booking_id" db:"booking_id"`
	AgentID   uuid.UUID        `json:"agent_id"   db:"agent_id"`
	Amount    float64          `json:"amount"     db:"amount"`
	TDS       float64          `json:"tds"        db:"tds"`
	NetAmount float64          `json:"net_amount" db:"net_amount"`
	Status    CommissionStatus `json:"status"     db:"status"`
	PaidAt    *time.Time       `json:"paid_at"    db:"paid_at"`
	CreatedAt time.Time        `json:"created_at" db:"created_at"`
	UpdatedAt time.Time        `json:"updated_at" db:"updated_at"`
}

// Area represents a micro-market or locality tracked for price trends.
type Area struct {
	ID             uuid.UUID       `json:"id"              db:"id"`
	Name           string          `json:"name"            db:"name"`
	Slug           string          `json:"slug"            db:"slug"`
	City           string          `json:"city"            db:"city"`
	State          string          `json:"state"           db:"state"`
	Description    *string         `json:"description"     db:"description"`
	Location       string          `json:"location"        db:"location"`        // PostGIS geometry (POINT or POLYGON)
	PriceTrend     json.RawMessage `json:"price_trend"     db:"price_trend"`     // JSONB
	Infrastructure json.RawMessage `json:"infrastructure"  db:"infrastructure"`  // JSONB
	CreatedAt      time.Time       `json:"created_at"      db:"created_at"`
	UpdatedAt      time.Time       `json:"updated_at"      db:"updated_at"`
}

// SiteVisit records a physical site visit by an agent with an investor.
type SiteVisit struct {
	ID        uuid.UUID       `json:"id"         db:"id"`
	LeadID    uuid.UUID       `json:"lead_id"    db:"lead_id"`
	AgentID   uuid.UUID       `json:"agent_id"   db:"agent_id"`
	ProjectID uuid.UUID       `json:"project_id" db:"project_id"`
	Feedback  *string         `json:"feedback"   db:"feedback"`
	Photos    json.RawMessage `json:"photos"     db:"photos"`    // JSONB array of URLs
	Duration  *int            `json:"duration"   db:"duration"`  // minutes
	Rating    *int            `json:"rating"     db:"rating"`    // 1-5
	VisitedAt time.Time       `json:"visited_at" db:"visited_at"`
	CreatedAt time.Time       `json:"created_at" db:"created_at"`
}

// Event is an immutable audit-log entry for every significant action.
type Event struct {
	ID         uuid.UUID       `json:"id"          db:"id"`
	ActorID    uuid.UUID       `json:"actor_id"    db:"actor_id"`
	ActorRole  string          `json:"actor_role"   db:"actor_role"`
	Action     string          `json:"action"       db:"action"`
	EntityType string          `json:"entity_type"  db:"entity_type"`
	EntityID   uuid.UUID       `json:"entity_id"    db:"entity_id"`
	Payload    json.RawMessage `json:"payload"      db:"payload"` // JSONB
	CreatedAt  time.Time       `json:"created_at"   db:"created_at"`
}
