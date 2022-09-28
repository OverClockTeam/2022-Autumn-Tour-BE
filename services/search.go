package services

import (
	"github.com/yanyiwu/gojieba"
	"strings"
)

func CutStr(s string) []string {
	j := gojieba.NewJieba()
	return j.CutForSearch(s, true)
}

func IsInSearchList(s string, searchList []string) bool {
	for _, str := range searchList {
		if strings.Contains(s, str) {
			return true
		}
	}
	return false
}
