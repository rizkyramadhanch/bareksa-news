package api

import(
	"github.com/gin-gonic/gin"
	"bareksa-news/modules/tags/repositories"
	// "bareksa-news/modules/tags/models"
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