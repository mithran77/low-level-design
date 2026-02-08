package task

import (
	"fmt"
	"time"
)

type TaskPriority int
type TaskType string

const (
	Low TaskPriority = iota
	Medium
	High
)

const (
	FixedDelay TaskType = "FixedDelay"
	OneTime    TaskType = "Once"
	FixedRate  TaskType = "FixedRate"
)

type Task struct {
	ID             int
	Name           string
	Creator        string
	MaxRetries     int
	Priority       TaskPriority
	Type           TaskType
	RecurringDelay int
	CreatedAt      time.Time
	ExecutionTime  time.Time
}

func NewTask(id int, name string, creator string, retries int, priority TaskPriority, task_type TaskType, recurring_delay int) *Task {
	return &Task{
		ID:             id,
		Name:           name,
		Creator:        creator,
		MaxRetries:     retries,
		Priority:       priority,
		Type:           task_type,
		RecurringDelay: recurring_delay,
		CreatedAt:      time.Now(),
	}
}

func (t *Task) Run() {
	if t.Type == FixedRate {
		t.UpdateExecutionTime(time.Now().Add(time.Duration(t.RecurringDelay) * time.Second))
		task_scheduler.Schedule(t)
	}
	t.Execute()
	if t.Type == FixedDelay {
		t.UpdateExecutionTime(time.Now().Add(time.Duration(t.RecurringDelay) * time.Second))
		task_scheduler.Schedule(t)
	}
}

func (t *Task) Execute() {
	fmt.Printf("%v\n", t)
}

func (t *Task) UpdateExecutionTime(new_time time.Time) {
	t.ExecutionTime = new_time
}
