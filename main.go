package main

import (
	"github.com/gin-contrib/cors"
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
	r.StaticFile("app.log", "./tmp/app.log")
	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"message": "Page not found"})
	})

	newsRoute := r.Group("/news")
	tagsRoute := r.Group("/tags")
	//News
	r.GET("/", responseToUser)
	newsRoute.GET("/", NewsController.List)
	newsRoute.GET("/detail/:id", NewsController.GetOne)
	newsRoute.POST("/update/:id", NewsController.UpdateOne)
	newsRoute.GET("/status/:status", NewsController.Status)
	newsRoute.GET("/topic/:topic", NewsController.Topic)
	newsRoute.POST("/add", NewsController.Add)
	//Tags
	tagsRoute.GET("/", TagsController.List)
	tagsRoute.GET("/detail/:id", TagsController.Get)

	r.Run()
}

//ensure server is working properly, test with JSON response
func responseToUser(c *gin.Context) {
	c.JSON(200, gin.H{
		"Status" : "OK",
	})
}



