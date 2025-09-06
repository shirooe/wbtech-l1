package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()

		defer func() {
			// обработка паники
			if r := recover(); r != nil {
				fmt.Println("Обработана паника")
			}
		}()
		// вызов panic
		fmt.Println("Вызов паники")
		panic("Паника")
	}(wg)

	wg.Wait()
	fmt.Println("Программа завершила работу")
}
