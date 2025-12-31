server:
	go run src/main.go src/server.go src/client.go server

clientA: 
	go run src/main.go src/server.go src/client.go Client1
clientB:
	go run src/main.go src/server.go src/client.go Client2
