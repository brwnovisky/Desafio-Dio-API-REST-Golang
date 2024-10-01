package services

import "sort"
import "Desafio-Dio-API-REST-Golang/helpers"

type IdGenService struct {
	IdTop  int   `json:"Id_top"`
	IdPool []int `json:"Id_pool"`
}

func (i *IdGenService) NewId() int {
	newId := i.getPoolId()

	if newId == 0 {
		newId = i.IdTop
		i.IdTop++
	}

	return newId
}

func (i *IdGenService) getPoolId() int {
	poolId := 0

	if len(i.IdPool) > 0 {
		poolId = i.IdPool[0]
		i.IdPool = helpers.RemoveSliceElement(i.IdPool, i.IdPool[0])
	}

	return poolId
}

func (i *IdGenService) UpdatePoolId(deletedId int) {
	i.IdPool = append(i.IdPool, deletedId)
	sort.Ints(i.IdPool)
}
