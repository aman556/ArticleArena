package controller

import (
	"net/http"

	. "github.com/aman556/ArticleArena/backend/controller/Geeksforgeeks"
	. "github.com/aman556/ArticleArena/backend/controller/medium"
	. "github.com/aman556/ArticleArena/backend/scrapper/utils"
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
