package validators

import (
	"testing"

	"github.com/degache-go/degache/types"
)

func TestValidatePhoneNumber(t *testing.T) {
	tests := []struct {
		name     string
		phone    string
		options  types.PhoneNumberValidationOptions
		expected bool
	}{
		{"Valid Ooredoo number", "20123456", types.PhoneNumberValidationOptions{}, true},
		{"Valid Orange number", "40123456", types.PhoneNumberValidationOptions{}, true},
		{"Valid Telecom number", "90123456", types.PhoneNumberValidationOptions{}, true},
		{"Valid with international prefix", "+21620123456", types.PhoneNumberValidationOptions{}, true},
		{"Valid with spaces", "20 123 456", types.PhoneNumberValidationOptions{}, true},
		{"Invalid prefix", "10123456", types.PhoneNumberValidationOptions{}, false},
		{"Too short", "2012345", types.PhoneNumberValidationOptions{}, false},
		{"Too long", "201234567", types.PhoneNumberValidationOptions{}, false},
		{"Empty string", "", types.PhoneNumberValidationOptions{}, false},
		{"With letters", "2012345a", types.PhoneNumberValidationOptions{}, false},

		// Strict mode tests
		{"Valid strict", "20123456", types.PhoneNumberValidationOptions{Strict: true}, true},
		{"Valid strict with prefix", "+21620123456", types.PhoneNumberValidationOptions{Strict: true}, true},
		{"Invalid strict with spaces", "20 123 456", types.PhoneNumberValidationOptions{Strict: true}, false},
		{"Invalid strict with extra chars", "20-123-456", types.PhoneNumberValidationOptions{Strict: true}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ValidatePhoneNumber(tt.phone, tt.options)
			if result != tt.expected {
				t.Errorf("ValidatePhoneNumber(%q, %+v) = %v, want %v", tt.phone, tt.options, result, tt.expected)
			}
		})
	}
}

func TestGetCarrierInfo(t *testing.T) {
	tests := []struct {
		name         string
		phone        string
		expectedName string
		shouldBeNil  bool
	}{
		{"Ooredoo number", "20123456", "Ooredoo Tunisia", false},
		{"Orange number", "40123456", "Orange Tunisia", false},
		{"Telecom number", "90123456", "Tunisie Telecom", false},
		{"Invalid number", "10123456", "", true},
		{"Empty string", "", "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GetCarrierInfo(tt.phone)
			if tt.shouldBeNil {
				if result != nil {
					t.Errorf("GetCarrierInfo(%q) = %+v, want nil", tt.phone, result)
				}
			} else {
				if result == nil {
					t.Errorf("GetCarrierInfo(%q) = nil, want non-nil", tt.phone)
				} else if result.Carrier.Name != tt.expectedName {
					t.Errorf("GetCarrierInfo(%q).Carrier.Name = %q, want %q", tt.phone, result.Carrier.Name, tt.expectedName)
				}
			}
		})
	}
}

func BenchmarkValidatePhoneNumber(b *testing.B) {
	phone := "20123456"
	for i := 0; i < b.N; i++ {
		ValidatePhoneNumber(phone)
	}
}
