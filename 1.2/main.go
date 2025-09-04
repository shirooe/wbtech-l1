package main

import (
	"fmt"
	"sync"
)

func main() {
	example := []int{2, 4, 6, 8, 10}

	wg := sync.WaitGroup{}
	wg.Add(len(example))

	for i, num := range example {
		// Возведение в квадрат числа и запись в массив по индексу
		// для упорядоченного вывода
		go func(i, n int) {
			example[i] = n * n
			wg.Done()
		}(i, num)
	}
	wg.Wait()

	// Вывод
	fmt.Println(example)
}
