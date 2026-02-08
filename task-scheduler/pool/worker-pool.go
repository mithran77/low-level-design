package pool

import (
	"fmt"
	"sync"
	"task-scheduler/utils"
)

var (
	pool_once sync.Once
	pool      *WorkerPool
)

type WorkerPool struct {
	workers     int
	workerQueue chan utils.Runnable
	wg          sync.WaitGroup
}

func NewWorkerPool(workers int) *WorkerPool {
	pool_once.Do(func() {
		pool = &WorkerPool{
			workers:     workers,
			workerQueue: make(chan utils.Runnable),
		}

		for i := range workers {
			pool.wg.Add(1)
			fmt.Println("Starting Worker", i)

			go func() {
				defer pool.wg.Done()
				for r := range pool.workerQueue {
					r.Run()
				}
			}()
		}

	})
	return pool
}

func (p *WorkerPool) Add(r utils.Runnable) {
	p.workerQueue <- r
}

func (p *WorkerPool) Wait() {
	p.wg.Wait()
	close(p.workerQueue)
}
