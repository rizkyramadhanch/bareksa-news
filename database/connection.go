package database

import (
	"os"
	"database/sql"
	_ "github.com/lib/pq"
	"fmt"
	"bareksa-news/config"
	"bareksa-news/library/logger"
)

// DB for querying and transaction
var DB *sql.DB

func init() {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		os.Getenv("host"), os.Getenv("port"), os.Getenv("user"), os.Getenv("password"), os.Getenv("name"))

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
