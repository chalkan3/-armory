package scheduler

var database = map[uint]*Job{
	1: NewJob(1, "new-node", "{}"),
}
