package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	out := map[string]int{
		"quarter_of_a_dozen": 3,
		"half_of_a_dozen":    6,
		"dozen":              12,
		"small_gross":        120,
		"gross":              144,
		"great_gross":        1728,
	}
	return out
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return map[string]int{}
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	if val, ok := units[unit]; ok {
		bill[item] += val
		return true
	}
	return false
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	cur, ok := GetItem(bill, item)
	if !ok {
		return false
	}
	if val, ok := units[unit]; ok {
		switch {
		case val < cur:
			bill[item] = cur - val
			return true
		case cur < val:
			return false
		default:
			delete(bill, item)
			return true
		}
	}
	return false
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	val, ok := bill[item]
	return val, ok
}
