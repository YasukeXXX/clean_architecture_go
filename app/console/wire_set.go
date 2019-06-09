package console

import (
  "github.com/google/wire"
  "github.com/YasukeXXX/clean_architecture_go/use_cases/user"
  "github.com/YasukeXXX/clean_architecture_go/repositories"
  "github.com/YasukeXXX/clean_architecture_go/infrastructures/memory"
  "github.com/YasukeXXX/clean_architecture_go/interactors"
  "github.com/YasukeXXX/clean_architecture_go/controllers"
  "github.com/YasukeXXX/clean_architecture_go/presenters/user"
)

var WireSet = wire.NewSet(
  wire.Bind(new(repositories.UserRepository), new(memory.UserRepository)),
  wire.Bind(new(use_cases.UserCreatePresenter), new(presenters.ConsolePresenter)),
  wire.Bind(new(use_cases.UserUseCase), new(interactors.UserInteractor)),
  memory.NewUserRepository,
  presenters.NewConsolePresenter,
  interactors.NewUserInteractor,
  controllers.NewUserController)

type Controller = controllers.UserController
