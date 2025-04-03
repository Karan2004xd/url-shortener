package config

import (
	"os"
)

const (
	AppName = "tinyurl"
	DatabaseName = "tinyurl.db"
	BaseUrl = "http://localhost:8080/"
)

func GetRootDir() string {
	return os.Getenv("ROOT_DIR")
}

func GetSchemaPath() string {
	return os.Getenv("ROOT_DIR") + "db/schema.sql"
}

func GetJwtSecretKey() string {
	return os.Getenv("URL_SHORTNER_SECRET_KEY")
}
