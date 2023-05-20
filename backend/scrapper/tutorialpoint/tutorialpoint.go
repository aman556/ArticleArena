package main

import (
	"fmt"
	"strconv"

	. "github.com/aman556/ArticleArena/backend/scrapper/utils"

	"github.com/gocolly/colly"
)

func main() {

	var userInfo User
	c := colly.NewCollector()

	c.OnHTML(`div[class="col-sm-12 col-md-8 col-xl-6 rounded-3 qa-content"]`, func(e *colly.HTMLElement) {
		//userInfo.UserName = e.ChildText(`span[class="qa_count"]:nth-child(1)`)
		userInfo.ArticleCount, _ = strconv.Atoi(e.ChildText(`span[class="qa_count"]:nth-child(2)`))
		var localArticle Article
		localArticle.ArticleTitle = e.Request.AbsoluteURL(e.ChildText(`a[target="_blank"]`))
		localArticle.ArtilceLink = e.Request.AbsoluteURL(e.ChildAttr(`a[target="_blank"]`, "href"))

		userInfo.ArticleData = append(userInfo.ArticleData, localArticle)

	})

	c.Visit("https://www.tutorialspoint.com/authors/aman-sharma-166063070448")

	for _, a := range userInfo.ArticleData {
		fmt.Println(a.ArticleTitle)
		fmt.Println()
		fmt.Println(a.ArtilceLink)
	}

}
