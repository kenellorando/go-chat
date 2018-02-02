package main

import (
	"fmt"
	"net"
	"bufio"
	"log"
	"os"
)

const (
	LADDR = "localhost:25624"
)

func main() {
	go server()


	log.Print("Attempting to connect to " + LADDR)
	// connect to this socket
	conn, err := net.Dial("tcp", LADDR)
	if err != nil {
		log.Print("Failed to connect.")
	} else {
		log.Print("Successfully connected.")
	}


	for { 
	  // read in input from stdin
	  reader := bufio.NewReader(os.Stdin)
	  fmt.Print("Text to send: ")
	  text, _ := reader.ReadString('\n')
	  // send to socket
	  fmt.Fprintf(conn, text + "\n")


		// Server reply
	  message, _ := bufio.NewReader(conn).ReadString('\n')
	  fmt.Print("Message from server: "+message)
	}
}

func server() {
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