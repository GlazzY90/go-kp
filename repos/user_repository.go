package repos

import "go-api/models"

type UserRepository interface {
  Save(user *models.User) error
  FindByEmail(email string) (*models.User, error)
}
