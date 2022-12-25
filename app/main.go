package main

import "fmt"

type number []int32

func main() {

	numbers := number{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, n := range numbers {
		checkNumber(n)
	}
}

func checkNumber(n int32) {
	if n%2 == 0 {
		fmt.Println(n, "is even")
	} else {
		fmt.Println(n, "is odd")
	}
}
