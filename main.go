package main

import (
	"github.com/gin-contrib/cors"
	// "github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"bareksa-news/modules/news/api"
	TagsAPI "bareksa-news/modules/tags/api"
)

func main() {
	r := gin.Default()

	NewsController := api.NewsController{}
	TagsController := TagsAPI.TagsController{}
	//Setup CORS
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"*"}
	r.Use(cors.New(corsConfig))
	r.StaticFile("/app.log", "./tmp/app.log")
	//News
	r.GET("/", responseToUser)
	r.GET("/news", NewsController.List)
	r.GET("/news/:id", NewsController.GetOne)
	r.POST("/news/update", NewsController.UpdateOne)
	//Tags
	r.GET("/tags", TagsController.List)

	r.Run()
}

//ensure server is working properly, test with JSON response
func responseToUser(c *gin.Context) {
	c.JSON(200, gin.H{
		"Status" : "OK",
	})
}



