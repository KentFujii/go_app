package main

import (
	"net/http"
)

func main() {
	http.ListenAndServe(":50004", nil)

}
