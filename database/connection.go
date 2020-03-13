package database

import (
	"os"
	"database/sql"
	_ "github.com/lib/pq"
	"fmt"
	"bareksa-news/config"
	"bareksa-news/library/logger"
	"github.com/gin-gonic/gin"
	// "strconv"
)

// DB for querying and transaction
var DB *sql.DB

var psqlInfo string

func init() {
	/*
		Uncommand code below to running in production mode
	*/
	gin.SetMode(gin.ReleaseMode)

	if gin.Mode() == "debug" {
		logger.Log.Println("[APP] DEBUG MODE")
		psqlInfo = fmt.Sprintf("host=%s port=%d user=%s "+
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
	} else if gin.Mode() == "release" {
		logger.Log.Println("[APP] RELEASE MODE")

		fmt.Println(os.Getenv("host"))
		
		psqlInfo = fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		os.Getenv("host"), os.Getenv("port"), os.Getenv("user"), os.Getenv("password"), os.Getenv("name"))

		var err error
		DB, err = sql.Open("postgres", psqlInfo)

		if err != nil {
			panic(err)
		}

		if err = DB.Ping(); err != nil {
			panic(err)
		}

		logger.Log.Println("[APP] Database successfully connected!")
	}	 else {
		os.Exit(500)
	}
}
