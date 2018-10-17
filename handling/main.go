package main

import (
	"net/http"
	"app/simplest"
)

func main() {
	http.HandleFunc("/", simplest.Handler)
	http.ListenAndServe(":50002", nil)
}
