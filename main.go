package main

import (
	"fmt"
	"net"
	"bufio"
	"log"
	"os"
)

const (
	SERVER_ADDR = "localhost:25624"
)

func main() {
	// Start server listener
	go server()

	// Connect to given address
	log.Print("[CLIENT] Attempting to connect to " + SERVER_ADDR)
	conn, err := net.Dial("tcp", SERVER_ADDR)
	if err != nil {
		log.Print("[CLIENT] Failed to connect.")
	} else {
		log.Print("[CLIENT] Successfully connected.")
	}


	for { 
	  // read in input from stdin
	  reader := bufio.NewReader(os.Stdin)
	  fmt.Print("Text to send: ")
	  text, _ := reader.ReadString('\n')
	  // send to socket
	  fmt.Fprintf(conn, text + "\n")


		// Print the server reply
	  message, _ := bufio.NewReader(conn).ReadString('\n')
	  log.Print("[CLIENT] Received reply from server: " + message)
	}
}

func server() {
	log.Println("[SERVER] Starting server at " + SERVER_ADDR)

	// Start listening
	ln, err := net.Listen("tcp", SERVER_ADDR)
	if err != nil {
		log.Print("[SERVER] Failed to start server.")
	} else {
		log.Print("[SERVER] Server started. Awaiting connections on " + SERVER_ADDR)
		
	}

	// accept connection on port
	conn, _ := ln.Accept()

	for {
		// will listen for message to process ending in newline (\n)
		message, _ := bufio.NewReader(conn).ReadString('\n')
		// output message received
		log.Print("[SERVER] Message received:", string(message))


		// send new string back to client
		conn.Write([]byte("[SERVER] Confirmation reply from server!" + "\n"))
	}
	
}