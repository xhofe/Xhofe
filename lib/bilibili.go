package lib

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/Xhofe/Xhofe/utils"
	"net/http"
)

func Bilibili(url string) string {
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.93 Safari/537.36")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("http get error", err)
		return ""
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	if resp.StatusCode != 200 {
		fmt.Printf("status code error: %d %s\n", resp.StatusCode, resp.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	var topics []string
	var urls []string
	// Find the review items
	doc.Find(".rank-list-wrap > .rank-list > .rank-item").Each(func(i int, selection *goquery.Selection) {
		a := selection.Find(".info > a")
		topic := a.Text()
		href, _ := a.Attr("href")
		fmt.Println(i+1, topic, href)
		topics = append(topics, topic)
		urls = append(urls, href)
	})
	return utils.GenerateContent(topics, urls)
}
