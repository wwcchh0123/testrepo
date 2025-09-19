package main

import "fmt"

// search 在有序数组中搜索目标值，返回目标值的下标，如果不存在则返回-1
// nums: 有序（升序）整型数组
// target: 目标值
// 返回值: 目标值在数组中的下标，不存在则返回-1
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

// 主函数用于测试
func main() {
	testBinarySearch()
}

// 测试函数
func testBinarySearch() {
	// 示例1: nums = [-1,0,3,5,9,12], target = 9
	nums1 := []int{-1, 0, 3, 5, 9, 12}
	target1 := 9
	result1 := search(nums1, target1)
	fmt.Printf("示例1: nums = %v, target = %d\n", nums1, target1)
	fmt.Printf("输出: %d\n", result1)
	fmt.Printf("解释: %d 出现在 nums 中并且下标为 %d\n\n", target1, result1)
	
	// 示例2: nums = [-1,0,3,5,9,12], target = 2
	nums2 := []int{-1, 0, 3, 5, 9, 12}
	target2 := 2
	result2 := search(nums2, target2)
	fmt.Printf("示例2: nums = %v, target = %d\n", nums2, target2)
	fmt.Printf("输出: %d\n", result2)
	fmt.Printf("解释: %d 不存在 nums 中因此返回 %d\n\n", target2, result2)
	
	// 额外测试用例
	fmt.Println("额外测试用例:")
	
	// 测试边界情况
	nums3 := []int{5}
	fmt.Printf("单元素数组 [5]，查找 5: %d\n", search(nums3, 5))
	fmt.Printf("单元素数组 [5]，查找 1: %d\n", search(nums3, 1))
	
	// 测试更大的数组
	nums4 := []int{-9999, -100, 0, 1, 2, 3, 4, 5, 100, 9999}
	fmt.Printf("数组 %v，查找 0: %d\n", nums4, search(nums4, 0))
	fmt.Printf("数组 %v，查找 -9999: %d\n", nums4, search(nums4, -9999))
	fmt.Printf("数组 %v，查找 9999: %d\n", nums4, search(nums4, 9999))
	fmt.Printf("数组 %v，查找 50: %d\n", nums4, search(nums4, 50))
}