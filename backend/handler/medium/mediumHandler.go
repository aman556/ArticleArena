package mediumhandler

import (
	"net/http"

	. "github.com/aman556/ArticleArena/backend/scrapper/medium"
	"github.com/gin-gonic/gin"
)

func GetMediumUserInfo(c *gin.Context) {

	c.IndentedJSON(http.StatusOK, MediumUserDataScarpping("https://medium.com/@ChindaVibhor"))
}
