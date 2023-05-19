package main

import (

	//"github.com/go-co-op/gocron"

	. "github.com/aman556/ArticleArena/handler/Geeksforgeeks"
	. "github.com/aman556/ArticleArena/handler/medium"
	"github.com/gin-gonic/gin"
)

func main() {
	//my_scheduler := gocron.NewScheduler(time.UTC)

	// my_scheduler.Every(5).Seconds().Do(GeeksForGeeksUserDataScarpping("https://auth.geeksforgeeks.org/user/aman55/articles#"))
	// my_scheduler.Every(5).Seconds().Do(MediumUserDataScarpping("https://medium.com/@ChindaVibhor"))
	// my_scheduler.StartBlocking()

	router := gin.Default()
	router.GET("/geeksforgeeksUserInfo", GetGeeksforGeeksUserInfo)
	router.GET("/mediumUserInfo", GetMediumUserInfo)

	router.Run("localhost:8081")
}
