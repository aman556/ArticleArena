package utils

type UserHandles struct {
	Name               string
	ArticleArenaHandle string
	UserHandleList     []UserHandle
}

type UserHandle struct {
	WebsiteName   string
	WebsiteHandle string
}

type UserInfo struct {
	Name               string
	ArticleArenaHandle string
	Email              string
	GithubUrl          string
	LinkedinUrl        string
}
