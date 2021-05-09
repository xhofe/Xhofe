package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func getHotTopic(url string) ([]string, []string, error) {
	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.93 Safari/537.36")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("http get error", err)
		return nil, nil, err
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
		topics = append(topics, topic)
		urls = append(urls, href)
	})
	return topics, urls, nil
}

func generateREADME(topics []string, urls []string) {
	readme := ""
	for index, topic := range topics {
		// n. XXX
		temp := fmt.Sprintf("%d. %s", index+1, topic)
		topic = strings.Replace(topic, " ", "%20", -1)
		temp += "[:link:](" + urls[index] + ")\n"

		if index == 10 {
			temp = fmt.Sprintf("<details>\n<summary>%d ~ %d</summary>\n\n%s", index+1, Min(index+10, len(topics)), temp)
		} else if index == len(topics)-1 {
			temp = fmt.Sprintf("%s</details>", temp)
		} else if index >= 11 && index%10 == 0 {
			temp = fmt.Sprintf("</details>\n<details>\n<summary>%d ~ %d</summary>\n\n%s", index+1, Min(index+10, len(topics)), temp)
		}

		readme += temp
	}

	writeStringToFile(readme)
}

func writeStringToFile(text string) {
	// write the whole body at once
	err := ioutil.WriteFile("README.md", []byte(text), 0644)
	if err != nil {
		panic(err)
	}
}

func main() {
	topics, urls, err := getHotTopic("https://www.bilibili.com/v/popular/rank/all")
	if err != nil {
		return
	}
	generateREADME(topics, urls)
}
