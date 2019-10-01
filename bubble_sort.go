package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 5, 4, 6}
	fmt.Println("Unsorted array: ", a)
	bubbleSort(a)
	fmt.Println("Sorted array: ", a)
}
func bubbleSort(a []int) {
	sorted := false
	end := len(a) - 1
	for {
		if sorted {
			break
		}
		fmt.Println(end)
		swapped := false
		for i := 0; i < len(a)-1; i++ {
			//fmt.Println(end, i)
			if a[i] > a[i+1] {
				swapped = true
				a[i], a[i+1] = a[i+1], a[i]
			}
		}
		if !swapped {
			sorted = true
		}
	}
}
