package main

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func downloadMusic(MusicName string) error {
	//musicname->urlcode
	urlname := url.QueryEscape(MusicName)

	//urlcode->mrid
	searchUrl := "https://www.kuwo.cn/api/www/search/searchMusicBykeyWord?key=" + urlname + "&pn=1&rn=30&httpsStatus=1&reqId=66186151-3989-11ed-8078-39fd02b59e03"
	req, _ := http.NewRequest("GET", searchUrl, nil)
	req.Header.Set("Cookie", "_ga=GA1.2.1737849527.1663585977; Hm_lvt_cdb524f42f0ce19b169a8071123a4797=1663585977,1663685364; _gid=GA1.2.5029194.1663685364; Hm_lpvt_cdb524f42f0ce19b169a8071123a4797=1663749389; kw_token=CXD5AR9O0Z5")
	req.Header.Set("csrf", "CXD5AR9O0Z5")
	req.Header.Set("Host", "www.kuwo.cn")
	req.Header.Set("Referer", "https://www.kuwo.cn/search/list?key="+urlname)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/105.0.0.0 Safari/537.36")
	firstRes, err := (&http.Client{}).Do(req)
	if err != nil {
		log.Printf(err.Error())
		return err
	}
	var firstResBytes []byte
	firstResBytes, _ = ioutil.ReadAll(firstRes.Body)
	defer firstRes.Body.Close()
	mridres := string(firstResBytes)
	if len(mridres) < 30 {
		fmt.Println("下载失败")
		return nil
	}
	start := strings.Index(mridres, "musicrid")
	mrid := mridres[start+17 : start+30]
	end := strings.Index(mrid, ",")
	if start == -1 || end == -1 {
		fmt.Println("下载失败")
		return nil
	}
	mrid = mrid[:end-1]

	//mrid->downloadpath
	filePath := "C:/Users/fuyik/Music/downloadmusic/" + MusicName + ".mp3"
	secondreq := "https://kuwo.cn/api/v1/www/music/playUrl?mid=" + mrid + "&type=music&httpsStatus=1&reqId=52f3c921-39ac-11ed-a443-91cfe5b56e50"
	req, _ = http.NewRequest("GET", secondreq, nil)
	req.Header.Set("Cookie", "_ga=GA1.2.1737849527.1663585977; Hm_lvt_cdb524f42f0ce19b169a8071123a4797=1663585977,1663685364; _gid=GA1.2.5029194.1663685364; Hm_lpvt_cdb524f42f0ce19b169a8071123a4797=1663749389; kw_token=CXD5AR9O0Z5")
	req.Header.Set("csrf", "CXD5AR9O0Z5")
	req.Header.Set("Host", "www.kuwo.cn")
	req.Header.Set("Referer", "https://www.kuwo.cn/search/list?key=%"+urlname)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/105.0.0.0 Safari/537.36")
	secondurl, err := (&http.Client{}).Do(req)
	var downloadurlBytes []byte
	downloadurlBytes, _ = ioutil.ReadAll(secondurl.Body)
	defer secondurl.Body.Close()
	download := string(downloadurlBytes)
	start = strings.Index(download, "url")
	end = strings.Index(download, "mp3")
	if start == -1 || end == -1 {
		fmt.Println("下载失败")
		return nil
	}
	//download
	downloadurl := download[start+6 : end+3]
	err = downLoad(downloadurl, filePath)
	if err != nil {
		log.Printf(err.Error())
	}
	return err
}

func downLoad(url string, filePath string) error {
	if url != "" && filePath != "" {
		res, err := http.Get(url)
		if err != nil {
			return err
		}
		f, err := os.Create(filePath)
		if err != nil {
			return err
		}
		io.Copy(f, res.Body)
		log.Println("下载成功！")
		defer f.Close()
		return nil
	} else {
		return errors.New("url or filePath is illegal")
	}
}
