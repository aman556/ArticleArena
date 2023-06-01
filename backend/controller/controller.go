package controller

import (
	//. "github.com/aman556/ArticleArena/backend/controller/Geeksforgeeks"
	//. "github.com/aman556/ArticleArena/backend/controller/medium"

	. "github.com/aman556/ArticleArena/backend/database"
	//. "github.com/aman556/ArticleArena/backend/scrapper/utils"
	. "github.com/aman556/ArticleArena/backend/utils"
	"github.com/gin-gonic/gin"
)

// func GetUserArticlesData(c *gin.Context) {
// 	userHandles = GetUserHandle("aman")
// 	var allArticlesData []AllArticles
// 	for userHandle := range userHandles {

// 		allArticlesData = append(allArticlesData, AllArticles{ArticleSite: "medium", ArticlesSiteData: GetMediumUserInfo(userHandleList["aman"][1][1])})

// 	}

// 	c.IndentedJSON(http.StatusOK, allArticlesData)
// }

func PostUserHandle(c *gin.Context) {
	var userData User
	userData.Name = "Aman Sharma"
	userData.ArticleArenaHandle = "aman"
	userData.UserHandleList = append(userData.UserHandleList, UserHandle{WebsiteName: "Geeksforgeeks", WebsiteHandle: "aman55"})
	userData.UserHandleList = append(userData.UserHandleList, UserHandle{WebsiteName: "Medium", WebsiteHandle: "amansharma14041998"})

	AddUserHandleInDB(userData)
}
