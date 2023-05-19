package geeksforgeekshandler

import (
	"net/http"

	. "github.com/aman556/ArticleArena/scrapper/Geeksforgeeks"
	"github.com/gin-gonic/gin"
)

func GetGeeksforGeeksUserInfo(c *gin.Context) {

	c.IndentedJSON(http.StatusOK, GeeksForGeeksUserDataScarpping("https://auth.geeksforgeeks.org/user/aman55/articles#"))
}
