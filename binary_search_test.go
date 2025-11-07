package main

import "testing"

func TestBinarySearch(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   int
	}{
		{
			name:   "示例1 - 找到目标值9",
			nums:   []int{-1, 0, 3, 5, 9, 12},
			target: 9,
			want:   4,
		},
		{
			name:   "示例2 - 目标值2不存在",
			nums:   []int{-1, 0, 3, 5, 9, 12},
			target: 2,
			want:   -1,
		},
		{
			name:   "单个元素 - 找到",
			nums:   []int{5},
			target: 5,
			want:   0,
		},
		{
			name:   "单个元素 - 未找到",
			nums:   []int{5},
			target: 3,
			want:   -1,
		},
		{
			name:   "目标值在开头",
			nums:   []int{1, 2, 3, 4, 5},
			target: 1,
			want:   0,
		},
		{
			name:   "目标值在末尾",
			nums:   []int{1, 2, 3, 4, 5},
			target: 5,
			want:   4,
		},
		{
			name:   "目标值在中间",
			nums:   []int{1, 2, 3, 4, 5},
			target: 3,
			want:   2,
		},
		{
			name:   "目标值小于最小值",
			nums:   []int{1, 2, 3, 4, 5},
			target: 0,
			want:   -1,
		},
		{
			name:   "目标值大于最大值",
			nums:   []int{1, 2, 3, 4, 5},
			target: 10,
			want:   -1,
		},
		{
			name:   "负数数组",
			nums:   []int{-9999, -5000, -1000, -100, -10},
			target: -1000,
			want:   2,
		},
		{
			name:   "包含负数和正数",
			nums:   []int{-50, -20, 0, 20, 50, 100},
			target: 0,
			want:   2,
		},
		{
			name:   "大数组",
			nums:   []int{-9999, -8000, -6000, -4000, -2000, 0, 2000, 4000, 6000, 8000, 9999},
			target: 6000,
			want:   8,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := BinarySearch(tt.nums, tt.target)
			if got != tt.want {
				t.Errorf("BinarySearch() = %v, want %v", got, tt.want)
			}
		})
	}
}

// BenchmarkBinarySearch 测试二分查找的性能
func BenchmarkBinarySearch(b *testing.B) {
	// 创建一个大数组
	nums := make([]int, 10000)
	for i := 0; i < 10000; i++ {
		nums[i] = i * 2
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		BinarySearch(nums, 5000)
	}
}
