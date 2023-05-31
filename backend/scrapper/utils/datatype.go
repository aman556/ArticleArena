package utils

type ArticleList struct {
	ArticleTitle string
	ArtilceLink  string
	//articleLikeCount  int
	//articleViewsCount int
}

type ArticleData struct {
	ArticleCount int
	ArticleList  []ArticleList
}

type ArticlePayload struct {
	Url             string
	ParentQuery     string
	ChildTitleQuery string
	ChildLinkQuery  []string
}

type AllArticles struct {
	ArticleSite      string
	ArticlesSiteData ArticleData
}
