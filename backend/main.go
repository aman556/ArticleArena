package main

import (

	//"github.com/go-co-op/gocron"

	. "github.com/aman556/ArticleArena/backend/controller"
	"github.com/aman556/ArticleArena/backend/database"
	"github.com/gin-gonic/gin"
)

func main() {
	database.InitDB()
	router := gin.Default()
	router.GET("/profile/:userid/getArticles", GetUserArticlesData)
	router.POST("/profile/:userid/postUserHandle", PostUserHandle)

	router.Run(":8081")
}
