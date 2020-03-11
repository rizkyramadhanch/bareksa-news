package repositories

import (
	"bareksa-news/database"
	// "database/sql"
	"bareksa-news/library/logger"
	"bareksa-news/modules/tags/models"
	// "strconv"
	// "errors"
)

type TagsRepositories struct {}

func (repo *TagsRepositories) All() (news []models.Tags, err error) {
	rows, err := database.DB.Query("select id, tag_name from tags")

	if err != nil {
		logger.Log.Println("[APP] Error : " + err.Error())
		return nil, err
	}
	defer rows.Close()

	arrTags := []models.Tags{}
	// Iterate rows
	for rows.Next() {
		tags := models.Tags{}

		// Set current row data to variable
		err = rows.Scan(
			&tags.ID,
			&tags.Name,
		)

		if err != nil {
			logger.Log.Println("[APP] Error : " + err.Error())
			return nil, err
		}
		
		arrTags = append(arrTags, tags)
	}
	err = rows.Err()
	if err != nil {
		logger.Log.Println("[APP] Error : " + err.Error())
		return nil, err
	}
	return arrTags, nil
}

func (repo *TagsRepositories) Get(id int) (news []models.Tags, err error) {
	rows, err := database.DB.Query("select id, tag_name from tags where news_id = $1", id)

	if err != nil {
		logger.Log.Println("[APP] Error : " + err.Error())
		return nil, err
	}
	defer rows.Close()

	arrTags := []models.Tags{}
	// Iterate rows
	for rows.Next() {
		tags := models.Tags{}

		// Set current row data to variable
		err = rows.Scan(
			&tags.ID,
			&tags.Name,
		)

		if err != nil {
			logger.Log.Println("[APP] Error : " + err.Error())
			return nil, err
		}
		
		arrTags = append(arrTags, tags)
	}
	err = rows.Err()
	if err != nil {
		logger.Log.Println("[APP] Error : " + err.Error())
		return nil, err
	}
	return arrTags, nil
}

func (repo *TagsRepositories) Detail(id int) (news []models.Tags, err error) {
	rows, err := database.DB.Query("select id, tag_name from tags where id = $1", id)

	if err != nil {
		logger.Log.Println("[APP] Error : " + err.Error())
		return nil, err
	}
	defer rows.Close()

	arrTags := []models.Tags{}
	// Iterate rows
	for rows.Next() {
		tags := models.Tags{}

		// Set current row data to variable
		err = rows.Scan(
			&tags.ID,
			&tags.Name,
		)

		if err != nil {
			logger.Log.Println("[APP] Error : " + err.Error())
			return nil, err
		}
		
		arrTags = append(arrTags, tags)
	}
	err = rows.Err()
	if err != nil {
		logger.Log.Println("[APP] Error : " + err.Error())
		return nil, err
	}
	return arrTags, nil
}