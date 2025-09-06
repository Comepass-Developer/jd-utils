// pkg/stringHelper/integer_str.go
package jstr

import "regexp"

// FormatPhoneStr remove all non-numeric characters from phone number
func FormatPhoneStr(phone string) string {
	re := regexp.MustCompile("[^0-9]+")
	cleanPhone := re.ReplaceAllString(phone, "")
	return cleanPhone
}
