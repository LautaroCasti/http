package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	socket, err := net.Listen("tcp", ":42069")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer socket.Close()

	for {
		conn, err := socket.Accept()
		if err != nil {
			return
		}
		fmt.Print("A connection has been accepted\n")
		for line := range getLinesChannel(conn) {
			fmt.Printf("%s", line)
		}

		fmt.Print("The connection has terminated\n")
	}
}

func getLinesChannel(f io.ReadCloser) <-chan string {
	ch := make(chan string)

	go func() {
		defer f.Close()
		defer close(ch)
		buf := make([]byte, 8)
		var current_line []byte

		for {
			n, err := f.Read(buf)
			if n > 0 {
				for _, b := range buf[:n] {
					current_line = append(current_line, b)
					if b == '\n' {
						ch <- string(current_line)
						current_line = current_line[:0]
					}
				}
			}
			if err == io.EOF {
				if len(current_line) > 0 {
					ch <- string(current_line)
				}
				break
			}
		}
	}()

	return ch
}
