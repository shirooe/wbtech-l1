package main

import "fmt"

func main() {
	s := []string{"cat", "cat", "dog", "cat", "tree"}

	fmt.Println(uniqueStrings(s))
}

func uniqueStrings(strings []string) []string {
	// словарь видели слово
	seen := make(map[string]struct{})
	result := make([]string, 0, len(strings))

	for _, str := range strings {
		// если слово еще не видели
		if _, exists := seen[str]; !exists {
			// запись пустой структуры для экономии памяти
			seen[str] = struct{}{}
			result = append(result, str)
		}
	}

	return result
}
