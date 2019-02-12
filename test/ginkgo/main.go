package ginkgo

import (
	"database/sql"
	"encoding/json"
	_ "github.com/lib/pq"
	"net/http"
	"path"
	"strconv"
)

func Db() (db *sql.DB) {
	db, err := sql.Open("postgres", "host=db port=5432 user=gp dbname=gp password=gp sslmode=disable")
	if err != nil {
		panic(err)
	}
	return db
}

type Text interface {
	retrieve(id int) (err error)
	create() (err error)
	update() (err error)
	delete() (err error)
}

// main handler function
func Process(t Text) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error
		switch r.Method {
		case "GET":
			err = handleGet(w, r, t)
		case "POST":
			err = handlePost(w, r, t)
		case "PUT":
			err = handlePut(w, r, t)
		case "DELETE":
			err = handleDelete(w, r, t)
		}
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

// Retrieve a post
// GET /ginkgo/1
func handleGet(w http.ResponseWriter, r *http.Request, post Text) (err error) {
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return
	}
	err = post.retrieve(id)
	if err != nil {
		return
	}
	output, err := json.MarshalIndent(post, "", "  ")
	if err != nil {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return
}

// Create a post
// POST /ginkgo/
func handlePost(w http.ResponseWriter, r *http.Request, post Text) (err error) {
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	json.Unmarshal(body, post)
	err = post.create()
	if err != nil {
		return
	}
	w.WriteHeader(200)
	return
}

// Update a post
// PUT /ginkgo/1
func handlePut(w http.ResponseWriter, r *http.Request, post Text) (err error) {
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return
	}
	err = post.retrieve(id)
	if err != nil {
		return
	}
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	json.Unmarshal(body, post)
	err = post.update()
	if err != nil {
		return
	}
	w.WriteHeader(200)
	return
}

// Delete a post
// DELETE /ginkgo/1
func handleDelete(w http.ResponseWriter, r *http.Request, post Text) (err error) {
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return
	}
	err = post.retrieve(id)
	if err != nil {
		return
	}
	err = post.delete()
	if err != nil {
		return
	}
	w.WriteHeader(200)
	return
}
