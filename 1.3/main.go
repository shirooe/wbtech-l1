package main

import (
	"fmt"
	"math/rand/v2"
	"os"
	"strconv"
	"sync"
	"time"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Введите количество воркеров.")
		return
	}

	count, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Введите число.")
		return
	}

	// Создаем канал
	jobs := make(chan int, count)
	wg := sync.WaitGroup{}

	for c := range count {
		// Создаем воркеров
		wg.Add(1)
		go worker(c, jobs, &wg)
	}

	// Заполняем канал
	go loop(jobs)

	// Ожидаем завершения воркеров
	wg.Wait()
	close(jobs)
}

func loop(jobs chan<- int) {
	for {
		// Заполняем канал рандомными числами
		jobs <- rand.IntN(100)
		// Задержка
		time.Sleep(time.Second)
	}
}

func worker(id int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	// Проходим по каналу
	for j := range jobs {
		fmt.Printf("Worker %d: %d\n", id, j)
	}
}
