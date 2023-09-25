package cat

import (
	"errors"
	"github.com/tidwall/gjson"
)

func (cat *Ciweimao) CatalogListApp(bookID string) []gjson.Result {
	var chapterListArray []gjson.Result
	for _, division := range cat.CatalogByBookIdNewApi(bookID).Get("data.chapter_list").Array() {
		for _, chapter := range division.Get("chapter_list").Array() {
			chapterListArray = append(chapterListArray, chapter)
		}
	}
	return chapterListArray
}

func (cat *Ciweimao) ChapterInfoApp(chapterId string) (string, error) {
	command := cat.ChapterCommandApi(chapterId)
	if command.Get("code").String() != "100000" {
		return "", errors.New("获取章节command失败,tips:" + command.Get("tip").String())
	}
	response := cat.ChapterInfoApi(chapterId, command.Get("data.command").String())
	if response.Get("code").String() != "100000" {
		return "", errors.New("获取章节内容失败,tips:" + response.Get("tip").String())
	}
	txtContent := DecodeText(response.Get("data.chapter_info.txt_content").String(), command.Get("data.command").String())
	if txtContent == "" {
		return "", errors.New("章节内容为空")
	}
	return txtContent, nil
}

func (cat *Ciweimao) UseGeetestApp(loginName string) bool {
	response := cat.UseGeetestInfoApi(loginName)
	if response.Get("data.need_use_geetest").String() == "0" {
		return true
	}
	return false
}
