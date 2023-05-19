package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	file := listFiles("testdata")

	for _, value := range file {
		fmt.Println(value)
	}
}

func listFiles(dirname string) []string {
	var dirs []string

	files, err := ioutil.ReadDir(dirname)
	if err != nil {
		log.Fatal(err)
	}
	var a = flag.Bool("a", false, "Show the files hidden")
	flag.Parse()
	for _, f := range files {
		if !(strings.HasPrefix(f.Name(), ".")) {
			dirs = append(dirs, f.Name())
		} else if *a {
			dirs = append(dirs, f.Name())
		}
	}

	return dirs
}
