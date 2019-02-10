package main

import (
	"net/http"
	"xml_parsing_unmarshal_1"
	"xml_parsing_unmarshal_2"
	"xml_parsing_decoder"
	"xml_creating_marshal"
	"xml_creating_encoder"
	"json_parsing_unmarshal"
	"json_parsing_decoder"
	"json_creating_marshal"
	"json_creating_encoder"
	"web_service"
)

func main() {
	http.HandleFunc("/xml_parsing_unmarshal_1", xml_parsing_unmarshal_1.Process)

	http.HandleFunc("/xml_parsing_unmarshal_2", xml_parsing_unmarshal_2.Process)

	http.HandleFunc("/xml_parsing_decoder", xml_parsing_decoder.Process)

	http.HandleFunc("/xml_creating_marshal", xml_creating_marshal.Process)

	http.HandleFunc("/xml_creating_encoder", xml_creating_encoder.Process)

	http.HandleFunc("/json_parsing_unmarshal", json_parsing_unmarshal.Process)

	http.HandleFunc("/json_parsing_decoder", json_parsing_decoder.Process)

	http.HandleFunc("/json_creating_marshal", json_creating_marshal.Process)

	http.HandleFunc("/json_creating_encoder", json_creating_encoder.Process)

	http.HandleFunc("/web_service/", web_service.Process)

	http.ListenAndServe(":50006", nil)
}
