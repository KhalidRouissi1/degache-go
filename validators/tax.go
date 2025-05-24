package validators

import (
	"fmt"
	"regexp"
)

// taxIDRegex is the regular expression for Tax ID validation
// Format: 7 digits followed by letter/letter/letter/3 digits
var taxIDRegex = regexp.MustCompile(`^\d{7}[A-Z]/[A-Z]/[A-Z]/\d{3}$`)

// ValidateTaxID validates a Tunisian Tax ID (Matricule Fiscal)
// A valid Tax ID follows the format: 7 digits + letter/letter/letter/3 digits
//
// Parameters:
//   - taxID: The Tax ID to validate
//
// Returns:
//   - bool: true if the Tax ID is valid, false otherwise
//
// Example:
//
//	isValid := ValidateTaxID("1234567A/P/M/000") // returns true
//	isValid := ValidateTaxID("123456A/P/M/000")  // returns false (not 7 digits)
//	isValid := ValidateTaxID("1234567A/P/M/00")  // returns false (not 3 digits at end)
func ValidateTaxID(taxID string) bool {
	if taxID == "" {
		return false
	}

	return taxIDRegex.MatchString(taxID)
}

// ValidateTaxIDWithDetails validates a Tax ID and returns detailed information
//
// Parameters:
//   - taxID: The Tax ID to validate
//
// Returns:
//   - bool: true if the Tax ID is valid
//   - string: error message if invalid, empty string if valid
//
// Example:
//
//	valid, msg := ValidateTaxIDWithDetails("1234567A/P/M/000")
//	if !valid {
//	    fmt.Println("Invalid Tax ID:", msg)
//	}
func ValidateTaxIDWithDetails(taxID string) (bool, string) {
	if taxID == "" {
		return false, "Tax ID cannot be empty"
	}

	if len(taxID) != 15 {
		return false, "Tax ID must be exactly 15 characters long"
	}

	if !taxIDRegex.MatchString(taxID) {
		return false, "Tax ID must follow format: 7 digits + letter/letter/letter/3 digits (e.g., 1234567A/P/M/000)"
	}

	return true, ""
}

// ExtractTaxIDComponents extracts components from a valid Tax ID
//
// Parameters:
//   - taxID: The Tax ID to extract components from
//
// Returns:
//   - map[string]string: components of the Tax ID (number, type1, type2, type3, sequence)
//   - error: error if Tax ID is invalid
//
// Example:
//
//	components, err := ExtractTaxIDComponents("1234567A/P/M/000")
//	if err == nil {
//	    fmt.Printf("Number: %s, Types: %s/%s/%s, Sequence: %s\n",
//	        components["number"], components["type1"], components["type2"], components["type3"], components["sequence"])
//	}
func ExtractTaxIDComponents(taxID string) (map[string]string, error) {
	if !ValidateTaxID(taxID) {
		return nil, fmt.Errorf("invalid Tax ID format")
	}

	components := make(map[string]string)
	components["number"] = taxID[:7]
	components["type1"] = string(taxID[7])
	components["type2"] = string(taxID[9])
	components["type3"] = string(taxID[11])
	components["sequence"] = taxID[13:]

	return components, nil
}
