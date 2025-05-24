package validators

import (
	"regexp"
)

// cinRegex is the regular expression for CIN validation
// 0 or 1 followed by 7 digits to match the Tunisian CIN format
var cinRegex = regexp.MustCompile(`^[01]\d{7}$`)

// ValidateCIN validates a Tunisian CIN (Carte d'Identité Nationale)
// A valid CIN is an 8-digit number starting with 0 or 1
//
// Parameters:
//   - cin: The CIN number to validate
//
// Returns:
//   - bool: true if the CIN is valid, false otherwise
//
// Example:
//
//	isValid := ValidateCIN("12345678") // returns true
//	isValid := ValidateCIN("22345678") // returns false (doesn't start with 0 or 1)
//	isValid := ValidateCIN("1234567")  // returns false (not 8 digits)
func ValidateCIN(cin string) bool {
	if cin == "" {
		return false
	}

	return cinRegex.MatchString(cin)
}

// ValidateCINWithDetails validates a CIN and returns detailed information
//
// Parameters:
//   - cin: The CIN number to validate
//
// Returns:
//   - bool: true if the CIN is valid
//   - string: error message if invalid, empty string if valid
//
// Example:
//
//	valid, msg := ValidateCINWithDetails("12345678")
//	if !valid {
//	    fmt.Println("Invalid CIN:", msg)
//	}
func ValidateCINWithDetails(cin string) (bool, string) {
	if cin == "" {
		return false, "CIN cannot be empty"
	}

	if len(cin) != 8 {
		return false, "CIN must be exactly 8 digits"
	}

	if !cinRegex.MatchString(cin) {
		return false, "CIN must start with 0 or 1 and contain only digits"
	}

	return true, ""
}
