package main

import (
	"fmt"
)

func main() {
	//input1 := bufio.NewScanner(os.Stdin)
	//input2 := bufio.NewScanner(os.Stdin)
	var n int
	fmt.Println("Unesite duzinu niza: ")
	//input1.Scan()
	//n, _ := strconv.Atoi(input1.Text())
	fmt.Scanf("%d", &n)
	a := make([]int, n)
	fmt.Println("Unesite elemente niza: ")
	for i := 0; i < n; i++ {
		fmt.Println("Unesite", i, ". element niza")
		fmt.Scanf("%d", &a[i])
	}
	//fmt.Scanf("%d", &a)
	fmt.Println("Elementi niza: ", a)
}
