package main

import "fmt"

// Fibonacci generates and prints the Fibonacci sequence up to n terms
func Fibonacci(n int) {
	if n <= 0 {
		fmt.Println("Please provide a positive number")
		return
	}

	fmt.Printf("Fibonacci sequence (%d terms):\n", n)

	if n >= 1 {
		fmt.Print("0")
	}
	if n >= 2 {
		fmt.Print(", 1")
	}

	a, b := 0, 1
	for i := 3; i <= n; i++ {
		next := a + b
		fmt.Printf(", %d", next)
		a, b = b, next
	}
	fmt.Println()
}

// FibonacciUpTo generates and prints Fibonacci numbers up to a maximum value
func FibonacciUpTo(max int) {
	if max < 0 {
		fmt.Println("Please provide a non-negative number")
		return
	}

	fmt.Printf("Fibonacci sequence up to %d:\n", max)

	a, b := 0, 1
	if max >= a {
		fmt.Print(a)
	}
	if max >= b && a != b {
		fmt.Print(", ", b)
	}

	for {
		next := a + b
		if next > max {
			break
		}
		fmt.Printf(", %d", next)
		a, b = b, next
	}
	fmt.Println()
}