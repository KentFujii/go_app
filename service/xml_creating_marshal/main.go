package xml_creating_marshal

import (
	"net/http"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"runtime"
	"path"
)

type Post struct {
	XMLName xml.Name `xml:"post"`
	Id      string   `xml:"id,attr"`
	Content string   `xml:"content"`
	Author  Author   `xml:"author"`
}

type Author struct {
	Id   string `xml:"id,attr"`
	Name string `xml:",chardata"`
}

func Process(_ http.ResponseWriter, _ *http.Request) {
	post := Post{
		Id:      "1",
		Content: "Hello World!",
		Author: Author{
			Id:   "2",
			Name: "Sau Sheong",
		},
	}

	output, err := xml.MarshalIndent(&post, "", "\t")
	if err != nil {
		fmt.Println("Error marshalling to XML:", err)
		return
	}
	err = ioutil.WriteFile(filePath("post.xml"), []byte(xml.Header + string(output)), 0644)
	if err != nil {
		fmt.Println("Error writing XML to file:", err)
		return
	}

	b, err := ioutil.ReadFile(filePath("post.xml"))
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
