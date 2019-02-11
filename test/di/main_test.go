package di

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"strings"
)

func TestGetPost(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/di/", Process(&FakePost{}))

	writer := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/di/1", nil)
	mux.ServeHTTP(writer, request)

	if writer.Code != 200 {
		t.Errorf("Response code is %v", writer.Code)
	}
	var post Post
	json.Unmarshal(writer.Body.Bytes(), &post)
	if post.Id != 1 {
		t.Errorf("Cannot retrieve JSON post")
	}
}

func TestPutPost(t *testing.T) {
	mux := http.NewServeMux()
	post := &FakePost{}
	mux.HandleFunc("/di/", Process(post))

	writer := httptest.NewRecorder()
	json := strings.NewReader(`{"content":"Updated post","author":"Sau Sheong"}`)
	request, _ := http.NewRequest("PUT", "/di/1", json)
	mux.ServeHTTP(writer, request)

	if writer.Code != 200 {
		t.Error("Response code is %v", writer.Code)
	}

	if post.Content != "Updated post" {
		t.Error("Content is not correct", post.Content)
	}
}
