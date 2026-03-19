// ---------------------------------------------------------------------------
// Phone Validation
// ---------------------------------------------------------------------------

/**
 * Validates an Indian mobile phone number.
 *
 * Rules:
 *   - Must be exactly 10 digits (after stripping country code / trunk prefix)
 *   - Must start with 6, 7, 8, or 9
 *   - Accepts raw 10 digits, with 0 prefix, or with +91/91 prefix
 *
 * @returns Object with `valid` flag and optional `error` message
 */
export function validatePhone(phone: string): ValidationResult {
  if (!phone || typeof phone !== 'string') {
    return { valid: false, error: 'Phone number is required' };
  }

  let cleaned = phone.trim().replace(/[\s\-()]/g, '');

  // Strip country code
  if (cleaned.startsWith('+91')) {
    cleaned = cleaned.slice(3);
  } else if (cleaned.startsWith('91') && cleaned.length > 10) {
    cleaned = cleaned.slice(2);
  }

  // Strip trunk prefix
  if (cleaned.startsWith('0') && cleaned.length === 11) {
    cleaned = cleaned.slice(1);
  }

  if (!/^\d{10}$/.test(cleaned)) {
    return { valid: false, error: 'Phone number must be 10 digits' };
  }

  if (!/^[6-9]/.test(cleaned)) {
    return {
      valid: false,
      error: 'Indian mobile numbers must start with 6, 7, 8, or 9',
    };
  }

  return { valid: true };
}

// ---------------------------------------------------------------------------
// RERA Number Validation (Maharashtra)
// ---------------------------------------------------------------------------

/**
 * Validates a MahaRERA registration number.
 *
 * MahaRERA project numbers follow the pattern:
 *   P52000XXXXXX   (projects — "P" prefix, "52000" state code, followed by digits)
 *
 * MahaRERA agent numbers follow the pattern:
 *   A52000XXXXXX   (agents — "A" prefix)
 *
 * The full pattern allows for alphanumeric characters after the state code
 * as MahaRERA sometimes uses mixed formats.
 *
 * @param reraNumber - The RERA registration number to validate
 * @param type - Whether to validate as 'project' or 'agent' number
 */
export function validateRERA(
  reraNumber: string,
  type: 'project' | 'agent' = 'project',
): ValidationResult {
  if (!reraNumber || typeof reraNumber !== 'string') {
    return { valid: false, error: 'RERA number is required' };
  }

  const cleaned = reraNumber.trim().toUpperCase();

  if (type === 'project') {
    // Project: P52000 followed by alphanumeric chars (typically 6-10 more chars)
    if (!/^P52000[A-Z0-9]{5,15}$/.test(cleaned)) {
      return {
        valid: false,
        error:
          'Invalid MahaRERA project number. Expected format: P52000XXXXXX',
      };
    }
  } else {
    // Agent: A52000 followed by alphanumeric chars
    if (!/^A52000[A-Z0-9]{5,15}$/.test(cleaned)) {
      return {
        valid: false,
        error: 'Invalid MahaRERA agent number. Expected format: A52000XXXXXX',
      };
    }
  }

  return { valid: true };
}

// ---------------------------------------------------------------------------
// PAN Validation
// ---------------------------------------------------------------------------

/**
 * Validates an Indian Permanent Account Number (PAN).
 *
 * PAN format: ABCDE1234F
 *   - First 5: letters (4th letter indicates holder type)
 *   - Next 4: digits
 *   - Last 1: letter (check digit)
 *
 * 4th character indicates type:
 *   P = Individual, C = Company, H = HUF, F = Firm, T = Trust, etc.
 */
export function validatePAN(pan: string): ValidationResult {
  if (!pan || typeof pan !== 'string') {
    return { valid: false, error: 'PAN is required' };
  }

  const cleaned = pan.trim().toUpperCase();

  if (!/^[A-Z]{5}[0-9]{4}[A-Z]$/.test(cleaned)) {
    return {
      valid: false,
      error: 'Invalid PAN format. Expected format: ABCDE1234F',
    };
  }

  // Validate 4th character is a valid type code
  const typeChar = cleaned[3];
  const validTypes = [
    'A', // Association of Persons
    'B', // Body of Individuals
    'C', // Company
    'F', // Firm
    'G', // Government
    'H', // HUF
    'J', // Artificial Juridical Person
    'L', // Local Authority
    'P', // Individual (Person)
    'T', // Trust
  ];

  if (!validTypes.includes(typeChar)) {
    return {
      valid: false,
      error: `Invalid PAN type code '${typeChar}' at position 4`,
    };
  }

  return { valid: true };
}

// ---------------------------------------------------------------------------
// GSTIN Validation
// ---------------------------------------------------------------------------

/**
 * Validates an Indian GST Identification Number (GSTIN).
 *
 * GSTIN format: 22AAAAA0000A1Z5 (15 characters)
 *   - Positions 1-2: State code (01-38)
 *   - Positions 3-12: PAN of the holder
 *   - Position 13: Entity number (1-9, then A-Z for >9 registrations)
 *   - Position 14: 'Z' by default
 *   - Position 15: Check digit (alphanumeric)
 */
export function validateGST(gstin: string): ValidationResult {
  if (!gstin || typeof gstin !== 'string') {
    return { valid: false, error: 'GSTIN is required' };
  }

  const cleaned = gstin.trim().toUpperCase();

  if (cleaned.length !== 15) {
    return { valid: false, error: 'GSTIN must be exactly 15 characters' };
  }

  // Full pattern
  const gstPattern =
    /^([0-3][0-9])[A-Z]{5}[0-9]{4}[A-Z][0-9A-Z][Z][0-9A-Z]$/;

  if (!gstPattern.test(cleaned)) {
    return {
      valid: false,
      error: 'Invalid GSTIN format. Expected format: 22AAAAA0000A1Z5',
    };
  }

  // Validate state code (01-38)
  const stateCode = parseInt(cleaned.slice(0, 2), 10);
  if (stateCode < 1 || stateCode > 38) {
    return { valid: false, error: 'Invalid state code in GSTIN' };
  }

  // Validate embedded PAN (positions 3-12)
  const embeddedPan = cleaned.slice(2, 12);
  const panResult = validatePAN(embeddedPan);
  if (!panResult.valid) {
    return { valid: false, error: 'Invalid PAN embedded in GSTIN' };
  }

  return { valid: true };
}

// ---------------------------------------------------------------------------
// Email Validation
// ---------------------------------------------------------------------------

/**
 * Validates an email address.
 *
 * Uses a practical regex that covers most real-world email addresses
 * without being overly strict or permissive.
 */
export function validateEmail(email: string): ValidationResult {
  if (!email || typeof email !== 'string') {
    return { valid: false, error: 'Email is required' };
  }

  const cleaned = email.trim().toLowerCase();

  if (cleaned.length > 254) {
    return { valid: false, error: 'Email address is too long' };
  }

  // RFC 5322 simplified — covers 99.9% of valid email addresses
  const emailPattern =
    /^[a-z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-z0-9](?:[a-z0-9-]{0,61}[a-z0-9])?(?:\.[a-z0-9](?:[a-z0-9-]{0,61}[a-z0-9])?)*\.[a-z]{2,}$/;

  if (!emailPattern.test(cleaned)) {
    return { valid: false, error: 'Invalid email address' };
  }

  return { valid: true };
}

// ---------------------------------------------------------------------------
// Aadhaar Number Validation (basic format check)
// ---------------------------------------------------------------------------

/**
 * Validates an Aadhaar number format.
 *
 * Rules:
 *   - 12 digits
 *   - Cannot start with 0 or 1
 *
 * Note: This only validates the format, not the Verhoeff checksum.
 */
export function validateAadhaar(aadhaar: string): ValidationResult {
  if (!aadhaar || typeof aadhaar !== 'string') {
    return { valid: false, error: 'Aadhaar number is required' };
  }

  const cleaned = aadhaar.trim().replace(/[\s-]/g, '');

  if (!/^\d{12}$/.test(cleaned)) {
    return { valid: false, error: 'Aadhaar number must be 12 digits' };
  }

  if (/^[01]/.test(cleaned)) {
    return {
      valid: false,
      error: 'Aadhaar number cannot start with 0 or 1',
    };
  }

  return { valid: true };
}

// ---------------------------------------------------------------------------
// Shared Types
// ---------------------------------------------------------------------------

export interface ValidationResult {
  valid: boolean;
  error?: string;
}
