package main

import (
	"fmt"
	"log"
	"net/http"
)

type handler func(w http.ResponseWriter, r *http.Request)

func loggerDecorator(h handler) handler {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Begin loggerDecorator")
		defer func() {
			log.Println("End loggerDecorator")
		}()
		h(w, r)
	}
}

func root(w http.ResponseWriter, r *http.Request) {
	log.Println("Root")
	fmt.Fprint(w, "root")
}

var Root = loggerDecorator(root)

func main() {
	http.HandleFunc("/", Root)
	http.ListenAndServe(":8080", nil)
}
