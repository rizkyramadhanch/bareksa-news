package api

import(
	"github.com/gin-gonic/gin"
	"strconv"
	"bareksa-news/modules/tags/repositories"
	// "fmt"
	// "strconv"
	// "regexp"
)

type TagsController struct {}


func (controller *TagsController) List(ctx *gin.Context){
	repo := repositories.TagsRepositories{}

	listTags, err := repo.All()
	if err != nil {
		ctx.JSON(400, gin.H{
			"message" : "[NEWS CONTROLLER] Getting list error",
			"error"   : err.Error(),
		})
		return
	}
	
	ctx.JSON(400, gin.H{
		"message" : "Getting list of news successfully",
		"data"   : listTags,
	})
	return
	
}

//received news_id
func (controller *TagsController) Get(ctx *gin.Context){
	repo := repositories.TagsRepositories{}
	s := ctx.Param("id")
	i, e := strconv.Atoi(s)
	if e != nil {
		ctx.JSON(400, gin.H{
			"message" : "[TAGS CONTROLLER] Getting detail tag error",
			"error"   : e.Error(),
		})
		return
	}
	listTags, err := repo.Get(i)
	if err != nil {
		ctx.JSON(400, gin.H{
			"message" : "[TAGS CONTROLLER] Getting list error",
			"error"   : err.Error(),
		})
		return
	}
	
	ctx.JSON(200, gin.H{
		"message" : "Getting list of news successfully",
		"data"   : listTags,
	})
	return
	
}

func (controller *TagsController) Detail(ctx *gin.Context){
	repo := repositories.TagsRepositories{}
	s := ctx.Param("id")
	i, e := strconv.Atoi(s)
	if e != nil {
		ctx.JSON(400, gin.H{
			"message" : "[TAGS CONTROLLER] Getting detail tag error",
			"error"   : e.Error(),
		})
		return
	}
	listTags, err := repo.Detail(i)
	if err != nil {
		ctx.JSON(400, gin.H{
			"message" : "[TAGS CONTROLLER] Getting list error",
			"error"   : err.Error(),
		})
		return
	}
	
	ctx.JSON(200, gin.H{
		"message" : "Getting list of news successfully",
		"data"   : listTags,
	})
	return
	
}