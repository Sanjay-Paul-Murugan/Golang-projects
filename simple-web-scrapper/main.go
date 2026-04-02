package main

import (
	"fmt"

	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	res, err := http.Get("https://books.toscrape.com/")

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		panic("server not available")
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		panic(err)
	}

	title := doc.Find("title").Text()
	fmt.Println(title)

	doc.Find("article.product_pod h3 a").Each(func(i int, s *goquery.Selection) {
		title, exists := s.Attr("title")

		if exists {
			fmt.Println(title)
		}
	})
}
