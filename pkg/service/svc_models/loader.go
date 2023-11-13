package svc_models

import "math/rand"

type Loader struct {
	Capacity   int32
	IsDrinking bool
	Fatigue    int
	Salary     int32
}

func NewLoader() *Loader {
	capacity := rand.Int31n(30-5) + 5
	if capacity < 5 {
		capacity += 5
	}
	drunk := rand.Intn(2) == 1
	var penalty int32

	if drunk {
		penalty = 2500
	}
	salary := rand.Int31n(30000-10000) + 10000
	salary -= penalty

	if salary < 10000 {
		salary += 10000
	}
	return &Loader{Capacity: capacity, IsDrinking: drunk, Salary: salary}
}
