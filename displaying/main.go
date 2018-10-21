package main

import (
	"net/http"
	"trigger_template"
	"random_number"
	"iterator"
)

func main() {
	http.HandleFunc("/trigger_template", trigger_template.Process)
	http.HandleFunc("/random_number", random_number.Process)
	http.HandleFunc("/iterator", iterator.Process)

	http.ListenAndServe(":50004", nil)

}
