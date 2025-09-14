package main

import "fmt"

func main() {
	tests := [][]int{
		{2, 1, 5, 3, 4},
		{5, 1, 6, 8, 2},
		{9, 3, 5, 7, 1},
	}

	for _, test := range tests {
		fmt.Printf("Исходный массив: \t%v\n", test)
		fmt.Printf("Отсортированный массив: %v\n", quickSort(test))
	}
}

func quickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	pivot := arr[0]
	left := make([]int, 0, len(arr)/2)
	right := make([]int, 0, len(arr)/2)

	for _, v := range arr[1:] {
		if v < pivot {
			left = append(left, v)
		} else {
			right = append(right, v)
		}
	}

	return append(append(quickSort(left), pivot), quickSort(right)...)
}
