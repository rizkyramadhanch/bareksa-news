package repositories

import (
	"bareksa-news/database"
	// "database/sql"
	"bareksa-news/library/logger"
	"bareksa-news/modules/tags/models"
	"strconv"
	"errors"
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

func (repo *TagsRepositories) GetByNews(id int) (news []models.Tags, err error) {
	rows, err := database.DB.Query("select a.tags_id, c.tag_name from news_tags a join news b on a.news_id = b.id join tags c on a.tags_id = c.id  where news_id = $1", id)

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

func (repo *TagsRepositories) Detail(id int) (tags models.Tags, err error) {
	row := database.DB.QueryRow("select id, tag_name from tags where id = $1", id)
	
	errScan := row.Scan(&tags.ID,&tags.Name)
	if errScan != nil {
		logger.Log.Println(errScan)
		return tags, errScan
	}

	return tags, nil
}

func (repo *TagsRepositories) Update(form models.Tags) (result string, err error) {
	rows, err := database.DB.Query("update tags set tag_name = $1 where id = $2 returning id", form.Name, form.ID)
	var checker int
	for rows.Next(){
		errScan := rows.Scan(&checker)
		if errScan != nil {
			logger.Log.Println("[APP] Error : failed to update tag " + errScan.Error())
			return "Failed", err
		}
	}
	s := strconv.Itoa(checker)
	errNotFound := errors.New("Tag with id not found")
	if s == "0" {
		logger.Log.Println("[APP] Error : " + errNotFound.Error())
		return "Tag ID not found", errNotFound
	} else {
		return "Tag ID " + s + " successfully updated", nil
	}
}

func (repo *TagsRepositories) Delete(id int) (result string, err error) {
	_, errExec := database.DB.Exec("delete from tags where id = $2 returning id", id)
	if errExec != nil {
		logger.Log.Println("[APP] Error : " + errExec.Error())
		return "Query delete failed	", errors.New("Failed to deleted tag")
	}

	return "Tag has been deleted", nil
}

func (repo *TagsRepositories) Add(form models.Tags) (result string, err error) {
	_, err = database.DB.Exec("insert into tags (tag_name) values ($1)", form.Name)
	if err != nil {
		logger.Log.Println("[APP] Error : failed to add a tag " + err.Error())
		return "Failed", err
	}
	return "A news with title " + form.Name + "has been created", nil
}