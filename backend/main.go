package main

import (

	//"github.com/go-co-op/gocron"

	. "github.com/aman556/ArticleArena/backend/controller"
	"github.com/aman556/ArticleArena/backend/database"
	"github.com/gin-gonic/gin"
)

func init() {
	database.InitDB()
}

func main() {
	router := gin.Default()
	//router.GET("/profile/:userid", GetUserArticlesData)
	router.POST("/profile/:userid/?gfghandle=%s/mediumhandle=%s", PostUserHandle)

	router.Run(":8081")
}
