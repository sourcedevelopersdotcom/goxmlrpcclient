package goxmlrpcclient

import (
	"bytes"
	"log"
	"net/http"
)

type ClientConnection struct {
	url string
}

func (clientConnection *ClientConnection) XmlRpcCall(method string, args interface{}) (reply []byte, err error) {

	buf, errmsh = MarshalXML(method, args)

	if errmsh != nil {
		log.Fatal(errmsh)
	}

	response, err := http.Post(tlc.url, "text/xml", bytes.NewBuffer(buf))

	if err != nil {
		return
	}

	defer response.Body.Close()

	err = UnmarshalXML(reader, &reply)
	return reply, err
}
