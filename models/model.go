package models

type Todo struct {
	ID     uint   `json:"id"`
	Title  string `json:"title" form:"title"`
	Status string `json:"status" form:"status"`
}
