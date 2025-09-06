package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	var i int64
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		for {
			// добавление к счетчику через атомарную операцию
			c := atomic.AddInt64(&i, 1)
			fmt.Printf("Текущее значение: %d\n", c)
			time.Sleep(time.Second * 1)

			// условие для выхода
			if c == 10 {
				fmt.Printf("Значение достигло условия: %d\n", c)
				wg.Done()
				break
			}
		}
	}(wg)

	wg.Wait()
	fmt.Println("Программа завершила работу")
}
