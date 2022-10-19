package main

import "fmt"

// Recursive fibonacci function
func fib(n int) int {
	if n <= 0 {
		return 0
	} else if n == 1 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}

// adds two intergers and returns the sum
func add(x, y int) int {
	return x+y
}

func main() {
	name := "Scooby"
	z := add(1,2)
	fmt.Printf("Ave mundo! %s\n", name)
	fmt.Println(fib(20))
	fmt.Println("Ave Maria")
}
