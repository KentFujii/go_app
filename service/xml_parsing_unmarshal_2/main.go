package xml_parsing_unmarshal_2

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
	XMLName  xml.Name  `xml:"post"`
	Id       string    `xml:"id,attr"`
	Content  string    `xml:"content"`
	Author   Author    `xml:"author"`
	Xml      string    `xml:",innerxml"`
	Comments []Comment `xml:"comments>comment"`
}

type Author struct {
	Id   string `xml:"id,attr"`
	Name string `xml:",chardata"`
}

type Comment struct {
	Id      string `xml:"id,attr"`
	Content string `xml:"content"`
	Author  Author `xml:"author"`
}

func Process(_ http.ResponseWriter, _ *http.Request) {
	xmlFile, err := os.Open(filePath("post.xml"))
	if err != nil {
		fmt.Println("Error opening XML file:", err)
		return
	}
	defer xmlFile.Close()
	xmlData, err := ioutil.ReadAll(xmlFile)
	if err != nil {
		fmt.Println("Error reading XML data:", err)
		return
	}

	var post Post
	xml.Unmarshal(xmlData, &post)
	fmt.Println(post.XMLName.Local)
	fmt.Println(post.Id)
	fmt.Println(post.Content)
	fmt.Println(post.Author)
	fmt.Println(post.Xml)
	fmt.Println(post.Author.Id)
	fmt.Println(post.Author.Name)
	fmt.Println(post.Comments)
	fmt.Println(post.Comments[0].Id)
	fmt.Println(post.Comments[0].Content)
	fmt.Println(post.Comments[0].Author)
	fmt.Println(post.Comments[1].Id)
	fmt.Println(post.Comments[1].Content)
	fmt.Println(post.Comments[1].Author)
}

func filePath(name string) string {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("No caller information")
	}
	return path.Dir(filename) + "/" + name
}
