package geeksforgeeksScrapper

import (
	. "github.com/aman556/ArticleArena/backend/scrapper/utils"
)

var childLinkQuery = []string{`a`, "href"}
var geeksforgeeksPayload = UserPayload{
	Url:             "",
	ParentQuery:     `div[class="card-content black-text"]`,
	ChildTitleQuery: `a`,
	ChildLinkQuery:  childLinkQuery,
}

func GeeksForGeeksUserArticleDataScarpping(url string) ArticleData {
	var articleInfo ArticleData
	geeksforgeeksPayload.Url = url

	articleInfo.ArticleList = ArticleDataSinglePage(geeksforgeeksPayload)
	articleInfo.ArticleCount = len(articleInfo.ArticleList)

	return articleInfo
}
