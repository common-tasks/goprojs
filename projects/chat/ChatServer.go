package chat

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func StartServer() {
	fmt.Println("chat server starting")
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer listener.Close()
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			conn.Close()
			continue
		}
		go handleConnection(conn)
	}

}
func handleConnection(conn net.Conn) {
	defer conn.Close()
	fmt.Println("connected to ", conn.RemoteAddr())

	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		text := scanner.Text()

		if strings.TrimSpace(string(text)) == "QUIT" {
			fmt.Println("closing")
			return
		}
		fmt.Printf("%s %s\n", conn.RemoteAddr(), text)
	}

}
