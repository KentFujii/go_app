package json_creating_encoder

import (
	"net/http"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"runtime"
	"path"
)

type Post struct {
	Id       int       `json:"id"`
	Content  string    `json:"content"`
	Author   Author    `json:"author"`
	Comments []Comment `json:"comments"`
}

type Author struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Comment struct {
	Id      int    `json:"id"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

func Process(_ http.ResponseWriter, _ *http.Request) {
	post := Post{
		Id:      1,
		Content: "Hello World!",
		Author: Author{
			Id:   2,
			Name: "Sau Sheong",
		},
		Comments: []Comment{
			Comment{
				Id:      1,
				Content: "Have a great day!",
				Author:  "Adam",
			},
			Comment{
				Id:      2,
				Content: "How are you today?",
				Author:  "Betty",
			},
		},
	}

	jsonFile, err := os.Create(filePath("post.json"))
	if err != nil {
		fmt.Println("Error creating JSON file:", err)
		return
	}
	jsonWriter := io.Writer(jsonFile)
	encoder := json.NewEncoder(jsonWriter)
	encoder.SetIndent("", "  ")
	err = encoder.Encode(&post)
	if err != nil {
		fmt.Println("Error encoding JSON to file:", err)
		return
	}
	b, err := ioutil.ReadFile(filePath("post.json"))
	if err != nil {
		fmt.Print(err)
	}
	str := string(b)
	fmt.Println(str)
}

func filePath(name string) string {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("No caller information")
	}
	return path.Dir(filename) + "/" + name
}
