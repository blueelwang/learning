package test

import (
	"math"
)

// CulCost 测试数学运算性能
func CulCost() {
	for i := 0; i < 100000000; i++ {
		fi := float64(i)
		t := math.Cos(fi)
		t = math.Sin(fi)
		_ = t
	}
}
