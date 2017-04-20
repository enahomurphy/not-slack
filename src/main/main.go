package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	var isHost bool

	flag.BoolVar(&isHost, "listen", false, "Listens for a specific port")
	flag.Parse()

	if isHost {
		connIP := os.Args[2]
		runHost(connIP)
	} else {
		connIP := os.Args[1]
		runGuest(connIP)
		fmt.Println("is guest")
	}
}

const port = "8080"

func runHost(ip string) {
	ipAndPort := ip + ":" + port
	listener, listenerErr := net.Listen("tcp", ipAndPort)
	if listenerErr != nil {
		log.Fatal("Error:", listenerErr)
	}

	conn, acceptErr := listener.Accept()
	if acceptErr != nil {
		log.Fatal("Error:", acceptErr)
	}

	reader := bufio.NewReader(conn)

	message, messageErr := reader.ReadString('\n')
	if messageErr != nil {
		log.Fatal("Error:", messageErr)
	}

	fmt.Println(message)
}

func runGuest(ip string) {

}
