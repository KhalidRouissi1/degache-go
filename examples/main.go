package main

import (
	"fmt"
	"log"
	"time"

	"github.com/degache-go/degache"
	"github.com/degache-go/degache/formatters"
	"github.com/degache-go/degache/types"
	"github.com/degache-go/degache/validators"
)

func main() {
	fmt.Println("🌴 degache-go Examples")
	fmt.Println("======================")

	// CIN Validation Examples
	fmt.Println("\n📋 CIN Validation:")
	testCINs := []string{"12345678", "01234567", "22345678", "1234567"}
	for _, cin := range testCINs {
		isValid := degache.ValidateCIN(cin)
		fmt.Printf("  CIN %s: %t\n", cin, isValid)
	}

	// Phone Number Validation and Formatting Examples
	fmt.Println("\n📱 Phone Number Validation and Formatting:")
	testPhones := []string{"20123456", "+21620123456", "50987654", "40555666", "99888777"}
	for _, phone := range testPhones {
		isValid := degache.ValidatePhoneNumber(phone)
		fmt.Printf("  Phone %s: %t", phone, isValid)

		if isValid {
			formatted, err := degache.FormatPhoneNumber(phone)
			if err == nil {
				fmt.Printf(" → %s", formatted)
			}

			// Get carrier info
			carrierInfo := degache.GetCarrierInfo(phone)
			if carrierInfo != nil {
				fmt.Printf(" (%s)", carrierInfo.Carrier.Name)
			}
		}
		fmt.Println()
	}

	// Phone Number Strict Validation
	fmt.Println("\n📱 Phone Number Strict Validation:")
	strictTestPhones := []string{"20123456", "20 123 456", "+21620123456"}
	for _, phone := range strictTestPhones {
		isValid := validators.ValidatePhoneNumber(phone, types.PhoneNumberValidationOptions{Strict: true})
		fmt.Printf("  Phone %s (strict): %t\n", phone, isValid)
	}

	// Tax ID Validation Examples
	fmt.Println("\n💼 Tax ID Validation:")
	testTaxIDs := []string{"1234567A/P/M/000", "1234567B/C/D/123", "123456A/P/M/000", "1234567A/P/M/00"}
	for _, taxID := range testTaxIDs {
		isValid := degache.ValidateTaxID(taxID)
		fmt.Printf("  Tax ID %s: %t\n", taxID, isValid)
	}

	// RIB Validation Examples
	fmt.Println("\n🏦 RIB Validation:")
	testRIBs := []string{"01234567890123456789", "02345678901234567890", "99234567890123456789"}
	for _, rib := range testRIBs {
		isValid := degache.ValidateRIB(rib)
		fmt.Printf("  RIB %s: %t", rib, isValid)

		if isValid {
			bankInfo := degache.GetBankFromRIB(rib)
			if bankInfo != nil {
				fmt.Printf(" (%s)", bankInfo.Bank.Name)
			}
		}
		fmt.Println()
	}

	// Postal Code Validation Examples
	fmt.Println("\n📮 Postal Code Validation:")
	testPostalCodes := []string{"1000", "2000", "3000", "9999", "123"}
	for _, postal := range testPostalCodes {
		isValid := degache.ValidatePostalCode(postal)
		fmt.Printf("  Postal Code %s: %t", postal, isValid)

		if isValid {
			gov := degache.GetGovernorateFromPostalCode(postal)
			if gov != nil {
				fmt.Printf(" (%s, %s)", gov.Name, gov.Region)
			}
		}
		fmt.Println()
	}

	// Car Plate Validation Examples
	fmt.Println("\n🚗 Car Plate Validation:")
	testCarPlates := []string{"123 تونس 4567", "RS 123 تونس", "123تونس4567", "456 تونس 7890"}
	for _, plate := range testCarPlates {
		isValid := degache.ValidateCarPlate(plate)
		fmt.Printf("  Car Plate %s: %t", plate, isValid)

		if isValid {
			plateInfo := degache.GetCarPlateInfo(plate)
			if plateInfo != nil {
				fmt.Printf(" (Type: %s)", plateInfo.Type)
			}
		}
		fmt.Println()
	}

	// Currency Formatting Examples
	fmt.Println("\n💰 Currency Formatting:")
	testAmounts := []float64{1234.567, 1000000.123, 50.5, 0.999}
	for _, amount := range testAmounts {
		formatted := degache.FormatCurrency(amount, types.CurrencyFormatOptions{Symbol: true})
		fmt.Printf("  %.3f → %s\n", amount, formatted)
	}

	// Currency with code
	fmt.Println("\n💰 Currency with Code:")
	for _, amount := range testAmounts {
		formatted := degache.FormatCurrency(amount, types.CurrencyFormatOptions{Code: true})
		fmt.Printf("  %.3f → %s\n", amount, formatted)
	}

	// Date Formatting Examples
	fmt.Println("\n📅 Date Formatting:")
	now := time.Now()
	fmt.Printf("  Full Date: %s\n", degache.FormatDate(now))
	fmt.Printf("  Short Date: %s\n", formatters.FormatDateShort(now))
	fmt.Printf("  Numeric Date: %s\n", formatters.FormatDateNumeric(now))
	fmt.Printf("  Date Time: %s\n", formatters.FormatDateTime(now))
	fmt.Printf("  Time 24h: %s\n", formatters.FormatTime(now))
	fmt.Printf("  Time 12h: %s\n", formatters.FormatTime12Hour(now))

	// Comprehensive Validation Example
	fmt.Println("\n🔍 Comprehensive Validation:")
	testData := map[string]string{
		"cin":      "12345678",
		"phone":    "20123456",
		"taxID":    "1234567A/P/M/000",
		"rib":      "01234567890123456789",
		"postal":   "1000",
		"carPlate": "123 تونس 4567",
	}

	results := degache.IsValidTunisianData(testData)
	for field, isValid := range results {
		status := "❌"
		if isValid {
			status = "✅"
		}
		fmt.Printf("  %s %s: %s\n", status, field, testData[field])
	}

	// Constants Examples
	fmt.Println("\n📊 Available Constants:")
	fmt.Printf("  Country Code: %s\n", degache.CountryCode)
	fmt.Printf("  Valid Prefixes: %v\n", degache.ValidPrefixes)
	fmt.Printf("  Number of Carriers: %d\n", len(degache.Carriers))
	fmt.Printf("  Number of Banks: %d\n", len(degache.Banks))
	fmt.Printf("  Number of Governorates: %d\n", len(degache.Governorates))

	// Error Handling Examples
	fmt.Println("\n⚠️  Error Handling Examples:")

	// Invalid phone number formatting
	_, err := formatters.FormatPhoneNumber("invalid")
	if err != nil {
		fmt.Printf("  Phone formatting error: %s\n", err)
	}

	// Invalid currency parsing
	_, err = formatters.ParseCurrency("invalid currency")
	if err != nil {
		fmt.Printf("  Currency parsing error: %s\n", err)
	}

	// Detailed validation
	valid, msg := validators.ValidateCINWithDetails("123")
	if !valid {
		fmt.Printf("  CIN validation error: %s\n", msg)
	}

	fmt.Println("\n🎉 Examples completed successfully!")
}

func init() {
	// Set up logging
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}
