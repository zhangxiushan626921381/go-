package main

import (
	"log"
	"net/http"
)

func viewHandler(writer http.ResponseWriter, request *http.Request) {
	message := []byte("hello web")
	_, err := writer.Write(message)
	if err != nil {
		log.Fatal(err)
	}
}
func main() {
	http.HandleFunc("/5", viewHandler)
	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err)
}
