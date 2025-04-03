package db

import (
	"database/sql"
	"errors"
	"fmt"
	"url-shortner/config"
	"url-shortner/utils"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func loadSchema() string {
	data, err := utils.ReadFile(config.GetSchemaPath())

	if err != nil {
		panic(err)
	}
	return data
}

func createSchema() {
	data := loadSchema()
	_, err := db.Exec(data)

	if err != nil {
		panic("Cannot create the schema of the database")
	}
}

func InitDb() {
	tempDb, err := sql.Open("sqlite3", config.DatabaseName)

	if err != nil {
		panic("Cannot connect to the database")
	}

	tempDb.SetMaxIdleConns(25)
	tempDb.SetMaxIdleConns(25)

	db = tempDb
	createSchema()
}

func prepareAndExecute(query string, args ...any) (sql.Result, error) {
	stmt, err := db.Prepare(query)

	defer stmt.Close()

	if err != nil {
		return nil, errors.New(fmt.Sprint("Error while preparing query.", err))
	}

	result, err := stmt.Exec(args...)

	if err != nil {
		return nil, errors.New(fmt.Sprint("Error while executing query.", err))
	}
	return result, nil 
}

func Insert(query string, args ...any) (sql.Result, error) {
	return prepareAndExecute(query, args...)
}

func Select(query string, args ...any) (*sql.Rows, error) {
	rows, err := db.Query(query, args...)

	if err != nil {
		return nil, err
	}
	return rows, nil 
}

func Update(query string, args ...any) (sql.Result, error) {
	return prepareAndExecute(query, args...)
}

func Delete(query string, args ...any) (sql.Result, error) {
	return prepareAndExecute(query, args...)
}
