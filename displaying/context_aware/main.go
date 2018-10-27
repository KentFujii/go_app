package context_aware

import (
	"html/template"
	"net/http"
	"runtime"
	"path"
)

func Process(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles(templatePath("tmpl.html"))
	content := `I asked: <i>"What's up?"</i>`
	t.Execute(w, content)
}

func templatePath(name string) string {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("No caller information")
	}
	return path.Dir(filename) + "/" + name
}
