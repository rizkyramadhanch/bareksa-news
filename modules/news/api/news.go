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
	
	ctx.JSON(200, gin.H{
		"message" : "Getting list of news successfully",
		"data"   : listNews,
	})
	return
	
}

func (controller *NewsController) Detail(ctx *gin.Context){
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
	ctx.JSON(200, gin.H{
		"message" : "Getting news detail succesfully",
		"data"   : result,
	})
	return
}

func (controller *NewsController) Update(ctx *gin.Context){
	repo := repositories.NewsRepositories{}
	form := models.CreateNews{}
	id := ctx.Param("id")
	i, e := strconv.Atoi(id)
	if e != nil {
		ctx.JSON(400, gin.H{
			"message" : "[NEWS CONTROLLER] Binding update data error",
			"error"   : e.Error(),
		})
		return
	}
	errBind := ctx.ShouldBindJSON(&form)
	if errBind != nil {
		ctx.JSON(400, gin.H{
			"message" : "[NEWS CONTROLLER] Binding  news data error",
			"error"   : errBind.Error(),
		})
		return
	}
	result, err := repo.Update(i, form)
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

func (controller *NewsController) Status(ctx *gin.Context){
	repo := repositories.NewsRepositories{}
	status := ctx.Param("status")
	i, e := strconv.Atoi(status)
	if e != nil {
		ctx.JSON(400, gin.H{
			"message" : "[NEWS CONTROLLER] Getting list by status error",
			"error"   : e.Error(),
		})
		return
	}
	listNews, err := repo.Status(i)
	if err != nil {
		ctx.JSON(400, gin.H{
			"message" : "[NEWS CONTROLLER] Getting list by status error",
			"error"   : err.Error(),
		})
		return
	}
	if len(listNews) > 0 {
		ctx.JSON(200, gin.H{
			"message" : "Getting list of news by status " + listNews[0].Status + " successfully",
			"data"   : listNews,
		})
		return
	} else {
		ctx.JSON(400, gin.H{
			"message" 	: "News with status " + status + " not found",
			"data"   	: "No Data",
		})
		return
	}
}

func (controller *NewsController) Topic(ctx *gin.Context){
	repo := repositories.NewsRepositories{}
	topic := ctx.Param("topic")
	i, e := strconv.Atoi(topic)
	if e != nil {
		ctx.JSON(400, gin.H{
			"message" : "[NEWS CONTROLLER] Getting news by topic error",
			"error"   : e.Error(),
		})
		return
	}
	listNews, err := repo.Topic(i)
	if err != nil {
		ctx.JSON(400, gin.H{
			"message" : "[NEWS CONTROLLER] Getting list by status error",
			"error"   : err.Error(),
		})
		return
	}
	if len(listNews) > 0 {
		ctx.JSON(200, gin.H{
			"message" : "Getting list of news by status " + listNews[0].Status + " successfully",
			"data"   : listNews,
		})
		return
	} else {
		ctx.JSON(400, gin.H{
			"message" 	: "News with status " + topic + " not found",
			"data"   	: "No Data",
		})
		return
	}
}

func (controller *NewsController) Add(ctx *gin.Context){
	repo := repositories.NewsRepositories{}
	form := models.CreateNews{}
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

func (controller *NewsController) Delete(ctx *gin.Context){
	repo := repositories.NewsRepositories{}
	id := ctx.Param("id")
	i, e := strconv.Atoi(id)
	if e != nil {
		ctx.JSON(400, gin.H{
			"message" : "[NEWS CONTROLLER] Binding update data error",
			"error"   : e.Error(),
		})
		return
	}
	result, err := repo.Delete(i)
	if err != nil {
		ctx.JSON(400, gin.H{
			"message" : "Delete news error",
			"error"   : err,
		})
		return
	}
	ctx.JSON(200, gin.H{
		"message" : "Delete news succesfully",
		"data"   : result,
	})
	return
}

func (controller *NewsController) NewsTag(ctx *gin.Context){
	repo := repositories.NewsRepositories{}
	form := models.TagsToNews{}
	e := ctx.ShouldBindJSON(&form)
	if e != nil {
		ctx.JSON(400, gin.H{
			"message" : "[NEWS CONTROLLER] Binding adding tag to news data error",
			"error"   : e.Error(),
		})
		return
	}
	result, err := repo.NewsTag(form.NewsID, form.TagID)
	if err != nil {
		ctx.JSON(400, gin.H{
			"status" : "Failed",
			"error"   : err,
		})
		return
	}
	ctx.JSON(200, gin.H{
		"status" : "Adding tag to news succesfully",
		"data"   : result,
	})
	return
}

