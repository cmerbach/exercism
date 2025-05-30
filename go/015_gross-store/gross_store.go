package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	return map[string]int{
		"quarter_of_a_dozen"	: 3,
		"half_of_a_dozen"			:	6,
		"dozen"								:	12,
		"small_gross"					: 120,
		"gross"								: 144,
		"great_gross"					: 1728,
	}
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return make(map[string]int)
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	quantity, exists := units[unit]
	if !exists {
		return false
	}
	
	bill[item] += quantity
	return true
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
		// Check if item exists in bill
		currentQty, itemExists := bill[item]
		if !itemExists {
			return false
		}
		
		// Check if unit exists in units map
		unitQty, unitExists := units[unit]
		if !unitExists {
			return false
		}
		
		newQty := currentQty - unitQty
		if newQty < 0 {
			return false
		} else if newQty == 0 {
			delete(bill, item)
		} else {
			bill[item] = newQty
		}
		
		return true
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	quantity, exists := bill[item]
	return quantity, exists
}
