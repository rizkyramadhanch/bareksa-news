package api

import(
	"github.com/gin-gonic/gin"
	"bareksa-news/modules/news/repositories"
	"bareksa-news/modules/news/models"
	// "fmt"
	"strconv"
	// "regexp"
)

type NewsController struct {}


func (controller *NewsController) List(ctx *gin.Context){
	repo := repositories.NewsRepositories{}

	listNews, err := repo.All()
	if err != nil {
		ctx.JSON(400, gin.H{
			"message" : "[NEWS CONTROLLER] Getting list error",
			"error"   : err.Error(),
		})
		return
	}
	
	ctx.JSON(400, gin.H{
		"message" : "Getting list of news successfully",
		"data"   : listNews,
	})
	return
	
}

func (controller *NewsController) GetOne(ctx *gin.Context){
	repo := repositories.NewsRepositories{}
	ID := ctx.Param("id")
	i, e := strconv.Atoi(ID)
	if e != nil {
		ctx.JSON(400, gin.H{
			"message" : "[NEWS CONTROLLER] Get ID error",
			"error"   : e.Error(),
		})
		return
	}
	result, err := repo.GetOne(i)
	if err != nil {
		ctx.JSON(400, gin.H{
			"message" : "Getting news failed",
			"error"   : err,
		})
		return
	}
	ctx.JSON(400, gin.H{
		"message" : "Getting news detail succesfully",
		"data"   : result,
	})
	return
}

func (controller *NewsController) UpdateOne(ctx *gin.Context){
	repo := repositories.NewsRepositories{}
	form := models.News{}
	errBind := ctx.ShouldBindJSON(&form)
	if errBind != nil {
		ctx.JSON(400, gin.H{
			"message" : "[NEWS CONTROLLER] Binding  news data error",
			"error"   : errBind.Error(),
		})
		return
	}
	result, err := repo.Update(form)
	if err != nil {
		ctx.JSON(400, gin.H{
			"message" : "Failed to update news",
			"error"   : err,
		})
		return
	}
	ctx.JSON(200, gin.H{
		"message" : "Update news successfully",
		"data"   : result,
	})
	return
}

// func removeLBR(text string) string {
//     re := regexp.MustCompile(`\x{000D}\x{000A}|[\x{000A}\x{000B}\x{000C}\x{000D}\x{0085}\x{2028}\x{2029}]`)
//     return re.ReplaceAllString(text, ``)
// }

