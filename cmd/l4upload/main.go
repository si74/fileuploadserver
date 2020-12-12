package main

import (
	"fmt"
	"log"
	"net"
)

func main() {

	// TODO(sneha): use flags for variables
	addr := ":8080"
	numWorkers := 10

	// NOTE: Under the hood this is making multiple syscalls (as are the calls to closer the listener and accept an incoming connection).
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("unable to open tcp listener on port: %s", addr)
	}
	defer func() {
		fmt.Printf("closing listener on port: %s", addr)
		listener.Close()
	}()

	conSema := make(chan struct{}, numWorkers)
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("error accepting tcp connection: %s", addr)
		}
		conSema <- struct{}{}
		go func(conn net.Conn) {
			buff := make([]byte, 1500) // TODO (sneha): variable for buffer size
			for {
				fmt.Printf("accepting tcp connection from: %s", conn.RemoteAddr().String())
				_, err := conn.Read(buff)
				// TODO (sneha): bit-shifting to read parts of the http packet
				if err != nil {
					// TODO(sneha): handle connection being closed differently (switch case for error types)
					log.Printf("error reading from conn: %v", conn.RemoteAddr().String())
					break
				}
				fmt.Println(string(buff))
			}

			// Defer to run before goroutine exits - release worker and close connection
			defer func() {
				fmt.Printf("closing tcp connection: %s", conn.RemoteAddr().String())
				<-conSema
				err := conn.Close()
				if err != nil {
					fmt.Printf("error closing connection for: %v", conn.RemoteAddr().String())
				}
			}()

		}(conn)

	}

}
