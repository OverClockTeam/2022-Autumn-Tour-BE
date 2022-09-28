package services

import (
	"github.com/yanyiwu/gojieba"
)

var SensitiveWords = []string{"sb", "SB", "Sb", "sB", "nt", "NT", "傻逼", "脑瘫"}

func isSensitive(content string) bool {
	for _, sensitiveWord := range SensitiveWords {
		if content == sensitiveWord {
			return true
		}
	}
	return false
}

func Audit(s string) (hasSensitiveWord bool, result []string) {
	j := gojieba.NewJieba()
	for _, sensitiveWord := range SensitiveWords {
		j.AddWord(sensitiveWord)
	}
	contents := j.CutAll(s)
	for _, content := range contents {
		if isSensitive(content) {
			hasSensitiveWord = true
			result = append(result, content)
		}
	}
	return
}
