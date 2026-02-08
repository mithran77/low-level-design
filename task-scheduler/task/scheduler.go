package task

import (
	"container/heap"
	"sync"
	"task-scheduler/config"
	"time"
)

type Scheduler interface {
	Schedule(task *Task)
}

var (
	scheduler_once sync.Once
	task_scheduler *TaskScheduler
)

type TaskScheduler struct {
	mtx   *sync.Mutex
	queue *TaskHeap
	cond  *sync.Cond
	pool  *config.WorkerPool
}

func NewTaskScheduler(pool *config.WorkerPool) *TaskScheduler {
	scheduler_once.Do(func() {
		task_scheduler = &TaskScheduler{
			mtx:   &sync.Mutex{},
			queue: &TaskHeap{},
			pool:  pool,
		}
		task_scheduler.cond = sync.NewCond(task_scheduler.mtx)

		go task_scheduler.Execute()
	})
	return task_scheduler
}

func (ts *TaskScheduler) Schedule(task *Task) {
	ts.mtx.Lock()
	defer ts.mtx.Unlock()

	heap.Push(ts.queue, task)
	ts.cond.Signal()
}

func (ts *TaskScheduler) Stop(task *Task) {
	ts.mtx.Lock()
	defer ts.mtx.Unlock()

	for i, item := range *ts.queue {
		if item == task {
			heap.Remove(ts.queue, i)
			break
		}
	}
}

func (ts *TaskScheduler) Execute() {
	for {
		ts.mtx.Lock()
		if ts.queue.Len() == 0 {
			ts.cond.Wait()
		}

		task := ts.queue.Pop().(*Task)
		if task.ExecutionTime.After(time.Now()) {
			heap.Push(ts.queue, task)
			ts.mtx.Unlock()
			time.Sleep(time.Until(task.ExecutionTime))
			continue
		}

		ts.mtx.Unlock()
		ts.pool.Add(task)
	}
}
