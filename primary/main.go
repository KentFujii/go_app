package main

import (
	"fmt"
	"net/http"
	"os"
	"log"
)

func handler(writer http.ResponseWriter, request *http.Request) {
	log.Println("Hello World")
	fmt.Fprintf(writer, "Hello World, %s!\n", request.URL.Path[1:])
}

func main() {
	f, err := os.OpenFile("/var/log/primary.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("error opening file :", err.Error())
	}
	log.SetOutput(f)
	http.HandleFunc("/", handler)
	http.ListenAndServe(":50001", nil)
}
