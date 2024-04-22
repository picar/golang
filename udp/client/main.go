package main

import (
	"fmt"
	"net"
)

func main() {
	addr, err := net.ResolveUDPAddr("udp", ":8080")
	if err != nil {
		fmt.Println("Error resolving UDP address:", err)
		return
	}
	fmt.Println("UDP Addr: ", addr)

	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		fmt.Println("Error connecting: ", err)
		return
	}
	defer conn.Close()
	fmt.Println("Connection: ", conn)

	message := []byte("Hello World!")

	_, err = conn.Write(message)
	if err != nil {
		fmt.Println("Error sending data:", err)
		return
	}

	fmt.Println("Message sent successfully:", string(message))
}
