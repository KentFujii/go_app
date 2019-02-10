package main

import (
	"net/http"
	"unittest"
)

func main() {
	http.HandleFunc("/unittest", unittest.Process)

	http.ListenAndServe(":50007", nil)
}
