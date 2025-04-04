package models

import (
	"database/sql"
	"errors"
	"fmt"
	"url-shortner/db"
	"url-shortner/utils"
)

type CustomUrl struct {
	Name string `json:"name"`
	Url `json:"url"`
}

func (url *CustomUrl) Scan(rows *sql.Rows) error {
	return rows.Scan(
		&url.Name, &url.Id, &url.ShortUrl,
		&url.LongUrl, &url.CreatedOn, &url.UserId)
}

func (url *CustomUrl) New() utils.RowScanner {
	return &CustomUrl {}
}

func CreateCustomUrl(name string, urlId int64) error {
	_, err := db.Insert(db.InsertNewCustomUrl, name, urlId)
	return err
}

func UpdateCustomUrl(name string, urlId int64) error {
	_, err := db.Update(db.UpdateCustomUrl, name, urlId)
	return err
}

func GetAllCustomUrls(userId int64) ([]*CustomUrl, error) {
	rows, err := db.Select(db.SelectAllCustomUrlsByUserId, userId)

	if err != nil {
		return []*CustomUrl {}, errors.New("Unable to fetch urls")
	}

	urls, err := utils.GetArrayFromRows[*CustomUrl](rows)

	if err != nil {
		fmt.Println(err)
		return []*CustomUrl {}, errors.New("Unable to fetch urls")
	}
	return urls, nil
}
