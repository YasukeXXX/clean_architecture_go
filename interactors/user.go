package interactors

import (
  "github.com/YasukeXXX/clean_architecture_go/models"
  "github.com/YasukeXXX/clean_architecture_go/repositories"
  "github.com/YasukeXXX/clean_architecture_go/use_cases/user"
)

type UserInteractor struct {
  Presenter use_cases.UserCreatePresenter
  Repository repositories.UserRepository
}

func NewUserInteractor(presenter use_cases.UserCreatePresenter, repository repositories.UserRepository) UserInteractor {
  return UserInteractor{ Presenter: presenter, Repository: repository }
}

func (interactor UserInteractor) Create(input use_cases.UserCreateInput) {
  user := models.User{ Name: input.Name }
  interactor.Repository.Save(&user)
  output := use_cases.NewUserCreateOutput(user.Id, user.CreatedAt)
  interactor.Presenter.Render(output)
}
