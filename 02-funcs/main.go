package main

import "fmt"

func main() {
	summation := AddTwoNumbers(2, 3)
	fmt.Printf("Sum of 2 + 3 = %d\n", summation)

	fib := Fibonacci(8)
	fmt.Printf("8th Fibonacci Number is %d\n", fib)
}

//AddTwoNumbers sums the given numbers
func AddTwoNumbers(num int, otherNum int) int {
	return num + otherNum
}

//Fibonacci - Caculates the Nth fibonacci number
func Fibonacci(n int) int {
	if n < 2 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}
