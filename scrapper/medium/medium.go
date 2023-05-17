package maedium

import (
	. "github.com/aman556/ArticleArena/scrapper/utils"
)

var childLinkQuery = []string{`a[class="af ag ah ai aj ak al am an ao ap aq ar as at"]`, "href"}
var mediumPayload = UserPayload{
	Url:             "",
	ParentQuery:     `article`,
	ChildTitleQuery: `h2`,
	ChildLinkQuery:  childLinkQuery,
}

func MediumUserDataScarpping(url string) User {
	var userInfo User
	mediumPayload.Url = url

	userInfo.ArticleData = ArticleDataSinglePage(mediumPayload)
	userInfo.ArticleCount = len(userInfo.ArticleData)

	return userInfo
}
