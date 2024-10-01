package controllers

import (
	"Desafio-Dio-API-REST-Golang/Context"
	"Desafio-Dio-API-REST-Golang/Models"
	"Desafio-Dio-API-REST-Golang/entities"
	"Desafio-Dio-API-REST-Golang/services"
	"encoding/json"
	"net/http"
)

type Client struct {
	context *Context.ClientDbContext
}

func NewClient(context *Context.ClientDbContext) *Client {
	return &Client{context: context}
}

func (h Client) Get(controllerData *entities.ControllerData) {
	(*controllerData).ResponseBody.Data = h.context.Clients
	if len(h.context.Clients) == 0 {
		(*controllerData).ResponseBody.Messages =
			append((*controllerData).ResponseBody.Messages, "Client List empty")
	}
}

func (h Client) Post(controllerData *entities.ControllerData) {
	var newClient Models.Person

	err := json.NewDecoder((*(*controllerData).Request).Body).Decode(&newClient)

	if err != nil {
		(*controllerData).ResponseBody.Messages = append((*controllerData).
			ResponseBody.Messages, "Wrong Client object")
		(*controllerData.Writer).WriteHeader(http.StatusBadRequest)
		return
	}

	newClient.Id = h.context.IdGen.NewId()

	h.context.Clients = append(h.context.Clients, newClient)

	(*controllerData).ResponseBody.Data =
		append((*controllerData).ResponseBody.Data, newClient)

	services.SaveDatabase(h.context.DatabaseName, h.context)

	(*controllerData.Writer).WriteHeader(http.StatusCreated)
}

func (h Client) Put(controllerData *entities.ControllerData) {
	(*controllerData.Writer).WriteHeader(http.StatusBadRequest)
}

func (h Client) Delete(controllerData *entities.ControllerData) {
	(*controllerData.Writer).WriteHeader(http.StatusBadRequest)
}
