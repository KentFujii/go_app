package main

import (
	"net/http"
	"httptest_1"
	"httptest_2"
)

func main() {
	http.HandleFunc("/httptest_1/", httptest_1.Process)

	http.HandleFunc("/httptest_2/", httptest_2.Process)

	http.ListenAndServe(":50007", nil)
}
