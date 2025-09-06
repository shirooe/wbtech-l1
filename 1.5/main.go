package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var n int
	fmt.Printf("Введите количество секунд выполнения программы: ")
	fmt.Scan(&n)

	now := time.Now()
	t := time.After(time.Duration(n) * time.Second)

	data := make(chan int)
	wg := sync.WaitGroup{}

	wg.Add(2)
	go writer(t, data, &wg)
	go reader(data, &wg)

	wg.Wait()
	fmt.Println("Программа завершила работу")
	fmt.Printf("Время выполнения программы: %.2f\n", time.Since(now).Seconds())
}

func writer(t <-chan time.Time, data chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	defer close(data)

	var count int

	for {
		select {
		case <-t:
			return
		case data <- count:
			count++
		}
	}
}

func reader(data <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for n := range data {
		fmt.Printf("Значение: %d\n", n)
	}
}
