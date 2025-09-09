package main

import "fmt"

type TestCase struct {
	a, b int
}

func main() {
	test := []TestCase{
		{1, 2},
		{2, 3},
		{3, 4},
	}

	for _, t := range test {
		fmt.Printf("Значения до изменений: \na: %d, b: %d\n", t.a, t.b)
		a, b := Swap(t.a, t.b)
		fmt.Printf("Значения после изменений: \na: %d, b: %d\n", a, b)
	}
}

func Swap(a, b int) (int, int) {
	a = a ^ b
	b = a ^ b
	a = a ^ b
	return a, b
}
