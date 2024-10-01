package entities

import "Desafio-Dio-API-REST-Golang/Models"

type ClientResponseBody struct {
	Data     []Models.Person `json:"data"`
	Messages []string        `json:"messages"`
}
