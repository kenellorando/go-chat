package main

import (
	"fmt"
	"net"
	"bufio"
	"log"
	"os"
)

const (
	LOCAL_ADDR = "localhost:25624"
)

func main() {
	scan := bufio.NewReader(os.Stdin)
	fmt.Print("Enter target IP -\n   Examples: localhost:25624, 198.37.25.198:25624\n   > ")
	serverAddr, _, err := scan.ReadLine()
	serverAddrStr := string(serverAddr)


	// Start server listener
	go server()

	// Connect to given address
	log.Print("[CLIENT] Attempting to connect to " + serverAddrStr)
	conn, err := net.Dial("tcp", serverAddrStr)
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
	log.Println("[SERVER] Starting server at " + LOCAL_ADDR)

	// Start listening
	ln, err := net.Listen("tcp", LOCAL_ADDR)
	if err != nil {
		log.Print("[SERVER] Failed to start server.")
	} else {
		log.Print("[SERVER] Server started. Awaiting connections on " + LOCAL_ADDR)
		
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
