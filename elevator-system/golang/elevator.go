package main

import (
	"fmt"
	"sync"
	"time"
)

type Elevator struct {
	id                int
	capacity          int
	Current_floor     int
	current_direction Direction
	requests          []*Request
	mu                sync.RWMutex
}

func NewElevator(id int, capacity int) *Elevator {
	return &Elevator{
		id:                id,
		capacity:          capacity,
		Current_floor:     0,
		current_direction: DirectionUp,
		requests:          make([]*Request, 0),
	}
}

func (e *Elevator) AddRequest(r *Request) {
	e.mu.Lock()
	defer e.mu.Unlock()

	if len(e.requests) < e.capacity {
		e.requests = append(e.requests, r)
		fmt.Sprintf("Elevator %d added request: %d to %d", e.id, r.source_floor, r.destination_floor)
	}
}

func (e *Elevator) GetNextRequest() *Request {
	e.mu.Lock()
	defer e.mu.Unlock()

	next_req := e.requests[len(e.requests)-1]
	e.requests = e.requests[:len(e.requests)-1]
	return next_req
}

func (e *Elevator) ProcessRequests() {
	for true {
		request := e.GetNextRequest()
		e.ProcessRequest(request)
	}
}

func (e *Elevator) ProcessRequest(r *Request) {
	start_floor := e.Current_floor
	end_floor := r.destination_floor

	if start_floor < end_floor {
		e.current_direction = DirectionUp
		for i := start_floor; i < end_floor+1; i++ {
			e.Current_floor = i
			fmt.Sprintf("Elevator %s reached floor %d", e.id, e.Current_floor)
			time.Sleep(1) // Simulating elevator movement
		}
	} else if start_floor > end_floor {
		e.current_direction = DirectionDown
		for i := start_floor; i < end_floor+1; i-- {
			e.Current_floor = i
			fmt.Sprintf("Elevator %s reached floor %d", e.id, e.Current_floor)
			time.Sleep(1) // Simulating elevator movement
		}
	}
}

func (e *Elevator) run() {
	e.ProcessRequests()
}
