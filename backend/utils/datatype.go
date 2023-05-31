package utils

type User struct {
	Name               string
	ArticleArenaHandle string
	UserHandleList     []UserHandle
}

type UserHandle struct {
	WebsiteName   string
	WebsiteHandle string
}
