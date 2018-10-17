package simplest

import (
	"net/http"
	"fmt"
)

func Handler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hello World, %s!", request.URL.Path[1:])
}
