package types

import "github.com/degache-go/degache/constants"

// CIN represents a Tunisian CIN (Carte d'Identité Nationale)
// An 8-digit number starting with 0 or 1
type CIN string

// PhoneNumber represents a Tunisian phone number
// 8 digits starting with 2-9
type PhoneNumber string

// PostalCode represents a Tunisian postal code
// 4 digits
type PostalCode string

// TaxID represents a Tunisian company tax ID (Matricule Fiscal)
// Format: 7 digits followed by letter/letter/letter/3 digits
type TaxID string

// RIB represents a Tunisian RIB (Relevé d'Identité Bancaire)
// 20 digits
type RIB string

// PhoneNumberValidationOptions contains options for phone number validation
type PhoneNumberValidationOptions struct {
	// Strict enforces strict format validation:
	// - No spaces or special characters allowed
	// - Must be exactly 8 digits or with +216 prefix
	// - Must start with a valid carrier prefix
	Strict bool
}

// CarPlateValidationOptions contains options for car plate validation
type CarPlateValidationOptions struct {
	// Type specifies the type of car plate (standard, special)
	Type string
	// Strict enforces strict format validation
	Strict bool
}

// CurrencyFormatOptions contains options for currency formatting
type CurrencyFormatOptions struct {
	// Symbol indicates whether to include the currency symbol
	Symbol bool
	// Code indicates whether to include the currency code
	Code bool
}

// CarPlateInfo contains information about a car plate
type CarPlateInfo struct {
	Type       string
	Components CarPlateComponents
}

// CarPlateComponents contains the components of a car plate
type CarPlateComponents struct {
	Prefix string
	Region string
	Suffix string
}

// Address represents a Tunisian address
type Address struct {
	Street      string
	City        string
	PostalCode  PostalCode
	Governorate string
}

// ValidationResult represents the result of a validation operation
type ValidationResult struct {
	Valid   bool
	Message string
}

// CarrierInfo contains information about a mobile carrier
type CarrierInfo struct {
	Carrier constants.Carrier
	Prefix  string
}

// BankInfo contains information about a bank
type BankInfo struct {
	Bank constants.Bank
	Code string
}
