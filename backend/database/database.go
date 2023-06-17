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

	globaldb = db
}

func AddUserHandleInDB(userData UserHandles) {
	query := "INSERT INTO `UserHandles` (`ArticleArenaHandle`, `GeeksforgeeksHandle`, `MediumHandle`, `TutorialpointHandle`) VALUES(?,?,?,?,?)"
	insert, err := globaldb.Query(query, userData.ArticleArenaHandle, userData.UserHandleList[0].WebsiteHandle, userData.UserHandleList[1].WebsiteHandle, userData.UserHandleList[2].WebsiteHandle)

	// if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}
	// be careful deferring Queries if you are using transactions
	defer insert.Close()
}

func AddUserInfoInDB(userInfo UserInfo) {
	query := "INSERT INTO `UserInfo` (`UserName`, `ArticleArenaHandle`, `Email`, `GithubUrl`, `LinkedinUrl`) VALUES(?,?,?,?,?)"
	insert, err := globaldb.Query(query, userInfo.Name, userInfo.ArticleArenaHandle, userInfo.Email, userInfo.GithubUrl, userInfo.LinkedinUrl)

	// if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}
	// be careful deferring Queries if you are using transactions
	defer insert.Close()
}

func SelectUserInfoInDB(userHandle string) {

}

func SelectUserHandleInDB(userHandle string) {

}
