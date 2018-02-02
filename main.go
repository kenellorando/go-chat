package main

import (
	"net"
	"bufio"
	"fmt"
	"os"
	"log"
)

const (
	LADDR = "127.0.0.1:25624"
)

func main() {
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