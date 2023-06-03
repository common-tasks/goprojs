package chat

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func Client() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	defer conn.Close()

	go readConnection(conn)
	for {
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		conn.Write([]byte(text))

	}

}
func readConnection(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
