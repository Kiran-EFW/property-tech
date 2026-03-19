import {
  GST_UNDER_CONSTRUCTION,
  REGISTRATION_RATE,
  REGISTRATION_MAX,
} from './constants';
import { formatINR } from './formatting';

// ---------------------------------------------------------------------------
// Input / Output Types
// ---------------------------------------------------------------------------

export interface CostCalculatorInput {
  /** Base sale price per square foot (INR) */
  basePricePerSqft: number;

  /** Carpet area in square feet */
  carpetAreaSqft: number;

  /** Floor number (0 = ground). Used to compute floor rise premium */
  floorNumber: number;

  /** Floor rise charge per floor per sqft (INR). Defaults to 0 if omitted */
  floorRisePerSqftPerFloor?: number;

  /** GST rate as a decimal (e.g., 0.05 for 5%). Defaults to 5% under-construction */
  gstRate?: number;

  /** Stamp duty rate as a decimal (e.g., 0.06 for 6%) */
  stampDutyRate: number;

  /** Registration rate as a decimal. Defaults to 1% capped at 30,000 */
  registrationRate?: number;

  /** Registration charge cap in INR. Defaults to 30,000 */
  registrationMax?: number;

  /** Annual maintenance charge per sqft (INR). Defaults to 0 */
  maintenancePerSqft?: number;

  /** Number of months of advance maintenance. Defaults to 24 (2 years) */
  maintenanceMonths?: number;

  /** Legal and documentation charges (INR). Defaults to 0 */
  legalCharges?: number;

  /** Parking cost (INR). Defaults to 0 */
  parkingCost?: number;
}

export interface CostBreakdownItem {
  /** Human-readable label */
  label: string;
  /** Amount in INR */
  amount: number;
  /** Amount formatted in Indian numbering */
  formatted: string;
}

export interface CostBreakdown {
  /** Base cost = base price per sqft x carpet area */
  baseCost: CostBreakdownItem;

  /** Floor rise premium = floor rise per sqft per floor x floor number x carpet area */
  floorRise: CostBreakdownItem;

  /** Effective price per sqft after floor rise */
  effectivePricePerSqft: number;

  /** Agreement value = base cost + floor rise */
  agreementValue: CostBreakdownItem;

  /** GST on agreement value */
  gst: CostBreakdownItem;

  /** Stamp duty on agreement value */
  stampDuty: CostBreakdownItem;

  /** Registration charges (capped) */
  registration: CostBreakdownItem;

  /** Legal and documentation charges */
  legalCharges: CostBreakdownItem;

  /** Advance maintenance deposit */
  maintenance: CostBreakdownItem;

  /** Parking cost */
  parking: CostBreakdownItem;

  /** Total all-inclusive cost */
  total: CostBreakdownItem;

  /** The GST rate used */
  gstRate: number;

  /** The stamp duty rate used */
  stampDutyRate: number;

  /** The registration rate used */
  registrationRate: number;
}

// ---------------------------------------------------------------------------
// Calculator
// ---------------------------------------------------------------------------

function makeItem(label: string, amount: number): CostBreakdownItem {
  return {
    label,
    amount: Math.round(amount),
    formatted: formatINR(Math.round(amount)),
  };
}

/**
 * Calculates all-inclusive property cost with a full breakdown.
 *
 * Returns every component (base, floor rise, GST, stamp duty, registration,
 * maintenance, legal, parking) individually along with the grand total.
 * All amounts are rounded to the nearest rupee and formatted with Indian
 * numbering (lakhs / crores).
 */
export function calculateAllInclusiveCost(
  input: CostCalculatorInput,
): CostBreakdown {
  const {
    basePricePerSqft,
    carpetAreaSqft,
    floorNumber,
    floorRisePerSqftPerFloor = 0,
    gstRate = GST_UNDER_CONSTRUCTION,
    stampDutyRate,
    registrationRate = REGISTRATION_RATE,
    registrationMax = REGISTRATION_MAX,
    maintenancePerSqft = 0,
    maintenanceMonths = 24,
    legalCharges = 0,
    parkingCost = 0,
  } = input;

  // --- Base cost ---
  const baseCostAmount = basePricePerSqft * carpetAreaSqft;

  // --- Floor rise ---
  const floorRiseAmount =
    floorRisePerSqftPerFloor * floorNumber * carpetAreaSqft;

  // --- Agreement value (base on which GST, stamp duty, registration are levied) ---
  const agreementValueAmount = baseCostAmount + floorRiseAmount;

  // --- Effective price per sqft ---
  const effectivePricePerSqft =
    carpetAreaSqft > 0 ? agreementValueAmount / carpetAreaSqft : 0;

  // --- GST ---
  const gstAmount = agreementValueAmount * gstRate;

  // --- Stamp duty ---
  const stampDutyAmount = agreementValueAmount * stampDutyRate;

  // --- Registration (capped) ---
  const registrationComputed = agreementValueAmount * registrationRate;
  const registrationAmount = Math.min(registrationComputed, registrationMax);

  // --- Maintenance deposit ---
  const monthlyMaintenance = maintenancePerSqft * carpetAreaSqft;
  const maintenanceAmount = monthlyMaintenance * maintenanceMonths;

  // --- Total ---
  const totalAmount =
    agreementValueAmount +
    gstAmount +
    stampDutyAmount +
    registrationAmount +
    legalCharges +
    maintenanceAmount +
    parkingCost;

  return {
    baseCost: makeItem('Base Cost', baseCostAmount),
    floorRise: makeItem('Floor Rise Premium', floorRiseAmount),
    effectivePricePerSqft: Math.round(effectivePricePerSqft),
    agreementValue: makeItem('Agreement Value', agreementValueAmount),
    gst: makeItem(`GST (${(gstRate * 100).toFixed(0)}%)`, gstAmount),
    stampDuty: makeItem(
      `Stamp Duty (${(stampDutyRate * 100).toFixed(0)}%)`,
      stampDutyAmount,
    ),
    registration: makeItem('Registration Charges', registrationAmount),
    legalCharges: makeItem('Legal & Documentation', legalCharges),
    maintenance: makeItem(
      `Maintenance (${maintenanceMonths} months advance)`,
      maintenanceAmount,
    ),
    parking: makeItem('Parking', parkingCost),
    total: makeItem('Total All-Inclusive Cost', totalAmount),
    gstRate,
    stampDutyRate,
    registrationRate,
  };
}

/**
 * Estimates monthly EMI using the standard reducing balance formula.
 *
 *   EMI = P x r x (1+r)^n / ((1+r)^n - 1)
 *
 * @param principal - Loan amount in INR
 * @param annualRate - Annual interest rate as decimal (e.g., 0.085 for 8.5%)
 * @param tenureYears - Loan tenure in years
 * @returns Monthly EMI rounded to nearest rupee
 */
export function calculateEMI(
  principal: number,
  annualRate: number,
  tenureYears: number,
): number {
  if (principal <= 0 || annualRate <= 0 || tenureYears <= 0) return 0;

  const monthlyRate = annualRate / 12;
  const months = tenureYears * 12;
  const compounded = Math.pow(1 + monthlyRate, months);

  const emi = (principal * monthlyRate * compounded) / (compounded - 1);
  return Math.round(emi);
}
