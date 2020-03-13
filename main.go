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
	newsRoute.GET("/list", NewsController.List)
	newsRoute.GET("/detail/:id", NewsController.Detail)
	newsRoute.POST("/update/:id", NewsController.Update)
	newsRoute.GET("/status/:status", NewsController.Status)
	newsRoute.GET("/topic/:topic", NewsController.Topic)
	newsRoute.POST("/add", NewsController.Add)
	newsRoute.POST("/tag/add", NewsController.NewsTag)
	newsRoute.DELETE("/delete/:id", NewsController.Delete)
	//Tags
	tagsRoute.GET("/", TagsController.List)
	tagsRoute.POST("/add", TagsController.Add)
	tagsRoute.GET("/detail/:id", TagsController.Detail)
	tagsRoute.POST("/update", TagsController.Update)
	tagsRoute.DELETE("/delete/:id", TagsController.Delete)

	r.Run()
}

//ensure server is working properly, test with JSON response
func responseToUser(c *gin.Context) {
	c.JSON(200, gin.H{
		"Status" : "OK",
	})
}



