package geeksforgeeksScrapper

import (
	. "github.com/aman556/ArticleArena/scrapper/utils"
)

var childLinkQuery = []string{`a`, "href"}
var geeksforgeeksPayload = UserPayload{
	Url:             "",
	ParentQuery:     `div[class="card-content black-text"]`,
	ChildTitleQuery: `a`,
	ChildLinkQuery:  childLinkQuery,
}

func GeeksForGeeksUserDataScarpping(url string) User {
	var userInfo User
	geeksforgeeksPayload.Url = url

	userInfo.ArticleData = ArticleDataSinglePage(geeksforgeeksPayload)
	userInfo.ArticleCount = len(userInfo.ArticleData)

	return userInfo
}
