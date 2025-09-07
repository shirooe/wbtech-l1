package main

import (
	"fmt"
)

func main() {
	// пример
	x := []int{1, 2, 3, 4, 5}
	writer := make(chan int, len(x))
	processsed := make(chan int, len(x))
 
	go worker(writer, x)
	go worker2(writer, processsed)

	// блокируемся на чтение из канала processsed
	for v := range processsed {
		fmt.Printf("Чтение из канала processsed: %d\n", v)
	}

	fmt.Println("Завершение чтения из канала processsed")
}

func worker(writer chan<- int, x []int) {
	defer close(writer)
	// запись в канал
	for _, v := range x {
		writer <- v
	}
}

func worker2(writer <-chan int, processsed chan<- int) {
	defer close(processsed)
	// чтение и преобразование
	for v := range writer {
		processsed <- v * 2
	}
}
