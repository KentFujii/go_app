package mosaic

import (
	"testing"
)

// func TestProcess(t *testing.T) {
// 	mux := http.NewServeMux()
// 	mux.HandleFunc("/mosaic", Process)

// 	writer := httptest.NewRecorder()
// 	request, _ := http.NewRequest("POST", "/mosaic", nil)
// 	mux.ServeHTTP(writer, request)

// 	if writer.Code != 200 {
// 		t.Errorf("Response code is %v", writer.Code)
// 	}
// 	var post Post
// 	json.Unmarshal(writer.Body.Bytes(), &post)
// 	if post.Id != 1 {
// 		t.Errorf("Cannot retrieve JSON post")
// 	}
// }


func TestCut(t *testing.T) {
}
