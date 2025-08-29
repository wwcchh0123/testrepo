package main

import (
	"fmt"
	"os"
	"strconv"
)

// This file provides a standalone demo of the Fibonacci functionality
// Run with: go run fibonacci.go fibonacci_demo.go

func runFibonacciDemo() {
	fmt.Println("=== Fibonacci Sequence Demo ===")
	
	// Demo 1: Print first 10 terms
	fmt.Println("\nDemo 1: First 10 terms")
	Fibonacci(10)
	
	// Demo 2: Print first 15 terms
	fmt.Println("\nDemo 2: First 15 terms")
	Fibonacci(15)
	
	// Demo 3: Print numbers up to 100
	fmt.Println("\nDemo 3: Numbers up to 100")
	FibonacciUpTo(100)
	
	// Demo 4: Handle command line argument if provided
	if len(os.Args) > 1 {
		if n, err := strconv.Atoi(os.Args[1]); err == nil {
			fmt.Printf("\nDemo 4: First %d terms (from command line)\n", n)
			Fibonacci(n)
		}
	}
}

func main() {
	// Check if this is being run as fibonacci demo
	if len(os.Args) > 0 && len(os.Args[0]) > 0 {
		runFibonacciDemo()
	}
}