package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		fmt.Println("Выполнение горутины")

		// выход из горутины
		runtime.Goexit()
	}(wg)

	wg.Wait()
	fmt.Println("Программа завершила работу")
}
