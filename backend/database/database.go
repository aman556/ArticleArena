package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"

	. "github.com/aman556/ArticleArena/backend/database/DAO"
	. "github.com/aman556/ArticleArena/backend/utils"
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

func AddUserHandleInDB(userData User) {
	query := "INSERT INTO Users VALUES (" + userData.Name + "," + userData.ArticleArenaHandle + "," + userData.UserHandleList[0].WebsiteHandle + "," + userData.UserHandleList[0].WebsiteHandle + ")"
	insert, err := globaldb.Query(query)

	// if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}
	// be careful deferring Queries if you are using transactions
	defer insert.Close()
}
