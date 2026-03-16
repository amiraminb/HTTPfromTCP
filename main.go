package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

func getLinesChannel(f io.ReadCloser) <-chan string {
	ch := make(chan string)
	go func() {
		defer close(ch)
		defer f.Close()
		buf := make([]byte, 8)
		current_line := ""
		for {
			n, err := f.Read(buf)
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
	f, err := os.Open("./message.txt")
	if err != nil {
		log.Fatal(err)
	}

	for line := range getLinesChannel(f) {
		fmt.Printf("read: %s\n", line)
	}
}
