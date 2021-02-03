package external

import(
	"bytes"
	"net/http"
	"io/ioutil"
)


func baseExternal(data []byte, method string, addr string) ([]byte, error){
	var response *http.Response
	var err error
	switch method {
	case "post":
		response, err = http.Post(addr, "application/json", bytes.NewBuffer(data))
	default:
		response, err = http.Get(addr)
	}

	resp, _ := ioutil.ReadAll(response.Body)

	return resp, err
}