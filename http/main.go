package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func handleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	io.WriteString(w, "Root!!!")
}

func handleLogin(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	io.WriteString(w, "Login!!!")
}

func main() {
	fmt.Println("Starting server")
	http.HandleFunc("/", handleRoot)
	http.HandleFunc("/login", handleLogin)
	s := &http.Server{
		Addr: ":8080",
	}
	err := s.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Stopping server")
}
