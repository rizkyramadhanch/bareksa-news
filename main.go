package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	//Setup CORS
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"*"}
	r.Use(cors.New(corsConfig))

	r.GET("/", responseToUser)

	r.Run()
}

//ensure server is working properly, test with JSON response
func responseToUser(c *gin.Context) {
	c.JSON(200, gin.H{
		"Status" : "OK",
	})
}
