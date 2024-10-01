package entities

import "net/http"

type ControllerData struct {
	Writer       *http.ResponseWriter
	Request      *http.Request
	ResponseBody ClientResponseBody
}
