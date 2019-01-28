package goxmlrpcclient

import (
	"encoding/xml"
	"fmt"
	"os"
	"testing"
)

type MethodCall struct {
	XMLName    xml.Name   `xml:"methodCall"`
	MethodName MethodName `xml:"methodName"`
	Parameters []Params   `xml:"params"`
}

type MethodName struct {
	XMLName xml.Name `xml:"methodName"`
}

type Params struct {
	XMLName   xml.Name `xml:"params"`
	Parameter []Param  `xml:"param"`
}

type Param struct {
	XMLName xml.Name `xml:"param"`
	Value   Value    `xml:"value"`
}

type Value struct {
	XMLName     xml.Name   `xml:"value"`
	StringValue StringType `xml:"string"`
}

type StringType struct {
	XMLName xml.Name `xml:"string"`
}

func TestXmlRpcCall(t *testing.T) {
	xmlFile, err := os.Open("xmlTest.xml")
	if err != nil {
		fmt.Println(err)
	}
}
