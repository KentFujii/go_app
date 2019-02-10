package main

import (
	"net/http"
	"headers"
	"body"
	"form"
	"fileupload"
	"formfile"
	"write"
	"cookie"
	"cookie_flash"
)

func main() {
	http.HandleFunc("/headers", headers.Headers)

	// curl -id "first_name=Kent&last_name=Fujii" localhost:50003/body
	http.HandleFunc("/body", body.Body)

	// curl -id "hello=aaa&post=123" http://localhost:50003/process?thread=3
	http.HandleFunc("/form", form.Process)

	// curl -i http://localhost:50003/fileupload -F "uploaded=@test.txt"
	http.HandleFunc("/fileupload", fileupload.Process)

	// curl -i http://localhost:50003/formfile -F "uploaded=@test.txt"
	http.HandleFunc("/formfile", formfile.Process)

	http.HandleFunc("/write", write.WriteExample)
	http.HandleFunc("/writeheader", write.WriteHeaderExample)
	http.HandleFunc("/redirect", write.HeaderExample)
	http.HandleFunc("/json", write.JsonExample)

	http.HandleFunc("/set_cookie", cookie.SetCookie)
	http.HandleFunc("/get_cookie", cookie.GetCookie)

	http.HandleFunc("/set_message", cookie_flash.SetMessage)
	http.HandleFunc("/show_message", cookie_flash.ShowMessage)

	http.ListenAndServe(":50003", nil)
}
