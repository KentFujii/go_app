package json_creating_marshal

import (
	"net/http"
	"encoding/json"
	"fmt"
	"io/ioutil"
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

	output, err := json.MarshalIndent(&post, "", "  ")
	if err != nil {
		fmt.Println("Error marshalling to JSON:", err)
		return
	}
	err = ioutil.WriteFile(filePath("post.json"), output, 0644)
	if err != nil {
		fmt.Println("Error writing JSON to file:", err)
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
