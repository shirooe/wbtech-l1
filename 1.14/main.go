package main

import "fmt"

func main() {
	test := []any{
		1,
		"string",
		true,
		make(chan string),
		make(chan int),
		make(chan bool),
	}

	for _, v := range test {
		fmt.Printf("Тип значения: %s\n", determine(v))
	}
}

func determine(v any) string {
	switch v.(type) {
	case int:
		return "int"
	case string:
		return "string"
	case bool:
		return "bool"
	case chan string:
		return "chan string"
	case chan int:
		return "chan int"
	case chan bool:
		return "chan bool"
	default:
		return "unknown"
	}
}
