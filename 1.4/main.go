package main

import (
	"context"
	"fmt"
	"math/rand/v2"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
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

	// Создаем контекст
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Создаем сигнал
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	// Отслеживаем сигнал
	go func() {
		<-sig
		cancel() // Отменяем контекст
	}()

	for c := range count {
		// Создаем воркеров
		wg.Add(1)
		go worker(ctx, c, jobs, &wg)
	}

	// Заполняем канал
	wg.Add(1)
	go func() {
		defer wg.Done()
		loop(ctx, jobs)
		// Закрываем канал при завершении
		close(jobs)
	}()

	// Ожидаем завершения воркеров
	wg.Wait()
}

func loop(ctx context.Context, jobs chan<- int) {
	for {
		select {
		case jobs <- rand.IntN(100):
			time.Sleep(time.Second)
		case <-ctx.Done():
			return
		}
	}
}

func worker(ctx context.Context, id int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	// Проходим по каналу
	for {
		select {
		case j, ok := <-jobs:
			// Проверяем, что канал не закрыт
			if !ok {
				return
			}
			fmt.Printf("Воркер %d: %d\n", id, j)
		case <-ctx.Done():
			fmt.Printf("Воркер %d завершил работу\n", id)
			return
		}
	}
}
