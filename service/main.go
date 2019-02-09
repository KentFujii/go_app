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
	http.HandleFunc("/xml_creating_encoder", xml_creating_encoder.Process)
	// curl localhost:50006/json_parsing_unmarshal
	http.HandleFunc("/json_parsing_unmarshal", json_parsing_unmarshal.Process)
	// curl localhost:50006/json_parsing_decoder
	http.HandleFunc("/json_parsing_decoder", json_parsing_decoder.Process)
	http.ListenAndServe(":50006", nil)
}
