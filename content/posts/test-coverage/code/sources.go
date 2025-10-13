package main

// CalcDiscount calculates discount based on user age and activity status.
// It returns a discount multiplier: 1.0 for no discount, 0.9 for 10% discount and so on.
// It might return values greater than 1.0 for negative discounts.
func CalcDiscount(userAge int, isUserActive bool) float64 {
	discountMul := 1.0

	if userAge < 18 {
		discountMul *= 0.9 // 10% discount for minors
	}

	if !isUserActive {
		discountMul *= 0.8 // 20% discount for inactive users
	}

	return discountMul
}
