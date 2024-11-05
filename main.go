package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println(": go run main.go <Target IP> <Target Port>")
		return
	}

	targetIP := os.Args[1]
	targetPort := os.Args[2]

	for {
		conn, err := net.Dial("tcp", targetIP+":"+targetPort)
		if err != nil {
			fmt.Println("Connection error:", err)
			return
		}
		conn.Close()
		fmt.Printf("Sent: %s:%s\n", targetIP, targetPort)
		time.Sleep(1 * time.Millisecond) 
	}
}
