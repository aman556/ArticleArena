package geeksforgeekshandler

import (
	. "github.com/aman556/ArticleArena/backend/scrapper/Geeksforgeeks"
	. "github.com/aman556/ArticleArena/backend/scrapper/utils"
)

func GetGeeksforGeeksUserInfo(gfghandle string) ArticleData {
	url := "https://auth.geeksforgeeks.org/user/" + gfghandle + "/articles#"
	return GeeksForGeeksUserArticleDataScarpping(url)
}
