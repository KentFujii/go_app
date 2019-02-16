package main

import (
	"net/http"
	"mosaic"
)

func main() {
	http.HandleFunc("/mosaic", mosaic.Process)

	http.ListenAndServe(":50008", nil)
}
