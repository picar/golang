package main

import "log"

type loggable func() error

func loggerDecorator(w loggable) loggable {
	return func() error {
		log.Println("Begin loggerDecorator")
		defer func() {
			log.Println("End loggerDecorator")
		}()
		return w()
	}
}

func wrapped() error {
	log.Println("Inside wrapped")
	return nil
}

var Wrapped = loggerDecorator(wrapped)

func main() {
	Wrapped()
	Wrapped()
	Wrapped()
}
