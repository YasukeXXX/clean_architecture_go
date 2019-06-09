package controllers

import (
  "github.com/YasukeXXX/clean_architecture_go/use_cases/user"
)

type UserController struct {
  Interactor use_cases.UserUseCase
}

func NewUserController(interactor use_cases.UserUseCase) UserController {
  return UserController{ Interactor: interactor }
}

func (c UserController) Create(name string) {
  input := use_cases.NewUserCreateInput(name)
  c.Interactor.Create(input)
}
