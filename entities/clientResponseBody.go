package entities

import "Desafio-Dio-API-REST-Golang/models"

type ClientResponseBody struct {
	Data     []models.Person `json:"data"`
	Messages []string        `json:"messages"`
}
