// The package contains a function that will take a year and report if it is a leap year.
package leap

const testVersion = 2

// IsLeapYear is a function that takes a year and reports if it is a leap year.
func IsLeapYear(year int) bool {
	if year%4 != 0 {
		return false
	}
	if year%100 == 0 && year%400 != 0 {
		return false
	}
	return true
}
