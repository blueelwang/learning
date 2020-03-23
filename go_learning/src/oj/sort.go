package oj

import (
	"fmt"
)

func DemoSort() {
	array := []int{1, 3, 36, 39, 2, 5, 82, 94, 4, 8, 68, 6, 83, 46, 92, 77, 30, 57, 21, 0, 73}
	fmt.Println(array)
	headSort(array)
	fmt.Println(array)
}

func headSort(array []int) {
	len := len(array)
	for i := len / 2 - 1; i >= 0; i-- {
		rebuild(array, len - 1, i)
	}
	for i := len - 1; i > 0; i-- {
		array[i], array[0] = array[0], array[i]
		rebuild(array, i - 1, 0)
	}
}

func rebuild(array []int, boundary int, root int) {
	for root < boundary {
		leftChild := (root + 1) * 2 - 1
		rightChild := leftChild + 1
		if rightChild <= boundary && array[rightChild] > array[root] && array[rightChild] > array[leftChild] {
			array[rightChild], array[root] = array[root], array[rightChild]
			root = rightChild
		} else if leftChild <= boundary && array[leftChild] > array[root] {
			array[leftChild], array[root] = array[root], array[leftChild]
			root = leftChild
		} else {
			break;
		}
	}
}
