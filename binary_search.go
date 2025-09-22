package main

import "fmt"

// search 在有序数组中搜索目标值，返回其索引，如果不存在则返回-1
func search(nums []int, target int) int {
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

// 测试函数
func testBinarySearch() {
	fmt.Println("=== 二分查找算法测试 ===")
	
	// 示例1测试
	nums1 := []int{-1, 0, 3, 5, 9, 12}
	target1 := 9
	result1 := search(nums1, target1)
	fmt.Printf("示例1: nums = %v, target = %d → 输出: %d\n", nums1, target1, result1)
	
	// 示例2测试
	nums2 := []int{-1, 0, 3, 5, 9, 12}
	target2 := 2
	result2 := search(nums2, target2)
	fmt.Printf("示例2: nums = %v, target = %d → 输出: %d\n", nums2, target2, result2)
	
	// 边界测试
	fmt.Println("\n=== 边界测试 ===")
	
	// 测试单个元素
	nums3 := []int{5}
	target3 := 5
	result3 := search(nums3, target3)
	fmt.Printf("单元素命中: nums = %v, target = %d → 输出: %d\n", nums3, target3, result3)
	
	// 测试单个元素不匹配
	nums4 := []int{5}
	target4 := 3
	result4 := search(nums4, target4)
	fmt.Printf("单元素未命中: nums = %v, target = %d → 输出: %d\n", nums4, target4, result4)
	
	// 测试目标值在数组开头
	nums5 := []int{1, 2, 3, 4, 5}
	target5 := 1
	result5 := search(nums5, target5)
	fmt.Printf("首位查找: nums = %v, target = %d → 输出: %d\n", nums5, target5, result5)
	
	// 测试目标值在数组末尾
	nums6 := []int{1, 2, 3, 4, 5}
	target6 := 5
	result6 := search(nums6, target6)
	fmt.Printf("末位查找: nums = %v, target = %d → 输出: %d\n", nums6, target6, result6)
	
	// 测试负数
	nums7 := []int{-10, -5, -1, 0, 3, 7}
	target7 := -5
	result7 := search(nums7, target7)
	fmt.Printf("负数查找: nums = %v, target = %d → 输出: %d\n", nums7, target7, result7)
}

func main() {
	testBinarySearch()
}