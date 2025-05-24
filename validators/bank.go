package validators

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/degache-go/degache/constants"
	"github.com/degache-go/degache/types"
)

// ribRegex is the regular expression for RIB validation
// 20 digits
var ribRegex = regexp.MustCompile(`^\d{20}$`)

// ValidateRIB validates a Tunisian RIB (Relevé d'Identité Bancaire)
// A valid RIB is a 20-digit number
//
// Parameters:
//   - rib: The RIB number to validate
//
// Returns:
//   - bool: true if the RIB is valid, false otherwise
//
// Example:
//
//	isValid := ValidateRIB("12345678901234567890") // returns true
//	isValid := ValidateRIB("1234567890123456789")  // returns false (not 20 digits)
func ValidateRIB(rib string) bool {
	if rib == "" {
		return false
	}

	if !ribRegex.MatchString(rib) {
		return false
	}

	// Additional validation: check if bank code exists
	bankCode := rib[:2]
	_, exists := constants.Banks[bankCode]
	return exists
}

// ValidateRIBWithDetails validates a RIB and returns detailed information
//
// Parameters:
//   - rib: The RIB number to validate
//
// Returns:
//   - bool: true if the RIB is valid
//   - string: error message if invalid, empty string if valid
//
// Example:
//
//	valid, msg := ValidateRIBWithDetails("12345678901234567890")
//	if !valid {
//	    fmt.Println("Invalid RIB:", msg)
//	}
func ValidateRIBWithDetails(rib string) (bool, string) {
	if rib == "" {
		return false, "RIB cannot be empty"
	}

	if len(rib) != 20 {
		return false, "RIB must be exactly 20 digits"
	}

	if !ribRegex.MatchString(rib) {
		return false, "RIB must contain only digits"
	}

	// Check if bank code exists
	bankCode := rib[:2]
	_, exists := constants.Banks[bankCode]
	if !exists {
		return false, "Bank code not recognized"
	}

	return true, ""
}

// GetBankFromRIB extracts bank information from a RIB
//
// Parameters:
//   - rib: The RIB number to extract bank information from
//
// Returns:
//   - *types.BankInfo: bank information or nil if invalid
//
// Example:
//
//	bankInfo := GetBankFromRIB("12345678901234567890")
//	if bankInfo != nil {
//	    fmt.Printf("Bank: %s\n", bankInfo.Bank.Name)
//	}
func GetBankFromRIB(rib string) *types.BankInfo {
	if !ValidateRIB(rib) {
		return nil
	}

	bankCode := rib[:2]
	bank, exists := constants.Banks[bankCode]
	if !exists {
		return nil
	}

	return &types.BankInfo{
		Bank: bank,
		Code: bankCode,
	}
}

// ExtractRIBComponents extracts components from a valid RIB
//
// Parameters:
//   - rib: The RIB to extract components from
//
// Returns:
//   - map[string]string: components of the RIB (bankCode, branchCode, accountNumber, key)
//   - error: error if RIB is invalid
//
// Example:
//
//	components, err := ExtractRIBComponents("12345678901234567890")
//	if err == nil {
//	    fmt.Printf("Bank: %s, Branch: %s, Account: %s, Key: %s\n",
//	        components["bankCode"], components["branchCode"], components["accountNumber"], components["key"])
//	}
func ExtractRIBComponents(rib string) (map[string]string, error) {
	valid, msg := ValidateRIBWithDetails(rib)
	if !valid {
		return nil, fmt.Errorf("invalid RIB: %s", msg)
	}

	components := make(map[string]string)
	components["bankCode"] = rib[:2]
	components["branchCode"] = rib[2:5]
	components["accountNumber"] = rib[5:18]
	components["key"] = rib[18:]

	return components, nil
}

// ValidateRIBChecksum validates the RIB checksum (simplified version)
// Note: This is a basic implementation. Real RIB validation may require more complex algorithms
//
// Parameters:
//   - rib: The RIB to validate checksum for
//
// Returns:
//   - bool: true if checksum is valid, false otherwise
func ValidateRIBChecksum(rib string) bool {
	if !ValidateRIB(rib) {
		return false
	}

	// Extract the key (last 2 digits)
	key, err := strconv.Atoi(rib[18:])
	if err != nil {
		return false
	}

	// Simple modulo 97 check (this is a simplified version)
	// Real implementation would require the full IBAN algorithm
	accountPart := rib[:18]
	accountNum, err := strconv.ParseInt(accountPart, 10, 64)
	if err != nil {
		return false
	}

	calculatedKey := int(accountNum % 97)
	return calculatedKey == key
}
