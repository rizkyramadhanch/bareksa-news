package repositories

import (
	"bareksa-news/database"
	// "database/sql"
	"bareksa-news/library/logger"
	"bareksa-news/modules/news/models"
	TagsRepo "bareksa-news/modules/tags/repositories"
	"strconv"
	"errors"
	"fmt"
)

type NewsRepositories struct {}

var TagRepo TagsRepo.TagsRepositories



func (repo *NewsRepositories) All() (news []models.News, err error) {
	rows, err := database.DB.Query("select a.id, a.news_title, a.news_content, b.topic_name, c.name from news a join topics b on a.topic_id = b.id join status c on a.status_id = c.id")

	if err != nil {
		logger.Log.Println("[APP] Error : " + err.Error())
		return nil, err
	}
	defer rows.Close()

	var arrNews []models.News
	// Iterate rows
	for rows.Next() {
		nNews := models.News{}
		// Set current row data to variable
		err = rows.Scan(
			&nNews.ID,
			&nNews.Title,
			&nNews.Description,
			&nNews.Topic,
			&nNews.Status,
		)
		if err != nil {
			logger.Log.Println("[APP] Error : " + err.Error())
			return nil, err
		}
		
		arrTags, errGetTags := TagRepo.Get(nNews.ID)
		if errGetTags != nil {
			logger.Log.Println("[APP] Error : " + errGetTags.Error())
			return nil, errGetTags
		}

			for _, each := range arrTags {
				nNews.Tags = append(nNews.Tags, each.Name)
			}

		arrNews = append(arrNews, nNews)
	}
	if err != nil {
		logger.Log.Println("[APP] Error : " + err.Error())
		return nil, err
	}

	
	return arrNews, nil
}

func (repo *NewsRepositories) Update(id int, form models.CreateNews) (result string, err error) {
	rows, err := database.DB.Query("update news set news_title = $1, news_content=$2, topic_id=$3, status_id=$4 where id = $5 returning id", form.Title, form.Description, form.TopicID, form.StatusID, id)
	var checker int
	for rows.Next(){
		errScan := rows.Scan(&checker)
		if errScan != nil {
			logger.Log.Println("[APP] Error : failed to update news " + errScan.Error())
			return "Failed", err
		}
	}
	s := strconv.Itoa(checker)
	errNotFound := errors.New("News with id not found")
	if s == "0" {
		logger.Log.Println("[APP] Error : " + errNotFound.Error())
		return "News ID not found", errNotFound
	} else {
		return "News ID " + s + " successfully updated", nil
	}
}

func (repo *NewsRepositories) GetOne(ID int) (result models.News, err error) {
	errRow := database.DB.QueryRow("select a.id, a.news_title, a.news_content, b.topic_name, c.name from news a join topics b on a.topic_id = b.id join status c on a.status_id = c.id where a.id= $1", ID).Scan(&result.ID,&result.Title,&result.Description, &result.Topic, &result.Status)
	if errRow != nil {
		logger.Log.Println("[APP] Error : " + errRow.Error())
		return result, errRow
	}
	arrTags, errGetTags := TagRepo.Get(ID)
		if errGetTags != nil {
			logger.Log.Println("[APP] Error : " + errGetTags.Error())
			return result, errGetTags
		}
		for _, each := range arrTags {
			result.Tags = append(result.Tags, each.Name)
		}
	return result, nil
}

func (repo *NewsRepositories) Status(status int) (news []models.News, err error) {
	rows, err := database.DB.Query("select a.id, a.news_title, a.news_content, b.topic_name, c.name from news a join topics b on a.topic_id = b.id join status c on a.status_id = c.id where a.status_id = $1", status)

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
			&news.Topic,
			&news.Status,
		)
		if err != nil {
			logger.Log.Println("[APP] Error : " + err.Error())
			return nil, err
		}
		
		arrTags, errGetTags := TagRepo.Get(news.ID)
		if errGetTags != nil {
			logger.Log.Println("[APP] Error : " + errGetTags.Error())
			return nil, errGetTags
		}
			for _, each := range arrTags {
				news.Tags = append(news.Tags, each.Name)
			}
		arrNews = append(arrNews, news)
	}
	return arrNews, nil
}

func (repo *NewsRepositories) Topic(topic int) (news []models.News, err error) {
	rows, err := database.DB.Query("select a.id, a.news_title, a.news_content, b.topic_name, c.name from news a join topics b on a.topic_id = b.id join status c on a.status_id = c.id where a.topic_id = $1", topic)

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
			&news.Topic,
			&news.Status,
		)
		if err != nil {
			logger.Log.Println("[APP] Error : " + err.Error())
			return nil, err
		}
		
		arrTags, errGetTags := TagRepo.Get(news.ID)
		if errGetTags != nil {
			logger.Log.Println("[APP] Error : " + errGetTags.Error())
			return nil, errGetTags
		}
			for _, each := range arrTags {
				news.Tags = append(news.Tags, each.Name)
			}
		arrNews = append(arrNews, news)
	}
	return arrNews, nil
}

func (repo *NewsRepositories) Add(form models.CreateNews) (result string, err error) {
	
	fmt.Println(form.TopicID)
	_, err = database.DB.Exec("insert into news (news_title, news_content, topic_id, status_id) values ($1, $2, $3, $4)", form.Title, form.Description, form.TopicID, form.StatusID)
		if err != nil {
			logger.Log.Println("[APP] Error : failed to add a news " + err.Error())
			return "Failed", err
		}
	return "A news with title " + form.Title + "has been created" , nil
}











