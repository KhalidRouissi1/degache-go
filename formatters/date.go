package formatters

import (
	"fmt"
	"time"
)

// Tunisian month names in Arabic
var tunisianMonths = map[time.Month]string{
	time.January:   "جانفي",
	time.February:  "فيفري",
	time.March:     "مارس",
	time.April:     "أفريل",
	time.May:       "ماي",
	time.June:      "جوان",
	time.July:      "جويلية",
	time.August:    "أوت",
	time.September: "سبتمبر",
	time.October:   "أكتوبر",
	time.November:  "نوفمبر",
	time.December:  "ديسمبر",
}

// Tunisian weekday names in Arabic
var tunisianWeekdays = map[time.Weekday]string{
	time.Sunday:    "الأحد",
	time.Monday:    "الإثنين",
	time.Tuesday:   "الثلاثاء",
	time.Wednesday: "الأربعاء",
	time.Thursday:  "الخميس",
	time.Friday:    "الجمعة",
	time.Saturday:  "السبت",
}

// FormatDate formats a date in Tunisian style
//
// Parameters:
//   - date: The date to format
//
// Returns:
//   - string: formatted date in Tunisian style
//
// Example:
//
//	formatted := FormatDate(time.Now())
//	// Returns: "الإثنين، 15 جانفي 2024"
func FormatDate(date time.Time) string {
	weekday := tunisianWeekdays[date.Weekday()]
	day := date.Day()
	month := tunisianMonths[date.Month()]
	year := date.Year()

	return fmt.Sprintf("%s، %d %s %d", weekday, day, month, year)
}

// FormatDateShort formats a date in short Tunisian style (without weekday)
//
// Parameters:
//   - date: The date to format
//
// Returns:
//   - string: formatted date in short Tunisian style
//
// Example:
//
//	formatted := FormatDateShort(time.Now())
//	// Returns: "15 جانفي 2024"
func FormatDateShort(date time.Time) string {
	day := date.Day()
	month := tunisianMonths[date.Month()]
	year := date.Year()

	return fmt.Sprintf("%d %s %d", day, month, year)
}

// FormatDateNumeric formats a date in numeric format (DD/MM/YYYY)
//
// Parameters:
//   - date: The date to format
//
// Returns:
//   - string: formatted date in numeric format
//
// Example:
//
//	formatted := FormatDateNumeric(time.Now())
//	// Returns: "15/01/2024"
func FormatDateNumeric(date time.Time) string {
	return date.Format("02/01/2006")
}

// FormatDateTime formats a date and time in Tunisian style
//
// Parameters:
//   - dateTime: The date and time to format
//
// Returns:
//   - string: formatted date and time
//
// Example:
//
//	formatted := FormatDateTime(time.Now())
//	// Returns: "الإثنين، 15 جانفي 2024 في 14:30"
func FormatDateTime(dateTime time.Time) string {
	dateStr := FormatDate(dateTime)
	timeStr := dateTime.Format("15:04")

	return fmt.Sprintf("%s في %s", dateStr, timeStr)
}

// FormatDateTimeShort formats a date and time in short format
//
// Parameters:
//   - dateTime: The date and time to format
//
// Returns:
//   - string: formatted date and time in short format
//
// Example:
//
//	formatted := FormatDateTimeShort(time.Now())
//	// Returns: "15 جانفي 2024، 14:30"
func FormatDateTimeShort(dateTime time.Time) string {
	dateStr := FormatDateShort(dateTime)
	timeStr := dateTime.Format("15:04")

	return fmt.Sprintf("%s، %s", dateStr, timeStr)
}

// FormatTime formats time in 24-hour format
//
// Parameters:
//   - time: The time to format
//
// Returns:
//   - string: formatted time
//
// Example:
//
//	formatted := FormatTime(time.Now())
//	// Returns: "14:30"
func FormatTime(t time.Time) string {
	return t.Format("15:04")
}

// FormatTime12Hour formats time in 12-hour format with Arabic AM/PM
//
// Parameters:
//   - time: The time to format
//
// Returns:
//   - string: formatted time in 12-hour format
//
// Example:
//
//	formatted := FormatTime12Hour(time.Now())
//	// Returns: "2:30 مساءً" or "10:30 صباحاً"
func FormatTime12Hour(t time.Time) string {
	hour := t.Hour()
	minute := t.Minute()

	var period string
	var displayHour int

	if hour == 0 {
		displayHour = 12
		period = "صباحاً"
	} else if hour < 12 {
		displayHour = hour
		period = "صباحاً"
	} else if hour == 12 {
		displayHour = 12
		period = "مساءً"
	} else {
		displayHour = hour - 12
		period = "مساءً"
	}

	return fmt.Sprintf("%d:%02d %s", displayHour, minute, period)
}

// GetTunisianMonthName returns the Tunisian (Arabic) name for a month
//
// Parameters:
//   - month: The month
//
// Returns:
//   - string: Tunisian month name
//
// Example:
//
//	name := GetTunisianMonthName(time.January) // Returns: "جانفي"
func GetTunisianMonthName(month time.Month) string {
	return tunisianMonths[month]
}

// GetTunisianWeekdayName returns the Tunisian (Arabic) name for a weekday
//
// Parameters:
//   - weekday: The weekday
//
// Returns:
//   - string: Tunisian weekday name
//
// Example:
//
//	name := GetTunisianWeekdayName(time.Monday) // Returns: "الإثنين"
func GetTunisianWeekdayName(weekday time.Weekday) string {
	return tunisianWeekdays[weekday]
}
