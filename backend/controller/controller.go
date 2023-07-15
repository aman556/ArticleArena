package controller

import (
	. "github.com/aman556/ArticleArena/backend/controller/Geeksforgeeks"
	. "github.com/aman556/ArticleArena/backend/controller/medium"

	"net/http"

	. "github.com/aman556/ArticleArena/backend/database"
	. "github.com/aman556/ArticleArena/backend/scrapper/utils"

	. "github.com/aman556/ArticleArena/backend/utils"
	"github.com/gin-gonic/gin"
)

func GetUserArticlesData(c *gin.Context) {
	userHandleList := make(map[string][][]string)
	userHandleList["aman"] = [][]string{{"GeeksforGeeks", "aman55"}, {"medium", "amansharma14041998"}}

	var allArticlesData []AllArticles
	allArticlesData = append(allArticlesData, AllArticles{ArticleSite: userHandleList["aman"][0][0], ArticlesSiteData: GetGeeksforGeeksUserInfo(userHandleList["aman"][0][1])})
	allArticlesData = append(allArticlesData, AllArticles{ArticleSite: userHandleList["aman"][1][0], ArticlesSiteData: GetMediumUserInfo(userHandleList["aman"][1][1])})

	c.IndentedJSON(http.StatusOK, allArticlesData)
}

func PostUserHandles(c *gin.Context) {
	var userHandlesData UserHandles
	if err := c.BindJSON(&userHandlesData); err != nil {
		return
	}
	AddUserHandleInDB(userHandlesData)
}

func PostUserInfo(c *gin.Context) {
	var userInfo UserInfo
	if err := c.BindJSON(&userInfo); err != nil {
		return
	}
	AddUserInfoInDB(userInfo)
}

func GetUserInfoData(c *gin.Context) {

}
