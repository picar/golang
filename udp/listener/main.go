package main

import (
	"fmt"
	"net"
	"os"
)

func main() {

	addr, err := net.ResolveUDPAddr("udp", ":8080")

	if err != nil {
		fmt.Println("Error resolving UDP address:", err)
		return
	}

	fmt.Println("UDP Addr: ", addr)

	conn, err := net.ListenUDP("udp", addr)

	if err != nil {
		fmt.Println("Error listening: ", err)
		return
	}
	defer conn.Close()
	fmt.Println("Connection: ", conn)

	buff := make([]byte, 1024)

	for {
		size, addr, err := conn.ReadFromUDP(buff)
		if err != nil {
			fmt.Println("Error reading: ", err)
			continue
		}
		fmt.Fprintln(os.Stdout, []any{"Received %d bytes from %s", size, addr}...)
		fmt.Println(string(buff[:size]))
	}
}
