package cards

// FavoriteCards returns a slice with the cards 2, 6 and 9 in that order.
func FavoriteCards() []int {
	return []int{2,6,9}
}

// GetItem retrieves an item from a slice at given position.
// If the index is out of range, we want it to return -1.
func GetItem(slice []int, index int) int {
    if index < len(slice) && index >= 0 {
        return slice[index]
    }
    return -1
}

// SetItem writes an item to a slice at given position overwriting an existing value.
// If the index is out of range the value needs to be appended.
func SetItem(slice []int, index, value int) []int {
	if index < len(slice) && index >= 0 {
        slice[index] = value
        return slice
    }
    return append(slice, value)
}

// PrependItems adds an arbitrary number of values at the front of a slice.
func PrependItems(slice []int, values ...int) []int {
	a := []int{}
    a = append(a, values...)
    a = append(a, slice...)
    return a
}

// RemoveItem removes an item from a slice by modifying the existing slice.
func RemoveItem(slice []int, index int) []int {
    var tempSliceA []int
    var tempSliceB []int
    var tempSliceC []int
    lastItemIndex := len(slice) - 1
    // Remove the first item
    if index == 0 {
        index++
        tempSliceA = slice[index:]
        return tempSliceA
    } else if index == lastItemIndex { // Remove the last item
        tempSliceA = slice[:index]
        return tempSliceA
    } else if index > lastItemIndex || index < 0 { // Remove out of band item
        return slice
    } else {
        tempSliceA = slice[:index]
        tempSliceB = slice[(index+1):]
        tempSliceC = append(tempSliceA, tempSliceB...)
        return tempSliceC
    }
}
