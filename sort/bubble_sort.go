package sort

// BubbleSort performs the bubble sort algorithm on the given integer slice.
// It rearranges the elements in ascending order.
// The time complexity is O(n^2), where n is the number of elements in the slice.
// The space complexity is O(1) as it sorts the slice in-place.
func BubbleSort(arr []int) []int {
	// Outer loop for the number of passes through the array
	for i := 0; i < len(arr)-1; i++ {
		// Inner loop for comparing and swapping adjacent elements
		for j := 0; j < len(arr)-i-1; j++ {
			// Compare adjacent elements and swap them if they are in the wrong order
			if arr[j] > arr[j+1] {
				arr[j+1], arr[j] = arr[j], arr[j+1]
			}
		}
	}

	// Return the sorted array
	return arr
}
