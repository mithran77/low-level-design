package main

import (
	"math"
)

type ElevatorController struct {
	elevators []*Elevator
}

func NewElevatorController(num_elevators int, capacity int) *ElevatorController {
	elevator := &ElevatorController{}
	for i := range num_elevators {
		elevator.elevators = append(elevator.elevators, NewElevator(i+1, capacity))
	}
	return elevator
}

func (e *ElevatorController) RequestElevator() *Elevator {

}

func (e *ElevatorController) FindOptimalElevator(source_floor int, destination_floor int) *Elevator {
	optimal_elevator := nil
	min_dist := math.Inf(1)

	for elevator := range e.elevators {
		distance := math.Abs((source_floor - elevator.Current_floor))
		if distance < min_dist {
			min_dist = distance
			optimal_elevator = elevator
		}
	}
	return optimal_elevator
}
