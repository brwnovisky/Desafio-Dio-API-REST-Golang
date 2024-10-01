package interfaces

import (
	"Desafio-Dio-API-REST-Golang/entities"
)

type IController interface {
	Get(controllerData *entities.ControllerData)
	Post(controllerData *entities.ControllerData)
	Put(controllerData *entities.ControllerData)
	Delete(controllerData *entities.ControllerData)
}
