package main

import (
<<<<<<< HEAD
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
	// fmt.Println("test client")
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
=======
	"fmt"
)

func main() {
	fmt.Println("test client")

>>>>>>> 50b1a39 (initial commit)
}
