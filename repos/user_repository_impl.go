package repos

import (
	"go-api/config"
	"go-api/models"
)

type userRepo struct{}

func NewUserRepository() UserRepository {
    return &userRepo{}
}

func (r *userRepo) FindByEmail(email string) (*models.User, error) {
    var user models.User
    result := config.DB.Where("email = ?", email).First(&user)
    return &user, result.Error
}

func (r *userRepo) Save(user *models.User) error {
    return config.DB.Create(user).Error
}
