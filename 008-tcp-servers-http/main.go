package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	li, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Panicln(err)
	}
	defer li.Close()
	for {
		conn, err := li.Accept()
		if err != nil {
			log.Panicln(err)
		}
		go handler(conn)
	}
}
func handler(conn net.Conn) {
	defer conn.Close()
	scanner := bufio.NewScanner(conn)
	i := 0
	for scanner.Scan() {
		ln := scanner.Text()
		if i == 0 {
			fs := strings.Fields(ln)
			if len(fs[0]) > 0 {
				fmt.Fprintln(conn, fs[1])
			}

		}
		if ln == "" {
			break
		}
		i++
	}
}
