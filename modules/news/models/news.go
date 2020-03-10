package models

// News struct
type News struct {
	ID           int      `json:"id"`
	Title        string   `json:"title"`
	Description  string   `json:"description"`
	Status       string    `json:"status"`
}