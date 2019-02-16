package main

import (
	"net/http"
)

func main() {
	// http.HandleFunc("/channel_wait", channel_wait.Process)

	http.ListenAndServe(":50008", nil)
}
