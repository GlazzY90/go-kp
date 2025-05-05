package services

import (
	"errors"
	"go-api/models"
	"go-api/repos"
	"go-api/utils"
)

type AuthService interface {
  Register(user *models.User) error
  Login(email, pwd string) (*models.User, error)
}

type authService struct{ repo repos.UserRepository }

func NewAuthService(r repos.UserRepository) AuthService {
  return &authService{r}
}

func (s *authService) Register(u *models.User) error {
  h, err := utils.HashPassword(u.Password); if err != nil { return err }
  u.Password = h; return s.repo.Save(u)
}
func (s *authService) Login(email, pwd string) (*models.User, error) {
  u, err := s.repo.FindByEmail(email); if err != nil { return nil, err }
  if !utils.CheckPasswordHash(pwd, u.Password) {
    return nil, errors.New("invalid credentials")
  }
  return u, nil
}