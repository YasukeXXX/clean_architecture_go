package memory

import (
  "time"
  "github.com/YasukeXXX/clean_architecture_go/models"
)

var data Data

type Data []*models.User

type UserRepository struct {
}

func (user_repository UserRepository) Save(user *models.User) bool {
  user.CreatedAt = time.Now()
  user.Id = uint(len(data)) + 1
  data = append(data, user)
  return true
}

func NewUserRepository() UserRepository {
  return UserRepository{}
}
