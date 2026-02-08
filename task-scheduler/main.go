package main

import (
	"task-scheduler/config"
	"task-scheduler/task"
	"time"
)

func main() {
	pool := config.NewWorkerPool(15)

	scheduler := task.NewTaskScheduler(pool)
	service := task.NewTaskService(scheduler)

	task1 := service.CreateTask("Pikachu Capture", "Team Rocket", 3, task.High, task.FixedRate, 60, 180)
	task2 := service.CreateTask("Defeat Vilgax", "Ben 10", 3, task.Low, task.FixedDelay, 30, 120)
	task3 := service.CreateTask("Eat Ramen", "Naruto", 3, task.Low, task.OneTime, 10, 0)

	scheduler.Schedule(task1)
	scheduler.Schedule(task2)
	scheduler.Schedule(task3)

	go func() {
		time.Sleep(time.Duration(40) * time.Second)
		service.StopTask(task2.ID)
	}()

	pool.Wait()
}
