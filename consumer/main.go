package main

import (
    "bufio"
    "fmt"
    "net"
    "os"
)

func main() {
	addr := os.Getenv("SENSOR_ADDR")
	if addr == "" {
		addr = "sensor:9000"
	}
	conn, err := net.Dial("tcp", addr)
    if err != nil {
        fmt.Println("dial error:", err)
        os.Exit(1)
    }
    scanner := bufio.NewScanner(conn)
    for scanner.Scan() {
        fmt.Println("received:", scanner.Text())
    }
    if err := scanner.Err(); err != nil {
        fmt.Println("read error:", err)
    }
}
