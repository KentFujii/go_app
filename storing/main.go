package main

import (
	"net/http"
	"map_store"
	"read_write_files"
	"csv_store"
	"gob_store"
	"sql_store1"
	"sql_store2"
	"gorm_store"
)

func main() {
	http.HandleFunc("/map_store", map_store.Process)

	http.HandleFunc("/read_write_file", read_write_files.Process)

	http.HandleFunc("/csv_store", csv_store.Process)

	http.HandleFunc("/gob_store", gob_store.Process)

	http.HandleFunc("/sql_store1", sql_store1.Process)

	http.HandleFunc("/sql_store2", sql_store2.Process)

	http.HandleFunc("/gorm_store", gorm_store.Process)

	http.ListenAndServe(":50005", nil)
}
