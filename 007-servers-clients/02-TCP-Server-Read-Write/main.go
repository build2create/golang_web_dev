package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	li, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Panic(err)
	}
	defer li.Close()
	for {
		conn, err := li.Accept()
		if err != nil {
			log.Panic(err)
		}
		//Writing to the connection
		fmt.Fprintf(conn, "#############Hello there, welcome to TCP server####################\n")
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	//THis time we set a deadline
	err := conn.SetDeadline(time.Now().Add(10 * time.Second))
	if err != nil {
		log.Panic(err)
	}
	//Read from the connection we use bufio
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		//Writing back to the connection what we read
		fmt.Fprintf(conn, "I heard you say:: %s\n", ln)
	}
	defer conn.Close()
	//We never get to following line
	//There is no way for the reader to know that it has to stop reading
	// Scan() always returns true
	fmt.Println("Code got here")
}
