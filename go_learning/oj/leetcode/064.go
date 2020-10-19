package leetcode

import (
	"fmt"
	"math"
)

func MinPathSum() {
	fmt.Println(minPathSum([][]int{
		{1,3,1},
		{1,5,1},
		{4,2,1},
	}))
}

func minPathSum(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	m, n := len(grid), len(grid[0])

	sum := make([]int, n)
	for i := 0; i < n; i++ {
		sum[i] = -1
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if j == 0 {
				if (sum[j] < 0) {
					sum[j] = grid[i][j]
				} else {
					sum[j] += grid[i][j]
				}
			} else {
				if (sum[j] < 0) {
					sum[j] = grid[i][j] + sum[j - 1]
				} else {
					sum[j] = grid[i][j] + int(math.Min(float64(sum[j]), float64(sum[j - 1])));
				}
			}
		}
	}
	return sum[n - 1]

}

