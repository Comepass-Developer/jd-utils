// pkg/stringHelper/prefix_str.go
package jstr

import (
	"fmt"
	"time"
)

// PrefixStr generate prefix string with year and month
func PrefixStr(prefix string, year bool, unit int, number int) string {
	// Get current time
	now := time.Now()
	curY := now.Year()
	curM := int(now.Month())
	m := fmt.Sprintf("%02d", curM)
	y := fmt.Sprintf("%04d", curY)
	if year {
		y = fmt.Sprintf("%02d", curY%100)
	}
	// Concat prefix with year, month
	base := fmt.Sprintf("%s%s%s", prefix, y, m)

	// If init value, concat with serial
	if unit > 0 {
		serial := fmt.Sprintf("%0*d", unit, number)
		return base + serial
	}
	return base
}
