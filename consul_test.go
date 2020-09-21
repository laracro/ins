package ins

import (
	"io/ioutil"
	"testing"
)

var (
	Service = "HELLO-EXAMPLE"
	Port = 7878
	CheckHandle = "/_/checkConsul"
	Address = "http://consul.caiqiton.com"
)

func TestConsul_NewConsulClient(t *testing.T) {

	config := DefaultConfig()
	config.Address = Address
	consul,err := NewConsulClient(config)
	if err != nil {
		t.Log(err)
	}

	// do request
	params := map[string]string{}
	response,err := consul.DoRequest("service-demo","/_/checkConsul",params)
	if err != nil {
		t.Log(err)
	}
	body, _ := ioutil.ReadAll(response.Body)
	t.Log(string(body))

	//// register service
	//err = consul.Registerer(Service,Port,CheckHandle)
	//if err != nil {
	//	t.Log(err)
	//}
}