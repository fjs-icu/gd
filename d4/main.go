package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type StringResources struct {
	XMLName        xml.Name         `xml:"resources"`
	ResourceString []ResourceString `xml:"string"`
}

type ResourceString struct {
	XMLName    xml.Name `xml:"string"`
	StringName string   `xml:"name,attr"`
	InnerText  string   `xml:",innerxml"`
}

func main() {
	content, err := ioutil.ReadFile("studygolang.xml")
	if err != nil {
		log.Fatal(err)
	}
	var result StringResources
	err = xml.Unmarshal(content, &result)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(result)
	log.Println(result.ResourceString)
	for i, line := range result.ResourceString {
		log.Println(line.StringName + "===" + line.InnerText)

		//修改ApplicationName节点的内部文本innerText
		if strings.EqualFold(line.StringName, "ApplicationName") {
			fmt.Println("change innerText")

			//注意修改的不是line对象，而是直接使用result中的真实对象
			result.ResourceString[i].InnerText = "这是新的ApplicationName"
		}
	}

	//保存修改后的内容
	xmlOutPut, outPutErr := xml.MarshalIndent(result, "", "    ")
	if outPutErr == nil {
		//加入XML头
		headerBytes := []byte(xml.Header)
		//拼接XML头和实际XML内容
		xmlOutPutData := append(headerBytes, xmlOutPut...)
		//写入文件
		_ = xmlOutPutData
		os.WriteFile("aa3.xml", xmlOutPutData, 064)

		fmt.Println("OK~")
	} else {
		fmt.Println(outPutErr)
	}

}
