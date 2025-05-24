package validators

import (
	"regexp"
	"strings"

	"github.com/degache-go/degache/constants"
	"github.com/degache-go/degache/types"
)

// Regular expressions for phone number validation
var (
	phoneRegex         = regexp.MustCompile(`^[2-9]\d{7}$`)
	internationalRegex = regexp.MustCompile(`^\+216[2-9]\d{7}$`)
	strictPhoneRegex   = regexp.MustCompile(`^(?:\+216)?[2-9]\d{7}$`)
)

// ValidatePhoneNumber validates a Tunisian phone number
//
// Parameters:
//   - phoneNumber: The phone number to validate
//   - options: Validation options (optional)
//
// Returns:
//   - bool: true if the phone number is valid, false otherwise
//
// Example:
//
//	isValid := ValidatePhoneNumber("20123456", types.PhoneNumberValidationOptions{})
//	isValid := ValidatePhoneNumber("+21620123456", types.PhoneNumberValidationOptions{})
//	isValid := ValidatePhoneNumber("20 123 456", types.PhoneNumberValidationOptions{Strict: true}) // false
func ValidatePhoneNumber(phoneNumber string, options ...types.PhoneNumberValidationOptions) bool {
	if phoneNumber == "" {
		return false
	}

	var opts types.PhoneNumberValidationOptions
	if len(options) > 0 {
		opts = options[0]
	}

	// In strict mode, validate against the strict regex first
	if opts.Strict && !strictPhoneRegex.MatchString(phoneNumber) {
		return false
	}

	// Remove international prefix if present
	normalizedNumber := strings.TrimPrefix(phoneNumber, "+216")

	// Remove spaces and special characters if not in strict mode
	if !opts.Strict {
		normalizedNumber = regexp.MustCompile(`\D`).ReplaceAllString(normalizedNumber, "")
	}

	if !phoneRegex.MatchString(normalizedNumber) {
		return false
	}

	// Check if the prefix is valid
	prefix := string(normalizedNumber[0])
	for _, validPrefix := range constants.ValidPrefixes {
		if prefix == validPrefix {
			return true
		}
	}

	return false
}

// GetCarrierInfo gets carrier information from a phone number
//
// Parameters:
//   - phoneNumber: The phone number to check
//   - options: Validation options (optional)
//
// Returns:
//   - *types.CarrierInfo: carrier information or nil if invalid
//
// Example:
//
//	info := GetCarrierInfo("20123456")
//	if info != nil {
//	    fmt.Printf("Carrier: %s\n", info.Carrier.Name)
//	}
func GetCarrierInfo(phoneNumber string, options ...types.PhoneNumberValidationOptions) *types.CarrierInfo {
	var opts types.PhoneNumberValidationOptions
	if len(options) > 0 {
		opts = options[0]
	}

	if !ValidatePhoneNumber(phoneNumber, opts) {
		return nil
	}

	normalizedNumber := strings.TrimPrefix(phoneNumber, "+216")
	if !opts.Strict {
		normalizedNumber = regexp.MustCompile(`\D`).ReplaceAllString(normalizedNumber, "")
	}

	prefix := string(normalizedNumber[0])

	for _, carrier := range constants.Carriers {
		for _, carrierPrefix := range carrier.Prefixes {
			if prefix == carrierPrefix {
				return &types.CarrierInfo{
					Carrier: carrier,
					Prefix:  prefix,
				}
			}
		}
	}

	return nil
}

// ValidatePhoneNumberWithDetails validates a phone number and returns detailed information
//
// Parameters:
//   - phoneNumber: The phone number to validate
//   - options: Validation options (optional)
//
// Returns:
//   - bool: true if the phone number is valid
//   - string: error message if invalid, empty string if valid
//
// Example:
//
//	valid, msg := ValidatePhoneNumberWithDetails("20123456")
//	if !valid {
//	    fmt.Println("Invalid phone number:", msg)
//	}
func ValidatePhoneNumberWithDetails(phoneNumber string, options ...types.PhoneNumberValidationOptions) (bool, string) {
	if phoneNumber == "" {
		return false, "Phone number cannot be empty"
	}

	var opts types.PhoneNumberValidationOptions
	if len(options) > 0 {
		opts = options[0]
	}

	// In strict mode, validate against the strict regex first
	if opts.Strict && !strictPhoneRegex.MatchString(phoneNumber) {
		return false, "Phone number format is invalid in strict mode"
	}

	// Remove international prefix if present
	normalizedNumber := strings.TrimPrefix(phoneNumber, "+216")

	// Remove spaces and special characters if not in strict mode
	if !opts.Strict {
		normalizedNumber = regexp.MustCompile(`\D`).ReplaceAllString(normalizedNumber, "")
	}

	if len(normalizedNumber) != 8 {
		return false, "Phone number must be exactly 8 digits"
	}

	if !phoneRegex.MatchString(normalizedNumber) {
		return false, "Phone number must start with 2-9 and contain only digits"
	}

	// Check if the prefix is valid
	prefix := string(normalizedNumber[0])
	for _, validPrefix := range constants.ValidPrefixes {
		if prefix == validPrefix {
			return true, ""
		}
	}

	return false, "Phone number prefix is not valid for Tunisian carriers"
}
