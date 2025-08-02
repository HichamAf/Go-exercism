package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
    units := map[string]int{
        "quarter_of_a_dozen": 3, 
        "half_of_a_dozen": 6, 
        "dozen": 12,  
        "small_gross": 120, 
        "gross": 144, 
        "great_gross": 1728,
    }
	return units
}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return map[string]int{}
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
    unitQuantity, unitExists := units[unit]
    if !unitExists {
        return false // Return false if the given unit is not in the units map
    }
    // Check if the item already exists in the bill
    itemQuantity, itemExists := bill[item]
    if itemExists {
        // If item is already exists, increase its quantity by the amount that belongs to the provided unit.
        bill[item] = itemQuantity + unitQuantity
    } else {
        // If item doesn't exist, add the unit quantity
        bill[item] = unitQuantity
    }
    return true
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
    itemQuantity, billExists := bill[item]
    if !billExists {
        return false
    }

    unitQuantity, unitExists := units[unit]
    if !unitExists {
        return false
    }

    newItemQuantity := itemQuantity - unitQuantity
    
    if newItemQuantity < 0 {
        return false
    }
    if newItemQuantity == 0 {
        delete(bill, item)
        return true
    }
    bill[item] = newItemQuantity
    return true
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	itemQuantity, billExists := bill[item]
    if !billExists {
        return 0, false
    }
    return itemQuantity, true
}
