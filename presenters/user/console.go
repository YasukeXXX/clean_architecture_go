package presenters

import (
  "fmt"
  "github.com/YasukeXXX/clean_architecture_go/use_cases/user"
)

type ConsolePresenter struct {
}

func NewConsolePresenter() ConsolePresenter {
  return ConsolePresenter{}
}

func (presenter ConsolePresenter) Render(output use_cases.UserCreateOutput) {
  fmt.Println(output.Id)
}
