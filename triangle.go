package main

import (
	"errors"
	"fmt"
)

func CalculateTriangleArea(base, height float64) (float64, error) {
	if base <= 0 {
		return 0, errors.New("base must be greater than 0")
	}
	if height <= 0 {
		return 0, errors.New("height must be greater than 0")
	}
	
	area := 0.5 * base * height
	return area, nil
}

func main() {
	base := 10.0
	height := 5.0
	
	area, err := CalculateTriangleArea(base, height)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	
	fmt.Printf("Triangle area with base %.2f and height %.2f is: %.2f\n", base, height, area)
}
