package main

import (
	"net/http"
	"trigger_template"
	"random_number"
	"iterator"
	"set_dot"
	"include"
	"custom_function"
	"context_aware"
	"nested_1"
	"nested_2"
)

func main() {
	http.HandleFunc("/trigger_template", trigger_template.Process)
	http.HandleFunc("/random_number", random_number.Process)
	http.HandleFunc("/iterator", iterator.Process)
	http.HandleFunc("/set_dot", set_dot.Process)
	http.HandleFunc("/include", include.Process)
	http.HandleFunc("/custom_function", custom_function.Process)
	http.HandleFunc("/context_aware", context_aware.Process)
	http.HandleFunc("/nested_1", nested_1.Process)
	http.HandleFunc("/nested_2", nested_2.Process)

	http.ListenAndServe(":50004", nil)

}
