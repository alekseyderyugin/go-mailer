package infrastructure

import (
	"go-mailer/letters/domain"
	"time"
)

type Task struct {
	queue []*domain.Letter
}

type repository interface {
	GetNextForSend(limit uint, lockDuration time.Duration) []*domain.Letter
}

func NewTask(limit uint, repository repository) *Task {
	letters := repository.GetNextForSend(limit, time.Hour)
	return &Task{queue: letters}
}

func (task *Task) Next() *domain.Letter {
	length := task.Length()

	if task.queue == nil || length == 0 {
		return nil
	}

	first := task.queue[0]

	if length > 1 {
		task.queue = task.queue[1:]
	} else {
		task.queue = nil
	}

	return first
}

func (task *Task) Length() int {
	return len(task.queue)
}
