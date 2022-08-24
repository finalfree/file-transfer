package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Hello, Carl!\nHttp Server is running!")
	http.HandleFunc("/ip", GetIpHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
