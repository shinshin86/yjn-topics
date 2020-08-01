package main

import (
	"fmt"
	"time"

	"github.com/PuerkitoBio/goquery"
	. "github.com/logrusorgru/aurora"
)

func main() {
	urls := [9]string{"https://news.yahoo.co.jp/",
		"https://news.yahoo.co.jp/categories/domestic",
		"https://news.yahoo.co.jp/categories/world",
		"https://news.yahoo.co.jp/categories/business",
		"https://news.yahoo.co.jp/categories/entertainment",
		"https://news.yahoo.co.jp/categories/sports",
		"https://news.yahoo.co.jp/categories/it",
		"https://news.yahoo.co.jp/categories/science",
		"https://news.yahoo.co.jp/categories/local"}

	for _, url := range urls {
		doc, err := goquery.NewDocument(url)

		if err != nil {
			panic(err)
		}

		current := doc.Find("ul.yjnHeader_sub_cat > li.current")
		title := Bold(BrightYellow("=== " + current.Text() + " ==="))
		fmt.Println(title)

		selection := doc.Find(".topicsListItem")
		selection.Each(func(index int, s *goquery.Selection) {
			fmt.Println(Bold(BrightCyan(s.Text())))

			a := s.Find("a")
			attr, _ := a.Attr("href")
			fmt.Println(attr)
		})

		time.Sleep(1 * time.Second)
	}
}
