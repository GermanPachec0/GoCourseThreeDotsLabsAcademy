package main

import (
	"fmt"
	"log"
	"net/http"
)

var calls = []string{}
var stats = map[string]int{}

func main() {
	http.HandleFunc("/hello", hello)
	log.Fatal(http.ListenAndServe(":8080", nil))

}

func hello(w http.ResponseWriter, r *http.Request) {
	if r.URL.Query().Get("name") == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	calls = append(calls, r.URL.Query().Get("name"))
	stats[r.URL.Query().Get("name")] += 1

	fmt.Fprint(w, "Hello, ", r.URL.Query().Get("name"))
	fmt.Printf("calls: %#v\n", calls)
	fmt.Printf("stats: %#v\n\n", stats)
}
