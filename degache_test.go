package degache

import "testing"

func TestMainPackageExports(t *testing.T) {
	// Test that all main package exports work correctly

	// Test CIN validation
	if !ValidateCIN("12345678") {
		t.Error("ValidateCIN should validate correct CIN")
	}

	// Test phone validation
	if !ValidatePhoneNumber("20123456") {
		t.Error("ValidatePhoneNumber should validate correct phone")
	}

	// Test tax ID validation
	if !ValidateTaxID("1234567A/P/M/000") {
		t.Error("ValidateTaxID should validate correct tax ID")
	}

	// Test RIB validation
	if !ValidateRIB("01234567890123456789") {
		t.Error("ValidateRIB should validate correct RIB")
	}

	// Test postal code validation
	if !ValidatePostalCode("1000") {
		t.Error("ValidatePostalCode should validate correct postal code")
	}

	// Test car plate validation
	if !ValidateCarPlate("123 تونس 4567") {
		t.Error("ValidateCarPlate should validate correct car plate")
	}
}

func TestFormatters(t *testing.T) {
	// Test phone formatting
	formatted, err := FormatPhoneNumber("20123456")
	if err != nil {
		t.Errorf("FormatPhoneNumber should not return error: %v", err)
	}
	expected := "+216 20 123 456"
	if formatted != expected {
		t.Errorf("FormatPhoneNumber returned %q, expected %q", formatted, expected)
	}

	// Test currency formatting
	currencyFormatted := FormatCurrency(1234.567)
	if currencyFormatted == "" {
		t.Error("FormatCurrency should return formatted currency")
	}
}

func TestUtilityFunctions(t *testing.T) {
	// Test carrier info
	carrierInfo := GetCarrierInfo("20123456")
	if carrierInfo == nil {
		t.Error("GetCarrierInfo should return carrier info for valid phone")
	}
	if carrierInfo.Carrier.Name != "Ooredoo Tunisia" {
		t.Errorf("Expected Ooredoo Tunisia, got %s", carrierInfo.Carrier.Name)
	}

	// Test bank info
	bankInfo := GetBankFromRIB("01234567890123456789")
	if bankInfo == nil {
		t.Error("GetBankFromRIB should return bank info for valid RIB")
	}
	if bankInfo.Bank.Name != "Banque Centrale de Tunisie" {
		t.Errorf("Expected Banque Centrale de Tunisie, got %s", bankInfo.Bank.Name)
	}

	// Test governorate info
	govInfo := GetGovernorateFromPostalCode("1000")
	if govInfo == nil {
		t.Error("GetGovernorateFromPostalCode should return governorate info for valid postal code")
	}
	if govInfo.Name != "Tunis" {
		t.Errorf("Expected Tunis, got %s", govInfo.Name)
	}
}

func TestConstants(t *testing.T) {
	// Test that constants are accessible
	if CountryCode != "+216" {
		t.Errorf("CountryCode should be +216, got %s", CountryCode)
	}

	if len(ValidPrefixes) == 0 {
		t.Error("ValidPrefixes should not be empty")
	}

	if len(Carriers) == 0 {
		t.Error("Carriers should not be empty")
	}

	if len(Banks) == 0 {
		t.Error("Banks should not be empty")
	}

	if len(Governorates) == 0 {
		t.Error("Governorates should not be empty")
	}
}

func TestIsValidTunisianData(t *testing.T) {
	testData := map[string]string{
		"cin":      "12345678",
		"phone":    "20123456",
		"taxID":    "1234567A/P/M/000",
		"rib":      "01234567890123456789",
		"postal":   "1000",
		"carPlate": "123 تونس 4567",
	}

	results := IsValidTunisianData(testData)

	for field, expected := range map[string]bool{
		"cin":      true,
		"phone":    true,
		"taxID":    true,
		"rib":      true,
		"postal":   true,
		"carPlate": true,
	} {
		if results[field] != expected {
			t.Errorf("IsValidTunisianData[%s] = %t, expected %t", field, results[field], expected)
		}
	}
}

func BenchmarkMainPackageValidation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ValidateCIN("12345678")
		ValidatePhoneNumber("20123456")
		ValidateTaxID("1234567A/P/M/000")
		ValidateRIB("01234567890123456789")
		ValidatePostalCode("1000")
	}
}
