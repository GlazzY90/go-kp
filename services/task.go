package services

import (
	"errors"

	"go-api/models"
	"go-api/repos"
)

type TaskService interface {
	Create(task *models.Task) error
	GetAll(userID uint) ([]models.Task, error)
	GetByID(id, userID uint) (*models.Task, error)
	Update(id uint, updated *models.Task, userID uint) error
	Delete(id, userID uint) error
}

type taskService struct {
	repo repos.TaskRepository
}

// NewTaskService constructs a TaskService
func NewTaskService(r repos.TaskRepository) TaskService {
	return &taskService{repo: r}
}

// Create persists a new task
func (s *taskService) Create(task *models.Task) error {
	return s.repo.Create(task)
}

// GetAll returns all tasks belonging to the given user
func (s *taskService) GetAll(userID uint) ([]models.Task, error) {
	return s.repo.FindAllByUser(userID)
}

// GetByID fetches a task by its ID
func (s *taskService) GetByID(id, userID uint) (*models.Task, error) {
	task, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	if task.UserID != userID {
		return nil, errors.New("unauthorized: you do not own this task")
	}
	return task, nil
}

// Update modifies an existing task
func (s *taskService) Update(id uint, updated *models.Task, userID uint) error {
	task, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}
	if task.UserID != userID {
		return errors.New("unauthorized: you do not own this task")
	}

	// Apply updates
	task.Title = updated.Title
	task.Content = updated.Content
	task.Completed = updated.Completed

	return s.repo.Update(task)
}

// Delete removes a task
func (s *taskService) Delete(id, userID uint) error {
	task, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}
	if task.UserID != userID {
		return errors.New("unauthorized: you do not own this task")
	}
	return s.repo.Delete(id)
}
