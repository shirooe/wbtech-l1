package main

import (
	"fmt"
	"sync"
)

type ConcurrentMap struct {
	mu *sync.RWMutex
	m  map[string]any
}

func main() {
	wg := &sync.WaitGroup{}
	m := NewConcurrentMap()

	// заполнение мапы
	for i := range 10 {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			m.Set(fmt.Sprintf("ключ-%d", i), i)
		}(i)
	}

	wg.Wait()

	// чтение из мапы
	for i := range 10 {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			v, ok := m.Get(fmt.Sprintf("ключ-%d", i))
			if !ok {
				fmt.Printf("ключ-%d не найден\n", i)
				return
			}

			fmt.Printf("ключ-%d: значение - %v\n", i, v)
		}(i)
	}

	wg.Wait()
}

// конструктор
func NewConcurrentMap() *ConcurrentMap {
	return &ConcurrentMap{
		mu: &sync.RWMutex{},
		m:  make(map[string]any),
	}
}

// безопасное запись в мапу
func (c *ConcurrentMap) Set(key string, value any) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.m[key] = value
}

// безопасное чтение из мапы
func (c *ConcurrentMap) Get(key string) (any, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	value, ok := c.m[key]
	// возвращаем значение и флаг
	return value, ok
}
