package scheduler

import (
	"context"
	"log"
	"math/rand"
	"time"
)

const (
	min = 1
	max = 20
)

type Scheduler struct {
	database map[uint]*Job
	listners Listeners
}

func NewScheduler(listners Listeners) *Scheduler {
	return &Scheduler{
		database: database,
		listners: listners,
	}
}

func (s Scheduler) Schedule(event string, payload interface{}, runAt time.Time) {
	log.Print("🚀 Scheduling event ", event, " to run at ", runAt)
	id := rand.Intn(max-min) + min
	s.database[uint(id)] = NewJob(id, event, payload)
}

func (s Scheduler) AddListener(event string, listenFunc ListenFunc) {
	s.listners[event] = listenFunc
}

func (s Scheduler) checkDueEvents() []Event {
	events := []Event{}

	for _, job := range s.database {
		event := Event{
			ID:      uint(job.id),
			Name:    job.name,
			Payload: job.payload,
		}

		events = append(events, event)

	}

	return events
}

func (s Scheduler) callListners(event Event) {
	eventFN, ok := s.listners[event.Name]
	if ok {
		go eventFN(event.Payload)
		delete(s.database, event.ID)
	}
}

func (s Scheduler) CheckEventsInInterval(ctx context.Context, duration time.Duration) {
	ticker := time.NewTicker(duration)
	go func() {
		for {
			select {
			case <-ctx.Done():
				ticker.Stop()
				return
			case <-ticker.C:
				log.Println("⏰ Ticks Received...")
				events := s.checkDueEvents()
				for _, e := range events {
					s.callListners(e)
				}
			}

		}
	}()
}
