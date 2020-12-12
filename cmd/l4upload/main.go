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
	fmt.Printf("starting listener on port: %v \n", addr)
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("unable to open tcp listener on port: %s \n", addr)
	}
	defer func() {
		fmt.Printf("closing listener on port: %s \n", addr)
		listener.Close()
	}()

	conSema := make(chan struct{}, numWorkers)
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("error accepting tcp connection: %s \n", addr)
		}
		conSema <- struct{}{}
		go func(conn net.Conn) {
			fmt.Printf("accepting tcp connection from: %s \n", conn.RemoteAddr().String())
			for {
				buff := make([]byte, 1500) // TODO (sneha): variable for buffer size
				_, err := conn.Read(buff)
				// TODO(sneha): bit-shifting to read parts of the TCP packet
				// TODO(sneha): Figure out how an HTTP request with a large filesize would be spread over multiple TCP packets
				// assuming an MTU of 1500 bytes

				if err != nil {
					// TODO(sneha): handle connection being closed differently (switch case for error types)
					log.Printf("error reading from conn: %v \n", conn.RemoteAddr().String())
					break
				}
				fmt.Println(string(buff))
				fmt.Printf("request protocol: %v", string(buff[:16]))
				fmt.Printf("host: %v", string(buff[16:38]))
				fmt.Printf("user agent: %v", string(buff[38:64]))
				fmt.Printf("accept: %v", string(buff[64:77]))
				fmt.Printf("body: %v\n", string(buff[130:]))
			}

			// Defer to run before goroutine exits - release worker and close connection
			defer func() {
				fmt.Printf("closing tcp connection: %s \n", conn.RemoteAddr().String())
				<-conSema
				err := conn.Close()
				if err != nil {
					fmt.Printf("error closing connection for: %v \n", conn.RemoteAddr().String())
				}
			}()

		}(conn)

	}

}
