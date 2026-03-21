package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net"
)

func getLinesChannel(conn net.Conn) <-chan string {
	ch := make(chan string, 1)

	go func() {
		defer close(ch)
		defer conn.Close()
		buf := make([]byte, 8)
		current_line := ""
		for {
			n, err := conn.Read(buf)
			data := buf[:n]

			for {
				idx := bytes.IndexByte(data, '\n')

				if idx == -1 {
					current_line += string(data)
					break
				}

				current_line += string(data[:idx])
				ch <- current_line
				current_line = ""
				data = data[idx+1:]
			}

			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatal(err)
			}
		}
	}()
	return ch
}

func main() {
	listener, err := net.Listen("tcp", "localhost:42069")
	defer listener.Close()
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("A connection has been accepted...")

		ch := getLinesChannel(conn)
		for data := range ch {
			fmt.Println(data)
		}

		fmt.Println("The connection is closed...")
	}
}
