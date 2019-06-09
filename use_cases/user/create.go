package use_cases

import (
  "time"
)

type UserCreateInput struct {
  Name string
}

func NewUserCreateInput(name string) UserCreateInput {
  return UserCreateInput{ Name: name }
}

type UserCreateOutput struct {
  Id uint
  CreatedAt time.Time
}

func NewUserCreateOutput(id uint, createdAt time.Time) UserCreateOutput {
  return UserCreateOutput{ Id: id, CreatedAt: createdAt }
}

type UserCreatePresenter interface {
  Render(UserCreateOutput)
}
