package mediumScrapper

import (
	. "github.com/aman556/ArticleArena/backend/scrapper/utils"
)

var childLinkQuery = []string{`a[class="af ag ah ai aj ak al am an ao ap aq ar as at"]`, "href"}
var mediumPayload = ArticlePayload{
	Url:             "",
	ParentQuery:     `article`,
	ChildTitleQuery: `h2`,
	ChildLinkQuery:  childLinkQuery,
}

func MediumUserDataScarpping(url string) ArticleData {
	var articleInfo ArticleData
	mediumPayload.Url = url

	articleInfo.ArticleList = ArticleDataSinglePage(mediumPayload)
	articleInfo.ArticleCount = len(articleInfo.ArticleList)

	return articleInfo
}
