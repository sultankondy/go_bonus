package main

import (
	"fmt"
	"log"
	"net"
)

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

var openConnections = make(map[net.Conn]bool)
var newConnection = make(chan net.Conn)

func main() {
	fmt.Println("test server")
	ln, err := net.Listen("tcp", ":8080")
	logFatal(err)

	defer ln.Close()

}
