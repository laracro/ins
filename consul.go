package ins

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"net/http"
	"strings"
)

// consul
type Consul struct {
	*api.Client
}

// config
type consulConfig struct {
	*api.Config
}

func DefaultConfig() *consulConfig {
	return &consulConfig{Config:api.DefaultConfig()}
}

// new consul client
func NewConsulClient(config *consulConfig) (*Consul,error) {
	client,err := api.NewClient(config.Config)
	if err != nil {
		return nil,err
	}
	return &Consul{client},nil
}

// registerer
func (c *Consul) Registerer(service string,port int,checkHandle string) error {
	ip, err := GetTo4()
	if err != nil {
		return err
	}
	tags := strings.Split(service,"-")
	tags = append(tags,service)
	reg := &api.AgentServiceRegistration{
		ID: service,
		Name: service,
		Tags: tags,
		Port: port,
		Address: ip.String(),
	}
	reg.Check = &api.AgentServiceCheck{
		HTTP: fmt.Sprintf("http://%v:%d%v",reg.Address,reg.Port,checkHandle),
		Timeout: "3s",
		Interval: "5s",
		DeregisterCriticalServiceAfter: "30s",
	}
	if err := c.Agent().ServiceRegister(reg); err != nil {
		return err
	}
	return nil
}

// deregister
func (c *Consul) Deregister(service string) error {
	return c.Agent().ServiceDeregister(service)
}

// request
func (c *Consul) DoRequest(serviceName,handle string,params interface{}) (*http.Response,error) {
	service,_,err := c.Agent().Service(serviceName,nil)
	if err != nil {
		return nil,err
	}

	client := &http.Client{}
	addr := service.Address + "/" + handle
	request, err := http.NewRequest(http.MethodGet,addr,nil)
	if err != nil {
		return nil,err
	}

	// todo headers


	// do request
	response, err := client.Do(request)
	if err != nil {
		return nil,err
	}
	defer response.Body.Close()
	return response,nil
}