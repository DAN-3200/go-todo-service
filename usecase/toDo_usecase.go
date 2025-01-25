// Lógica de Negócio (dependem de Repositories e manipulam Models)
// O que a aplicação faz e como as models são usados
package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"app/db"
	"app/model"
	"app/repository"
)

type ToDo_UseCase struct {
	toDoRepository repository.ToDo_Repository
}

// -- Constructor
func Init(repository repository.ToDo_Repository) ToDo_UseCase {
	return ToDo_UseCase{
		toDoRepository: repository,
	}
}

// -- Methods
func (it *ToDo_UseCase) Create_ToDo(request model.ToDo) error {
	err := it.toDoRepository.Insert_ToDo_DB(request)
	if err != nil {
		return err
	}

	return nil
}

var ctx = context.Background()
var useRedis = db.Conn_Redis()

func (it *ToDo_UseCase) Read_ToDo(nID int) (model.ToDo, error) {

	// Consultar o Cache Redis
	var ToDo_Redis model.ToDo
	var value, err = useRedis.Get(ctx, strconv.Itoa(nID)).Result()
	if err != nil {
		fmt.Println("Erro ao Consultar no Redis", err)
	} else {
		json.Unmarshal([]byte(value), &ToDo_Redis)

		// fmt.Println("r:", ToDo_Redis)
		fmt.Println("Consulta direto do Redis")
		return ToDo_Redis, nil
	}

	// Consultar o Banco Postgres
	result, err := it.toDoRepository.Select_ToDo_DB(nID)
	if err != nil {
		return model.ToDo{}, err
	}

	var resultJSON, _ = json.Marshal(result)
	// Tenta salvar no Redis, se não der certo continua
	if err := useRedis.Set(ctx, strconv.Itoa(nID), resultJSON, 30*time.Second).Err(); err != nil {
		fmt.Println("Não foi possível salvar no Redis \n", err)
	}

	return result, nil
}

func (it *ToDo_UseCase) Read_ToDoAll() ([]model.ToDo, error) {
	var result, err = it.toDoRepository.Select_All_ToDo_DB()
	if err != nil {
		return []model.ToDo{}, err
	}
	return result, nil
}

func (it *ToDo_UseCase) Update_ToDo(request model.ToDo) error {
	err := it.toDoRepository.Update_ToDo_DB(request)
	if err != nil {
		return err
	}

	return nil
}

func (it *ToDo_UseCase) Delete_ToDo(nID int) error {
	err := it.toDoRepository.Delete_ToDo_DB(nID)
	if err != nil {
		return err
	}

	return nil
}
