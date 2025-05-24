package formatters

import "testing"

func TestFormatPhoneNumber(t *testing.T) {
	tests := []struct {
		name     string
		phone    string
		expected string
		hasError bool
	}{
		{"Valid phone", "20123456", "+216 20 123 456", false},
		{"Valid with prefix", "+21620123456", "+216 20 123 456", false},
		{"Valid with spaces", "20 123 456", "+216 20 123 456", false},
		{"Invalid phone", "10123456", "", true},
		{"Empty phone", "", "", true},
		{"Too short", "2012345", "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := FormatPhoneNumber(tt.phone)
			if tt.hasError {
				if err == nil {
					t.Errorf("FormatPhoneNumber(%q) expected error, got nil", tt.phone)
				}
			} else {
				if err != nil {
					t.Errorf("FormatPhoneNumber(%q) unexpected error: %v", tt.phone, err)
				}
				if result != tt.expected {
					t.Errorf("FormatPhoneNumber(%q) = %q, want %q", tt.phone, result, tt.expected)
				}
			}
		})
	}
}

func TestFormatPhoneNumberNational(t *testing.T) {
	tests := []struct {
		name     string
		phone    string
		expected string
		hasError bool
	}{
		{"Valid phone", "20123456", "20 123 456", false},
		{"Valid with prefix", "+21620123456", "20 123 456", false},
		{"Invalid phone", "10123456", "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := FormatPhoneNumberNational(tt.phone)
			if tt.hasError {
				if err == nil {
					t.Errorf("FormatPhoneNumberNational(%q) expected error, got nil", tt.phone)
				}
			} else {
				if err != nil {
					t.Errorf("FormatPhoneNumberNational(%q) unexpected error: %v", tt.phone, err)
				}
				if result != tt.expected {
					t.Errorf("FormatPhoneNumberNational(%q) = %q, want %q", tt.phone, result, tt.expected)
				}
			}
		})
	}
}

func TestFormatPhoneNumberCompact(t *testing.T) {
	tests := []struct {
		name     string
		phone    string
		expected string
		hasError bool
	}{
		{"Valid phone", "20123456", "+21620123456", false},
		{"Valid with spaces", "20 123 456", "+21620123456", false},
		{"Invalid phone", "10123456", "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := FormatPhoneNumberCompact(tt.phone)
			if tt.hasError {
				if err == nil {
					t.Errorf("FormatPhoneNumberCompact(%q) expected error, got nil", tt.phone)
				}
			} else {
				if err != nil {
					t.Errorf("FormatPhoneNumberCompact(%q) unexpected error: %v", tt.phone, err)
				}
				if result != tt.expected {
					t.Errorf("FormatPhoneNumberCompact(%q) = %q, want %q", tt.phone, result, tt.expected)
				}
			}
		})
	}
}

func TestNormalizePhoneNumber(t *testing.T) {
	tests := []struct {
		name     string
		phone    string
		expected string
		hasError bool
	}{
		{"Valid phone", "20123456", "20123456", false},
		{"With spaces", "20 123 456", "20123456", false},
		{"With prefix", "+216 20 123 456", "20123456", false},
		{"With dashes", "20-123-456", "20123456", false},
		{"Invalid phone", "10123456", "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := NormalizePhoneNumber(tt.phone)
			if tt.hasError {
				if err == nil {
					t.Errorf("NormalizePhoneNumber(%q) expected error, got nil", tt.phone)
				}
			} else {
				if err != nil {
					t.Errorf("NormalizePhoneNumber(%q) unexpected error: %v", tt.phone, err)
				}
				if result != tt.expected {
					t.Errorf("NormalizePhoneNumber(%q) = %q, want %q", tt.phone, result, tt.expected)
				}
			}
		})
	}
}
