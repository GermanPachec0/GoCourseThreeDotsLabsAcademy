package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"main/server"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httprate"
)

func main() {
	s := server.Server{
		TweetsRepository: &server.TweetsMemoryRepository{},
	}
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/tweets", s.ListTweets)
	r.With(httprate.LimitByIP(10, 1*time.Minute)).Post("/tweets", s.AddTweet)
	go spamTweets()
	log.Fatal(http.ListenAndServe(":8080", r))
}

func spamTweets() {
	url := "http://localhost:8080/tweets"
	for {
		time.Sleep(time.Millisecond * 100)
		addTweetPayload := server.Tweet{
			Message:  "ass",
			Location: "ass",
		}
		marshaledPayload, err := json.Marshal(addTweetPayload)
		if err != nil {
			panic(err)
		}

		resp, err := http.Post(url, "application/json", bytes.NewBuffer(marshaledPayload))
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}

		respBody := server.Response{}

		if err := json.Unmarshal(body, &respBody); err != nil {
			log.Println(err)
		} else {
			fmt.Sprintf("Added tweet %d", respBody.ID)

		}
	}
}
