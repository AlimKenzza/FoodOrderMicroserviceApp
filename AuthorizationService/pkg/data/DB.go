package data

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"gitlab.com/AlimKenzza/authorization/config"
)

var DB *sql.DB

//Connect to db
func Connect() {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", config.DB_USER, config.DB_PASSWORD, config.DB_NAME)

	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	DB = db
}
