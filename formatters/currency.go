package formatters

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/degache-go/degache/types"
)

const (
	// TunisianDinarSymbol is the symbol for Tunisian Dinar
	TunisianDinarSymbol = "دينار تونسي"
	// TunisianDinarCode is the ISO code for Tunisian Dinar
	TunisianDinarCode = "TND"
)

// FormatCurrency formats an amount in Tunisian Dinar with proper formatting
//
// Parameters:
//   - amount: The amount to format
//   - options: Formatting options (optional)
//
// Returns:
//   - string: formatted currency string
//
// Example:
//
//	formatted := FormatCurrency(1234.56, types.CurrencyFormatOptions{Symbol: true})
//	// Returns: "1.234,560 دينار تونسي"
//
//	formatted := FormatCurrency(1234.56, types.CurrencyFormatOptions{Code: true})
//	// Returns: "1.234,560 TND"
func FormatCurrency(amount float64, options ...types.CurrencyFormatOptions) string {
	var opts types.CurrencyFormatOptions
	if len(options) > 0 {
		opts = options[0]
	}

	// Round to 3 decimal places (millimes)
	rounded := math.Round(amount*1000) / 1000

	// Format the number with Tunisian locale (dot as thousands separator, comma as decimal)
	formatted := formatTunisianNumber(rounded)

	// Add currency symbol or code
	if opts.Symbol {
		return fmt.Sprintf("%s %s", formatted, TunisianDinarSymbol)
	} else if opts.Code {
		return fmt.Sprintf("%s %s", formatted, TunisianDinarCode)
	}

	// Default: include symbol
	return fmt.Sprintf("%s %s", formatted, TunisianDinarSymbol)
}

// FormatCurrencyCompact formats currency in a compact format (no symbol/code)
//
// Parameters:
//   - amount: The amount to format
//
// Returns:
//   - string: formatted amount without currency symbol
//
// Example:
//
//	formatted := FormatCurrencyCompact(1234.56)
//	// Returns: "1.234,560"
func FormatCurrencyCompact(amount float64) string {
	rounded := math.Round(amount*1000) / 1000
	return formatTunisianNumber(rounded)
}

// ParseCurrency parses a formatted currency string back to float64
//
// Parameters:
//   - currencyStr: The formatted currency string to parse
//
// Returns:
//   - float64: parsed amount
//   - error: error if parsing fails
//
// Example:
//
//	amount, err := ParseCurrency("1.234,560 دينار تونسي")
//	// Returns: 1234.56, nil
func ParseCurrency(currencyStr string) (float64, error) {
	// Remove currency symbols and codes
	cleaned := strings.ReplaceAll(currencyStr, TunisianDinarSymbol, "")
	cleaned = strings.ReplaceAll(cleaned, TunisianDinarCode, "")
	cleaned = strings.TrimSpace(cleaned)

	// Convert Tunisian format to standard format
	// Replace dots (thousands separator) with empty string
	cleaned = strings.ReplaceAll(cleaned, ".", "")
	// Replace comma (decimal separator) with dot
	cleaned = strings.ReplaceAll(cleaned, ",", ".")

	// Parse as float
	amount, err := strconv.ParseFloat(cleaned, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid currency format: %s", currencyStr)
	}

	return amount, nil
}

// formatTunisianNumber formats a number according to Tunisian locale
// Uses dot as thousands separator and comma as decimal separator
func formatTunisianNumber(amount float64) string {
	// Split into integer and decimal parts
	intPart := int64(math.Abs(amount))
	decPart := math.Abs(amount) - float64(intPart)

	// Format integer part with dot as thousands separator
	intStr := formatWithThousandsSeparator(intPart, ".")

	// Format decimal part (3 decimal places for millimes)
	decStr := fmt.Sprintf("%.3f", decPart)[2:] // Remove "0."

	// Combine with comma as decimal separator
	result := fmt.Sprintf("%s,%s", intStr, decStr)

	// Add negative sign if needed
	if amount < 0 {
		result = "-" + result
	}

	return result
}

// formatWithThousandsSeparator formats an integer with thousands separator
func formatWithThousandsSeparator(num int64, separator string) string {
	str := strconv.FormatInt(num, 10)

	// Add thousands separator
	if len(str) <= 3 {
		return str
	}

	var result strings.Builder
	for i, digit := range str {
		if i > 0 && (len(str)-i)%3 == 0 {
			result.WriteString(separator)
		}
		result.WriteRune(digit)
	}

	return result.String()
}

// ConvertMillimesToDinars converts millimes to dinars
//
// Parameters:
//   - millimes: Amount in millimes
//
// Returns:
//   - float64: Amount in dinars
//
// Example:
//
//	dinars := ConvertMillimesToDinars(1500) // Returns: 1.5
func ConvertMillimesToDinars(millimes int64) float64 {
	return float64(millimes) / 1000.0
}

// ConvertDinarsToMillimes converts dinars to millimes
//
// Parameters:
//   - dinars: Amount in dinars
//
// Returns:
//   - int64: Amount in millimes
//
// Example:
//
//	millimes := ConvertDinarsToMillimes(1.5) // Returns: 1500
func ConvertDinarsToMillimes(dinars float64) int64 {
	return int64(math.Round(dinars * 1000))
}
