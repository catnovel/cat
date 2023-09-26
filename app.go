package cat

import (
	"encoding/json"
	"errors"
	"github.com/catnovelapi/cat/cattemplate"
)

func (cat *CiweimaoApp) CatalogListApp(bookId string) ([]cattemplate.CatalogueInfoListTemplate, error) {
	var chapterListArray []cattemplate.CatalogueInfoListTemplate
	result := cat.api.CatalogByBookIdNewApi(bookId)
	if result.Get("code").String() != "100000" {
		return nil, errors.New("获取书籍目录失败,tips:" + result.Get("tip").String())
	}
	if len(result.Get("data.chapter_list").Array()) == 0 {
		return nil, errors.New("书籍目录为空, bookId:" + bookId)
	}
	err := json.Unmarshal([]byte(result.Get("data.chapter_list").String()), &chapterListArray)
	if err != nil {
		return nil, errors.New("书籍目录解析失败:" + err.Error())
	}
	return chapterListArray, nil
}

func (cat *CiweimaoApp) ChapterInfoApp(chapterId string) (string, error) {
	command := cat.api.ChapterCommandApi(chapterId)
	if command.Get("code").String() != "100000" {
		return "", errors.New("获取章节command失败,tips:" + command.Get("tip").String())
	}
	commandKey := command.Get("data.command").String()
	response := cat.api.ChapterInfoApi(chapterId, commandKey)
	if response.Get("code").String() != "100000" {
		return "", errors.New("获取章节内容失败,tips:" + response.Get("tip").String())
	}
	txtContent, err := cat.api.DecodeEncryptText(response.Get("data.chapter_info.txt_content").String(), commandKey)
	if err != nil {
		return "", err
	}
	return txtContent, nil
}

func (cat *CiweimaoApp) RetryChapterInfoApp(chapterId string) string {
	for i := 0; i < 5; i++ {
		content, err := cat.ChapterInfoApp(chapterId)
		if err == nil {
			return content
		}
	}
	return ""
}

func (cat *CiweimaoApp) NewSearchByKeywordApp(keyword string, page int) ([]cattemplate.BookInfoTemplate, error) {
	var bookList []cattemplate.BookInfoTemplate
	result := cat.api.SearchByKeywordApi(keyword, page, 0)
	if result.Get("code").String() != "100000" {
		return nil, errors.New("获取搜索书籍列表失败,tips:" + result.Get("tip").String())
	}
	if len(result.Get("data.book_list").Array()) == 0 {
		return nil, errors.New("搜索书籍列表为空, keyword:" + keyword)
	}
	err := json.Unmarshal([]byte(result.Get("data.book_list").String()), &bookList)
	if err != nil {
		return nil, errors.New("搜索书籍列表解析失败:" + err.Error())
	}
	return bookList, nil
}

func (cat *CiweimaoApp) BookInfoAPP(bookId string) (*cattemplate.BookInfoTemplate, error) {
	result := cat.api.BookInfoApi(bookId)
	if result.Get("code").String() != "100000" {
		return nil, errors.New("获取书籍信息失败,tips:" + result.Get("tip").String())
	}
	if result.Get("data.book_info.book_name").String() == "" {
		return nil, errors.New("书籍信息为空, bookId:" + bookId)
	}
	var bookInfo cattemplate.BookInfoTemplate
	err := json.Unmarshal([]byte(result.Get("data.book_info").String()), &bookInfo)
	if err != nil {
		return nil, errors.New("书籍信息解析失败:" + err.Error())
	}
	return &bookInfo, nil
}

func (cat *CiweimaoApp) UseGeetestApp(loginName string) bool {
	response := cat.api.UseGeetestInfoApi(loginName)
	if response.Get("data.need_use_geetest").String() == "0" {
		return true
	}
	return false
}

func (cat *CiweimaoApp) AutoRegV2App() (*cattemplate.AutoRegV2Template, error) {
	var regResult cattemplate.AutoRegV2Template
	response := cat.api.AutoRegV2Api()
	if response.Get("code").String() != "100000" {
		return nil, errors.New("AutoRegV2App error:" + response.Get("tip").String())
	}
	err := json.Unmarshal([]byte(response.Get("data").String()), &regResult)
	if err != nil {
		return nil, errors.New("AutoReg 解析失败:" + err.Error())
	}
	return &regResult, nil

}
