package main

import "fmt"

func main() {
	var n int64 = 5
	setBIT(&n, 1, 0)
	fmt.Println(n) // 4

	setBIT(&n, 2, 1)
	fmt.Println(n) // 6

	setBIT(&n, 1, 1)
	fmt.Println(n) // 7
}

func setBIT(n *int64, i int64, bit int64) {
	if i < 1 || i > 64 {
		fmt.Println("Позиция бита должна быть в диапазоне от 1 до 64")
		return
	}

	if bit < 0 || bit > 1 {
		fmt.Println("Значение бита должно быть 0 или 1")
		return
	}

	// сдвигаемся на нужную позицию
	i--

	switch bit {
	case 0:
		// очистка бита
		*n &= ^(int64(1) << i)
	case 1:
		// установка бита
		*n |= int64(1) << i
	}
}
