package main

import (
	"net"
	"log"
	"fmt"
	"io"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:58000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	fmt.Println("connection arrived!\n")
	defer conn.Close()
	for {
		_, err := io.WriteString(conn, time.Now().Format("15時04分05秒\n"))
		if err != nil {
			log.Print(err)
			return
		}
		time.Sleep(1 * time.Second)
	}
}
