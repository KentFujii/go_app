package random_number

import (
	"html/template"
	"math/rand"
	"net/http"
	"time"
	"runtime"
	"path"
)

func Process(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles(templatePath("tmpl.html"))
	rand.Seed(time.Now().Unix())
	t.Execute(w, rand.Intn(10) > 5)
}

func templatePath(name string) string {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("No caller information")
	}
	return path.Dir(filename) + "/" + name
}
