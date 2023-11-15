package models

type Player struct {
	ID       int    `json:"id" db:"id"`
	Login    string `json:"login"`
	Password string `json:"password"`
	Class    string `json:"class"`
}
