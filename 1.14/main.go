package main

import (
	"fmt"
	"reflect"
)

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
	default:
		if reflect.TypeOf(v).Kind() == reflect.Chan {
			return "chan"
		}
		return "unknown"
	}
}
