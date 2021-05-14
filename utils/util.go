package utils

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func WriteStringToFile(text string) {
	// write the whole body at once
	err := ioutil.WriteFile("README.md", []byte(text), 0644)
	if err != nil {
		panic(err)
	}
}

func GenerateContent(topics []string, urls []string) string {
	readme := ""
	for index, topic := range topics {
		// n. XXX
		temp := fmt.Sprintf("%d. %s", index+1, topic)
		topic = strings.Replace(topic, " ", "%20", -1)
		temp += " [:link:](" + urls[index] + ")\n"

		if index == 10 {
			temp = fmt.Sprintf("<details>\n<summary>%d ~ %d</summary>\n\n%s", index+1, Min(index+10, len(topics)), temp)
		} else if index == len(topics)-1 {
			temp = fmt.Sprintf("%s</details>", temp)
		} else if index >= 11 && index%10 == 0 {
			temp = fmt.Sprintf("</details>\n<details>\n<summary>%d ~ %d</summary>\n\n%s", index+1, Min(index+10, len(topics)), temp)
		}

		readme += temp
	}
	return readme
}

func GenerateReadme(titles []string, contents ...string) string {
	readme := ""
	readme += contents[0]
	for i := 1; i < len(titles); i++ {
		readme += fmt.Sprintf("<details>\n<summary>%s</summary>\n\n%s</details>", titles[i], contents[i])
	}
	return readme
}
