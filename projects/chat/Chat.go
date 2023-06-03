package chat

import (
	"fmt"
	"sync"
)

func Chatter() {
	var wg sync.WaitGroup

	wg.Add(2)

	go serve(&wg)
	go client(&wg)

	wg.Wait()
	fmt.Println("all task done")

}
func client(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("starting client")
	ChatClient()
	fmt.Println("end client")
}
func serve(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("starting chat server")
	StartServer()
	fmt.Println(" end chat server")
}
