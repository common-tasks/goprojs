package main

import (
	"fmt"
	// "projects/urlshortener"
	"projects/loadbalancer"
)

func main() {
	fmt.Println("main function")
	// urlshortener.Shorten()
	loadbalancer.LoadBalance()
	fmt.Println("end of main function")
}
