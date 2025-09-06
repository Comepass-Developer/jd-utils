package jslice

// Include check if value is included in slice
func Include[T comparable](arr []T, val T) bool {
	// Loop and check value is included
	for _, v := range arr {
		if v == val {
			return true
		}
	}
	return false
}
