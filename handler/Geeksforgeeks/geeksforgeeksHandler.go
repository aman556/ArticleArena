package geeksforgeekshandler

import (
	"net/http"

	. "github.com/aman556/ArticleArena/scrapper/Geeksforgeeks"
	"github.com/gin-gonic/gin"
)

func GetGeeksforGeeksUserInfo(c *gin.Context) {
	url := "https://auth.geeksforgeeks.org/user/" + c.Param("userid") + "/articles#"
	c.IndentedJSON(http.StatusOK, GeeksForGeeksUserDataScarpping(url))
}
