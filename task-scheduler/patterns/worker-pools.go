package patterns

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("Worker ID:", id, "started")
		time.Sleep(time.Second)
		fmt.Println(id, "Completed")
		results <- j * 2
	}
}

func main() {
	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// Create 3 workers
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// Create 5 Jobs
	for i := 1; i <= 5; i++ {
		jobs <- i
	}
	close(jobs)

	// Read Results
	for range 5 {
		<-results
		// fmt.Println("Res:", <-results)
	}
}
