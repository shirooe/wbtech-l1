package main

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	var i int64
	wg := &sync.WaitGroup{}

	// создание контекста с отменой
	ctx, cancel := context.WithCancel(context.Background())

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		for {
			select {
			case <-ctx.Done():
				// сигнал о завершении контекста
				wg.Done()
				fmt.Println("Получили сигнал о завершении контекста")
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
	cancel()

	wg.Wait()
	fmt.Println("Программа завершила работу")
}
