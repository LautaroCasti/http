package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	addr, err := net.ResolveUDPAddr("udp", ":42069")
	if err != nil {
		log.Fatal(err)
		return
	}

	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer conn.Close()

	rd := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(">  ")
		line, err := rd.ReadString('\n')
		if err != nil {
			log.Fatal(err)
			return
		}

		_, err = conn.Write([]byte(line))
		if err != nil {
			log.Fatal(err)
			return
		}
	}
}
