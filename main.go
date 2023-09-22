package main

import (
	"fmt"
)

func minmax(a, b int) (int, int) {
	if a < b {
		return a, b
	}
	return b, a
}

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	num1 := 10
	num2 := 5

	min, max := minmax(num1, num2)

	fmt.Printf("Minimum: %d\n", min)
	fmt.Printf("Maximum: %d\n", max)

	v1, v2 := swap("kapoor", "vinay")
	fmt.Println(v1, v2)

}
