// ---------------------------------------------------------------------------
// Indian Numbering Format
// ---------------------------------------------------------------------------

/**
 * Formats a number in Indian numbering system with the Rupee symbol.
 *
 * Indian numbering groups digits as: last 3, then every 2 digits.
 *   1,25,00,000 (1 crore 25 lakhs) instead of 12,500,000
 *
 * @example
 *   formatINR(12500000)  // "₹1,25,00,000"
 *   formatINR(75000)     // "₹75,000"
 *   formatINR(1500.5)    // "₹1,501" (rounds to nearest rupee)
 */
export function formatINR(amount: number): string {
  const rounded = Math.round(amount);
  const isNegative = rounded < 0;
  const absolute = Math.abs(rounded);
  const str = absolute.toString();

  if (str.length <= 3) {
    return `${isNegative ? '-' : ''}₹${str}`;
  }

  // Split into last 3 digits and the rest
  const lastThree = str.slice(-3);
  const remaining = str.slice(0, -3);

  // Group remaining digits in pairs from right
  const pairs: string[] = [];
  for (let i = remaining.length; i > 0; i -= 2) {
    const start = Math.max(0, i - 2);
    pairs.unshift(remaining.slice(start, i));
  }

  const formatted = `${pairs.join(',')},${lastThree}`;
  return `${isNegative ? '-' : ''}₹${formatted}`;
}

/**
 * Formats amount in compact crore notation.
 *
 * @example
 *   formatCrores(12500000)  // "₹1.25 Cr"
 *   formatCrores(250000000) // "₹25 Cr"
 */
export function formatCrores(amount: number): string {
  const crores = amount / 1_00_00_000;
  if (crores >= 100) {
    return `₹${Math.round(crores)} Cr`;
  }
  if (crores >= 10) {
    return `₹${crores.toFixed(1).replace(/\.0$/, '')} Cr`;
  }
  return `₹${crores.toFixed(2).replace(/\.?0+$/, '')} Cr`;
}

/**
 * Formats amount in compact lakh notation.
 *
 * @example
 *   formatLakhs(8500000) // "₹85 L"
 *   formatLakhs(1250000) // "₹12.5 L"
 *   formatLakhs(75000)   // "₹0.75 L"
 */
export function formatLakhs(amount: number): string {
  const lakhs = amount / 1_00_000;
  if (lakhs >= 100) {
    return `₹${Math.round(lakhs)} L`;
  }
  if (lakhs >= 10) {
    return `₹${lakhs.toFixed(1).replace(/\.0$/, '')} L`;
  }
  return `₹${lakhs.toFixed(2).replace(/\.?0+$/, '')} L`;
}

/**
 * Formats amount in the most appropriate compact form.
 * Uses crores for amounts >= 1 crore, lakhs otherwise.
 *
 * @example
 *   formatCompactINR(12500000) // "₹1.25 Cr"
 *   formatCompactINR(8500000)  // "₹85 L"
 *   formatCompactINR(75000)    // "₹75,000"
 */
export function formatCompactINR(amount: number): string {
  const absolute = Math.abs(amount);
  if (absolute >= 1_00_00_000) {
    return formatCrores(amount);
  }
  if (absolute >= 1_00_000) {
    return formatLakhs(amount);
  }
  return formatINR(amount);
}

// ---------------------------------------------------------------------------
// Area Formatting
// ---------------------------------------------------------------------------

/**
 * Formats area in square feet with locale-aware grouping.
 *
 * @example
 *   formatArea(1250)  // "1,250 sq ft"
 *   formatArea(650)   // "650 sq ft"
 */
export function formatArea(sqft: number): string {
  const rounded = Math.round(sqft);
  return `${rounded.toLocaleString('en-IN')} sq ft`;
}

/**
 * Converts square feet to square meters.
 */
export function sqftToSqm(sqft: number): number {
  return sqft * 0.092903;
}

/**
 * Converts square meters to square feet.
 */
export function sqmToSqft(sqm: number): number {
  return sqm / 0.092903;
}

// ---------------------------------------------------------------------------
// Phone Number Formatting
// ---------------------------------------------------------------------------

/**
 * Normalizes an Indian phone number to the +91XXXXXXXXXX format.
 *
 * Handles common input variations:
 *   - "9876543210"       -> "+919876543210"
 *   - "09876543210"      -> "+919876543210"
 *   - "+919876543210"    -> "+919876543210"
 *   - "919876543210"     -> "+919876543210"
 *   - "98765 43210"      -> "+919876543210"
 *   - "98765-43210"      -> "+919876543210"
 *   - "+91-98765-43210"  -> "+919876543210"
 *
 * @returns Normalized phone string or null if invalid
 */
export function formatPhone(phone: string): string | null {
  // Strip all non-digit characters except leading +
  let cleaned = phone.trim();

  // Remember if it started with +
  const hadPlus = cleaned.startsWith('+');

  // Remove everything except digits
  cleaned = cleaned.replace(/\D/g, '');

  // Remove leading 91 country code if present
  if (cleaned.startsWith('91') && cleaned.length > 10) {
    cleaned = cleaned.slice(2);
  }

  // Remove leading 0 (trunk prefix)
  if (cleaned.startsWith('0') && cleaned.length === 11) {
    cleaned = cleaned.slice(1);
  }

  // Must be exactly 10 digits and start with 6-9
  if (cleaned.length !== 10 || !/^[6-9]/.test(cleaned)) {
    return null;
  }

  return `+91${cleaned}`;
}

/**
 * Formats a phone number for display: +91 98765 43210
 */
export function formatPhoneDisplay(phone: string): string {
  const normalized = formatPhone(phone);
  if (!normalized) return phone;

  const digits = normalized.slice(3); // remove +91
  return `+91 ${digits.slice(0, 5)} ${digits.slice(5)}`;
}

// ---------------------------------------------------------------------------
// Price per sqft display
// ---------------------------------------------------------------------------

/**
 * Formats price per square foot.
 *
 * @example
 *   formatPricePerSqft(12500) // "₹12,500/sq ft"
 */
export function formatPricePerSqft(pricePerSqft: number): string {
  return `${formatINR(pricePerSqft)}/sq ft`;
}
