package main

import (
	"github.com/Xhofe/Xhofe/lib"
	"github.com/Xhofe/Xhofe/utils"
)

func main() {
	zhihu := lib.Zhihu("https://www.zhihu.com/api/v3/feed/topstory/hot-lists/total?limit=50&desktop=true")
	bilibili := lib.Bilibili("https://www.bilibili.com/v/popular/rank/all")
	readme := utils.GenerateReadme([]string{
		"zhihu",
		"bilibili",
	}, zhihu,bilibili)
	utils.WriteStringToFile(readme)
}
