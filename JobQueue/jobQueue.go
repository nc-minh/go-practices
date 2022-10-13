package main

import (
	"fmt"
	"time"
)

type Job interface {
	Process()
}

type Worker struct {
	WorkerID   int
	Done       chan bool
	JobRunning chan Job
}

func NewWorker(workerID int, jobChan chan Job) *Worker {
	return &Worker{
		WorkerID:   workerID,
		Done:       make(chan bool),
		JobRunning: jobChan,
	}
}

func (w *Worker) Run() {
	fmt.Println("Worker: ", w.WorkerID, "started")
	go func() {
		for {
			select {
			case job := <-w.JobRunning:
				fmt.Println("Job running on worker: ", w.WorkerID)
				job.Process()
			case <-w.Done:
				fmt.Println("Stop worker: ", w.WorkerID)
				return
			}
		}
	}()
}

func (w *Worker) StopWorker() {
	go func() {
		w.Done <- true
	}()
}

type JobQueue struct {
	JobRunning chan Job
	Workers    []*Worker
	Done       chan bool
}

func NewJobQueue(numWorkers int) *JobQueue {
	workers := make([]*Worker, numWorkers, numWorkers)
	jobRunning := make(chan Job)

	for i := 0; i < numWorkers; i++ {
		workers[i] = NewWorker(i, jobRunning)
	}

	return &JobQueue{
		JobRunning: jobRunning,
		Workers:    workers,
		Done:       make(chan bool),
	}
}

func (jq *JobQueue) Push(job Job) {
	jq.JobRunning <- job
}

func (jq *JobQueue) Stop() {
	go func() {
		jq.Done <- true
	}()
}

func (jq *JobQueue) Start() {
	go func() {
		for _, worker := range jq.Workers {
			worker.Run()
		}
	}()

	go func() {
		for {
			select {
			case <-jq.Done:
				for _, worker := range jq.Workers {
					worker.StopWorker()
				}
				return
			}
		}
	}()
}

type Sender struct {
	Email string
}

func (s Sender) Process() {
	fmt.Println("Send email to: ", s.Email)
}

func main() {
	emails := []string{
		"ncminh1@gmail.com",
		"ncminh2@gmail.com",
		"ncminh3@gmail.com",
		"ncminh4@gmail.com",
		"ncminh5@gmail.com",
		"ncminh6@gmail.com",
		"ncminh7@gmail.com",
		"ncminh8@gmail.com",
	}

	jobQueue := NewJobQueue(4)
	jobQueue.Start()

	for _, email := range emails {
		sender := Sender{Email: email} //this is a job
		jobQueue.Push(sender)
	}

	time.AfterFunc(time.Second*2, func() {
		jobQueue.Stop()
	})

	time.Sleep(time.Second * 5)
}
