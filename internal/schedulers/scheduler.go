package schedulers

import (
	"log"
	"time"
)

type Scheduler struct {
	tasks []func()
}

func NewScheduler() *Scheduler {
	return &Scheduler{}
}

func (s *Scheduler) AddTask(task func()) {
	s.tasks = append(s.tasks, task)
}

func (s *Scheduler) Start(interval time.Duration) {
	log.Println("Scheduler started...")
	ticker := time.NewTicker(interval)

	go func() {
		for range ticker.C {
			for _, task := range s.tasks {
				go task()
			}
		}
	}()
}
