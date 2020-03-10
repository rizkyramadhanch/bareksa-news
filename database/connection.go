package database

import (
	"database/sql"
	_ "github.com/lib/pq"
	"fmt"
	"bareksa-news/config"
	"bareksa-news/library/logger"
)

// DB for querying and transaction
var DB *sql.DB

func init() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		config.Config.DB.Host, config.Config.DB.Port, config.Config.DB.User, config.Config.DB.Password, config.Config.DB.Name)

	var err error
	fmt.Println(config.Config.DB.Name)
	DB, err = sql.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	}

	if err = DB.Ping(); err != nil {
		panic(err)
	}

	logger.Log.Println("[APP] Database successfully connected!")
}
