package contexts

import (
	"Desafio-Dio-API-REST-Golang/models"
	"Desafio-Dio-API-REST-Golang/services"
)

type ClientDbContext struct {
	DatabaseName string                `json:"database_name"`
	Clients      []models.Person       `json:"clients"`
	IdGen        services.IdGenService `json:"id_gen"`
}

func NewDatabaseContext(databaseName string) *ClientDbContext {
	clientDbContext := ClientDbContext{
		DatabaseName: databaseName,
		Clients:      make([]models.Person, 0),
		IdGen:        services.IdGenService{IdTop: 1, IdPool: make([]int, 0)},
	}

	services.LoadDatabase(databaseName, &clientDbContext)

	return &clientDbContext
}
