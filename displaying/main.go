package main

import (
	"net/http"
	"trigger_template"
	"random_number"
	"iterator"
	"set_dot"
	"include"
	"custom_function"
)

func main() {
	http.HandleFunc("/trigger_template", trigger_template.Process)
	http.HandleFunc("/random_number", random_number.Process)
	http.HandleFunc("/iterator", iterator.Process)
	http.HandleFunc("/set_dot", set_dot.Process)
	http.HandleFunc("/include", include.Process)
	http.HandleFunc("/custom_function", custom_function.Process)

	http.ListenAndServe(":50004", nil)

}
