package api

import(
	"github.com/gin-gonic/gin"
	"strconv"
	"bareksa-news/modules/tags/repositories"
	"bareksa-news/modules/tags/models"
	// "fmt"
	// "strconv"
	// "regexp"
)

type TagsController struct {}

func (controller *TagsController) Add(ctx *gin.Context){
	repo := repositories.TagsRepositories{}
	form := models.Tags{}
	errBind := ctx.ShouldBindJSON(&form)
	if errBind != nil {
		ctx.JSON(400, gin.H{
			"message" : "Bind data while creating news error",
			"error"   : errBind,
		})
		return
	}
	result, err := repo.Add(form)
	if err != nil {
		ctx.JSON(400, gin.H{
			"message" : "Create news error",
			"error"   : err,
		})
		return
	}
	ctx.JSON(200, gin.H{
		"message" : "Create a news succesfully",
		"data"   : result,
	})
	return
}


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

func (controller *TagsController) Update(ctx *gin.Context){
	repo := repositories.TagsRepositories{}
	form := models.Tags{}
	errBind := ctx.ShouldBindJSON(&form)
	if errBind != nil {
		ctx.JSON(400, gin.H{
			"message" : "[NEWS CONTROLLER] Binding  tag data error",
			"error"   : errBind.Error(),
		})
		return
	}
	result, err := repo.Update(form)
	if err != nil {
		ctx.JSON(400, gin.H{
			"message" : "Failed to update tag",
			"error"   : err,
		})
		return
	}
	ctx.JSON(200, gin.H{
		"message" : "Update tag successfully",
		"data"   : result,
	})
	return
}

func (controller *TagsController) Delete(ctx *gin.Context){
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
	
	result, err := repo.Delete(i)
	if err != nil {
		ctx.JSON(400, gin.H{
			"status" : "Failed to delete tag",
			"data"   : err,
		})
		return
	}
	ctx.JSON(200, gin.H{
		"status" : "Delete tag successfully",
		"data"   : result,
	})
	return
}