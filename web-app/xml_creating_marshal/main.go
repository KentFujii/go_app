package xml_creating_marshal

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
	xmlStr := string(xmlData)
	fmt.Println(xmlStr)


}

func filePath(name string) string {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("No caller information")
	}
	return path.Dir(filename) + "/" + name
}
