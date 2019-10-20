package main

import (
	"fmt"
)

func quickSort(array []int) {
	if len(array) <= 1 {
		return
	}
	pivot := array[0]
	last_small := 0
	for i := last_small + 1; i < len(array); i++ {
		if array[i] < pivot {
			last_small += 1
			array[last_small], array[i] = array[i], array[last_small]
		}
	}
	array[last_small], array[0] = array[0], array[last_small]
	quickSort(array[:last_small])
	quickSort(array[last_small+1:])
}

func main() {
	example := []int{10, 6, 8, 1, 4, 7, 2, 9, 0, 7, 4, 2, 6}
	fmt.Println("Before sort: ", example)
	quickSort(example)
	fmt.Println("After sort: ", example)
}
