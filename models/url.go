package models

import (
	"database/sql"
	"errors"
	"time"
	"url-shortner/db"
	"url-shortner/internal"
	"url-shortner/utils"
)

type Url struct {
	Id int64 `json:"id"`
	LongUrl string `json:"long_url"`
	ShortUrl string `json:"short_url"`
	CreatedOn time.Time
	UserId int64 `json:"user_id"`
}

func (url *Url) Scan(rows *sql.Rows) error {
	return rows.Scan(
		&url.Id, &url.LongUrl, &url.ShortUrl, &url.CreatedOn, &url.UserId)
}

func (url *Url) New() utils.RowScanner {
	return &Url {}
}

func deleteUrlFromBaseTable(id int64) error {
	_, err := db.Delete(db.DeleteUrlById, id)
	return err
}

func (url *Url) GenerateShortUrl() error {
	shortUrl := url.ShortUrl

	url.Id = internal.GenerateId()
	url.ShortUrl = internal.GetBase62Encoding(url.Id)

	_, err := db.Insert(
		db.InsertNewUrl, url.Id, url.ShortUrl, url.LongUrl, url.UserId)

	if err != nil {
		return errors.New("Unable to create new url.")
	}

	if shortUrl == "" {
		shortUrl = url.ShortUrl
	}

	err = CreateCustomUrl(shortUrl, url.Id)

	if err != nil {
		deleteUrlFromBaseTable(url.Id)
		return errors.New("Unable to create new custom url.")
	}
	return nil
}

func GetLongUrl(shortUrl string) (*Url, error) {
	var url *Url

	rows, err := db.Select(db.SelectUrlByShortUrl, shortUrl)

	if err != nil {
		return nil, err
	}

	urls, err := utils.GetArrayFromRows[*Url](rows)

	if err != nil {
		return nil, err
	}

	if len(urls) == 0 {
		return nil, errors.New("No matching url found.")
	}

	url = urls[0]
	return url, nil
}

func (url *Url) UpdateUrl(existingShortUrl string) (error) {
	fetchedUrl, err := GetLongUrl(existingShortUrl)

	if err != nil {
		return err
	}

	if url.UserId != fetchedUrl.UserId {
		return errors.New("The Url does not belong to this user")
	}

	if fetchedUrl.ShortUrl != "" {
		err := UpdateCustomUrl(fetchedUrl.ShortUrl, fetchedUrl.Id)

		if err != nil {
			return errors.New("Unable to update custom url name")
		}
	}

	if fetchedUrl.LongUrl != "" {
		_, err := db.Update(
			db.UpdateUrl, fetchedUrl.LongUrl, fetchedUrl.Id, fetchedUrl.ShortUrl)

		if err != nil {
			return errors.New("Unable to update the long url")
		}
	}
	return nil
}
