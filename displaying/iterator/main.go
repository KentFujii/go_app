package iterator

import (
	"html/template"
	"net/http"
	"runtime"
	"path"
)

func Process(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles(templatePath("tmpl.html"))
	daysOfWeek := []string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}
	t.Execute(w, daysOfWeek)
}

func templatePath(name string) string {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("No caller information")
	}
	return path.Dir(filename) + "/" + name
}
