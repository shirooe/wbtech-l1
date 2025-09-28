package main

import "fmt"

type TestCase struct {
	array  []int
	target int
}

func main() {
	tests := []TestCase{
		{[]int{1, 2, 3, 4, 5}, 2}, // 1
		{[]int{1, 2, 3, 4, 5}, 5}, // 4
		{[]int{1, 2, 3, 4, 5}, 0}, // -1
	}

	for _, test := range tests {
		result := BinarySearch(test.array, test.target)

		if result == -1 {
			fmt.Println("Элемент не найден")
		} else {
			fmt.Printf("Элемент под индексом %d\n", result)
		}
	}
}

func BinarySearch(arr []int, target int) int {
	low, high := 0, len(arr)-1

	for low <= high {
		mid := (low + high) / 2

		switch {
		case arr[mid] == target:
			return mid
		case arr[mid] > target:
			high = mid - 1
		default:
			low = mid + 1
		}
	}

	return -1
}
