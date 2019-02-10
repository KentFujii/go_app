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

	http.HandleFunc("/body", body.Body)

	http.HandleFunc("/form", form.Process)

	http.HandleFunc("/fileupload", fileupload.Process)

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
