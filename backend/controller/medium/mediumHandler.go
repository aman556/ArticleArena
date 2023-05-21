package mediumhandler

import (
	. "github.com/aman556/ArticleArena/backend/scrapper/medium"
	. "github.com/aman556/ArticleArena/backend/scrapper/utils"
)

func GetMediumUserInfo(mediumHandle string) ArticleData {
	url := "https://medium.com/@" + mediumHandle
	return MediumUserDataScarpping(url)
}
