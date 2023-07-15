package main

import (

	//"github.com/go-co-op/gocron"

	. "github.com/aman556/ArticleArena/backend/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/profile/:userid/getArticles", GetUserArticlesData)
	router.POST("/profile/:userid/postUserHandles", PostUserHandles)
	router.POST("/profile/:userid/postUserInfo", PostUserInfo)
	// router.GET("/profile/:userid/getUserHandles", GetUserHandles)
	// router.GET("/profile/:userid/getUserInfo", GetUserInfo)

	router.Run(":8081")
}
