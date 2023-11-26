package main

import (
	"fmt"

	"github.com/EniRox001/go_algorithms/sort"
)

func main() {
	sortedArr := sort.BubbleSort([]int{1, 5, 4, 2, 3, 10, 29, 1, 2, 2, 5, 3, 5, 2, 6, 8, 9, 22, 3, 1, 32})
	fmt.Println(sortedArr)
}
