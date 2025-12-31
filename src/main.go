package main

import (
	"fmt"
	"os"
	"strings"
)

type Chat struct {
	Name string
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage:")
		fmt.Println("  server: go run main.go server")
		fmt.Println("  client: go run main.go <name>")
		return
	}

	mode := strings.ToLower(os.Args[1])
	if mode == "server" {
		StartServer()
	}

	chat := Chat{Name: os.Args[1]}
	chat.ConnectClient()

}
