package main

import "fmt"

// search performs binary search on a sorted array to find the target value
// Returns the index of target if found, otherwise returns -1
// Time complexity: O(log n), Space complexity: O(1)
func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	
	for left <= right {
		// Use (left + right) / 2 but avoid potential overflow
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

// runTests executes test cases to verify the binary search implementation
func runTests() {
	fmt.Println("=== 二分查找算法测试 ===")
	
	// Test case 1: Example from problem
	nums1 := []int{-1, 0, 3, 5, 9, 12}
	target1 := 9
	result1 := search(nums1, target1)
	fmt.Printf("示例1: nums = %v, target = %d → 输出: %d", nums1, target1, result1)
	if result1 == 4 {
		fmt.Println(" ✅")
	} else {
		fmt.Println(" ❌")
	}
	
	// Test case 2: Example from problem
	nums2 := []int{-1, 0, 3, 5, 9, 12}
	target2 := 2
	result2 := search(nums2, target2)
	fmt.Printf("示例2: nums = %v, target = %d → 输出: %d", nums2, target2, result2)
	if result2 == -1 {
		fmt.Println(" ✅")
	} else {
		fmt.Println(" ❌")
	}
	
	fmt.Println("\n=== 边界测试 ===")
	
	// Test case 3: Single element - found
	nums3 := []int{5}
	target3 := 5
	result3 := search(nums3, target3)
	fmt.Printf("单元素命中: nums = %v, target = %d → 输出: %d", nums3, target3, result3)
	if result3 == 0 {
		fmt.Println(" ✅")
	} else {
		fmt.Println(" ❌")
	}
	
	// Test case 4: Single element - not found
	nums4 := []int{5}
	target4 := 3
	result4 := search(nums4, target4)
	fmt.Printf("单元素未命中: nums = %v, target = %d → 输出: %d", nums4, target4, result4)
	if result4 == -1 {
		fmt.Println(" ✅")
	} else {
		fmt.Println(" ❌")
	}
	
	// Test case 5: First element
	nums5 := []int{1, 2, 3, 4, 5}
	target5 := 1
	result5 := search(nums5, target5)
	fmt.Printf("首位查找: nums = %v, target = %d → 输出: %d", nums5, target5, result5)
	if result5 == 0 {
		fmt.Println(" ✅")
	} else {
		fmt.Println(" ❌")
	}
	
	// Test case 6: Last element
	nums6 := []int{1, 2, 3, 4, 5}
	target6 := 5
	result6 := search(nums6, target6)
	fmt.Printf("末位查找: nums = %v, target = %d → 输出: %d", nums6, target6, result6)
	if result6 == 4 {
		fmt.Println(" ✅")
	} else {
		fmt.Println(" ❌")
	}
	
	// Test case 7: Negative numbers
	nums7 := []int{-10, -5, -1, 0, 3, 7}
	target7 := -5
	result7 := search(nums7, target7)
	fmt.Printf("负数查找: nums = %v, target = %d → 输出: %d", nums7, target7, result7)
	if result7 == 1 {
		fmt.Println(" ✅")
	} else {
		fmt.Println(" ❌")
	}
}

func main() {
	runTests()
}