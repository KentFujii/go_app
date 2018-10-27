package nested_1

import (
	"html/template"
	"net/http"
	"runtime"
	"path"
)

func Process(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles(templatePath("layout.html"))
	t.ExecuteTemplate(w, "layout", "")
}

func templatePath(name string) string {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("No caller information")
	}
	return path.Dir(filename) + "/" + name
}
