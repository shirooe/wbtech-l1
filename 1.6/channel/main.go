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
	ch := make(chan struct{}, 1)

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		for {
			select {
			case <-ch:
				// уведомление о завершении работы
				wg.Done()
				fmt.Println("Программа получила уведомление о завершении работы")
				return
			default:
				// добавление к счетчику через атомарную операцию
				c := atomic.AddInt64(&i, 1)
				time.Sleep(time.Second * 1)
				fmt.Printf("Текущее значение: %d\n", c)
			}
		}
	}(wg)

	// имитация завершения программы
	time.Sleep(time.Second * 5)
	ch <- struct{}{}

	wg.Wait()
	fmt.Println("Программа завершила работу")
}
