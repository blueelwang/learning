package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Printf("00000000000000000000")
	<-time.After(0 * time.Second)
	fmt.Printf("11111111111111111")
	arr := []int{}
	fmt.Printf("%+v\n", IntSliceChunk(arr, 12))
}

func IntSliceChunk(data []int, size int) (res [][]int) {
	l := len(data)
	if l == 0 || size <= 0 {
		return
	}
	res = make([][]int, 0, l / size + 1)
	for i := 0; i < l; i += size {
		s, e := i, i + size
		if e > l {
			e = l
		}
		res = append(res, data[s:e])
	}
	return
}
