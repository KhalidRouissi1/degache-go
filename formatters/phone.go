package formatters

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/degache-go/degache/constants"
	"github.com/degache-go/degache/validators"
)

// FormatPhoneNumber formats a Tunisian phone number with country code and proper spacing
//
// Parameters:
//   - phoneNumber: The phone number to format
//
// Returns:
//   - string: formatted phone number with country code and proper spacing
//   - error: error if phone number is invalid
//
// Example:
//
//	formatted, err := FormatPhoneNumber("20123456")
//	// Returns: "+216 20 123 456", nil
//
//	formatted, err := FormatPhoneNumber("+21620123456")
//	// Returns: "+216 20 123 456", nil
func FormatPhoneNumber(phoneNumber string) (string, error) {
	if !validators.ValidatePhoneNumber(phoneNumber) {
		return "", fmt.Errorf("invalid phone number: %s", phoneNumber)
	}

	// Remove all non-digit characters except +
	cleaned := regexp.MustCompile(`[^\d+]`).ReplaceAllString(phoneNumber, "")

	// Remove international prefix if present
	cleaned = strings.TrimPrefix(cleaned, "+216")

	// Ensure we have exactly 8 digits
	if len(cleaned) != 8 {
		return "", fmt.Errorf("phone number must be exactly 8 digits after cleaning")
	}

	// Format as: +216 XX XXX XXX
	return fmt.Sprintf("%s %s %s %s",
		constants.CountryCode,
		cleaned[:2],
		cleaned[2:5],
		cleaned[5:]), nil
}

// FormatPhoneNumberNational formats a phone number in national format (without country code)
//
// Parameters:
//   - phoneNumber: The phone number to format
//
// Returns:
//   - string: formatted phone number in national format
//   - error: error if phone number is invalid
//
// Example:
//
//	formatted, err := FormatPhoneNumberNational("20123456")
//	// Returns: "20 123 456", nil
func FormatPhoneNumberNational(phoneNumber string) (string, error) {
	if !validators.ValidatePhoneNumber(phoneNumber) {
		return "", fmt.Errorf("invalid phone number: %s", phoneNumber)
	}

	// Remove all non-digit characters except +
	cleaned := regexp.MustCompile(`[^\d+]`).ReplaceAllString(phoneNumber, "")

	// Remove international prefix if present
	cleaned = strings.TrimPrefix(cleaned, "+216")

	// Ensure we have exactly 8 digits
	if len(cleaned) != 8 {
		return "", fmt.Errorf("phone number must be exactly 8 digits after cleaning")
	}

	// Format as: XX XXX XXX
	return fmt.Sprintf("%s %s %s",
		cleaned[:2],
		cleaned[2:5],
		cleaned[5:]), nil
}

// FormatPhoneNumberCompact formats a phone number in compact format (no spaces)
//
// Parameters:
//   - phoneNumber: The phone number to format
//
// Returns:
//   - string: formatted phone number in compact format
//   - error: error if phone number is invalid
//
// Example:
//
//	formatted, err := FormatPhoneNumberCompact("20 123 456")
//	// Returns: "+21620123456", nil
func FormatPhoneNumberCompact(phoneNumber string) (string, error) {
	if !validators.ValidatePhoneNumber(phoneNumber) {
		return "", fmt.Errorf("invalid phone number: %s", phoneNumber)
	}

	// Remove all non-digit characters except +
	cleaned := regexp.MustCompile(`[^\d+]`).ReplaceAllString(phoneNumber, "")

	// Remove international prefix if present
	cleaned = strings.TrimPrefix(cleaned, "+216")

	// Ensure we have exactly 8 digits
	if len(cleaned) != 8 {
		return "", fmt.Errorf("phone number must be exactly 8 digits after cleaning")
	}

	// Format as: +216XXXXXXXX
	return constants.CountryCode + cleaned, nil
}

// NormalizePhoneNumber normalizes a phone number by removing all formatting
//
// Parameters:
//   - phoneNumber: The phone number to normalize
//
// Returns:
//   - string: normalized phone number (8 digits only)
//   - error: error if phone number is invalid
//
// Example:
//
//	normalized, err := NormalizePhoneNumber("+216 20 123 456")
//	// Returns: "20123456", nil
func NormalizePhoneNumber(phoneNumber string) (string, error) {
	if !validators.ValidatePhoneNumber(phoneNumber) {
		return "", fmt.Errorf("invalid phone number: %s", phoneNumber)
	}

	// Remove all non-digit characters except +
	cleaned := regexp.MustCompile(`[^\d+]`).ReplaceAllString(phoneNumber, "")

	// Remove international prefix if present
	cleaned = strings.TrimPrefix(cleaned, "+216")

	// Ensure we have exactly 8 digits
	if len(cleaned) != 8 {
		return "", fmt.Errorf("phone number must be exactly 8 digits after cleaning")
	}

	return cleaned, nil
}
