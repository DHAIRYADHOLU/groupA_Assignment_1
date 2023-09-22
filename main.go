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

func multiplyAndDivide(a, b int) (int, float64) {
    product := a * b
    quotient := float64(a) / float64(b)
    return product, quotient
}

func IsDivisibleBy7(num int) bool {
	return num%7 == 0
}

func princeadd(no1, no2 int) int 
{
	return no1 + no2
}


func main() {
	
	num1 := 10
	num2 := 5

	min, max := minmax(num1, num2)

	fmt.Printf("Minimum: %d\n", min)
	fmt.Printf("Maximum: %d\n", max)
//////////////////////////////////////////////////////////////////
	v1, v2 := swap("kapoor", "vinay")
	fmt.Println(v1, v2)
//////////////////////////////////////////////////////////////////
	x := 10
    y := 3
    result1, result2 := multiplyAndDivide(x, y)
    fmt.Printf("Product: %d, Quotient: %.2f\n", result1, result2)
///////////////////////////////////////////////////////////////////
	var number int
	fmt.Print("Enter a number: ")
	fmt.Scan(&number)
		if IsDivisibleBy7(number) 
		{
			fmt.Printf("%d is divisible by 7\n", number)
		} 
		else 
		{
			fmt.Printf("%d is not divisible by 7\n", number)
		}
///////////////////////////////////////////////////////////////////
	result := princeadd(5, 3)
		fmt.Printf("The sum is: %d\n", result)
		
}
