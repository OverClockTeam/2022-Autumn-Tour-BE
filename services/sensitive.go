package services

import (
	"github.com/importcjj/sensitive"
)

//var SensitiveWords = []string{"sb", "SB", "Sb", "sB", "nt", "NT", "傻逼", "脑瘫"}

func Audit(s string) (hasSensitiveWord bool, result []string) {
	filter := sensitive.New()
	filter.LoadNetWordDict("https://raw.githubusercontent.com/importcjj/sensitive/master/dict/dict.txt")
	result = filter.FindAll(s)
	hasSensitiveWord = result != nil
	return
}
