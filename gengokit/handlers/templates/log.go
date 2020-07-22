
package templates

const Log = `
package handlers

import(
	"seller/basis/log"
	"github.com/go-kit/kit/endpoint"
)

const (
	SERVICE_NAME = "{{.Service.Name}}Service"
)

var LogInstance *log.Log

func init(){
	if LogInstance != nil {
		LogInstance = log.GetInstance(SERVICE_NAME)
	}
}
func LogServer() (func(string, endpoint.Endpoint) endpoint.Endpoint) {
	return log.LogServer(SERVICE_NAME)
}
`
