package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Lyric struct {
	lineLyric string
	time      string
}

func downloadlyrics(mrid string, filepath string) error {
	lyricsUrl := "http://m.kuwo.cn/newh5/singles/songinfoandlrc?musicId=" + mrid + "&httpsStatus=1&reqId=5d3b3540-3a41-11ed-8e54-774f6359d4c7"
	req, _ := http.NewRequest("GET", lyricsUrl, nil)
	req.Header.Set("Cookie", "_ga=GA1.2.1737849527.1663585977; Hm_lvt_cdb524f42f0ce19b169a8071123a4797=1663585977,1663685364; _gid=GA1.2.5029194.1663685364; Hm_lpvt_cdb524f42f0ce19b169a8071123a4797=1663749389; kw_token=CXD5AR9O0Z5")
	req.Header.Set("csrf", "CXD5AR9O0Z5")
	req.Header.Set("Host", "www.kuwo.cn")
	//req.Header.Set("Referer", "https://www.kuwo.cn/search/list?key="+urlname)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/105.0.0.0 Safari/537.36")
	//firstRes, err := (&http.Client{}).Do(req)
	resp, err := http.Get(lyricsUrl)
	if err != nil {
		log.Printf(err.Error())
		return err
	}

	var firstResBytes []byte
	firstResBytes, _ = ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	mridres := string(firstResBytes)
	var result map[string]interface{}
	err = json.Unmarshal([]byte(mridres), &result)
	if err != nil {
		fmt.Println(err.Error())
	}
	var lrc = result["data"].(map[string]interface{})["lrclist"].([]interface{})
	var lyrics []Lyric

	for _, value := range lrc {
		var d = value.(map[string]interface{})
		lyrics = append(lyrics, Lyric{d["lineLyric"].(string), d["time"].(string)})
	}

	fmt.Println(lyrics)

	return nil
}
