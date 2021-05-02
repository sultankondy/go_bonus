package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	// client
	connection, err := net.Dial("tcp", "localhost:8080")
	logFatal(err)

	defer connection.Close()
	fmt.Println("Enter your name")

	reader := bufio.NewReader(os.Stdin)
	name, err := reader.ReadString('\n')

	logFatal(err)

	name = strings.Trim(name, " \r\n")

	helloMSG := fmt.Sprintf("Hello %s\n", name)

	connection.Write([]byte(helloMSG))

}
