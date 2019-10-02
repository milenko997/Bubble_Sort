package main

import "fmt"

func main() {
	a := []int{2, 1, 3, 4, 5, 6}
	fmt.Println("Unsorted array: ", a)
	bubbleSort(a)
	fmt.Println("Sorted array: ", a)
}
func bubbleSort(a []int) {
	end := len(a) - 1
	for {
		swapped := false
		fmt.Println(end)
		for i := 0; i < len(a)-1; i++ {
			//fmt.Println(end, i)
			if a[i] > a[i+1] {
				swapped = true
				a[i], a[i+1] = a[i+1], a[i]
			}
		}
		if !swapped {
			break
		}
	}
}
