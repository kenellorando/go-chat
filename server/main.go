package main

import (
	"fmt"
	"net"
	"bufio"
	"log"
)

const (
	LADDR = "localhost:25624"
)

func main() {
	log.Println("Starting server at " + LADDR)

	// Start listening
	ln, err := net.Listen("tcp", LADDR)
	if err != nil {
		log.Print("Failed to start server.")
	} else {
		log.Print("Server started. Awaiting connections...")
	}

	// accept connection on port
	conn, _ := ln.Accept()


	for {
		// will listen for message to process ending in newline (\n)
		message, _ := bufio.NewReader(conn).ReadString('\n')
		// output message received
		fmt.Print("Message Received:", string(message))


		// send new string back to client
		conn.Write([]byte("Message received at server!" + "\n"))
	}
}