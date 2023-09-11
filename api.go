package cat

import (
	"github.com/tidwall/gjson"
	"log"
	"strconv"
)

func (cat *Ciweimao) AccountInfoApi() gjson.Result {
	return cat.post("/reader/get_my_info", nil)
}

func (cat *Ciweimao) CatalogByBookIDApi(bookID string) gjson.Result {
	return cat.post("/chapter/get_updated_chapter_by_division_new", map[string]string{"book_id": bookID})
}

func (cat *Ciweimao) NewCatalogByBookIDApi(bookID string) []gjson.Result {
	var chapterListArray []gjson.Result
	for _, division := range cat.CatalogByBookIDApi(bookID).Get("data.chapter_list").Array() {
		for _, chapter := range division.Get("chapter_list").Array() {
			chapterListArray = append(chapterListArray, chapter)
		}
	}
	return chapterListArray
}

func (cat *Ciweimao) BookInfoApi(bookId string) gjson.Result {
	return cat.post("/book/get_info_by_id", map[string]string{"book_id": bookId})
}

func (cat *Ciweimao) SearchByKeywordApi(keyword string, page int) gjson.Result {
	return cat.post("/bookcity/get_filter_search_book_list", map[string]string{"count": "10", "page": strconv.Itoa(page), "category_index": "0", "key": keyword})
}

func (cat *Ciweimao) SignupApi(account, password string) (string, string) {
	response := cat.post("/signup/login", map[string]string{"login_name": account, "passwd": password})
	if response.Get("code").String() != "100000" {
		log.Println("登录失败,tips:" + response.Get("tip").String())
	}
	return response.Get("data.reader_info.account").String(), response.Get("data.login_token").String()
}
func (cat *Ciweimao) ChapterCommandApi(chapterId string) gjson.Result {
	return cat.post("/chapter/get_chapter_command", map[string]string{"chapter_id": chapterId})
}

func (cat *Ciweimao) ChapterInfoApi(chapterId string, command string) string {
	response := cat.post("/chapter/get_cpt_ifm", map[string]string{"chapter_id": chapterId, "chapter_command": command})
	if response.Get("data.chapter_info.txt_content").String() != "" {
		return DecodeText(response.Get("data.chapter_info.txt_content").String(), command)
	}
	log.Println("chapterId:", chapterId, "获取失败,error:", response.Get("tip").String())
	return ""
}

func (cat *Ciweimao) useGeetestApi(loginName string) bool {
	response := cat.post("/signup/use_geetest", map[string]string{"login_name": loginName}, NoDecode())
	if response.Get("data.need_use_geetest").String() == "0" {
		return true
	}
	return false
}

func (cat *Ciweimao) BookShelfIdListApi() gjson.Result {
	return cat.post("/bookshelf/get_shelf_list", nil)

}
func (cat *Ciweimao) BookShelfListApi(shelfId string) gjson.Result {
	return cat.post("/bookshelf/get_shelf_book_list", map[string]string{"shelf_id": shelfId, "last_mod_time": "0", "direction": "prev"})
}
func (cat *Ciweimao) BookmarkListApi(bookID string, page string) gjson.Result {
	return cat.post("/book/get_bookmark_list", map[string]string{"count": "10", "book_id": bookID, "page": page})
}

func (cat *Ciweimao) DivisionListApi(bookID string) gjson.Result {
	return cat.post("/book/get_division_list", map[string]string{"book_id": bookID})
}

func (cat *Ciweimao) TsukkomiNumApi(chapterID string) gjson.Result {
	return cat.post("/chapter/get_tsukkomi_num", map[string]string{"chapter_id": chapterID})
}

func (cat *Ciweimao) BdaudioInfoApi(bookID string) gjson.Result {
	return cat.post("/reader/bdaudio_info", map[string]string{"book_id": bookID})
}

func (cat *Ciweimao) AddReadbookApi(bookID string, readTimes string, getTime string) gjson.Result {
	return cat.post("/reader/add_readbook", map[string]string{"book_id": bookID, "readTimes": readTimes, "getTime": getTime})
}

func (cat *Ciweimao) SetLastReadChapterApi(lastReadChapterID string, bookID string) gjson.Result {
	return cat.post("/bookshelf/set_last_read_chapter", map[string]string{"last_read_chapter_id": lastReadChapterID, "book_id": bookID})
}
func (cat *Ciweimao) PostPrivacyPolicyVersionApi() gjson.Result {
	return cat.post("/setting/privacy_policy_version", map[string]string{"privacy_policy_version": "1"})
}

func (cat *Ciweimao) PostPropInfoApi() gjson.Result {
	return cat.post("/reader/get_prop_info", nil)
}

func (cat *Ciweimao) MetaDataApi() gjson.Result {
	return cat.post("/meta/get_meta_data", nil)
}

func (cat *Ciweimao) VersionApi() gjson.Result {
	return cat.post("/setting/get_version", nil)
}

func (cat *Ciweimao) StartpageUrlListApi() gjson.Result {
	return cat.post("/setting/get_startpage_url_list", nil)
}

func (cat *Ciweimao) ThirdPartySwitchApi() gjson.Result {
	return cat.post("/setting/thired_party_switch", nil)
}
