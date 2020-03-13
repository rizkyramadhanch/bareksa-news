package models

// News struct
type Tags struct {
	ID          int      	`json:"id"`
	Name        string   	`json:"name"`
}

type IsTagged struct {
	ID int `json:"id"`
}
