package recursion

import (
	"reflect"
)

// NestedAddition recursively calculates the sum of integers in a nested slice.
// It traverses through the elements of the input slice, and if an element is
// an integer, it adds it to the sum. If an element is a nested slice, it
// recursively calls itself to calculate the sum of the nested slice.
func NestedAddition(arr []interface{}) int {
	// Initialize the sum to zero
	sum := 0

	// Iterate over each element in the slice
	for _, v := range arr {
		// Check if the element is a slice
		if reflect.TypeOf(v).Kind() == reflect.Slice {
			// If it's a slice, recursively call NestedAddition
			// to calculate the sum of the nested slice
			sum += NestedAddition(v.([]interface{}))
		} else {
			// If it's an integer, add it to the sum
			sum += v.(int)
		}
	}

	// Return the calculated sum
	return sum
}

