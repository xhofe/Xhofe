package lib

import (
	"encoding/json"
	"fmt"
	"github.com/Xhofe/Xhofe/utils"
	"io/ioutil"
	"net/http"
	"strings"
)

type ZhihuResp struct {
	Data []struct {
		Target struct {
			Title string `json:"title"`
			Url   string `json:"url"`
		} `json:"target"`
	} `json:"data"`
}

func Zhihu(url string) string {
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.93 Safari/537.36")
	req.Header.Set("Referer", "https://www.zhihu.com/hot")
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
	var body []byte
	if body, err = ioutil.ReadAll(resp.Body); err != nil {
		return ""
	}
	var zhihuResp ZhihuResp
	if err = json.Unmarshal(body, &zhihuResp); err != nil {
		return ""
	}
	var topics []string
	var urls []string
	for i, item := range zhihuResp.Data {
		topic := item.Target.Title
		topics = append(topics, topic)
		temp := strings.Split(item.Target.Url, "/")
		href := fmt.Sprintf("https://www.zhihu.com/question/%s", temp[len(temp)-1])
		urls = append(urls, href)
		fmt.Println(i, topic, href)
	}
	return utils.GenerateContent(topics,urls)
}
