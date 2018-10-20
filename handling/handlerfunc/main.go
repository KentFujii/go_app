package handlerfunc

import (
	"fmt"
	"net/http"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello!")
}

func World(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "World!")
}
