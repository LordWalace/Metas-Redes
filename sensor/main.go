package main

import (
	"fmt"
	"math/rand"
	"net"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	ln, err := net.Listen("tcp", ":9000")
	if err != nil {
		fmt.Println("listen error:", err)
		return
	}
	fmt.Println("sensor listening on :9000")
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("accept error:", err)
			continue
		}
		go handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		v := rand.Intn(100)
		fmt.Fprintf(c, "%d\n", v)
		time.Sleep(1 * time.Second)
	}
}
