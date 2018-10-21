package trigger_template

import (
	"html/template"
	"net/http"
	"runtime"
	"path"
)

func Process(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles(templatePath("tmpl.html"))
	t.Execute(w, "Hello World!")
}

func templatePath(name string) string {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("No caller information")
	}
	return path.Dir(filename) + "/" + name
}
