// Lógica de aplicação (dependem de Repositories e manipulam Models)
package usecase

import (
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

func (it *ToDo_UseCase) Read_ToDo(nID int) (model.ToDo, error) {
	result, err := it.toDoRepository.Select_ToDo_DB(nID)
	if err != nil {
		return model.ToDo{}, err
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