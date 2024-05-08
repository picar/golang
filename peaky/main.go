package main

import "peaky/apiserver"

func main() {

	s := apiserver.NewServer(apiserver.Server{Address: ":8080"})
	s.Start()
}
