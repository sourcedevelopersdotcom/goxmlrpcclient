package goxmlrpcclient

//XmlRpcCaller interface ...
type XmlRpcCaller interface {
	XmlRpcCall(method string, args interface{}) (reply []byte, err error)
}
