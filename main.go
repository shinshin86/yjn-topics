package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/PuerkitoBio/goquery"
	. "github.com/logrusorgru/aurora"
	"github.com/manifoldco/promptui"
)

func flagUsage() {
	usageText := `CLI tool of fetch Yahoo Japan News Topics.
	
	Usage:
	yjn-topics [select or help (optional arguments)]
	
	The commands are:
	no option: Display all category news
	select: Display specific category news
	help: Display usage of this tool`

	fmt.Fprintf(os.Stderr, "%s\n\n", usageText)
}

func main() {
	flag.Usage = flagUsage

	if len(os.Args) == 1 {
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
			news(url)

			time.Sleep(1 * time.Second)
		}
	} else if len(os.Args) == 2 && os.Args[1] == "select" {
		prompt := promptui.Select{
			Label: "Select News Category",
			Items: []string{"主要", "国内", "国際", "経済", "エンタメ", "スポーツ", "IT", "科学", "地域"},
		}

		_, category, err := prompt.Run()

		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			os.Exit(1)
		}

		url, err := specificNewsURL(category)

		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		news(url)
	} else if len(os.Args) == 2 && os.Args[1] == "help" {
		flag.Usage()
	} else {
		fmt.Println("Invalid option")
		fmt.Println("==============")
		flag.Usage()
		os.Exit(1)
	}
}

func specificNewsURL(category string) (url string, err error) {
	switch category {
	case "主要":
		return "https://news.yahoo.co.jp/", nil
	case "国内":
		return "https://news.yahoo.co.jp/categories/domestic", nil
	case "国際":
		return "https://news.yahoo.co.jp/categories/world", nil
	case "経済":
		return "https://news.yahoo.co.jp/categories/business", nil
	case "エンタメ":
		return "https://news.yahoo.co.jp/categories/entertainment", nil
	case "スポーツ":
		return "https://news.yahoo.co.jp/categories/sports", nil
	case "IT":
		return "https://news.yahoo.co.jp/categories/it", nil
	case "科学":
		return "https://news.yahoo.co.jp/categories/science", nil
	case "地域":
		return "https://news.yahoo.co.jp/categories/local", nil
	default:
		return "", errors.New("Invalid category")
	}
}

func news(url string) {
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
}
