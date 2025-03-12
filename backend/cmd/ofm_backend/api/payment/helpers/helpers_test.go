package helpers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateCardNumber(t *testing.T) {
	tests := []struct {
		cardNumber string
		name       string
		isValid    bool
	}{
		{
			cardNumber: "4539148803436467",
			name:       "Valid card",
			isValid:    true,
		},
		{
			cardNumber: "1234567891011121",
			name:       "Invalid card",
			isValid:    false,
		},
		{
			cardNumber: "",
			name:       "Empty card number",
			isValid:    false,
		},
		{
			cardNumber: "1234",
			name:       "Invalid card (too short)",
			isValid:    false,
		},
		{
			cardNumber: "453914880343646745391488034364674539148803436467",
			name:       "Invalid card (too long)",
			isValid:    false,
		},
		{
			cardNumber: "453914880343646A",
			name:       "Card with non-numeric characters",
			isValid:    false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if test.isValid {
				assert.True(t, ValidateCardNumber(test.cardNumber), "Expected valid card number for: "+test.cardNumber)
			} else {
				assert.False(t, ValidateCardNumber(test.cardNumber), "Expected invalid card number for: "+test.cardNumber)
			}
		})
	}
}
