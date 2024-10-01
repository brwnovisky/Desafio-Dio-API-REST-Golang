package controllers

import (
	"Desafio-Dio-API-REST-Golang/Context"
	"Desafio-Dio-API-REST-Golang/Models"
	"Desafio-Dio-API-REST-Golang/entities"
	"Desafio-Dio-API-REST-Golang/helpers"
	"Desafio-Dio-API-REST-Golang/services"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type ClientId struct {
	context *Context.ClientDbContext
}

func NewClientId(context *Context.ClientDbContext) *ClientId {
	return &ClientId{context: context}
}

func (h ClientId) Get(controllerData *entities.ControllerData) {
	vars := mux.Vars(controllerData.Request)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		panic(err)
	}

	for _, person := range h.context.Clients {

		var person1 = person
		if person1.Id == id {
			(*controllerData).ResponseBody.Data =
				append((*controllerData).ResponseBody.Data, person1)

			(*controllerData.Writer).WriteHeader(http.StatusOK)
			return
		}
	}

	if len(h.context.Clients) == 0 {
		(*controllerData).ResponseBody.Messages =
			append((*controllerData).ResponseBody.Messages, "Client List empty")
		(*controllerData.Writer).WriteHeader(http.StatusNotFound)
		return
	}

	(*controllerData).ResponseBody.Messages =
		append((*controllerData).ResponseBody.Messages, "Client not found")
	(*controllerData.Writer).WriteHeader(http.StatusNotFound)
}

func (h ClientId) Post(controllerData *entities.ControllerData) {
	(*controllerData.Writer).WriteHeader(http.StatusBadRequest)
}

func (h ClientId) Put(controllerData *entities.ControllerData) {
	vars := mux.Vars(controllerData.Request)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		panic(err)
	}

	var target *Models.Person = nil

	for i := 0; i < len(h.context.Clients); i++ {
		if (h.context.Clients)[i].Id == id {
			target = &(h.context.Clients)[i]
			break
		}
	}

	if target == nil {
		(*controllerData).ResponseBody.Messages =
			append((*controllerData).ResponseBody.Messages, "Client not found")
	}

	var updatedClient Models.Person

	err = json.NewDecoder((*(*controllerData).Request).Body).Decode(&updatedClient)

	if err != nil {
		(*controllerData).ResponseBody.Messages = append((*controllerData).
			ResponseBody.Messages, "Wrong Client object")
		(*controllerData.Writer).WriteHeader(http.StatusBadRequest)
		return
	}

	target.Name = updatedClient.Name
	target.Age = updatedClient.Age

	(*controllerData).ResponseBody.Data =
		append((*controllerData).ResponseBody.Data, *target)

	services.SaveDatabase(h.context.DatabaseName, h.context)

	(*controllerData.Writer).WriteHeader(http.StatusOK)
}

func (h ClientId) Delete(controllerData *entities.ControllerData) {

	vars := mux.Vars(controllerData.Request)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		panic(err)
	}

	var target *Models.Person = nil

	for _, person := range h.context.Clients {
		if person.Id == id {
			target = &person
			break
		}
	}

	if target == nil {
		(*controllerData).ResponseBody.Messages =
			append((*controllerData).ResponseBody.Messages, "Client not found")
	}

	h.context.Clients = helpers.RemoveSliceElement(h.context.Clients, *target)

	(*controllerData).ResponseBody.Messages =
		append((*controllerData).ResponseBody.Messages,
			"Client id "+strconv.Itoa(id)+" has been removed")

	h.context.IdGen.UpdatePoolId(id)

	services.SaveDatabase(h.context.DatabaseName, h.context)

	(*controllerData.Writer).WriteHeader(http.StatusOK)
}
