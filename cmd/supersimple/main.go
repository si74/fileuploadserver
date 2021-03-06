package main

import (
	"fmt"
	"net"
	"time"
)

func main() {

	fmt.Println("creating connection")
	conn, err := net.Dial("tcp", "127.0.0.1:80")
	if err != nil {
		fmt.Printf("trouble making connection: %v", err)
	}

	go func() {
		for {
			conn.Write([]byte("string"))
			time.Sleep(2 * time.Second)
		}
	}()

	for {
		buff := make([]byte, 1500)
		fmt.Println("attempting to read...")
		n, err := conn.Read(buff)
		if err != nil {
			fmt.Printf("error reading: %v", err)
		}
		fmt.Println("we've made it here")
		fmt.Printf("length of buff: %v", n)
		fmt.Println(string(buff))
	}
}
