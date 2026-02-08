package task

import (
	"fmt"
	"sync"
)

type TaskService struct {
	mtx       sync.Mutex
	task_map  map[int]*Task
	scheduler *TaskScheduler
}

func NewTaskService() *TaskService {
	return &TaskService{
		task_map:  make(map[int]*Task),
		scheduler: task_scheduler,
	}
}

func (ts *TaskService) CreateTask(name string, creator string, retries int, priority TaskPriority, task_type TaskType, recurring_delay int) *Task {
	ts.mtx.Lock()
	defer ts.mtx.Unlock()

	id := len(ts.task_map)
	ts.task_map[id] = NewTask(id, name, creator, retries, priority, task_type, recurring_delay)

	return ts.task_map[id]
}

func (ts *TaskService) StopTask(id int) bool {
	ts.mtx.Lock()
	defer ts.mtx.Unlock()

	task, ok := ts.task_map[id]
	if !ok {
		fmt.Println("No task with ID", id)
		return true
	}

	ts.scheduler.Stop(task)
	return true
}
