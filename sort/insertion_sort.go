package sort

// InsertionSort performs the insertion sort algorithm on the given integer slice.
// It rearranges the elements in ascending order.
// The time complexity is O(n^2), where n is the number of elements in the slice.
// The space complexity is O(1) as it sorts the slice in-place.
func InsertionSort(arr []int) []int {
	// i is the index variable used to traverse the array
	// key is the current element being compared and moved to its correct position
	// j is the index used to traverse the array in reverse to find the correct position for key
	var i, key, j int

	// Traverse the array starting from the second element
	for i = 1; i < len(arr); i++ {
		// Set the key to the current element being considered
		key = arr[i]
		// Set j to the index before i
		j = i - 1

		// Move elements greater than key to one position ahead of their current position
		// until the correct position for key is found
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j = j - 1
		}

		// Place key at its correct position in the sorted part of the array
		arr[j+1] = key
	}

	// Return the sorted array
	return arr
}
