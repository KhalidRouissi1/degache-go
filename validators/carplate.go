package validators

import (
	"regexp"
	"strings"

	"github.com/degache-go/degache/types"
)

// Regular expressions for car plate validation
var (
	// Standard format: XXX تونس XXXX
	standardCarPlateRegex = regexp.MustCompile(`^\d{3}\s*تونس\s*\d{4}$`)
	// Special format: RS XXX تونس
	specialCarPlateRegex = regexp.MustCompile(`^RS\s*\d{3}\s*تونس$`)
	// Strict format (no extra spaces)
	strictStandardCarPlateRegex = regexp.MustCompile(`^\d{3} تونس \d{4}$`)
	strictSpecialCarPlateRegex  = regexp.MustCompile(`^RS \d{3} تونس$`)
)

// ValidateCarPlate validates a Tunisian car plate
// Supports both standard (XXX تونس XXXX) and special (RS XXX تونس) formats
//
// Parameters:
//   - carPlate: The car plate to validate
//   - options: Validation options (optional)
//
// Returns:
//   - bool: true if the car plate is valid, false otherwise
//
// Example:
//
//	isValid := ValidateCarPlate("123 تونس 4567", types.CarPlateValidationOptions{})
//	isValid := ValidateCarPlate("RS 123 تونس", types.CarPlateValidationOptions{Type: "special"})
//	isValid := ValidateCarPlate("123  تونس  4567", types.CarPlateValidationOptions{Strict: true}) // false
func ValidateCarPlate(carPlate string, options ...types.CarPlateValidationOptions) bool {
	if carPlate == "" {
		return false
	}

	var opts types.CarPlateValidationOptions
	if len(options) > 0 {
		opts = options[0]
	}

	// Normalize spaces if not in strict mode
	normalizedPlate := carPlate
	if !opts.Strict {
		// Replace multiple spaces with single space
		normalizedPlate = regexp.MustCompile(`\s+`).ReplaceAllString(strings.TrimSpace(carPlate), " ")
	}

	// Check based on type
	switch opts.Type {
	case "special":
		if opts.Strict {
			return strictSpecialCarPlateRegex.MatchString(normalizedPlate)
		}
		return specialCarPlateRegex.MatchString(normalizedPlate)
	case "standard", "":
		if opts.Strict {
			return strictStandardCarPlateRegex.MatchString(normalizedPlate)
		}
		return standardCarPlateRegex.MatchString(normalizedPlate)
	default:
		return false
	}
}

// ValidateCarPlateWithDetails validates a car plate and returns detailed information
//
// Parameters:
//   - carPlate: The car plate to validate
//   - options: Validation options (optional)
//
// Returns:
//   - bool: true if the car plate is valid
//   - string: error message if invalid, empty string if valid
//
// Example:
//
//	valid, msg := ValidateCarPlateWithDetails("123 تونس 4567")
//	if !valid {
//	    fmt.Println("Invalid car plate:", msg)
//	}
func ValidateCarPlateWithDetails(carPlate string, options ...types.CarPlateValidationOptions) (bool, string) {
	if carPlate == "" {
		return false, "Car plate cannot be empty"
	}

	var opts types.CarPlateValidationOptions
	if len(options) > 0 {
		opts = options[0]
	}

	// Check if it contains the required Arabic text
	if !strings.Contains(carPlate, "تونس") {
		return false, "Car plate must contain 'تونس'"
	}

	// Normalize spaces if not in strict mode
	normalizedPlate := carPlate
	if !opts.Strict {
		normalizedPlate = regexp.MustCompile(`\s+`).ReplaceAllString(strings.TrimSpace(carPlate), " ")
	}

	// Check based on type
	switch opts.Type {
	case "special":
		if opts.Strict {
			if !strictSpecialCarPlateRegex.MatchString(normalizedPlate) {
				return false, "Special car plate must follow format: RS XXX تونس (with exact spacing)"
			}
		} else {
			if !specialCarPlateRegex.MatchString(normalizedPlate) {
				return false, "Special car plate must follow format: RS XXX تونس"
			}
		}
	case "standard", "":
		if opts.Strict {
			if !strictStandardCarPlateRegex.MatchString(normalizedPlate) {
				return false, "Standard car plate must follow format: XXX تونس XXXX (with exact spacing)"
			}
		} else {
			if !standardCarPlateRegex.MatchString(normalizedPlate) {
				return false, "Standard car plate must follow format: XXX تونس XXXX"
			}
		}
	default:
		return false, "Invalid car plate type specified"
	}

	return true, ""
}

// GetCarPlateInfo extracts information from a valid car plate
//
// Parameters:
//   - carPlate: The car plate to extract information from
//   - options: Validation options (optional)
//
// Returns:
//   - *types.CarPlateInfo: car plate information or nil if invalid
//
// Example:
//
//	info := GetCarPlateInfo("123 تونس 4567")
//	if info != nil {
//	    fmt.Printf("Type: %s, Prefix: %s, Region: %s, Suffix: %s\n",
//	        info.Type, info.Components.Prefix, info.Components.Region, info.Components.Suffix)
//	}
func GetCarPlateInfo(carPlate string, options ...types.CarPlateValidationOptions) *types.CarPlateInfo {
	var opts types.CarPlateValidationOptions
	if len(options) > 0 {
		opts = options[0]
	}

	if !ValidateCarPlate(carPlate, opts) {
		return nil
	}

	// Normalize spaces
	normalizedPlate := carPlate
	if !opts.Strict {
		normalizedPlate = regexp.MustCompile(`\s+`).ReplaceAllString(strings.TrimSpace(carPlate), " ")
	}

	// Determine type and extract components
	if strings.HasPrefix(normalizedPlate, "RS") {
		// Special plate
		parts := strings.Fields(normalizedPlate)
		if len(parts) >= 3 {
			return &types.CarPlateInfo{
				Type: "special",
				Components: types.CarPlateComponents{
					Prefix: parts[0] + " " + parts[1], // "RS XXX"
					Region: parts[2],                  // "تونس"
					Suffix: "",                        // No suffix for special plates
				},
			}
		}
	} else {
		// Standard plate
		parts := strings.Fields(normalizedPlate)
		if len(parts) >= 3 {
			return &types.CarPlateInfo{
				Type: "standard",
				Components: types.CarPlateComponents{
					Prefix: parts[0], // "XXX"
					Region: parts[1], // "تونس"
					Suffix: parts[2], // "XXXX"
				},
			}
		}
	}

	return nil
}
