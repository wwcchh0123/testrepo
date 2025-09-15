package main

import "fmt"

// SearchTarget finds the index of target in the sorted array nums.
// Returns the index if found, otherwise returns -1.
func SearchTargetStandalone(nums []int, target int) int {
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
	// Test case 1: target exists
	nums1 := []int{-1, 0, 3, 5, 9, 12}
	target1 := 9
	result1 := SearchTargetStandalone(nums1, target1)
	fmt.Printf("Test 1 - nums: %v, target: %d, result: %d, expected: 4\n", nums1, target1, result1)
	
	// Test case 2: target does not exist
	nums2 := []int{-1, 0, 3, 5, 9, 12}
	target2 := 2
	result2 := SearchTargetStandalone(nums2, target2)
	fmt.Printf("Test 2 - nums: %v, target: %d, result: %d, expected: -1\n", nums2, target2, result2)
	
	// Additional test cases
	nums3 := []int{5}
	target3 := 5
	result3 := SearchTargetStandalone(nums3, target3)
	fmt.Printf("Test 3 - nums: %v, target: %d, result: %d, expected: 0\n", nums3, target3, result3)
	
	nums4 := []int{}
	target4 := 1
	result4 := SearchTargetStandalone(nums4, target4)
	fmt.Printf("Test 4 - nums: %v, target: %d, result: %d, expected: -1\n", nums4, target4, result4)
}