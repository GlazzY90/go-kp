package repos

import (
	"go-api/config"
	"go-api/models"
)

type taskRepo struct{}

func NewTaskRepository() TaskRepository { return &taskRepo{} }
func (r *taskRepo) Create(t *models.Task) error {
  return config.DB.Create(t).Error
}
func (r *taskRepo) FindAllByUser(uid uint) ([]models.Task, error) {
  var ts []models.Task
  return ts, config.DB.Where("user_id = ?", uid).Find(&ts).Error
}
func (r *taskRepo) FindByID(id uint) (*models.Task, error) {
  var t models.Task
  return &t, config.DB.First(&t, id).Error
}
func (r *taskRepo) Update(t *models.Task) error {
  return config.DB.Save(t).Error
}
func (r *taskRepo) Delete(id uint) error {
  return config.DB.Delete(&models.Task{}, id).Error
}
