package degache

import (
	"fmt"
	"testing"
	"time"

	"github.com/degache-go/degache/formatters"
	"github.com/degache-go/degache/types"
)

// Benchmark all main validation functions
func BenchmarkValidations(b *testing.B) {
	b.Run("CIN", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			ValidateCIN("12345678")
		}
	})

	b.Run("PhoneNumber", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			ValidatePhoneNumber("20123456")
		}
	})

	b.Run("PhoneNumberStrict", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			ValidatePhoneNumber("20123456", types.PhoneNumberValidationOptions{Strict: true})
		}
	})

	b.Run("TaxID", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			ValidateTaxID("1234567A/P/M/000")
		}
	})

	b.Run("RIB", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			ValidateRIB("01234567890123456789")
		}
	})

	b.Run("PostalCode", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			ValidatePostalCode("1000")
		}
	})

	b.Run("CarPlate", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			ValidateCarPlate("123 تونس 4567")
		}
	})
}

// Benchmark formatting functions
func BenchmarkFormatting(b *testing.B) {
	b.Run("PhoneNumber", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			FormatPhoneNumber("20123456")
		}
	})

	b.Run("Currency", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			FormatCurrency(1234.567)
		}
	})

	b.Run("Date", func(b *testing.B) {
		now := time.Now()
		for i := 0; i < b.N; i++ {
			FormatDate(now)
		}
	})
}

// Benchmark utility functions
func BenchmarkUtilities(b *testing.B) {
	b.Run("GetCarrierInfo", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			GetCarrierInfo("20123456")
		}
	})

	b.Run("GetBankFromRIB", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			GetBankFromRIB("01234567890123456789")
		}
	})

	b.Run("GetGovernorateFromPostalCode", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			GetGovernorateFromPostalCode("1000")
		}
	})
}

// Benchmark comprehensive validation
func BenchmarkComprehensiveValidation(b *testing.B) {
	testData := map[string]string{
		"cin":      "12345678",
		"phone":    "20123456",
		"taxID":    "1234567A/P/M/000",
		"rib":      "01234567890123456789",
		"postal":   "1000",
		"carPlate": "123 تونس 4567",
	}

	b.Run("IsValidTunisianData", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			IsValidTunisianData(testData)
		}
	})
}

// Benchmark memory allocations
func BenchmarkMemoryAllocations(b *testing.B) {
	b.Run("PhoneFormatting", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			formatters.FormatPhoneNumber("20123456")
		}
	})

	b.Run("CurrencyFormatting", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			formatters.FormatCurrency(1234.567)
		}
	})

	b.Run("ValidationOnly", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			ValidateCIN("12345678")
			ValidatePhoneNumber("20123456")
			ValidateTaxID("1234567A/P/M/000")
		}
	})
}

// Benchmark different input sizes
func BenchmarkInputSizes(b *testing.B) {
	// Test with various phone number formats
	phoneNumbers := []string{
		"20123456",        // 8 digits
		"+21620123456",    // with country code
		"20 123 456",      // with spaces
		"+216 20 123 456", // full format
	}

	for _, phone := range phoneNumbers {
		b.Run(fmt.Sprintf("Phone_%d_chars", len(phone)), func(b *testing.B) {
			for j := 0; j < b.N; j++ {
				ValidatePhoneNumber(phone)
			}
		})
	}
}

// Benchmark concurrent access
func BenchmarkConcurrentAccess(b *testing.B) {
	b.Run("ConcurrentValidation", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				ValidateCIN("12345678")
				ValidatePhoneNumber("20123456")
				ValidateTaxID("1234567A/P/M/000")
			}
		})
	})

	b.Run("ConcurrentFormatting", func(b *testing.B) {
		b.RunParallel(func(pb *testing.PB) {
			for pb.Next() {
				FormatPhoneNumber("20123456")
				FormatCurrency(1234.567)
			}
		})
	})
}

// Benchmark error cases
func BenchmarkErrorCases(b *testing.B) {
	b.Run("InvalidInputs", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			ValidateCIN("invalid")
			ValidatePhoneNumber("invalid")
			ValidateTaxID("invalid")
			ValidateRIB("invalid")
		}
	})
}
