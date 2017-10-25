package goxmlrpcclient

type Encoder interface {
	MarshalXML(method string, args interface{}) ([]byte, error)
}

func MarshalXML(method string, args interface{}) ([]byte, error) {

}
