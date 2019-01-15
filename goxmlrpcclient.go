package goxmlrpcclient

import (
	"bytes"
	"encoding/xml"
	"log"
	"net/http"
)

type ClientConnection struct {
	url string
}

func (clientConnection *ClientConnection) XmlRpcCall(request interface{}) (response []byte, err error) {

	// request := make([]byte, 1)
	// requestBuffer := bytes.NewBuffer(request)
	// encoder := xml.NewEncoder(requestBuffer)
	apiRequest, err := xml.Marshal(request)

	if err != nil {
		log.Fatal(err)
	}

	requestReader := bytes.NewBuffer(apiRequest)

	// if err := encoder.Encode(apiRequest); err != nil {
	// 	log.Fatal(err)
	// }

	apiResponse, err := http.Post(clientConnection.url, "text/xml", requestReader)

	if err != nil {
		return
	}

	defer apiResponse.Body.Close()

	apiResponse.Write(bytes.NewBuffer(response))
	// if err := xml.Unmarshal(apiResponse, response); err != nil {
	// 	log.Fatal(err)
	// }
	// err = UnmarshalXML(bytes.NewReader(reply), &reply)
	return response, err
}
