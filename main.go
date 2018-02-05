package main

import (
	"fmt"
	"net"
	"bufio"
	"log"
	"os"
)

const (
	LOCAL_ADDR = "localhost:50123"
)

func main() {
	// Start server listener

	done := make(chan bool)
	go server(done)
	<-done

	scan := bufio.NewReader(os.Stdin)
	fmt.Print("Enter target IP -\n   Examples: localhost:50123, 198.37.25.198:50123\n   > ")
	serverAddr, _, err := scan.ReadLine()
	serverAddrStr := string(serverAddr)

	tcpAddr, err := net.ResolveTCPAddr("tcp", serverAddrStr)




	

	// Connect to given address
	log.Print("[CLIENT] Attempting to connect to " + serverAddrStr)
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
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

func server(done chan bool) {
	log.Println("[SERVER] Starting server at " + LOCAL_ADDR)

	// Start listening
	ln, err := net.Listen("tcp", ":50123")
	if err != nil {
		log.Print("[SERVER] Failed to start server.")
	} else {
		log.Print("[SERVER] Server started. Awaiting connections on " + LOCAL_ADDR)	
	}

	done<-true

	for {
		// accept connection on port
		conn, _ := ln.Accept()
		go listen(conn)
	}
}

func listen(conn net.Conn) {
	// will listen for message to process ending in newline (\n)
	message, _ := bufio.NewReader(conn).ReadString('\n')
	// output message received
	log.Print("[SERVER] Message received:", string(message))


	// send new string back to client
	conn.Write([]byte("[SERVER] Confirmation reply from server!" + "\n"))

}
