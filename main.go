package main

import (
	"fmt"
	"lighten.top/lightning/api"
	"lighten.top/lightning/api/service"
	"log"
	"net/http"
	"os"
)

func main() {
	fmt.Println("Hello, Carl!\nHttp Server is running!")

	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/ip", api.GetIpHandler)
	http.HandleFunc("/shortMessage", api.ShortMessageHandler)
	http.HandleFunc("/upload", api.UploadFileHandler)

	preStart()

	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func preStart() {
	//create upload dir
	err := os.MkdirAll(service.UploadsDir, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
}
