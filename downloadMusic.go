package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
)

type List struct {
	mrid string
	name string
}

func downloadAll(searchName string) {
	for i := 1; i < 10; i++ {
		musiclist := searchMusic(searchName, strconv.Itoa(i))
		if musiclist == nil {
			log.Printf("查询失败")
		}
		{
			for _, value := range musiclist {
				err, _ := downLoad(value)
				if err != nil {
					fmt.Println(err.Error())
				}
				_ = downloadlyrics(value)
			}
			log.Println("Page" + strconv.Itoa(i) + "完成")
		}
	}
}

func downloadOne(searchName string) {
	musiclist := searchMusic(searchName, "1")
	if musiclist == nil {
		log.Printf("查询失败")
	}
	{
		_, _ = downLoad(musiclist[0])
		_ = downloadlyrics(musiclist[0])
	}
}

func searchMusic(SearchName string, PageNum string) []List {
	//musicname->urlcode
	urlname := url.QueryEscape(SearchName)

	//urlcode->mrid
	searchUrl := "https://www.kuwo.cn/api/www/search/searchMusicBykeyWord?key=" + urlname + "&pn=" + PageNum + "&rn=30&httpsStatus=1&reqId=66186151-3989-11ed-8078-39fd02b59e03"
	req, _ := http.NewRequest("GET", searchUrl, nil)
	req.Header.Set("Cookie", "_ga=GA1.2.1737849527.1663585977; Hm_lvt_cdb524f42f0ce19b169a8071123a4797=1663585977,1663685364; _gid=GA1.2.5029194.1663685364; Hm_lpvt_cdb524f42f0ce19b169a8071123a4797=1663749389; kw_token=CXD5AR9O0Z5")
	req.Header.Set("csrf", "CXD5AR9O0Z5")
	req.Header.Set("Host", "www.kuwo.cn")
	req.Header.Set("Referer", "https://www.kuwo.cn/search/list?key="+urlname)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/105.0.0.0 Safari/537.36")
	firstRes, err := (&http.Client{}).Do(req)
	if err != nil {
		log.Printf(err.Error())
		return nil
	}
	var firstResBytes []byte
	firstResBytes, _ = ioutil.ReadAll(firstRes.Body)
	defer firstRes.Body.Close()
	firstResString := string(firstResBytes)

	var result map[string]interface{}
	err = json.Unmarshal([]byte(firstResString), &result)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	var list = result["data"].(map[string]interface{})["list"].([]interface{})
	var musicrid []List
	for _, value := range list {
		var d = value.(map[string]interface{})
		musicrid = append(musicrid, List{d["musicrid"].(string), d["name"].(string)})
	}
	return musicrid
}

func downLoad(music List) (err error, fpath string) {
	music.mrid = music.mrid[6:]
	urlname := url.QueryEscape(music.name)
	secondreq := "https://kuwo.cn/api/v1/www/music/playUrl?mid=" + music.mrid + "&type=music&httpsStatus=1&reqId=52f3c921-39ac-11ed-a443-91cfe5b56e50"
	req, _ := http.NewRequest("GET", secondreq, nil)
	req.Header.Set("Cookie", "_ga=GA1.2.1737849527.1663585977; Hm_lvt_cdb524f42f0ce19b169a8071123a4797=1663585977,1663685364; _gid=GA1.2.5029194.1663685364; Hm_lpvt_cdb524f42f0ce19b169a8071123a4797=1663749389; kw_token=CXD5AR9O0Z5")
	req.Header.Set("csrf", "CXD5AR9O0Z5")
	req.Header.Set("Host", "www.kuwo.cn")
	req.Header.Set("Referer", "https://www.kuwo.cn/search/list?key=%"+urlname)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/105.0.0.0 Safari/537.36")
	secondurl, err := (&http.Client{}).Do(req)
	if err != nil {
		log.Printf(err.Error())
		return err, ""
	}
	var downloadurlBytes []byte
	downloadurlBytes, _ = ioutil.ReadAll(secondurl.Body)
	defer secondurl.Body.Close()
	download := string(downloadurlBytes)
	var result map[string]interface{}
	err = json.Unmarshal([]byte(download), &result)
	if err != nil {
		fmt.Println(err.Error())
	}
	if result["data"] == nil {
		fmt.Println(music.name + "歌曲下载失败")
		return nil, ""
	}
	var downloadurl = result["data"].(map[string]interface{})["url"].(string)

	filePath := "C:/Users/fuyik/Music/downloadmusic/" + music.name + ".mp3"

	res, err := http.Get(downloadurl)
	if err != nil {
		return err, ""
	}
	f, err := os.Create(filePath)
	if err != nil {
		return err, ""
	}
	io.Copy(f, res.Body)
	log.Println(music.name + "歌曲下载成功！")
	defer f.Close()
	return err, filePath
}
