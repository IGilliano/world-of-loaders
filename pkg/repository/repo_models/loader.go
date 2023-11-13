package repo_models

import "math/rand"

type Loader struct {
	ID         int  `json:"id" db:"p_id"`
	Capacity   int  `json:"capacity" db:"capacity"`
	IsDrinking bool `json:"is_drinking" db:"is_drinking"`
	Fatigue    int  `json:"fatigue" db:"fatigue"`
	Salary     int  `json:"salary" db:"salary"`
}

func NewLoader(id int) *Loader {
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
	return &Loader{ID: id, Capacity: int(capacity), IsDrinking: drunk, Salary: int(salary)}
}
