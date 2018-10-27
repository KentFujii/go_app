package nested_2

import (
	"html/template"
	"math/rand"
	"net/http"
	"time"
	"runtime"
	"path"
)

func Process(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().Unix())
	var t *template.Template
	if rand.Intn(10) > 5 {
		t, _ = template.ParseFiles(templatePath("layout.html"), templatePath("red_hello.html"))
	} else {
		t, _ = template.ParseFiles(templatePath("layout.html"), templatePath("blue_hello.html"))
	}
	t.ExecuteTemplate(w, "layout", "")
}

func templatePath(name string) string {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("No caller information")
	}
	return path.Dir(filename) + "/" + name
}
