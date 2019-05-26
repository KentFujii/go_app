package logging

import (
	"time"
	"github.com/go-kit/kit/log"
)

type stringService interface {
	Uppercase(string) (string, error)
	Count(string) int
}

type LoggingMiddleware struct {
	Logger log.Logger
	Next stringService
}

func (mw *LoggingMiddleware) Uppercase(s string) (output string, err error) {
	defer func(begin time.Time) {
		_ = mw.Logger.Log(
			"method", "uppercase",
			"input", s,
			"output", output,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())
	output, err = mw.Next.Uppercase(s)
	return
}

func (mw *LoggingMiddleware) Count(s string) (n int) {
	defer func(begin time.Time) {
		_ = mw.Logger.Log(
			"method", "count",
			"input", s,
			"n", n,
			"took", time.Since(begin),
		)
	}(time.Now())
	n = mw.Next.Count(s)
	return
}
