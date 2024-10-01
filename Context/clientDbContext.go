package Context

import (
	"Desafio-Dio-API-REST-Golang/Models"
	"Desafio-Dio-API-REST-Golang/services"
)

type ClientDbContext struct {
	DatabaseName string                `json:"database_name"`
	Clients      []Models.Person       `json:"clients"`
	IdGen        services.IdGenService `json:"id_gen"`
}

func NewDatabaseContext(databaseName string) *ClientDbContext {
	clientDbContext := ClientDbContext{
		DatabaseName: databaseName,
		Clients:      make([]Models.Person, 0),
		IdGen:        services.IdGenService{IdTop: 1, IdPool: make([]int, 0)},
	}

	services.LoadDatabase(databaseName, &clientDbContext)

	return &clientDbContext
}
