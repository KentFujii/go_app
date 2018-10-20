package main

import (
	"net/http"
	"handler"
	"multihandler"
	"handlerfunc"
	"chain_handlerfunc"
	"chain_handler"
)

func main() {
	http.Handle("/handler", &handler.MyHandler{})

	http.Handle("/multihandler_hello", &multihandler.HelloHandler{})
	http.Handle("/multihandler_world", &multihandler.WorldHandler{})

	http.HandleFunc("/handlerfunc_hello", handlerfunc.Hello)
	http.HandleFunc("/handlerfunc_world", handlerfunc.World)

	http.HandleFunc("/chain_handlerfunc", chain_handlerfunc.Log(chain_handlerfunc.Hello))

	http.Handle("/chain_handler", chain_handler.Protect(chain_handler.Log(chain_handler.HelloHandler{})))

	http.ListenAndServe(":50002", nil)
}
