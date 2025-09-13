package main

import (
	"bytes"
	"fmt"

	"github.com/PuerkitoBio/goquery"
	"github.com/yznx04/crawler/collect"
)

// Fetch tag v0.0.4

// tag v0.0.3
func main() {
	url := "http://www.thepaper.cn"
	baseFetch := collect.BaseFetch{}
	body, err := baseFetch.Get(url)

	if err != nil {
		fmt.Println("read body, error:", err)
	}
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(body))
	doc.Find("a.index_inherit__A1ImK").Each(func(i int, s *goquery.Selection) {
		title := s.Text()
		fmt.Println("title:", title)
	})

}
