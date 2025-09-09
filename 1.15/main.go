package main

import (
	"fmt"
	"strings"
)

/*
	строка justString записывает первые 100 символов
	строки v (формально будет ссылаться на адрес памяти строки v)
	то есть хранить всю большую строку в памяти
	gc не будет очисщать адрес v пока существует justString
	в то время как justString существует как глобальная переменная
*/

func someFunc(justString *string) {
	v := createHugeString(1 << 10)

	b := make([]byte, 100)
	// решением будет являться создания фиксированного количества байт и копирование в него
	copy(b, v)
	*justString = string(b)
}

func createHugeString(n int) string {
	return strings.Repeat("x", n)
}

func main() {
	var justString string
	someFunc(&justString)

	fmt.Println(justString)
}
