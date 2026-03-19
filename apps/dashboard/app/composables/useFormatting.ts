/**
 * Formatting utilities for Indian locale display
 */
export function useFormatting() {
  /**
   * Format amount in Indian numbering system with Rupee symbol
   * e.g., 125000 -> "₹1,25,000"
   */
  function formatINR(amount: number): string {
    if (amount === 0) return '₹0'
    const isNegative = amount < 0
    const abs = Math.abs(Math.round(amount))
    const str = abs.toString()

    if (str.length <= 3) return `${isNegative ? '-' : ''}₹${str}`

    // Last 3 digits
    let result = str.slice(-3)
    let remaining = str.slice(0, -3)

    // Group remaining digits in pairs (Indian system)
    while (remaining.length > 0) {
      const chunk = remaining.slice(-2)
      remaining = remaining.slice(0, -2)
      result = chunk + ',' + result
    }

    // Remove leading comma if any
    result = result.replace(/^,/, '')

    return `${isNegative ? '-' : ''}₹${result}`
  }

  /**
   * Format amount in short form: 50L, 1.2Cr etc.
   */
  function formatINRShort(amount: number): string {
    if (amount >= 10000000) {
      const cr = amount / 10000000
      return `₹${cr % 1 === 0 ? cr.toFixed(0) : cr.toFixed(1)}Cr`
    }
    if (amount >= 100000) {
      const lakh = amount / 100000
      return `₹${lakh % 1 === 0 ? lakh.toFixed(0) : lakh.toFixed(1)}L`
    }
    if (amount >= 1000) {
      const k = amount / 1000
      return `₹${k % 1 === 0 ? k.toFixed(0) : k.toFixed(1)}K`
    }
    return formatINR(amount)
  }

  /**
   * Format date as DD MMM YYYY
   * e.g., "2024-03-15" -> "15 Mar 2024"
   */
  function formatDate(date: string | Date): string {
    const d = typeof date === 'string' ? new Date(date) : date
    const months = ['Jan', 'Feb', 'Mar', 'Apr', 'May', 'Jun', 'Jul', 'Aug', 'Sep', 'Oct', 'Nov', 'Dec']
    return `${d.getDate()} ${months[d.getMonth()]} ${d.getFullYear()}`
  }

  /**
   * Format date as DD MMM YYYY, HH:MM AM/PM
   */
  function formatDateTime(date: string | Date): string {
    const d = typeof date === 'string' ? new Date(date) : date
    const dateStr = formatDate(d)
    let hours = d.getHours()
    const minutes = d.getMinutes().toString().padStart(2, '0')
    const ampm = hours >= 12 ? 'PM' : 'AM'
    hours = hours % 12 || 12
    return `${dateStr}, ${hours}:${minutes} ${ampm}`
  }

  /**
   * Format phone as +91 XXXXX XXXXX
   */
  function formatPhone(phone: string): string {
    // Strip everything except digits
    const digits = phone.replace(/\D/g, '')
    if (digits.length === 10) {
      return `+91 ${digits.slice(0, 5)} ${digits.slice(5)}`
    }
    if (digits.length === 12 && digits.startsWith('91')) {
      return `+91 ${digits.slice(2, 7)} ${digits.slice(7)}`
    }
    // Return as-is if format is unexpected
    return phone
  }

  /**
   * Format relative time from now
   * e.g., "2 hours ago", "Yesterday", "3 days ago"
   */
  function formatRelativeTime(date: string | Date): string {
    const d = typeof date === 'string' ? new Date(date) : date
    const now = new Date()
    const diffMs = now.getTime() - d.getTime()
    const diffSec = Math.floor(diffMs / 1000)
    const diffMin = Math.floor(diffSec / 60)
    const diffHr = Math.floor(diffMin / 60)
    const diffDay = Math.floor(diffHr / 24)

    if (diffSec < 60) return 'Just now'
    if (diffMin < 60) return `${diffMin} min ago`
    if (diffHr < 24) return `${diffHr} hour${diffHr > 1 ? 's' : ''} ago`
    if (diffDay === 1) return 'Yesterday'
    if (diffDay < 7) return `${diffDay} days ago`
    if (diffDay < 30) return `${Math.floor(diffDay / 7)} week${Math.floor(diffDay / 7) > 1 ? 's' : ''} ago`
    return formatDate(d)
  }

  /**
   * Get a percentage string
   */
  function formatPercent(value: number, decimals = 1): string {
    return `${value.toFixed(decimals)}%`
  }

  /**
   * Format carpet area with "sq ft" suffix
   * e.g., 450 -> "450 sq ft"
   */
  function formatCarpetArea(sqft: number): string {
    return `${sqft.toLocaleString('en-IN')} sq ft`
  }

  /**
   * Format a number as percentage string
   * e.g., 15.5 -> "15.5%"
   */
  function formatPercentage(value: number, decimals = 1): string {
    return `${value.toFixed(decimals)}%`
  }

  return {
    formatINR,
    formatINRShort,
    formatDate,
    formatDateTime,
    formatPhone,
    formatRelativeTime,
    formatPercent,
    formatCarpetArea,
    formatPercentage,
  }
}
