package main

import (
	"fmt"

	. "github.com/aman556/ArticleArena/scrapper/Geeksforgeeks"
	. "github.com/aman556/ArticleArena/scrapper/medium"
)

func main() {
	geeksforgeeksUserInfo := GeeksForGeeksUserDataScarpping("https://auth.geeksforgeeks.org/user/aman55/articles#")

	fmt.Println(geeksforgeeksUserInfo.ArticleCount)
	fmt.Println(geeksforgeeksUserInfo.ArticleData)

	mediumUserInfo := MediumUserDataScarpping("https://medium.com/@ChindaVibhor")

	fmt.Println(mediumUserInfo.ArticleCount)
	fmt.Println(mediumUserInfo.ArticleData)
}
