package utils

type Article struct {
	ArticleTitle string
	ArtilceLink  string
	//articleLikeCount  int
	//articleViewsCount int
}

type User struct {
	//userName     string
	ArticleCount int
	ArticleData  []Article
}

type UserPayload struct {
	Url             string
	ParentQuery     string
	ChildTitleQuery string
	ChildLinkQuery  []string
}
