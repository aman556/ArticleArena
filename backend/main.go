package main

import (

	//"github.com/go-co-op/gocron"

	. "github.com/aman556/ArticleArena/backend/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	// testing on main testing
	router.GET("/profile/:userid", GetUserArticlesData)

	router.Run(":8081")
}
