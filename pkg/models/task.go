package models

var TaskNames = []string{"Unload the car", "Apartment moving", "Furniture assembly", "Cargo delivery"}
var ItemNames = []string{"Books", "Furniture", "Boxes", "Clothes", "Grocery"}

type Task struct {
	ID        int    `json:"id" db:"id"`
	Name      string `json:"name" db:"name"`
	Item      string `json:"item" db:"item"`
	Weight    int    `json:"weight" db:"weight"`
	Available bool   `json:"available" db:"available"`
}

type TaskArchive struct {
	TaskID   int `json:"t_id" db:"t_id"`
	PlayerID int `json:"p_id" db:"p_id"`
}
