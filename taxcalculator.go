package taxcalculator

import (
	"math"
)

// Exclusive function takes main amount and taxRate (here, taxRate as parcent(%) value).
// If the value of amount is 100 and the value of taxRate is 15 then the return value is 15 (rounded value)
func Exclusive(amount, taxRate float64) float64 {
	tax := (taxRate * amount) / 100

	return math.Round(tax)
}

// Inclusive function takes main amount and taxRate (here, taxRate as parcent(%) value).
// If the value of amount is 100 and the value of taxRate is 15 then the return value is 13 (rounded value)
func Inclusive(amount, taxRate float64) float64 {
	tempAmount := Exclusive(100, taxRate) + 100

	tax := (amount * 100) / tempAmount

	return math.Round(amount - tax)
}
