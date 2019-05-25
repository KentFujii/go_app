package transport

type stringService interface {
	Uppercase(string) (string, error)
	Count(string) int
}

