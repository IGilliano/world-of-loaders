package svc_models

var TaskNames = []string{"Unload the car", "Apartment moving", "Furniture assembly", "Cargo delivery"}
var ItemNames = []string{"Books", "Furniture", "Boxes", "Clothes", "Grocery"}

type Task struct {
	Name   string
	Items  string
	Weight int32
	Cost   int32
}
