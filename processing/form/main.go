package form

import (
	"fmt"
	"net/http"
)

func Process(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Fprintln(w, r.Form)
}
