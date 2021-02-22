package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	f, err := os.OpenFile("./data.tmp", os.O_RDWR | os.O_CREATE | os.O_EXCL, 0664)
	time.Sleep(10 * time.Second)
	fmt.Printf("file: %+v, err: %s", f, err)
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
