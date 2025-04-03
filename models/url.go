package models

import (
	"database/sql"
	"errors"
	"time"
	"url-shortner/db"
	"url-shortner/internal"
)

type Url struct {
	Id int64 `json:"id"`
	LongUrl string `binding:"required" json:"long_url"`
	ShortUrl string `binding:"required" json:"short_url"`
	CreatedOn time.Time
	UserId int64 `json:"user_id"`
}

func getUrlArrayFromRows(rows *sql.Rows) ([]Url, error) {
	var urls []Url

	defer rows.Close()

	for rows.Next() {
		var url Url
		err := rows.Scan(
			&url.Id, &url.ShortUrl,
			&url.LongUrl, &url.CreatedOn,
			&url.UserId)

		if err != nil {
			return []Url {}, err
		}

		urls = append(urls, url)
	}
	return urls, nil
}

func (url *Url) GenerateShortUrl() error {
	url.Id = internal.GenerateId()
	url.ShortUrl = internal.GetBase62Encoding(url.Id)

	_, err := db.Insert(
		db.InsertNewUrl, url.Id, url.ShortUrl, url.LongUrl, url.UserId)

	if err != nil {
		return errors.New("Unable to create new url.")
	}
	return nil
}

func GetLongUrl(shortUrl string) (*Url, error) {
	var url *Url

	rows, err := db.Select(db.SelectUrlByShortUrl, shortUrl)

	if err != nil {
		return nil, err
	}

	urls, err := getUrlArrayFromRows(rows)

	if err != nil {
		return nil, err
	}

	if len(urls) == 0 {
		return nil, errors.New("No matching url found.")
	}

	url = &urls[0]
	return url, nil
}
