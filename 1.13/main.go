package main

import "fmt"

func main() {
	a, b := 0, 1
	fmt.Printf("Значение до замены: %d %d\n", a, b)
	a, b = Swap(a, b)
	fmt.Printf("Значение после замены: %d %d\n", a, b)

	c, d := 1, 2
	fmt.Printf("Значение до замены: %d %d\n", c, d)
	c, d = Swap(c, d)
	fmt.Printf("Значение после замены: %d %d\n", c, d)

	e, f := 2, 3
	fmt.Printf("Значение до замены: %d %d\n", e, f)
	e, f = Swap(e, f)
	fmt.Printf("Значение после замены: %d %d\n", e, f)
}

func Swap(a, b int) (int, int) {
	a = a ^ b
	b = a ^ b
	a = a ^ b
	return a, b
}
