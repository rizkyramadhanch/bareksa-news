package models

// News struct
type News struct {
	ID           int      			`json:"id"`
	Title        string   			`json:"title"`
	Description  string   			`json:"description"`
	Status       string    			`json:"status"`
	Topic		 string				`json:"topic"`
	Tags         []string 			`json:"tag"`
}

type CreateNews struct {
	Title        string   	`json:"title"`
	Description  string   	`json:"description"`
	StatusID     int    	`json:"status_id"`
	TopicID		 int		`json:"topic_id"`
	Tags         string 	`json:"tag"`
}


