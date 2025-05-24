package validators

import (
	"regexp"

	"github.com/degache-go/degache/constants"
)

// postalCodeRegex is the regular expression for postal code validation
// 4 digits
var postalCodeRegex = regexp.MustCompile(`^\d{4}$`)

// ValidatePostalCode validates a Tunisian postal code
// A valid postal code is a 4-digit number
//
// Parameters:
//   - postalCode: The postal code to validate
//
// Returns:
//   - bool: true if the postal code is valid, false otherwise
//
// Example:
//
//	isValid := ValidatePostalCode("1000") // returns true (Tunis)
//	isValid := ValidatePostalCode("123")  // returns false (not 4 digits)
func ValidatePostalCode(postalCode string) bool {
	if postalCode == "" {
		return false
	}

	if !postalCodeRegex.MatchString(postalCode) {
		return false
	}

	// Check if postal code exists in governorates
	for _, governorate := range constants.Governorates {
		if governorate.PostalCode == postalCode {
			return true
		}
	}

	// If not found in main governorates, accept any 4-digit code
	// as there might be sub-regions not listed
	return true
}

// ValidatePostalCodeWithDetails validates a postal code and returns detailed information
//
// Parameters:
//   - postalCode: The postal code to validate
//
// Returns:
//   - bool: true if the postal code is valid
//   - string: error message if invalid, empty string if valid
//
// Example:
//
//	valid, msg := ValidatePostalCodeWithDetails("1000")
//	if !valid {
//	    fmt.Println("Invalid postal code:", msg)
//	}
func ValidatePostalCodeWithDetails(postalCode string) (bool, string) {
	if postalCode == "" {
		return false, "Postal code cannot be empty"
	}

	if len(postalCode) != 4 {
		return false, "Postal code must be exactly 4 digits"
	}

	if !postalCodeRegex.MatchString(postalCode) {
		return false, "Postal code must contain only digits"
	}

	return true, ""
}

// GetGovernorateFromPostalCode gets governorate information from a postal code
//
// Parameters:
//   - postalCode: The postal code to check
//
// Returns:
//   - *constants.Governorate: governorate information or nil if not found
//
// Example:
//
//	gov := GetGovernorateFromPostalCode("1000")
//	if gov != nil {
//	    fmt.Printf("Governorate: %s, Region: %s\n", gov.Name, gov.Region)
//	}
func GetGovernorateFromPostalCode(postalCode string) *constants.Governorate {
	if !ValidatePostalCode(postalCode) {
		return nil
	}

	for _, governorate := range constants.Governorates {
		if governorate.PostalCode == postalCode {
			return &governorate
		}
	}

	return nil
}

// IsMainGovernoratePostalCode checks if a postal code belongs to a main governorate
//
// Parameters:
//   - postalCode: The postal code to check
//
// Returns:
//   - bool: true if the postal code belongs to a main governorate
//
// Example:
//
//	isMain := IsMainGovernoratePostalCode("1000") // returns true (Tunis)
//	isMain := IsMainGovernoratePostalCode("1001") // returns false (sub-region)
func IsMainGovernoratePostalCode(postalCode string) bool {
	return GetGovernorateFromPostalCode(postalCode) != nil
}
