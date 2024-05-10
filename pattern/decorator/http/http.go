package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type handleFunc func(w http.ResponseWriter, r *http.Request)

func withLog(h handleFunc) handleFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Begin loggerDecorator %v\n", r.URL.Path)
		h(w, r)
		log.Println("End loggerDecorator")
	}
}

func root(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Root")
}

func randomTime(w http.ResponseWriter, r *http.Request) {
	time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
	fmt.Fprint(w, "randomtime")
}

func main() {
	http.HandleFunc("/", withLog(root))
	http.HandleFunc("/randomtime", withLog(randomTime))
	http.ListenAndServe(":8080", nil)
}
