// +build wireinject

package main

import (
  "github.com/google/wire"
  app "github.com/YasukeXXX/clean_architecture_go/app/console"
)

func NewUserController() app.Controller {
  wire.Build(app.WireSet)
  return app.Controller{}
}
