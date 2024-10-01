package controllers

import (
	"Desafio-Dio-API-REST-Golang/entities"
	"net/http"
)

type Home struct {
}

func NewHome() *Home {
	return &Home{}
}

func (h Home) Get(controllerData *entities.ControllerData) {
	(*controllerData.Writer).WriteHeader(http.StatusOK)

	controllerData.ResponseBody.Messages =
		append(controllerData.ResponseBody.Messages, "Welcome to Go API!")

	controllerData.ResponseBody.Messages =
		append(controllerData.ResponseBody.Messages,
			"'/client' to get all clients")

	controllerData.ResponseBody.Messages =
		append(controllerData.ResponseBody.Messages,
			"'/client/{id}' to client operations")
}

func (h Home) Post(controllerData *entities.ControllerData) {
	(*controllerData.Writer).WriteHeader(http.StatusBadRequest)
}

func (h Home) Put(controllerData *entities.ControllerData) {
	(*controllerData.Writer).WriteHeader(http.StatusBadRequest)
}

func (h Home) Delete(controllerData *entities.ControllerData) {
	(*controllerData.Writer).WriteHeader(http.StatusBadRequest)
}
