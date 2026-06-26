package main

import (
	"go-tips/interfaceSize"
)

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}
	for i := 0; i < 10000; i++ {
		slice = append(slice, i)
	}

	interfaceSize.MeasureSize(slice)
}
