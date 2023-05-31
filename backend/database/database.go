package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"

	. "github.com/aman556/ArticleArena/backend/database/DAO"
)

var globaldb *sql.DB

func InitDB() {
	config := NewConfig()

	dbServerURL := config.DbUser + ":" + config.DbPass + "@tcp(" + config.DbServiceName + ":" + config.DbHostPort + ")/" + config.DbName

	db, err := sql.Open("mysql", dbServerURL)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	globaldb = db
}

func AddUserHandleInDB() {

}
