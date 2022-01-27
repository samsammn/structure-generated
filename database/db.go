package database

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
)

func NewMysqlDB() *sqlx.DB {
	dbConnection := os.Getenv("DB_CONNECTION")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USERNAME")
	dbPass := os.Getenv("DB_PASSWORD")

	parseTime := "true"
	loc := "Asia%2FJakarta"

	// conStr only for mysql connection
	conStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=%s&loc=%s",
		dbUser,
		dbPass,
		dbHost,
		dbPort,
		dbName,
		parseTime,
		loc,
	)

	Db, err := sqlx.Connect(dbConnection, conStr)
	if err != nil {
		fmt.Println("could not connect db, because: " + err.Error())
	}

	return Db
}
