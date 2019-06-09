package repositories

import (
  "github.com/YasukeXXX/clean_architecture_go/models"
)

type UserRepository interface {
  Save(user *models.User) bool
}
