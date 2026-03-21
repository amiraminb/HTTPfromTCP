package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	udpAddr, err := net.ResolveUDPAddr("udp", "localhost:42069")
	if err != nil {
		panic(err)
	}

	udpConn, err := net.DialUDP("udp", nil, udpAddr)
	if err != nil {
		log.Fatal(err)
	}
	defer udpConn.Close()

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")
		r, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		_, err = udpConn.Write([]byte(r))
		if err != nil {
			log.Fatal(err)
		}
	}
}
