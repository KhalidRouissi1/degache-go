package validators

import "testing"

func TestValidateCIN(t *testing.T) {
	tests := []struct {
		name     string
		cin      string
		expected bool
	}{
		{"Valid CIN starting with 0", "01234567", true},
		{"Valid CIN starting with 1", "12345678", true},
		{"Invalid CIN starting with 2", "22345678", false},
		{"Invalid CIN too short", "1234567", false},
		{"Invalid CIN too long", "123456789", false},
		{"Invalid CIN with letters", "1234567a", false},
		{"Empty CIN", "", false},
		{"Invalid CIN with spaces", "1234 5678", false},
		{"Invalid CIN with special chars", "1234-567", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ValidateCIN(tt.cin)
			if result != tt.expected {
				t.Errorf("ValidateCIN(%q) = %v, want %v", tt.cin, result, tt.expected)
			}
		})
	}
}

func TestValidateCINWithDetails(t *testing.T) {
	tests := []struct {
		name        string
		cin         string
		expectedOK  bool
		expectedMsg string
	}{
		{"Valid CIN", "12345678", true, ""},
		{"Empty CIN", "", false, "CIN cannot be empty"},
		{"Short CIN", "1234567", false, "CIN must be exactly 8 digits"},
		{"Long CIN", "123456789", false, "CIN must be exactly 8 digits"},
		{"Invalid format", "2234567a", false, "CIN must start with 0 or 1 and contain only digits"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ok, msg := ValidateCINWithDetails(tt.cin)
			if ok != tt.expectedOK {
				t.Errorf("ValidateCINWithDetails(%q) ok = %v, want %v", tt.cin, ok, tt.expectedOK)
			}
			if msg != tt.expectedMsg {
				t.Errorf("ValidateCINWithDetails(%q) msg = %q, want %q", tt.cin, msg, tt.expectedMsg)
			}
		})
	}
}

func BenchmarkValidateCIN(b *testing.B) {
	cin := "12345678"
	for i := 0; i < b.N; i++ {
		ValidateCIN(cin)
	}
}
