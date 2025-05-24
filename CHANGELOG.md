# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [1.0.0] - 2025-01-24

### Added
- Initial release of degache-go
- CIN (Carte d'Identité Nationale) validation
- Phone number validation and formatting with carrier detection
- Tax ID (Matricule Fiscal) validation
- RIB (bank account) validation with bank identification
- Postal code validation with governorate mapping
- Car plate validation (Arabic format support)
- Currency formatting in Tunisian Dinar
- Date formatting in Tunisian style (Arabic)
- Comprehensive constants for carriers, banks, and governorates
- Full test coverage
- Example usage code
- Documentation and README

### Features
- ✅ CIN validation (8-digit format starting with 0 or 1)
- ✅ Phone number validation with carrier detection (Ooredoo, Orange, Tunisie Telecom)
- ✅ Phone number formatting (international, national, compact)
- ✅ Tax ID validation (7 digits + letter/letter/letter/3 digits format)
- ✅ RIB validation with bank identification
- ✅ Postal code validation with governorate mapping
- ✅ Car plate validation (standard and special formats)
- ✅ Currency formatting with Tunisian locale
- ✅ Date formatting in Arabic
- ✅ Strict validation modes
- ✅ Detailed error messages
- ✅ Type-safe APIs
- ✅ High performance
- ✅ Comprehensive test coverage

### Technical Details
- Go 1.21+ compatibility
- Zero external dependencies
- Modular package structure
- Comprehensive error handling
- Performance optimized
- Memory efficient
- Thread-safe operations

### Package Structure
- `degache` - Main package with convenience exports
- `validators/` - All validation functions
- `formatters/` - All formatting functions
- `constants/` - Tunisian-specific constants
- `types/` - Type definitions and structs
- `examples/` - Usage examples

### Validation Coverage
- 28 Tunisian banks supported
- 3 mobile carriers (Ooredoo, Orange, Tunisie Telecom)
- 24 governorates with postal codes
- Standard and special car plate formats
- International and national phone number formats
- Strict and flexible validation modes
