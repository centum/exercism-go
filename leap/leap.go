// Package leap implements calendar functions
package leap

// IsLeapYear returns year is leap year.
func IsLeapYear(year int) bool {
	if year % 4 != 0 {
		return false
	}
	if year % 100 == 0 && year % 400 != 0 {
		return false
	}
	return true
}
