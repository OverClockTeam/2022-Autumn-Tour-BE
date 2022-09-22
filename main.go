package main

import (
	"math/rand"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func main() {
	//for i := 0; i < 999; i++ {
	//	res := RandStringRunes(2)
	//	fmt.Println(res)
	//	_ = downloadMusic(res)
	//}
	_ = downloadMusic("麻雀")
}
