package main

import (
	"fmt"

	"github.com/EniRox001/go_algorithms/recursion"
)

func main() {
	// rand.NewSource(time.Now().UnixNano())
	// unsortedArr := rand.Perm(10)
	// sortedArr := sort.InsertionSort(unsortedArr)
	// fmt.Println(sortedArr)

	// call nested additon function
	nestedSlice := []interface{}{1, 2, []interface{}{3, 4, 5}, 6}
	nestedAddition := recursion.NestedAddition(nestedSlice)
	fmt.Println(nestedAddition)
}
