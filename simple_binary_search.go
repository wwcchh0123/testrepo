package main

import "fmt"

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

func main() {
	fmt.Println("Testing Binary Search Algorithm:")
	
	// Test Case 1
	nums1 := []int{-1, 0, 3, 5, 9, 12}
	target1 := 9
	result1 := binarySearch(nums1, target1)
	fmt.Printf("Test 1: nums = %v, target = %d, result = %d (expected 4)\n", nums1, target1, result1)
	
	// Test Case 2
	nums2 := []int{-1, 0, 3, 5, 9, 13}
	target2 := 2
	result2 := binarySearch(nums2, target2)
	fmt.Printf("Test 2: nums = %v, target = %d, result = %d (expected -1)\n", nums2, target2, result2)
	
	if result1 == 4 && result2 == -1 {
		fmt.Println("All tests passed!")
	}
}
