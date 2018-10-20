package headers

import (
	"fmt"
	"net/http"
)

func Headers(w http.ResponseWriter, r *http.Request) {
	h := r.Header
	fmt.Fprintln(w, h)
}
