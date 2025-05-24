# 🌴 degache-go

**Tunisian Developer's Essential Utility Library for Go**

A Go port of the popular [degachejs](https://github.com/KhalidRouissi1/degachejs) library, providing essential utilities for Tunisian developers working with CIN validation, phone number formatting, tax IDs, and more.

## 🚀 Features

### CIN (Carte d'Identité Nationale) 🆔
- ✅ Validate Tunisian national ID numbers
- 🔒 Secure validation with proper format checking

### Phone Numbers 📱
- ✅ Support for all Tunisian carriers (Ooredoo, Orange, Tunisie Telecom)
- 🔄 International format conversion
- 📞 Smart formatting with country code
- 🏢 Carrier detection
- 🔒 Strict mode validation

### Tax ID (Matricule Fiscal) 💼
- ✅ Validate Tunisian tax identification numbers
- 📋 Proper format validation

### Currency Formatting 💰
- ✅ Format amounts in Tunisian Dinar
- 🌍 Localized formatting

### Bank Account (RIB) Validation 🏦
- ✅ Validate Tunisian bank account numbers
- 🏦 Bank identification from RIB

### Car Plates 🚗
- ✅ Validate Tunisian car plates (Arabic format support)
- 🔍 Extract plate information

### Postal Codes 📮
- ✅ Validate Tunisian postal codes
- 📍 Governorate mapping

## 📦 Installation

```bash
go get github.com/degache-go/degache
```

## 🛠️ Usage

```go
package main

import (
    "fmt"
    "github.com/degache-go/degache"
    "github.com/degache-go/degache/validators"
    "github.com/degache-go/degache/formatters"
)

func main() {
    // Validate CIN
    isValid := validators.ValidateCIN("12345678")
    fmt.Printf("CIN valid: %t\n", isValid)

    // Format phone number
    formatted, err := formatters.FormatPhoneNumber("20123456")
    if err == nil {
        fmt.Printf("Formatted phone: %s\n", formatted)
    }

    // Validate Tax ID
    isValidTax := validators.ValidateTaxID("1234567A/P/M/000")
    fmt.Printf("Tax ID valid: %t\n", isValidTax)
}
```

## 🤝 Contributing

We welcome contributions from the Tunisian developer community! Whether it's:

- 🐛 Bug fixes
- ✨ New features
- 📚 Documentation improvements
- 🧪 Test cases
- 💡 Feature suggestions

## 📄 License

degache-go is MIT licensed.

## 🏆 Production Ready

- ✅ Comprehensive input validation
- ⚡ Optimized performance
- 🧪 High test coverage
- 📚 Detailed documentation
- 🔒 Type-safe APIs

---

Built with ❤️ for the Tunisian developer community 🇹🇳

*Go port of the original [degachejs](https://github.com/KhalidRouissi1/degachejs) by Khalid Rouissi*
