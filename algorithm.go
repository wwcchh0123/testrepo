package main

import "fmt"

// binarySearch finds the index of target in a sorted array nums
// Returns the index if target exists, otherwise returns -1
func binarySearch(nums []int, target int) int {
	left, right := 0, len(nums)-1
	
	for left <= right {
		mid := left + (right-left)/2
		
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	
	return -1
}

// testBinarySearch runs the test cases from the problem
func testBinarySearch() {
	fmt.Println("Testing Binary Search Algorithm:")
	
	// Test Case 1: nums = [-1,0,3,5,9,12], target = 9
	nums1 := []int{-1, 0, 3, 5, 9, 12}
	target1 := 9
	result1 := binarySearch(nums1, target1)
	fmt.Printf("Test 1: nums = %v, target = %d\n", nums1, target1)
	fmt.Printf("Expected: 4, Got: %d\n", result1)
	fmt.Printf("Result: %s\n\n", getTestResult(result1, 4))
	
	// Test Case 2: nums = [-1,0,3,5,9,12], target = 2
	nums2 := []int{-1, 0, 3, 5, 9, 12}
	target2 := 2
	result2 := binarySearch(nums2, target2)
	fmt.Printf("Test 2: nums = %v, target = %d\n", nums2, target2)
	fmt.Printf("Expected: -1, Got: %d\n", result2)
	fmt.Printf("Result: %s\n\n", getTestResult(result2, -1))
}

func getTestResult(got, expected int) string {
	if got == expected {
		return "PASS ✓"
	}
	return "FAIL ✗"
}