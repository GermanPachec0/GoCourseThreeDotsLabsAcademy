package main

import "fmt"

var products = map[int]string{
	1: "Book",
	2: "Video Course",
	3: "Lecture",
	4: "Talk",
	5: "Training",
}

func main() {
	ids := Keys(products)
	names := Values(products)

	fmt.Println("Prouct IDs:", ids)
	fmt.Println("Product names:", names)
}

func Keys(prods map[int]string) []int {
	var myKeys []int

	for key := range prods {
		myKeys = append(myKeys, key)
	}
	return myKeys
}

func Values(prods map[int]string) []string {

	var myValues []string

	for _, product := range prods {
		myValues = append(myValues, product)
	}

	return myValues
}
