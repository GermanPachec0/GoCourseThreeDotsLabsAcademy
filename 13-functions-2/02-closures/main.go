package main

import "fmt"

func WordGenerator(words []string) func() string {
	index := -1
	return func() string {
		index += 1
		if index > len(words)-1 {
			index = 0
			return words[0]
		}
		return words[index]
	}
}

func main() {
	continents := []string{
		"Africa",
		"Antarctica",
		"Asia",
		"Australia",
		"Europe",
		"North America",
		"South America",
	}

	generator := WordGenerator(continents)

	for i := 0; i < 10; i++ {
		fmt.Println(generator())
	}
}
