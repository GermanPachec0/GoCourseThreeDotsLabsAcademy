package main

import "fmt"

func main() {
	x := "World!"
	y := "Hello,"

	x, y = Swap(x, y)

	fmt.Println(x, y)
}

func Swap(first, second string) (string, string) {
	return second, first
}
