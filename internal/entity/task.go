package entity

import (
	"github.com/google/uuid"
)

type Task struct {
	UserID   uuid.UUID `json:"user_id" length:"8"`
	Email    string    `json:"email"`
	TaskName string    `json:"task_name"`
}

func NewTask(userId uuid.UUID, taskName, email string) *Task {
	return &Task{
		UserID:   userId,
		Email:    email,
		TaskName: taskName,
	}
}
