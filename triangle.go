package main

import (
	"errors"
	"fmt"
)

// CalculateTriangleArea 计算三角形面积,使用公式: 面积 = 1/2 × 底边 × 高
//
// 参数:
//   base - 三角形的底边长度,必须大于 0
//   height - 三角形的高度,必须大于 0
//
// 返回值:
//   float64 - 计算得到的面积
//   error - 如果输入无效则返回错误
func CalculateTriangleArea(base, height float64) (float64, error) {
	if base <= 0 {
		return 0, errors.New("base must be greater than 0")
	}
	if height <= 0 {
		return 0, errors.New("height must be greater than 0")
	}

	// 检查数值边界以防止溢出
	const maxDimension = 1e100 // 合理的几何上限
	if base > maxDimension || height > maxDimension {
		return 0, errors.New("base and height values are too large")
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
