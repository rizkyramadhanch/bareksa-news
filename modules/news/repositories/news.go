package repositories

import (
	"bareksa-news/database"
	// "database/sql"
	"bareksa-news/library/logger"
	"bareksa-news/modules/news/models"
	TagsRepo "bareksa-news/modules/tags/repositories"
	"strconv"
	"errors"
)

type NewsRepositories struct {}

var TagRepo TagsRepo.TagsRepositories



func (repo *NewsRepositories) All() (news []models.News, err error) {
	rows, err := database.DB.Query("select id, news_title, news_content, status from news")

	if err != nil {
		logger.Log.Println("[APP] Error : " + err.Error())
		return nil, err
	}
	defer rows.Close()

	arrNews := []models.News{}
	// Iterate rows
	for rows.Next() {
		news := models.News{}

		// Set current row data to variable
		err = rows.Scan(
			&news.ID,
			&news.Title,
			&news.Description,
			&news.Status,
		)


		if err != nil {
			logger.Log.Println("[APP] Error : " + err.Error())
			return nil, err
		}
		
		arrTags, errGetTags := TagRepo.All()
		if errGetTags != nil {
			logger.Log.Println("[APP] Error : " + errGetTags.Error())
			return nil, errGetTags
		}

		news.Tags = arrTags
		arrNews = append(arrNews, news)
	}
	err = rows.Err()
	if err != nil {
		logger.Log.Println("[APP] Error : " + err.Error())
		return nil, err
	}
	return arrNews, nil
}

func (repo *NewsRepositories) Update(form models.News) (result string, err error) {
	var id int
	rows, err := database.DB.Query("update news set news_title = $1, news_content=$2, status=$3 where id = $4 returning id", form.Title, form.Description, form.Status, form.ID)
	for rows.Next(){
		errScan := rows.Scan(&id)
		if errScan != nil {
			logger.Log.Println("[APP] Error : failed to update news " + errScan.Error())
			return "Failed", err
		}
	}
	s := strconv.Itoa(id)
	errNotFound := errors.New("News not found")
	if s == "0" {
		logger.Log.Println("[APP] Error : " + errNotFound.Error())
		return "News ID not found", errNotFound 
	} else {
		return "News ID " + s + " successfully updated", nil
	}
}

func (repo *NewsRepositories) GetOne(ID int) (result models.News, err error) {
	errRow := database.DB.QueryRow("select id, news_title, news_content, status from news where id= $1", ID).Scan(&result.ID,&result.Title,&result.Description,&result.Status)
	if errRow != nil {
		logger.Log.Println("[APP] Error : failed to get specific news ")
		return result, errRow
	}
	arrTags, errGetTags := TagRepo.All()
		if errGetTags != nil {
			logger.Log.Println("[APP] Error : " + errGetTags.Error())
			return result, errGetTags
		}
	result.Tags = arrTags
	return result, nil
}

func (repo *NewsRepositories) Status(status string) (news []models.News, err error) {
	rows, err := database.DB.Query("select id, news_title, news_content, status from news where status = $1", status)

	if err != nil {
		logger.Log.Println("[APP] Error : " + err.Error())
		return nil, err
	}
	defer rows.Close()

	arrNews := []models.News{}
	// Iterate rows
	for rows.Next() {
		news := models.News{}

		// Set current row data to variable
		err = rows.Scan(
			&news.ID,
			&news.Title,
			&news.Description,
			&news.Status,
		)


		if err != nil {
			logger.Log.Println("[APP] Error : " + err.Error())
			return nil, err
		}
		
		arrTags, errGetTags := TagRepo.All()
		if errGetTags != nil {
			logger.Log.Println("[APP] Error : " + errGetTags.Error())
			return nil, errGetTags
		}

		news.Tags = arrTags
		arrNews = append(arrNews, news)
	}
	err = rows.Err()
	if err != nil {
		logger.Log.Println("[APP] Error : " + err.Error())
		return nil, err
	}
	return arrNews, nil
}











