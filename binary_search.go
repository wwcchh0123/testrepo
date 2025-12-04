package main

// BinarySearch 在有序数组中查找目标值，返回其索引
// 如果目标值不存在，返回 -1
//
// 参数:
//   nums: 升序排列的整型数组（元素不重复）
//   target: 要查找的目标值
//
// 返回值:
//   目标值在数组中的索引，如果不存在则返回 -1
//
// 时间复杂度: O(log n)
// 空间复杂度: O(1)
func BinarySearch(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for left <= right {
		// 使用 left + (right-left)/2 防止整数溢出
		mid := left + (right-left)/2

		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			// 目标在右半部分
			left = mid + 1
		} else {
			// 目标在左半部分
			right = mid - 1
		}
	}

	// 未找到目标值
	return -1
}
