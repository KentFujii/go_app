package xml_creating_encoder

import (
	"net/http"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
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

	xmlFile, err := os.Create(filePath("post.xml"))
	if err != nil {
		fmt.Println("Error creating XML file:", err)
		return
	}
	encoder := xml.NewEncoder(xmlFile)
	encoder.Indent("", "\t")
	err = encoder.Encode(&post)
	if err != nil {
		fmt.Println("Error encoding XML to file:", err)
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
