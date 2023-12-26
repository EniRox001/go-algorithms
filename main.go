package main

import (
	"math/rand"
	"fmt"
	"time"

	"github.com/EniRox001/go_algorithms/sort"
)

func main() {
	rand.NewSource(time.Now().UnixNano())
	unsortedArr := rand.Perm(10)
	sortedArr := sort.InsertionSort(unsortedArr)
	fmt.Println(sortedArr)
}
