package controllers

import (
	"Desafio-Dio-API-REST-Golang/entities"
	"Desafio-Dio-API-REST-Golang/interfaces"
	"Desafio-Dio-API-REST-Golang/models"
	"encoding/json"
	"net/http"
)

type RequestHandler struct {
	interfaces.IController
}

func (rH *RequestHandler) RequestHandler(writer http.ResponseWriter, request *http.Request) {

	controllerData := entities.ControllerData{
		Writer:  &writer,
		Request: request,
		ResponseBody: entities.ClientResponseBody{
			Data:     make([]models.Person, 0),
			Messages: make([]string, 0),
		},
	}

	switch request.Method {
	case "GET":
		rH.Get(&controllerData)
	case "POST":
		rH.Post(&controllerData)
	case "PUT":
		rH.Put(&controllerData)
	case "DELETE":
		rH.Delete(&controllerData)
	default:
		writer.WriteHeader(http.StatusBadRequest)
	}

	jsonBytes, err := json.Marshal(controllerData.ResponseBody)
	if err != nil {
		panic(err)
	}

	writer.Write(jsonBytes)
}
