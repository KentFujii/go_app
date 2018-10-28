package main

import (
	"net/http"
	"map_store"
	"read_write_files"
	"csv_store"
	"gob_store"
	"gorm_store"
)

func main() {
	// curl localhost:50005/map_store
	http.HandleFunc("/map_store", map_store.Process)
	// curl localhost:50005/read_write_file
	http.HandleFunc("/read_write_file", read_write_files.Process)
	// curl localhost:50005/csv_store
	http.HandleFunc("/csv_store", csv_store.Process)
	// curl localhost:50005/gob_store
	http.HandleFunc("/gob_store", gob_store.Process)
	// curl localhost:50005/gorm_store
	http.HandleFunc("/gorm_store", gorm_store.Process)

	http.ListenAndServe(":50005", nil)
}
