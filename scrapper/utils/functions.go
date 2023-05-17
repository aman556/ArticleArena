package utils

import "github.com/gocolly/colly"

func ArticleDataSinglePage(userPayload UserPayload) []Article {
	var artilceUtil []Article
	c := colly.NewCollector()

	c.OnHTML(userPayload.ParentQuery, func(e *colly.HTMLElement) {
		var localArticle Article
		localArticle.ArticleTitle = e.ChildText(userPayload.ChildTitleQuery)
		localArticle.ArtilceLink = e.Request.AbsoluteURL(e.ChildAttr(userPayload.ChildLinkQuery[0], userPayload.ChildLinkQuery[1]))

		artilceUtil = append(artilceUtil, localArticle)
	})

	c.Visit(userPayload.Url)
	return artilceUtil
}
