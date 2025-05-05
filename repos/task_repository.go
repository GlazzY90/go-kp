package repos

import "go-api/models"

type TaskRepository interface {
  Create(task *models.Task) error
  FindAllByUser(userID uint) ([]models.Task, error)
  FindByID(id uint) (*models.Task, error)
  Update(task *models.Task) error
  Delete(id uint) error
}