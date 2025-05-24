// Package degache provides utilities for Tunisian developers
// including validation and formatting for CIN, phone numbers, tax IDs, and more.
//
// This is a Go port of the popular degachejs library.
//
// Example usage:
//
//	import "github.com/degache-go/degache"
//	import "github.com/degache-go/degache/validators"
//	import "github.com/degache-go/degache/formatters"
//
//	// Validate CIN
//	isValid := validators.ValidateCIN("12345678")
//
//	// Format phone number
//	formatted, err := formatters.FormatPhoneNumber("20123456")
//
//	// Validate Tax ID
//	isValidTax := validators.ValidateTaxID("1234567A/P/M/000")
package degache

import (
	"github.com/degache-go/degache/constants"
	"github.com/degache-go/degache/formatters"
	"github.com/degache-go/degache/types"
	"github.com/degache-go/degache/validators"
)

// Version of the degache-go library
const Version = "1.0.0"

// Re-export commonly used validators for convenience
var (
	// ValidateCIN validates a Tunisian CIN (Carte d'Identité Nationale)
	ValidateCIN = validators.ValidateCIN

	// ValidatePhoneNumber validates a Tunisian phone number
	ValidatePhoneNumber = validators.ValidatePhoneNumber

	// ValidateTaxID validates a Tunisian Tax ID (Matricule Fiscal)
	ValidateTaxID = validators.ValidateTaxID

	// ValidateRIB validates a Tunisian RIB (bank account number)
	ValidateRIB = validators.ValidateRIB

	// ValidatePostalCode validates a Tunisian postal code
	ValidatePostalCode = validators.ValidatePostalCode

	// ValidateCarPlate validates a Tunisian car plate
	ValidateCarPlate = validators.ValidateCarPlate
)

// Re-export commonly used formatters for convenience
var (
	// FormatPhoneNumber formats a Tunisian phone number
	FormatPhoneNumber = formatters.FormatPhoneNumber

	// FormatCurrency formats an amount in Tunisian Dinar
	FormatCurrency = formatters.FormatCurrency

	// FormatDate formats a date in Tunisian style
	FormatDate = formatters.FormatDate
)

// Re-export commonly used functions for convenience
var (
	// GetCarrierInfo gets carrier information from a phone number
	GetCarrierInfo = validators.GetCarrierInfo

	// GetBankFromRIB gets bank information from a RIB
	GetBankFromRIB = validators.GetBankFromRIB

	// GetCarPlateInfo gets information from a car plate
	GetCarPlateInfo = validators.GetCarPlateInfo

	// GetGovernorateFromPostalCode gets governorate from postal code
	GetGovernorateFromPostalCode = validators.GetGovernorateFromPostalCode
)

// Re-export constants for convenience
var (
	// Carriers contains all Tunisian mobile carriers
	Carriers = constants.Carriers

	// Banks contains all Tunisian banks
	Banks = constants.Banks

	// Governorates contains all Tunisian governorates
	Governorates = constants.Governorates

	// CountryCode is Tunisia's international calling code
	CountryCode = constants.CountryCode

	// ValidPrefixes contains all valid mobile prefixes
	ValidPrefixes = constants.ValidPrefixes
)

// Re-export types for convenience
type (
	// CIN represents a Tunisian CIN
	CIN = types.CIN

	// PhoneNumber represents a Tunisian phone number
	PhoneNumber = types.PhoneNumber

	// TaxID represents a Tunisian Tax ID
	TaxID = types.TaxID

	// RIB represents a Tunisian bank account number
	RIB = types.RIB

	// PostalCode represents a Tunisian postal code
	PostalCode = types.PostalCode

	// PhoneNumberValidationOptions contains options for phone validation
	PhoneNumberValidationOptions = types.PhoneNumberValidationOptions

	// CarPlateValidationOptions contains options for car plate validation
	CarPlateValidationOptions = types.CarPlateValidationOptions

	// CurrencyFormatOptions contains options for currency formatting
	CurrencyFormatOptions = types.CurrencyFormatOptions

	// CarrierInfo contains carrier information
	CarrierInfo = types.CarrierInfo

	// BankInfo contains bank information
	BankInfo = types.BankInfo

	// CarPlateInfo contains car plate information
	CarPlateInfo = types.CarPlateInfo

	// Address represents a Tunisian address
	Address = types.Address
)

// Utility functions

// IsValidTunisianData performs comprehensive validation of common Tunisian data
//
// Parameters:
//   - data: map containing the data to validate
//
// Returns:
//   - map[string]bool: validation results for each field
//
// Example:
//
//	results := IsValidTunisianData(map[string]string{
//	    "cin": "12345678",
//	    "phone": "20123456",
//	    "taxID": "1234567A/P/M/000",
//	})
func IsValidTunisianData(data map[string]string) map[string]bool {
	results := make(map[string]bool)

	if cin, exists := data["cin"]; exists {
		results["cin"] = ValidateCIN(cin)
	}

	if phone, exists := data["phone"]; exists {
		results["phone"] = ValidatePhoneNumber(phone)
	}

	if taxID, exists := data["taxID"]; exists {
		results["taxID"] = ValidateTaxID(taxID)
	}

	if rib, exists := data["rib"]; exists {
		results["rib"] = ValidateRIB(rib)
	}

	if postal, exists := data["postal"]; exists {
		results["postal"] = ValidatePostalCode(postal)
	}

	if plate, exists := data["carPlate"]; exists {
		results["carPlate"] = ValidateCarPlate(plate)
	}

	return results
}
