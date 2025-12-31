package main

import (
	"log"
	"net"
	"sync"
)

var (
	clients   = make(map[net.Conn]bool)
	clientsMu sync.Mutex
)

func StartServer() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln("Listener failed: ", err)
	}
	log.Println("Listening on port: ", ln.Addr().String())
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println("Failed to accept connection: ", err)
		}

		// conn is just a collection bytes in the form of []byte
		go func(c net.Conn) {
			defer func() {
				clientsMu.Lock()
				delete(clients, c)
				clientsMu.Unlock()
				c.Close()
			}()

			// Writes data to client connection
			data, err := c.Write([]byte("Welcome To Echo-Chat\n"))
			if err != nil {
				log.Println("Failed to write to connection: ", err, data)
				c.Close()
			}

			// How many bytes we can use
			buf := make([]byte, 1024)

			// Read Connection Data
			for {
				// Manage Client Connections
				clients[c] = true

				n, err := c.Read(buf)
				if err != nil {
					log.Println("User Left: ", err)
					return
				}

				// Send data back to client connection
				broadcast(conn, buf[:n])
			}
		}(conn)
	}
}

func broadcast(sender net.Conn, msg []byte) {
	defer clientsMu.Unlock()
	clientsMu.Lock()

	for client := range clients {
		if client == sender {
			continue
		}

		if _, err := client.Write([]byte(msg)); err != nil {
			log.Println("Failed to send msg to client: ", err)
		}
	}

}
