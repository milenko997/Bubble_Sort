package main

import "fmt"

func main() {
	a := []int{4, 9, 8, 2, 7}
	fmt.Println("Unsorted array: ",a)
	bubbleSort(a)
	fmt.Println("Sorted array: ", a)
}
func bubbleSort(a []int) {
	end := len(a) - 1
	for {
		if end == 0 {
			break
		}
		for i := 0; i < len(a)-1; i++ {
			if a[i] > a[i+1] {
				a[i], a[i+1] = a[i+1], a[i]
			}
		}
		end -= 1
	}
}
