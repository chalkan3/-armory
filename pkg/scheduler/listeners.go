package scheduler

type ListenFunc func(interface{})
type Listeners map[string]ListenFunc
