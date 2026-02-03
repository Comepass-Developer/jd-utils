package jnum

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"strings"
)

// GenerateDigits create a random string of digits base on amount number
// Ex: GenerateDigits(5) -> "12345"
func GenerateDigits(amount int) (string, error) {
	// Check amount is valid (no negative)
	if amount <= 0 {
		return "", fmt.Errorf("amount must be greater than 0")
	}

	// Use strings.Builder to build string
	var builder strings.Builder
	builder.Grow(amount)

	// Define random range
	max := big.NewInt(10)

	// Loop 'amount' times to get each digit
	for i := 0; i < amount; i++ {
		// Get a random number safely in the range [0, 9]
		n, err := rand.Int(rand.Reader, max)
		if err != nil {
			return "", err
		}

		// Append digit to builder
		builder.WriteString(n.String())
	}

	return builder.String(), nil
}
