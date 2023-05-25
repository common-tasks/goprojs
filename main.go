package main

import (
	"fmt"
	"path/filepath"
	// "projects/urlshortener"
	"projects/loadbalancer"
)

func main() {
	fmt.Println("main function")
	matches, err := filepath.Glob("novell-zenworks-fde-api-23.3.0*x86_64.msi")
	if(err!=nil){
		fmt.Println("err",err)
	}else{
		fmt.Println(matches)
	}
	if len(matches)>1{
		fmt.Println("le is >1")
	}
	// urlshortener.Shorten()
	loadbalancer.LoadBalance()
	fmt.Println("end of main function")
}
