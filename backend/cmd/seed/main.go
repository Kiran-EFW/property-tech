package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

// ---------------------------------------------------------------------------
// Deterministic UUIDs
// ---------------------------------------------------------------------------

// Users
var (
	userAdmin   = uuid.MustParse("a0000000-0000-0000-0000-000000000001")
	userRajesh  = uuid.MustParse("a0000000-0000-0000-0000-000000000002")
	userPriya   = uuid.MustParse("a0000000-0000-0000-0000-000000000003")
	userAmit    = uuid.MustParse("a0000000-0000-0000-0000-000000000004")
	userVikram  = uuid.MustParse("a0000000-0000-0000-0000-000000000005")
	userNeha    = uuid.MustParse("a0000000-0000-0000-0000-000000000006")
)

// Builders
var (
	builderLodha      = uuid.MustParse("b0000000-0000-0000-0000-000000000001")
	builderGodrej     = uuid.MustParse("b0000000-0000-0000-0000-000000000002")
	builderHiranandani = uuid.MustParse("b0000000-0000-0000-0000-000000000003")
	builderTata       = uuid.MustParse("b0000000-0000-0000-0000-000000000004")
	builderShapoorji  = uuid.MustParse("b0000000-0000-0000-0000-000000000005")
)

// Areas
var (
	areaPanvel      = uuid.MustParse("c0000000-0000-0000-0000-000000000001")
	areaDombivli    = uuid.MustParse("c0000000-0000-0000-0000-000000000002")
	areaKalyan      = uuid.MustParse("c0000000-0000-0000-0000-000000000003")
	areaUlwe        = uuid.MustParse("c0000000-0000-0000-0000-000000000004")
	areaKharghar    = uuid.MustParse("c0000000-0000-0000-0000-000000000005")
	areaTaloja      = uuid.MustParse("c0000000-0000-0000-0000-000000000006")
)

// Projects
var (
	projLodhaPalava    = uuid.MustParse("d0000000-0000-0000-0000-000000000001")
	projGodrejEmerald  = uuid.MustParse("d0000000-0000-0000-0000-000000000002")
	projHiranandaniFortune = uuid.MustParse("d0000000-0000-0000-0000-000000000003")
	projTataSerein     = uuid.MustParse("d0000000-0000-0000-0000-000000000004")
	projShapoorjiPark  = uuid.MustParse("d0000000-0000-0000-0000-000000000005")
	projLodhaBelmondo  = uuid.MustParse("d0000000-0000-0000-0000-000000000006")
	projGodrejNirvaan  = uuid.MustParse("d0000000-0000-0000-0000-000000000007")
	projHiranandaniZen = uuid.MustParse("d0000000-0000-0000-0000-000000000008")
	projTataVivati     = uuid.MustParse("d0000000-0000-0000-0000-000000000009")
	projShapoorjiVirar = uuid.MustParse("d0000000-0000-0000-0000-000000000010")
)

// Agents
var (
	agentRajesh = uuid.MustParse("e0000000-0000-0000-0000-000000000001")
	agentPriya  = uuid.MustParse("e0000000-0000-0000-0000-000000000002")
	agentAmit   = uuid.MustParse("e0000000-0000-0000-0000-000000000003")
)

// Project Units (40 units, numbered sequentially)
func unitID(n int) uuid.UUID {
	return uuid.MustParse(fmt.Sprintf("f0000000-0000-0000-0000-%012d", n))
}

// Leads
func leadID(n int) uuid.UUID {
	return uuid.MustParse(fmt.Sprintf("10000000-0000-0000-0000-%012d", n))
}

// Bookings
var (
	booking1 = uuid.MustParse("20000000-0000-0000-0000-000000000001")
	booking2 = uuid.MustParse("20000000-0000-0000-0000-000000000002")
	booking3 = uuid.MustParse("20000000-0000-0000-0000-000000000003")
)

// Commissions
var (
	commission1 = uuid.MustParse("30000000-0000-0000-0000-000000000001")
	commission2 = uuid.MustParse("30000000-0000-0000-0000-000000000002")
	commission3 = uuid.MustParse("30000000-0000-0000-0000-000000000003")
)

// ---------------------------------------------------------------------------
// Helper: raw JSON
// ---------------------------------------------------------------------------

func rawJSON(v interface{}) json.RawMessage {
	b, err := json.Marshal(v)
	if err != nil {
		log.Fatalf("json marshal: %v", err)
	}
	return b
}

// ---------------------------------------------------------------------------
// Main
// ---------------------------------------------------------------------------

func main() {
	// Load .env (ignore error if missing)
	_ = godotenv.Load()

	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL environment variable is required")
	}

	ctx := context.Background()

	pool, err := pgxpool.New(ctx, dbURL)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}
	defer pool.Close()

	if err := pool.Ping(ctx); err != nil {
		log.Fatalf("Unable to ping database: %v", err)
	}
	fmt.Println("Connected to database successfully")

	seedUsers(ctx, pool)
	seedBuilders(ctx, pool)
	seedAreas(ctx, pool)
	seedProjects(ctx, pool)
	seedProjectUnits(ctx, pool)
	seedAgents(ctx, pool)
	seedLeads(ctx, pool)
	seedBookings(ctx, pool)
	seedCommissions(ctx, pool)

	fmt.Println("\nSeed completed successfully!")
}

// ---------------------------------------------------------------------------
// 1. Users
// ---------------------------------------------------------------------------

func seedUsers(ctx context.Context, pool *pgxpool.Pool) {
	fmt.Print("Seeding users...")

	type user struct {
		ID           uuid.UUID
		Phone        string
		Email        string
		Name         string
		Role         string
		PasswordHash string
		IsActive     bool
		IsNRI        bool
	}

	// password_hash is a bcrypt hash of "password123" for dev use
	devHash := "$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy"

	users := []user{
		{userAdmin, "+919999000001", "admin@proptech.in", "Admin User", "admin", devHash, true, false},
		{userRajesh, "+919999000002", "rajesh.kumar@proptech.in", "Rajesh Kumar", "agent", devHash, true, false},
		{userPriya, "+919999000003", "priya.sharma@proptech.in", "Priya Sharma", "agent", devHash, true, false},
		{userAmit, "+919999000004", "amit.patel@proptech.in", "Amit Patel", "agent", devHash, true, false},
		{userVikram, "+919999000005", "vikram.singh@investor.in", "Vikram Singh", "investor", devHash, true, true},
		{userNeha, "+919999000006", "neha.gupta@investor.in", "Neha Gupta", "investor", devHash, true, false},
	}

	query := `
		INSERT INTO users (id, phone, email, name, role, password_hash, is_active, is_nri, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, now(), now())
		ON CONFLICT (id) DO NOTHING`

	for _, u := range users {
		if _, err := pool.Exec(ctx, query, u.ID, u.Phone, u.Email, u.Name, u.Role, u.PasswordHash, u.IsActive, u.IsNRI); err != nil {
			log.Fatalf("seed users: %v", err)
		}
	}
	fmt.Println(" done (6 users)")
}

// ---------------------------------------------------------------------------
// 2. Builders
// ---------------------------------------------------------------------------

func seedBuilders(ctx context.Context, pool *pgxpool.Pool) {
	fmt.Print("Seeding builders...")

	type builder struct {
		ID           uuid.UUID
		Name         string
		Slug         string
		ReraNumber   string
		PAN          string
		GST          string
		TrackRecord  string
		ContactPhone string
		ContactEmail string
		IsVerified   bool
	}

	builders := []builder{
		{builderLodha, "Lodha Group", "lodha-group", "P52000001234",
			"AABCL1234A", "27AABCL1234A1Z5",
			"30+ years, 50M+ sqft delivered across Mumbai Metropolitan Region",
			"+912248001234", "info@lodhagroup.com", true},
		{builderGodrej, "Godrej Properties", "godrej-properties", "P52000001235",
			"AABCG5678B", "27AABCG5678B1Z8",
			"Godrej legacy since 1897, 20M+ sqft delivered, presence in 12 cities",
			"+912267501235", "sales@godrejproperties.com", true},
		{builderHiranandani, "Hiranandani Group", "hiranandani-group", "P52000001236",
			"AABCH9012C", "27AABCH9012C1Z3",
			"35+ years, pioneers of Powai and Thane townships, 25M+ sqft delivered",
			"+912225701236", "contact@hiranandani.com", true},
		{builderTata, "Tata Housing", "tata-housing", "P52000001237",
			"AABCT3456D", "27AABCT3456D1Z6",
			"Tata Group company, premium developments across 8 cities, 25M+ sqft",
			"+912266001237", "sales@tatahousing.in", true},
		{builderShapoorji, "Shapoorji Pallonji", "shapoorji-pallonji", "P52000001238",
			"AABCS7890E", "27AABCS7890E1Z1",
			"150+ years legacy, 50M+ sqft across residential and commercial",
			"+912267801238", "realestate@shapoorji.in", true},
	}

	query := `
		INSERT INTO builders (id, name, slug, rera_number, pan, gst, track_record, contact_phone, contact_email, is_verified, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, now(), now())
		ON CONFLICT (id) DO NOTHING`

	for _, b := range builders {
		if _, err := pool.Exec(ctx, query,
			b.ID, b.Name, b.Slug, b.ReraNumber, b.PAN, b.GST,
			b.TrackRecord, b.ContactPhone, b.ContactEmail, b.IsVerified,
		); err != nil {
			log.Fatalf("seed builders: %v", err)
		}
	}
	fmt.Println(" done (5 builders)")
}

// ---------------------------------------------------------------------------
// 3. Areas
// ---------------------------------------------------------------------------

type priceTrend struct {
	Year   int `json:"year"`
	AvgPSF int `json:"avg_psf"`
}

type infrastructure struct {
	Metro     string   `json:"metro"`
	Airport   string   `json:"airport"`
	Highways  []string `json:"highways"`
	Hospitals []string `json:"hospitals"`
	Schools   []string `json:"schools"`
}

func seedAreas(ctx context.Context, pool *pgxpool.Pool) {
	fmt.Print("Seeding areas...")

	type area struct {
		ID             uuid.UUID
		Name           string
		Slug           string
		City           string
		State          string
		Description    string
		Lng            float64
		Lat            float64
		PriceTrend     []priceTrend
		Infrastructure infrastructure
	}

	areas := []area{
		{
			areaPanvel, "Panvel", "panvel", "Navi Mumbai", "Maharashtra",
			"Panvel is a rapidly developing CIDCO node in Navi Mumbai, strategically located near the upcoming Navi Mumbai International Airport. With excellent connectivity via the Mumbai-Pune Expressway and Sion-Panvel Expressway, it has become one of the most sought-after residential destinations in the Mumbai Metropolitan Region.",
			73.1059, 18.9894,
			[]priceTrend{
				{2020, 5200}, {2021, 5600}, {2022, 6100}, {2023, 6800}, {2024, 7500}, {2025, 8200},
			},
			infrastructure{
				Metro:     "Metro Line 1 (Belapur-Pendhar, operational by 2026)",
				Airport:   "Navi Mumbai International Airport (8 km, under construction)",
				Highways:  []string{"Mumbai-Pune Expressway", "Sion-Panvel Expressway", "NH-4"},
				Hospitals: []string{"MGM Hospital Panvel", "Apollo Clinic Panvel", "Vashi General Hospital"},
				Schools:   []string{"DY Patil International School", "Ryan International Panvel", "DAV Public School"},
			},
		},
		{
			areaDombivli, "Dombivli East", "dombivli-east", "Thane", "Maharashtra",
			"Dombivli East is an established residential hub with excellent rail connectivity to Mumbai. The upcoming Lodha Palava township and infrastructure development have positioned it as a premium affordable housing destination in the Thane-Mumbai corridor.",
			73.0863, 19.2094,
			[]priceTrend{
				{2020, 4800}, {2021, 5100}, {2022, 5500}, {2023, 6000}, {2024, 6700}, {2025, 7400},
			},
			infrastructure{
				Metro:     "Metro Line 4 extension (proposed)",
				Airport:   "Navi Mumbai International Airport (35 km)",
				Highways:  []string{"Kalyan-Shilphata Road", "NH-4 via Bhiwandi Bypass"},
				Hospitals: []string{"Acharya Shrimannarayan Hospital", "Lifeline Hospital", "Dombivli Medical Centre"},
				Schools:   []string{"Lodha World School", "Shree Saraswati Vidyalaya", "Swami Vivekanand School"},
			},
		},
		{
			areaKalyan, "Kalyan-Shilphata", "kalyan-shilphata", "Thane", "Maharashtra",
			"Kalyan-Shilphata corridor is emerging as a major growth zone in the Thane district. With the Kalyan-Dombivli-Shilphata road widening and proximity to the proposed metro corridor, this area offers excellent value-for-money residential options.",
			73.1322, 19.2437,
			[]priceTrend{
				{2020, 4200}, {2021, 4500}, {2022, 4900}, {2023, 5400}, {2024, 6000}, {2025, 6600},
			},
			infrastructure{
				Metro:     "Metro Line 12 (Kalyan-Taloja, proposed)",
				Airport:   "Navi Mumbai International Airport (40 km)",
				Highways:  []string{"Kalyan-Shilphata Road", "Mumbai-Nashik Highway"},
				Hospitals: []string{"Fortis Hospital Kalyan", "Kalyan City Hospital", "Vedant Hospital"},
				Schools:   []string{"Billabong High International", "Orchid International School", "St. Mary's Kalyan"},
			},
		},
		{
			areaUlwe, "Ulwe", "ulwe", "Navi Mumbai", "Maharashtra",
			"Ulwe is the closest residential node to the upcoming Navi Mumbai International Airport. With CIDCO's planned infrastructure including a coastal road and metro connectivity, Ulwe is positioned as the next premium micro-market in Navi Mumbai.",
			73.0297, 18.9768,
			[]priceTrend{
				{2020, 5500}, {2021, 5900}, {2022, 6400}, {2023, 7100}, {2024, 7900}, {2025, 8700},
			},
			infrastructure{
				Metro:     "Metro Line 1 (Belapur-Pendhar, station at Ulwe)",
				Airport:   "Navi Mumbai International Airport (3 km)",
				Highways:  []string{"Sion-Panvel Expressway", "Mumbai Trans Harbour Link (MTHL)"},
				Hospitals: []string{"NMMC Hospital Ulwe", "Seven Hills Clinic"},
				Schools:   []string{"Podar International School Ulwe", "DY Patil School Ulwe"},
			},
		},
		{
			areaKharghar, "Kharghar", "kharghar", "Navi Mumbai", "Maharashtra",
			"Kharghar is one of the most developed and premium nodes in Navi Mumbai. Known for the Central Park, Golf Course, and excellent social infrastructure, it commands the highest property prices in the Navi Mumbai corridor. Home to major IT parks and corporate offices.",
			73.0785, 19.0330,
			[]priceTrend{
				{2020, 8000}, {2021, 8500}, {2022, 9200}, {2023, 9800}, {2024, 10500}, {2025, 11200},
			},
			infrastructure{
				Metro:     "Metro Line 1 (Belapur-Pendhar, station at Kharghar)",
				Airport:   "Navi Mumbai International Airport (12 km)",
				Highways:  []string{"Sion-Panvel Expressway", "Palm Beach Road"},
				Hospitals: []string{"Kharghar CIDCO Hospital", "Apollo Hospital Kharghar", "MGM Hospital Vashi"},
				Schools:   []string{"DY Patil International School", "Apeejay School Kharghar", "Ryan International Kharghar"},
			},
		},
		{
			areaTaloja, "Taloja", "taloja", "Navi Mumbai", "Maharashtra",
			"Taloja is an emerging affordable housing destination in Navi Mumbai with CIDCO-planned infrastructure. The upcoming Taloja MIDC expansion and metro connectivity are expected to drive significant appreciation. Currently offers the most competitive pricing in the Navi Mumbai corridor.",
			73.1090, 19.0628,
			[]priceTrend{
				{2020, 3800}, {2021, 4100}, {2022, 4500}, {2023, 5000}, {2024, 5600}, {2025, 6200},
			},
			infrastructure{
				Metro:     "Metro Line 12 (Kalyan-Taloja, proposed)",
				Airport:   "Navi Mumbai International Airport (15 km)",
				Highways:  []string{"Panvel-Matheran Road", "Sion-Panvel Expressway"},
				Hospitals: []string{"Taloja CIDCO Hospital", "Medipoint Hospital"},
				Schools:   []string{"Kendriya Vidyalaya Taloja", "St. Xavier's Taloja"},
			},
		},
	}

	query := `
		INSERT INTO areas (id, name, slug, city, state, description, location, price_trend, infrastructure, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, ST_SetSRID(ST_MakePoint($7, $8), 4326), $9, $10, now(), now())
		ON CONFLICT (id) DO NOTHING`

	for _, a := range areas {
		if _, err := pool.Exec(ctx, query,
			a.ID, a.Name, a.Slug, a.City, a.State, a.Description,
			a.Lng, a.Lat, rawJSON(a.PriceTrend), rawJSON(a.Infrastructure),
		); err != nil {
			log.Fatalf("seed areas: %v", err)
		}
	}
	fmt.Println(" done (6 areas)")
}

// ---------------------------------------------------------------------------
// 4. Projects
// ---------------------------------------------------------------------------

func seedProjects(ctx context.Context, pool *pgxpool.Pool) {
	fmt.Print("Seeding projects...")

	type project struct {
		ID           uuid.UUID
		Name         string
		Slug         string
		ReraNumber   string
		BuilderID    uuid.UUID
		Description  string
		CarpetMin    float64
		CarpetMax    float64
		PriceMin     float64
		PriceMax     float64
		Lng          float64
		Lat          float64
		Address      string
		City         string
		State        string
		Pincode      string
		Status       string
		Amenities    []string
	}

	projects := []project{
		{
			projLodhaPalava, "Lodha Palava City", "lodha-palava-city", "P52000010001",
			builderLodha,
			"Lodha Palava City is a fully integrated smart township spread across 4500 acres in Dombivli. It features world-class amenities, schools, hospitals, and commercial spaces, making it a self-sustaining urban ecosystem.",
			400, 1200, 4500000, 15000000, 73.0870, 19.2100,
			"Lodha Palava City, Dombivli East, Thane", "Thane", "Maharashtra", "421204",
			"active",
			[]string{"Swimming Pool", "Gym", "Clubhouse", "Children Play Area", "Garden", "Jogging Track", "24/7 Security", "Power Backup", "Smart Home Features", "Shopping Complex"},
		},
		{
			projGodrejEmerald, "Godrej Emerald", "godrej-emerald", "P52000010002",
			builderGodrej,
			"Godrej Emerald is a premium residential project in Panvel offering luxurious 1, 2, and 3 BHK apartments. Strategically located near the upcoming Navi Mumbai International Airport with excellent connectivity.",
			500, 1000, 5500000, 12000000, 73.1065, 18.9900,
			"Godrej Emerald, Panvel, Navi Mumbai", "Navi Mumbai", "Maharashtra", "410206",
			"active",
			[]string{"Swimming Pool", "Gym", "Clubhouse", "Children Play Area", "Landscaped Garden", "Jogging Track", "24/7 Security", "Power Backup", "Indoor Games Room"},
		},
		{
			projHiranandaniFortune, "Hiranandani Fortune City", "hiranandani-fortune-city", "P52000010003",
			builderHiranandani,
			"Hiranandani Fortune City is a premium township project in Panvel featuring signature Hiranandani architecture. Spread across 600 acres, it offers a blend of residential and commercial spaces with world-class infrastructure.",
			450, 1100, 6000000, 14000000, 73.1050, 18.9880,
			"Hiranandani Fortune City, Panvel, Navi Mumbai", "Navi Mumbai", "Maharashtra", "410206",
			"active",
			[]string{"Swimming Pool", "Gym", "Clubhouse", "Children Play Area", "Garden", "Jogging Track", "24/7 Security", "Power Backup", "Tennis Court", "Amphitheatre"},
		},
		{
			projTataSerein, "Tata Serein", "tata-serein", "P52000010004",
			builderTata,
			"Tata Serein is a premium residential development in Panvel by Tata Housing. It offers thoughtfully designed 2 and 3 BHK apartments with panoramic views and top-of-the-line amenities.",
			550, 1200, 6500000, 18000000, 73.1070, 18.9910,
			"Tata Serein, Panvel, Navi Mumbai", "Navi Mumbai", "Maharashtra", "410206",
			"active",
			[]string{"Swimming Pool", "Gym", "Clubhouse", "Children Play Area", "Garden", "Jogging Track", "24/7 Security", "Power Backup", "Squash Court", "Business Centre"},
		},
		{
			projShapoorjiPark, "Shapoorji Park Central", "shapoorji-park-central", "P52000010005",
			builderShapoorji,
			"Shapoorji Park Central in Kalyan-Shilphata offers affordable luxury living with modern amenities. It features well-designed 1 and 2 BHK apartments in a green, well-planned township setting.",
			350, 900, 3500000, 9500000, 73.1330, 19.2440,
			"Shapoorji Park Central, Kalyan-Shilphata, Thane", "Thane", "Maharashtra", "421301",
			"active",
			[]string{"Swimming Pool", "Gym", "Clubhouse", "Children Play Area", "Garden", "Jogging Track", "24/7 Security", "Power Backup"},
		},
		{
			projLodhaBelmondo, "Lodha Belmondo", "lodha-belmondo", "P52000010006",
			builderLodha,
			"Lodha Belmondo is an ultra-premium golf-course living project in Kharghar offering spacious 2, 3, and 4 BHK residences. Features an 18-hole golf course, international school, and resort-style clubhouse.",
			600, 1500, 8000000, 25000000, 73.0790, 19.0335,
			"Lodha Belmondo, Kharghar, Navi Mumbai", "Navi Mumbai", "Maharashtra", "410210",
			"active",
			[]string{"Golf Course", "Swimming Pool", "Gym", "Clubhouse", "Children Play Area", "Landscaped Garden", "Jogging Track", "24/7 Security", "Power Backup", "Spa", "Concierge Service"},
		},
		{
			projGodrejNirvaan, "Godrej Nirvaan", "godrej-nirvaan", "P52000010007",
			builderGodrej,
			"Godrej Nirvaan in Kalyan offers well-crafted 1 and 2 BHK homes designed for modern families. With Godrej's signature quality and thoughtful planning, it brings premium living to the Kalyan corridor.",
			400, 950, 4000000, 10000000, 73.1325, 19.2435,
			"Godrej Nirvaan, Kalyan-Shilphata, Thane", "Thane", "Maharashtra", "421301",
			"active",
			[]string{"Swimming Pool", "Gym", "Clubhouse", "Children Play Area", "Garden", "Jogging Track", "24/7 Security", "Power Backup", "Multi-purpose Court"},
		},
		{
			projHiranandaniZen, "Hiranandani Zen", "hiranandani-zen", "P52000010008",
			builderHiranandani,
			"Hiranandani Zen in Ulwe brings the signature Hiranandani lifestyle to the airport-adjacent micro-market. Currently under construction, it offers 1, 2, and 3 BHK apartments with stunning views of the creek and upcoming airport.",
			500, 1100, 5000000, 13000000, 73.0300, 18.9770,
			"Hiranandani Zen, Ulwe, Navi Mumbai", "Navi Mumbai", "Maharashtra", "410206",
			"active",
			[]string{"Swimming Pool", "Gym", "Clubhouse", "Children Play Area", "Garden", "Jogging Track", "24/7 Security", "Power Backup", "Yoga Deck", "Sky Lounge"},
		},
		{
			projTataVivati, "Tata Vivati", "tata-vivati", "P52000010009",
			builderTata,
			"Tata Vivati in Taloja is an upcoming pre-launch project by Tata Housing. Designed for the aspiring homebuyer, it will offer compact and efficient 1 and 2 BHK apartments at competitive price points.",
			380, 850, 3000000, 7500000, 73.1095, 19.0630,
			"Tata Vivati, Taloja, Navi Mumbai", "Navi Mumbai", "Maharashtra", "410208",
			"draft",
			[]string{"Swimming Pool", "Gym", "Clubhouse", "Children Play Area", "Garden", "24/7 Security", "Power Backup"},
		},
		{
			projShapoorjiVirar, "Shapoorji Joyville", "shapoorji-joyville", "P52000010010",
			builderShapoorji,
			"Shapoorji Joyville in Taloja offers thoughtfully designed 1 and 2 BHK apartments in the affordable segment. With Shapoorji Pallonji's 150-year legacy of trust and quality, it promises value-for-money homes.",
			350, 800, 2800000, 7000000, 73.1085, 19.0625,
			"Shapoorji Joyville, Taloja, Navi Mumbai", "Navi Mumbai", "Maharashtra", "410208",
			"active",
			[]string{"Swimming Pool", "Gym", "Clubhouse", "Children Play Area", "Garden", "24/7 Security", "Power Backup", "Co-working Space"},
		},
	}

	query := `
		INSERT INTO projects (id, name, slug, rera_number, builder_id, description, carpet_area_min, carpet_area_max,
			price_min, price_max, location, address, city, state, pincode, status, amenities, media, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10,
			ST_SetSRID(ST_MakePoint($11, $12), 4326),
			$13, $14, $15, $16, $17, $18, '[]'::jsonb, now(), now())
		ON CONFLICT (id) DO NOTHING`

	for _, p := range projects {
		if _, err := pool.Exec(ctx, query,
			p.ID, p.Name, p.Slug, p.ReraNumber, p.BuilderID, p.Description,
			p.CarpetMin, p.CarpetMax, p.PriceMin, p.PriceMax,
			p.Lng, p.Lat, p.Address, p.City, p.State, p.Pincode,
			p.Status, rawJSON(p.Amenities),
		); err != nil {
			log.Fatalf("seed projects: %v", err)
		}
	}
	fmt.Println(" done (10 projects)")
}

// ---------------------------------------------------------------------------
// 5. Project Units
// ---------------------------------------------------------------------------

func seedProjectUnits(ctx context.Context, pool *pgxpool.Pool) {
	fmt.Print("Seeding project units...")

	type unit struct {
		ID         uuid.UUID
		ProjectID  uuid.UUID
		Floor      int
		UnitNumber string
		UnitType   string
		CarpetArea float64
		Price      float64
		Status     string
	}

	units := []unit{
		// Lodha Palava City - 4 units
		{unitID(1), projLodhaPalava, 3, "A-301", "1 BHK", 420, 4800000, "available"},
		{unitID(2), projLodhaPalava, 5, "B-501", "2 BHK", 680, 7800000, "available"},
		{unitID(3), projLodhaPalava, 8, "C-801", "3 BHK", 1050, 12000000, "booked"},
		{unitID(4), projLodhaPalava, 12, "D-1201", "2 BHK", 720, 8500000, "sold"},

		// Godrej Emerald - 4 units
		{unitID(5), projGodrejEmerald, 2, "A-201", "1 BHK", 520, 5800000, "available"},
		{unitID(6), projGodrejEmerald, 7, "B-701", "2 BHK", 750, 8500000, "available"},
		{unitID(7), projGodrejEmerald, 10, "C-1001", "3 BHK", 980, 11200000, "booked"},
		{unitID(8), projGodrejEmerald, 4, "A-401", "2 BHK", 700, 7800000, "available"},

		// Hiranandani Fortune City - 4 units
		{unitID(9), projHiranandaniFortune, 6, "T1-601", "1 BHK", 480, 6200000, "available"},
		{unitID(10), projHiranandaniFortune, 9, "T2-901", "2 BHK", 780, 9500000, "sold"},
		{unitID(11), projHiranandaniFortune, 14, "T1-1401", "3 BHK", 1080, 13500000, "available"},
		{unitID(12), projHiranandaniFortune, 3, "T2-301", "2 BHK", 740, 8800000, "available"},

		// Tata Serein - 4 units
		{unitID(13), projTataSerein, 5, "W1-501", "2 BHK", 620, 7200000, "available"},
		{unitID(14), projTataSerein, 11, "W2-1101", "3 BHK", 1100, 16500000, "available"},
		{unitID(15), projTataSerein, 8, "E1-801", "2 BHK", 680, 8000000, "booked"},
		{unitID(16), projTataSerein, 15, "E2-1501", "3 BHK", 1180, 17800000, "sold"},

		// Shapoorji Park Central - 4 units
		{unitID(17), projShapoorjiPark, 2, "A-201", "1 BHK", 370, 3800000, "available"},
		{unitID(18), projShapoorjiPark, 4, "B-401", "2 BHK", 650, 6500000, "available"},
		{unitID(19), projShapoorjiPark, 7, "A-701", "1 BHK", 380, 3900000, "available"},
		{unitID(20), projShapoorjiPark, 10, "B-1001", "2 BHK", 680, 7000000, "sold"},

		// Lodha Belmondo - 4 units
		{unitID(21), projLodhaBelmondo, 4, "G1-401", "2 BHK", 650, 9000000, "available"},
		{unitID(22), projLodhaBelmondo, 8, "G2-801", "3 BHK", 1100, 16000000, "available"},
		{unitID(23), projLodhaBelmondo, 12, "G1-1201", "3 BHK", 1350, 22000000, "booked"},
		{unitID(24), projLodhaBelmondo, 16, "G2-1601", "3 BHK", 1480, 24500000, "available"},

		// Godrej Nirvaan - 4 units
		{unitID(25), projGodrejNirvaan, 3, "A-301", "1 BHK", 420, 4200000, "available"},
		{unitID(26), projGodrejNirvaan, 6, "B-601", "2 BHK", 700, 7200000, "sold"},
		{unitID(27), projGodrejNirvaan, 9, "A-901", "2 BHK", 720, 7500000, "available"},
		{unitID(28), projGodrejNirvaan, 5, "B-501", "1 BHK", 440, 4500000, "available"},

		// Hiranandani Zen - 4 units
		{unitID(29), projHiranandaniZen, 4, "Z1-401", "1 BHK", 520, 5200000, "available"},
		{unitID(30), projHiranandaniZen, 8, "Z2-801", "2 BHK", 780, 8500000, "available"},
		{unitID(31), projHiranandaniZen, 11, "Z1-1101", "3 BHK", 1050, 12500000, "available"},
		{unitID(32), projHiranandaniZen, 6, "Z2-601", "2 BHK", 740, 8000000, "booked"},

		// Tata Vivati - 4 units
		{unitID(33), projTataVivati, 2, "V1-201", "1 BHK", 390, 3200000, "available"},
		{unitID(34), projTataVivati, 5, "V2-501", "2 BHK", 650, 5500000, "available"},
		{unitID(35), projTataVivati, 3, "V1-301", "1 BHK", 400, 3300000, "available"},
		{unitID(36), projTataVivati, 7, "V2-701", "2 BHK", 680, 5800000, "available"},

		// Shapoorji Joyville - 4 units
		{unitID(37), projShapoorjiVirar, 3, "J1-301", "1 BHK", 360, 2900000, "available"},
		{unitID(38), projShapoorjiVirar, 6, "J2-601", "2 BHK", 620, 5200000, "available"},
		{unitID(39), projShapoorjiVirar, 4, "J1-401", "1 BHK", 370, 3000000, "sold"},
		{unitID(40), projShapoorjiVirar, 8, "J2-801", "2 BHK", 650, 5500000, "available"},
	}

	query := `
		INSERT INTO project_units (id, project_id, floor, unit_number, unit_type, carpet_area, price, status, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, now(), now())
		ON CONFLICT (id) DO NOTHING`

	for _, u := range units {
		if _, err := pool.Exec(ctx, query,
			u.ID, u.ProjectID, u.Floor, u.UnitNumber, u.UnitType, u.CarpetArea, u.Price, u.Status,
		); err != nil {
			log.Fatalf("seed project_units: %v", err)
		}
	}
	fmt.Println(" done (40 units)")
}

// ---------------------------------------------------------------------------
// 6. Agents
// ---------------------------------------------------------------------------

func seedAgents(ctx context.Context, pool *pgxpool.Pool) {
	fmt.Print("Seeding agents...")

	type agent struct {
		ID         uuid.UUID
		UserID     uuid.UUID
		ReraNumber string
		PAN        string
		GST        string
		Tier       string
		IsActive   bool
	}

	agents := []agent{
		{agentRajesh, userRajesh, "RERA/AGT/001", "ABCPK1234A", "27ABCPK1234A1Z5", "gold", true},
		{agentPriya, userPriya, "RERA/AGT/002", "ABCPS5678B", "27ABCPS5678B1Z8", "silver", true},
		{agentAmit, userAmit, "RERA/AGT/003", "ABCPP9012C", "27ABCPP9012C1Z3", "bronze", true},
	}

	query := `
		INSERT INTO agents (id, user_id, rera_number, pan, gst, tier, is_active, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, now(), now())
		ON CONFLICT (id) DO NOTHING`

	for _, a := range agents {
		if _, err := pool.Exec(ctx, query,
			a.ID, a.UserID, a.ReraNumber, a.PAN, a.GST, a.Tier, a.IsActive,
		); err != nil {
			log.Fatalf("seed agents: %v", err)
		}
	}
	fmt.Println(" done (3 agents)")
}

// ---------------------------------------------------------------------------
// 7. Leads
// ---------------------------------------------------------------------------

func seedLeads(ctx context.Context, pool *pgxpool.Pool) {
	fmt.Print("Seeding leads...")

	type lead struct {
		ID         uuid.UUID
		InvestorID *uuid.UUID
		ProjectID  uuid.UUID
		AgentID    *uuid.UUID
		Source     string
		Status     string
		Phone      string
		Name       string
		Email      string
		Budget     float64
		Notes      string
		FollowUpAt *time.Time
	}

	// helper for optional UUIDs
	ref := func(id uuid.UUID) *uuid.UUID { return &id }
	ts := func(days int) *time.Time {
		t := time.Now().AddDate(0, 0, days)
		return &t
	}

	leads := []lead{
		// Converted leads (these will have bookings)
		{leadID(1), ref(userVikram), projLodhaPalava, ref(agentRajesh), "website", "converted",
			"+919876543001", "Vikram Singh", "vikram.singh@investor.in", 10000000,
			"NRI investor interested in 2BHK for investment purposes. Ready to book.", nil},
		{leadID(2), ref(userNeha), projGodrejEmerald, ref(agentPriya), "referral", "converted",
			"+919876543002", "Neha Gupta", "neha.gupta@investor.in", 12000000,
			"End user looking for family home. Site visit done, very positive.", nil},
		{leadID(3), ref(userVikram), projLodhaBelmondo, ref(agentAmit), "website", "converted",
			"+919876543003", "Vikram Singh", "vikram.singh@investor.in", 20000000,
			"Second investment property. Premium 3BHK with golf course view.", nil},

		// Active leads in pipeline
		{leadID(4), ref(userNeha), projHiranandaniFortune, ref(agentRajesh), "website", "site_visit",
			"+919876543004", "Neha Gupta", "neha.gupta@investor.in", 14000000,
			"Interested in Fortune City. Site visit scheduled for next week.", ts(3)},
		{leadID(5), nil, projTataSerein, ref(agentPriya), "whatsapp", "qualified",
			"+919876543005", "Sanjay Mehta", "sanjay.mehta@gmail.com", 16000000,
			"WhatsApp inquiry for 3BHK. Budget confirmed, sharing brochure.", ts(2)},
		{leadID(6), nil, projShapoorjiPark, ref(agentAmit), "walk_in", "negotiation",
			"+919876543006", "Anita Desai", "anita.desai@yahoo.com", 7000000,
			"Walk-in at Kalyan office. Negotiating on 2BHK pricing.", ts(1)},
		{leadID(7), nil, projGodrejNirvaan, ref(agentRajesh), "referral", "contacted",
			"+919876543007", "Rahul Joshi", "rahul.joshi@hotmail.com", 8000000,
			"Referred by Vikram. First call done, interested in 2BHK.", ts(5)},
		{leadID(8), nil, projHiranandaniZen, ref(agentPriya), "website", "new",
			"+919876543008", "Deepak Nair", "deepak.nair@gmail.com", 10000000,
			"Website lead for Zen Ulwe. Awaiting first contact.", ts(1)},
		{leadID(9), nil, projTataVivati, ref(agentAmit), "campaign", "new",
			"+919876543009", "Meera Iyer", "meera.iyer@gmail.com", 5000000,
			"Campaign lead from social media ad. Budget segment.", ts(1)},
		{leadID(10), nil, projShapoorjiVirar, ref(agentRajesh), "whatsapp", "contacted",
			"+919876543010", "Suresh Patil", "suresh.patil@gmail.com", 4500000,
			"WhatsApp inquiry for Joyville. Called back, interested.", ts(4)},

		// More mixed status leads
		{leadID(11), nil, projLodhaPalava, ref(agentPriya), "website", "qualified",
			"+919876543011", "Kavita Reddy", "kavita.reddy@gmail.com", 9000000,
			"Looking for 2BHK in Palava. Budget qualified. Scheduling visit.", ts(3)},
		{leadID(12), nil, projGodrejEmerald, ref(agentAmit), "referral", "site_visit",
			"+919876543012", "Manish Tiwari", "manish.tiwari@outlook.com", 11000000,
			"Site visit completed. Very impressed with location. Follow up needed.", ts(2)},
		{leadID(13), nil, projLodhaBelmondo, ref(agentRajesh), "website", "negotiation",
			"+919876543013", "Pooja Agarwal", "pooja.agarwal@gmail.com", 18000000,
			"Premium buyer. Negotiating 3BHK price with Lodha team.", ts(1)},
		{leadID(14), nil, projTataSerein, ref(agentPriya), "walk_in", "qualified",
			"+919876543014", "Ramesh Gupta", "ramesh.gupta@yahoo.com", 15000000,
			"Walk-in at Panvel office. Interested in 3BHK for own use.", ts(4)},
		{leadID(15), nil, projShapoorjiPark, ref(agentAmit), "whatsapp", "new",
			"+919876543015", "Sunita Sharma", "sunita.sharma@gmail.com", 6000000,
			"New inquiry via WhatsApp. First-time homebuyer.", ts(2)},

		// Lost leads
		{leadID(16), nil, projHiranandaniFortune, ref(agentRajesh), "website", "lost",
			"+919876543016", "Arun Saxena", "arun.saxena@gmail.com", 13000000,
			"Lost to competitor project. Budget was a mismatch.", nil},
		{leadID(17), nil, projGodrejNirvaan, ref(agentPriya), "campaign", "lost",
			"+919876543017", "Geeta Pillai", "geeta.pillai@yahoo.com", 7500000,
			"Campaign lead. Not responsive after initial contact.", nil},

		// Additional new leads
		{leadID(18), nil, projHiranandaniZen, ref(agentAmit), "referral", "contacted",
			"+919876543018", "Vivek Kulkarni", "vivek.kulkarni@gmail.com", 9500000,
			"Referral from existing customer. Interested in Ulwe location.", ts(3)},
		{leadID(19), nil, projTataVivati, ref(agentRajesh), "website", "new",
			"+919876543019", "Nisha Verma", "nisha.verma@hotmail.com", 5500000,
			"Website registration for Vivati pre-launch. Budget buyer.", ts(1)},
		{leadID(20), nil, projShapoorjiVirar, ref(agentPriya), "walk_in", "site_visit",
			"+919876543020", "Ajay Bhatt", "ajay.bhatt@gmail.com", 5000000,
			"Walk-in lead. Site visit done. Comparing with other projects.", ts(5)},
	}

	query := `
		INSERT INTO leads (id, investor_id, project_id, agent_id, source, status, phone, name, email, budget, notes, follow_up_at, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, now(), now())
		ON CONFLICT (id) DO NOTHING`

	for _, l := range leads {
		if _, err := pool.Exec(ctx, query,
			l.ID, l.InvestorID, l.ProjectID, l.AgentID, l.Source, l.Status,
			l.Phone, l.Name, l.Email, l.Budget, l.Notes, l.FollowUpAt,
		); err != nil {
			log.Fatalf("seed leads: %v", err)
		}
	}
	fmt.Println(" done (20 leads)")
}

// ---------------------------------------------------------------------------
// 8. Bookings
// ---------------------------------------------------------------------------

func seedBookings(ctx context.Context, pool *pgxpool.Pool) {
	fmt.Print("Seeding bookings...")

	type booking struct {
		ID             uuid.UUID
		ProjectID      uuid.UUID
		UnitID         uuid.UUID
		InvestorID     uuid.UUID
		AgentID        uuid.UUID
		BookingAmount  float64
		AgreementValue float64
		Status         string
	}

	bookings := []booking{
		// Vikram booked Lodha Palava 2BHK (unit D-1201, sold)
		{booking1, projLodhaPalava, unitID(4), userVikram, agentRajesh,
			200000, 8500000, "confirmed"},
		// Neha booked Godrej Emerald 3BHK (unit C-1001, booked)
		{booking2, projGodrejEmerald, unitID(7), userNeha, agentPriya,
			250000, 11200000, "confirmed"},
		// Vikram booked Lodha Belmondo 3BHK (unit G1-1201, booked)
		{booking3, projLodhaBelmondo, unitID(23), userVikram, agentAmit,
			500000, 22000000, "confirmed"},
	}

	query := `
		INSERT INTO bookings (id, project_id, unit_id, investor_id, agent_id, booking_amount, agreement_value, status, booked_at, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, now(), now(), now())
		ON CONFLICT (id) DO NOTHING`

	for _, b := range bookings {
		if _, err := pool.Exec(ctx, query,
			b.ID, b.ProjectID, b.UnitID, b.InvestorID, b.AgentID,
			b.BookingAmount, b.AgreementValue, b.Status,
		); err != nil {
			log.Fatalf("seed bookings: %v", err)
		}
	}
	fmt.Println(" done (3 bookings)")
}

// ---------------------------------------------------------------------------
// 9. Commissions
// ---------------------------------------------------------------------------

func seedCommissions(ctx context.Context, pool *pgxpool.Pool) {
	fmt.Print("Seeding commissions...")

	type commission struct {
		ID        uuid.UUID
		BookingID uuid.UUID
		AgentID   uuid.UUID
		Amount    float64
		TDS       float64
		NetAmount float64
		Status    string
		PaidAt    *time.Time
	}

	now := time.Now()

	// agreement_value * 3% brokerage, TDS = 5% of brokerage
	// Booking 1: 8,500,000 * 3% = 255,000; TDS = 12,750; Net = 242,250
	// Booking 2: 11,200,000 * 3% = 336,000; TDS = 16,800; Net = 319,200
	// Booking 3: 22,000,000 * 3% = 660,000; TDS = 33,000; Net = 627,000

	commissions := []commission{
		{commission1, booking1, agentRajesh, 255000, 12750, 242250, "paid", &now},
		{commission2, booking2, agentPriya, 336000, 16800, 319200, "approved", nil},
		{commission3, booking3, agentAmit, 660000, 33000, 627000, "pending", nil},
	}

	query := `
		INSERT INTO commissions (id, booking_id, agent_id, amount, tds, net_amount, status, paid_at, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, now(), now())
		ON CONFLICT (id) DO NOTHING`

	for _, c := range commissions {
		if _, err := pool.Exec(ctx, query,
			c.ID, c.BookingID, c.AgentID, c.Amount, c.TDS, c.NetAmount, c.Status, c.PaidAt,
		); err != nil {
			log.Fatalf("seed commissions: %v", err)
		}
	}
	fmt.Println(" done (3 commissions)")
}
