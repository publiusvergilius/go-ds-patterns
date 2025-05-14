package main

type Singleton interface {
	Increment()
}

type singleton struct {
	count int
}

var instance *singleton

func GetInstance () *singleton {	
	if instance == nil {
		instance = new(singleton)
	}
	return instance
}

func (s *singleton) Increment() int {
	s.count++
	return s.count
}


