package main

import "testing"

func TestSearchTarget(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		target   int
		expected int
	}{
		{
			name:     "Example 1: target exists",
			nums:     []int{-1, 0, 3, 5, 9, 12},
			target:   9,
			expected: 4,
		},
		{
			name:     "Example 2: target does not exist",
			nums:     []int{-1, 0, 3, 5, 9, 12},
			target:   2,
			expected: -1,
		},
		{
			name:     "Single element - found",
			nums:     []int{5},
			target:   5,
			expected: 0,
		},
		{
			name:     "Single element - not found",
			nums:     []int{5},
			target:   -5,
			expected: -1,
		},
		{
			name:     "Empty array",
			nums:     []int{},
			target:   1,
			expected: -1,
		},
		{
			name:     "Target at beginning",
			nums:     []int{1, 2, 3, 4, 5},
			target:   1,
			expected: 0,
		},
		{
			name:     "Target at end",
			nums:     []int{1, 2, 3, 4, 5},
			target:   5,
			expected: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := SearchTarget(tt.nums, tt.target)
			if result != tt.expected {
				t.Errorf("SearchTarget(%v, %d) = %d; want %d", tt.nums, tt.target, result, tt.expected)
			}
		})
	}
}