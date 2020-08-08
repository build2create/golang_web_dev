package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	li, err := net.Listen("tcp", "localhost:9000")
	if err != nil {
		log.Panic(err)
	}
	defer li.Close()
	for {
		conn, err := li.Accept()
		if err != nil {
			log.Panic(err)
		}
		go handle(conn)
	}
}
func handle(conn net.Conn) {
	defer conn.Close()
	fmt.Println("Welcome to in- memory database")
	fmt.Println("GET key")
	fmt.Println("SET key, value")
	fmt.Println("Example")
	fmt.Println("SET fav choclate")
	fmt.Println("GET fav")

	data := make(map[string]string)
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		line := scanner.Text()
		fs := strings.Fields(line)
		switch fs[0] {
		case "GET":
			fmt.Fprintf(conn, "Value is %s", data[fs[1]])

		case "SET":
			if len(fs) < 3 {
				fmt.Fprintf(conn, "Invalid number of arguments")
				continue
			}
			k := fs[1]
			v := fs[2]
			data[k] = v
		case "DEL":
			k := fs[1]
			delete(data, k)
		default:
			fmt.Fprintf(conn, "Invalid command")
		}
	}
}
