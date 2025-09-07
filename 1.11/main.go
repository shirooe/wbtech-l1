package main

import "fmt"

func main() {
	a := []int{1, 2, 3}
	b := []int{2, 3, 4}

	fmt.Println(intersect(a, b)) // [2 3]

	c := []int{1, 2, 2}
	d := []int{4, 5, 2, 1}

	fmt.Println(intersect(c, d)) // [2 1]
}

func intersect(a, b []int) []int {
	result := make([]int, 0)

	m := make(map[int]bool)

	// запись значений в мапу
	for _, i := range a {
		m[i] = true
	}

	for _, i := range b {
		// если значение есть в мапе (пересечение без дубликатов)
		if m[i] {
			result = append(result, i)
		}
	}

	return result
}
