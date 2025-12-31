package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func (c *Chat) ConnectClient() {
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Println("Failed to connect client: ", err)
	}
	// Introduce the Server to Client
	fmt.Fprintf(conn, "%s\n", c.Name)

	status, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		log.Println("Failed to get status: ", status)
	} else {
		log.Println(status)
	}

	fmt.Println("Welcome, ", c.Name)

	input := bufio.NewScanner(os.Stdin)

	// Create a connection here and provide feedback from server
	go func(connection net.Conn) {
		serverOutput := bufio.NewScanner(connection)
		for serverOutput.Scan() {
			// Get server message string
			msg := serverOutput.Text()
			fmt.Printf("\r%s\n%s: ", msg, c.Name)
		}
		if err := serverOutput.Err(); err != nil {
			log.Println("Failed to read msg: ", err)
		}
	}(conn)

	// Client Name
	fmt.Printf("%s: ", c.Name)

	// Loop Until Client Leaves Server. For input
	for input.Scan() {

		// Get the msg string
		text := input.Text()

		// Send with newline so server sees complete message
		if _, err := fmt.Fprintf(conn, "%s: %s\n", c.Name, text); err != nil {
			log.Println("Failed to send msg: ", err)
			break
		}

		// Reprint prompt after sending
		fmt.Printf("%s: ", c.Name)
	}

	if err := input.Err(); err != nil {
		log.Println("Input error:", err)
	}

}
