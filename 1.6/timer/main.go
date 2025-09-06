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
	// создание таймера
	t := time.NewTimer(time.Second * 5)

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		for {
			select {
			// сигнал о завершении таймера
			case <-t.C:
				wg.Done()
				fmt.Println("Истекло время таймера")
				return
			default:
				// добавление к счетчику через атомарную операцию
				c := atomic.AddInt64(&i, 1)
				time.Sleep(time.Second * 1)
				fmt.Printf("Текущее значение: %d\n", c)
			}
		}
	}(wg)

	wg.Wait()
	fmt.Println("Программа завершила работу")
}
