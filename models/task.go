// models/task.go
package models

import "time"

type Task struct {
  ID        uint      `gorm:"primaryKey"`
  Title     string    `gorm:"type:varchar(255);not null"`
  Content   string    `gorm:"type:text"`
  Completed bool
  UserID    uint      `gorm:"not null"`
  CreatedAt time.Time
  UpdatedAt time.Time
}
