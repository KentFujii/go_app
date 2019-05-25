package main

import (
	"log"
	"net/http"
	"service"
	"transport"
	httptransport "github.com/go-kit/kit/transport/http"
)

func main() {
	var svc service.String

	tpuc := transport.Uppercase{}
	tpc := transport.Count{}

	uppercaseHandler := httptransport.NewServer(
		tpuc.MakeUppercaseEndpoint(svc),
		tpuc.DecodeUppercaseRequest,
		tpuc.EncodeResponse,
	)

	countHandler := httptransport.NewServer(
		tpc.MakeCountEndpoint(svc),
		tpc.DecodeCountRequest,
		tpc.EncodeResponse,
	)

	http.Handle("/uppercase", uppercaseHandler)
	http.Handle("/count", countHandler)
	log.Fatal(http.ListenAndServe(":50009", nil))
}
