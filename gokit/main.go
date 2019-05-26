package main

import (
	"os"
	"net/http"
	"service"
	"transport"
	"logging"
	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"
)

// type stringService interface {
// 	Uppercase(string) (string, error)
// 	Count(string) int
// }

func main() {
	logger := log.NewLogfmtLogger(os.Stderr)
	svc := &logging.LoggingMiddleware{Logger: logger, Next: &service.String{}}

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

	logger.Log("msg", "HTTP", "addr", ":50009")
	logger.Log("err", http.ListenAndServe(":50009", nil))
}
