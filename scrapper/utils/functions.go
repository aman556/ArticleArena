package utils

import (
	"fmt"

	"github.com/gocolly/colly"
)

func ArticleDataSinglePage(userPayload UserPayload) []Article {
	var artilceUtil []Article
	c := colly.NewCollector()

	c.OnHTML(userPayload.ParentQuery, func(e *colly.HTMLElement) {
		var localArticle Article
		localArticle.ArticleTitle = e.ChildText(userPayload.ChildTitleQuery)
		localArticle.ArtilceLink = e.Request.AbsoluteURL(e.ChildAttr(userPayload.ChildLinkQuery[0], userPayload.ChildLinkQuery[1]))

		artilceUtil = append(artilceUtil, localArticle)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println(r.StatusCode)
	})
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})
	c.Visit(userPayload.Url)
	return artilceUtil
}
