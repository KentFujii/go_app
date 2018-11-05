package main

import (
	"net/http"
	"xml_parsing_unmarshal_1"
	"xml_parsing_unmarshal_2"
	"xml_parsing_decoder"
	"xml_creating_marshal"
)

func main() {
	// curl localhost:50006/xml_parsing_unmarshal_1
	http.HandleFunc("/xml_parsing_unmarshal_1", xml_parsing_unmarshal_1.Process)
	// curl localhost:50006/xml_parsing_unmarshal_2
	http.HandleFunc("/xml_parsing_unmarshal_2", xml_parsing_unmarshal_2.Process)
	// curl localhost:50006/xml_parsing_decoder
	http.HandleFunc("/xml_parsing_decoder", xml_parsing_decoder.Process)
	// curl localhost:50006/xml_creating_marshal
	http.HandleFunc("/xml_creating_marshal", xml_creating_marshal.Process)
	// curl localhost:50006/xml_creating_encoder
	http.ListenAndServe(":50006", nil)
}
