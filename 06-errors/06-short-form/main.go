package main

import (
	"fmt"
	"os"
)

func main() {
	if _, err := os.Stat("analysis.xlsx"); err != nil {
		fmt.Print(err)
	}
	fmt.Print("ok")

}
