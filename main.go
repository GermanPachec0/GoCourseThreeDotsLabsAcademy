package main

import "fmt"

/*
	go build main.go -> archivo ejecutable
	go run main.go
*/

func saludar(username string) {
	fmt.Println("Hola", username)
}

func main() {
	saludar("GERMAN")
}
