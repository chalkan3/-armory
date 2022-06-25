package scheduler

import "time"

type Job struct {
	id      int
	name    string
	payload interface{}
	runAT   time.Time
}

func NewJob(id int, name string, payload interface{}) *Job {
	return &Job{
		id:      id,
		name:    name,
		payload: payload,
	}
}

func (j *Job) ID() int              { return j.id }
func (j *Job) Name() string         { return j.name }
func (j *Job) Payload() interface{} { return j.payload }
func (j *Job) RunAT() time.Time     { return j.runAT }
