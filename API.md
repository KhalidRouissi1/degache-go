# degache-go API Documentation

This document provides comprehensive API documentation for the degache-go library.

## Table of Contents

- [Installation](#installation)
- [Quick Start](#quick-start)
- [Validators](#validators)
- [Formatters](#formatters)
- [Constants](#constants)
- [Types](#types)
- [Error Handling](#error-handling)

## Installation

```bash
go get github.com/degache-go/degache
```

## Quick Start

```go
package main

import (
    "fmt"
    "github.com/degache-go/degache"
)

func main() {
    // Validate CIN
    isValid := degache.ValidateCIN("12345678")
    fmt.Printf("CIN valid: %t\n", isValid)

    // Format phone number
    formatted, err := degache.FormatPhoneNumber("20123456")
    if err == nil {
        fmt.Printf("Formatted phone: %s\n", formatted)
    }
}
```

## Validators

### CIN Validation

#### `ValidateCIN(cin string) bool`

Validates a Tunisian CIN (Carte d'Identité Nationale).

**Parameters:**
- `cin`: 8-digit string starting with 0 or 1

**Returns:**
- `bool`: true if valid

**Example:**
```go
isValid := degache.ValidateCIN("12345678") // true
isValid := degache.ValidateCIN("22345678") // false
```

#### `ValidateCINWithDetails(cin string) (bool, string)`

Validates CIN with detailed error messages.

**Returns:**
- `bool`: validation result
- `string`: error message if invalid

### Phone Number Validation

#### `ValidatePhoneNumber(phoneNumber string, options ...PhoneNumberValidationOptions) bool`

Validates Tunisian phone numbers.

**Parameters:**
- `phoneNumber`: phone number string
- `options`: optional validation options

**Options:**
```go
type PhoneNumberValidationOptions struct {
    Strict bool // Enforce strict format validation
}
```

**Example:**
```go
// Basic validation
isValid := degache.ValidatePhoneNumber("20123456") // true

// Strict validation
isValid := degache.ValidatePhoneNumber("20 123 456", 
    degache.PhoneNumberValidationOptions{Strict: true}) // false
```

#### `GetCarrierInfo(phoneNumber string, options ...PhoneNumberValidationOptions) *CarrierInfo`

Gets carrier information from phone number.

**Returns:**
- `*CarrierInfo`: carrier details or nil if invalid

### Tax ID Validation

#### `ValidateTaxID(taxID string) bool`

Validates Tunisian Tax ID (Matricule Fiscal).

**Format:** 7 digits + letter/letter/letter/3 digits

**Example:**
```go
isValid := degache.ValidateTaxID("1234567A/P/M/000") // true
```

### RIB Validation

#### `ValidateRIB(rib string) bool`

Validates Tunisian RIB (bank account number).

**Parameters:**
- `rib`: 20-digit string

#### `GetBankFromRIB(rib string) *BankInfo`

Extracts bank information from RIB.

### Postal Code Validation

#### `ValidatePostalCode(postalCode string) bool`

Validates Tunisian postal codes.

#### `GetGovernorateFromPostalCode(postalCode string) *Governorate`

Gets governorate information from postal code.

### Car Plate Validation

#### `ValidateCarPlate(carPlate string, options ...CarPlateValidationOptions) bool`

Validates Tunisian car plates.

**Options:**
```go
type CarPlateValidationOptions struct {
    Type   string // "standard" or "special"
    Strict bool   // Enforce strict spacing
}
```

## Formatters

### Phone Number Formatting

#### `FormatPhoneNumber(phoneNumber string) (string, error)`

Formats phone number with international prefix.

**Returns:** `"+216 20 123 456"`

#### `FormatPhoneNumberNational(phoneNumber string) (string, error)`

Formats phone number without country code.

**Returns:** `"20 123 456"`

#### `FormatPhoneNumberCompact(phoneNumber string) (string, error)`

Formats phone number in compact format.

**Returns:** `"+21620123456"`

### Currency Formatting

#### `FormatCurrency(amount float64, options ...CurrencyFormatOptions) string`

Formats amounts in Tunisian Dinar.

**Options:**
```go
type CurrencyFormatOptions struct {
    Symbol bool // Include currency symbol
    Code   bool // Include currency code
}
```

**Example:**
```go
formatted := degache.FormatCurrency(1234.56, 
    degache.CurrencyFormatOptions{Symbol: true})
// Returns: "1.234,560 دينار تونسي"
```

### Date Formatting

#### `FormatDate(date time.Time) string`

Formats date in Tunisian style with Arabic text.

**Returns:** `"الإثنين، 15 جانفي 2024"`

## Constants

### Carriers

```go
var Carriers = map[string]Carrier{
    "OOREDOO": {Name: "Ooredoo Tunisia", Prefixes: []string{"2", "5"}},
    "ORANGE":  {Name: "Orange Tunisia", Prefixes: []string{"4"}},
    "TELECOM": {Name: "Tunisie Telecom", Prefixes: []string{"9"}},
}
```

### Banks

28 Tunisian banks with codes and names.

### Governorates

24 Tunisian governorates with postal codes and regions.

### Other Constants

- `CountryCode`: "+216"
- `ValidPrefixes`: []string{"2", "4", "5", "9"}

## Types

### Core Types

```go
type CIN string        // 8-digit CIN
type PhoneNumber string // 8-digit phone number
type TaxID string      // Tax ID format
type RIB string        // 20-digit RIB
type PostalCode string // 4-digit postal code
```

### Information Types

```go
type CarrierInfo struct {
    Carrier Carrier
    Prefix  string
}

type BankInfo struct {
    Bank Bank
    Code string
}

type CarPlateInfo struct {
    Type       string
    Components CarPlateComponents
}
```

## Error Handling

### Validation Errors

Most validation functions return boolean results. For detailed errors, use the `*WithDetails` variants:

```go
valid, msg := degache.ValidateCINWithDetails("123")
if !valid {
    fmt.Printf("Error: %s\n", msg)
}
```

### Formatting Errors

Formatting functions return errors for invalid inputs:

```go
formatted, err := degache.FormatPhoneNumber("invalid")
if err != nil {
    fmt.Printf("Formatting error: %s\n", err)
}
```

## Best Practices

1. **Always handle errors** from formatting functions
2. **Use strict mode** for user input validation
3. **Cache validation results** for performance
4. **Validate early** in your application flow
5. **Use detailed validation** for user feedback

## Performance

- All validators are optimized for performance
- Regular expressions are compiled once
- Constants are pre-computed
- Memory allocation is minimized

## Thread Safety

All functions in degache-go are thread-safe and can be used concurrently without additional synchronization.
