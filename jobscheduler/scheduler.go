package jobscheduler

import (
	"fmt"
	"time"
)

type Job struct {
	ID        string
	Payload   []byte
	Schedule  time.Time
	Completed bool
}

type Jobscheduler struct {
	Jobs chan Job
}

func NewJobScheduler() *Jobscheduler {
	return &Jobscheduler{Jobs: make(chan Job)}
}
func (js *Jobscheduler) Schedule(job Job) {
	go func() {
		time.Sleep(job.Schedule.Sub(time.Now()))
		js.Jobs <- job
	}()
}

type JobExecutor struct {
	jobs chan Job
}

func NewJobExecutor() *JobExecutor {
	return &JobExecutor{
		jobs: make(chan Job),
	}
}

func (e *JobExecutor) Start() {
	for job := range e.jobs {
		// execute the job
		job.Completed = true
		fmt.Println("Job completed", job.ID)
	}
}
type Node struct {
	ID     string
	Scheduler *Jobscheduler
	Executor  *JobExecutor
}

func (n *Node) DispatchJob(job Job) {
	n.Scheduler.Schedule(job)
}

