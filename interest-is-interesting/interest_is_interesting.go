package interest

// InterestRate returns the interest rate for the provided balance.
// - 3.213% for a balance less than `0` dollars (balance gets more negative).
// - 0.5% for a balance greater than or equal to `0` dollars, and less than `1000` dollars.
// - 1.621% for a balance greater than or equal to `1000` dollars, and less than `5000` dollars.
// - 2.475% for a balance greater than or equal to `5000` dollars.
func InterestRate(balance float64) float32 {
	switch {
	case balance < 0:
		return 3.213
	case balance >= 0 && balance < 1000:
		return 0.5
	case balance >= 1000 && balance < 5000:
		return 1.621
	case balance >= 5000:
		return 2.475
	}
	return 0.5
}

// Interest calculates the interest for the provided balance.
func Interest(balance float64) float64 {
	return balance * float64(InterestRate(balance)/100.0)

}

// AnnualBalanceUpdate calculates the annual balance update, taking into account the interest rate.
func AnnualBalanceUpdate(balance float64) float64 {
	return balance + Interest(balance)

}

// YearsBeforeDesiredBalance calculates the minimum number of years required to reach the desired balance.
func YearsBeforeDesiredBalance(balance, targetBalance float64) int {
	max_iter := int(1e10)
	for i := 0; i < max_iter; i++ {
		if balance >= targetBalance {
			return i
		}
		balance = AnnualBalanceUpdate(balance)
	}
	return -1
}
