package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []int{1, 2, 2, 4, 5}
	asd := sort.IntsAreSorted(a)
	if asd == true {
		fmt.Println("Array is alredy sorted")
	} else {
		fmt.Println("Sorted array: ")
		bubbleSort(a)
	}
}
func bubbleSort(a []int) {
	sort.Ints(a)
	fmt.Println(a)
}
