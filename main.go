package main

import (
	"log"
	"net/http"
	"study/server"
)

func main() {
	http.HandleFunc("/test", server.InterfaceTest)
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
