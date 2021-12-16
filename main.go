package main

import (
	"embed"
	"fmt"
	"io/fs"
	"math/rand"
	"strings"
	"time"
)

var marks = []string{"|", "/", "-", "\\"}

func mark(i int) string {
	return marks[i%4]
}

//go:embed topic.txt
var static embed.FS

func main() {
	cnt := 5000000000
	rand.Seed(time.Now().UnixNano())
	// file を読み取ります
	b1, err := fs.ReadFile(static, "topic.txt")
	if err != nil {
		panic(err)
	}
	arr01 := strings.Split(string(b1), "\n")
	for _, s := range arr01 {
		if s != "" {
			time.Sleep(time.Millisecond * 55)
			fmt.Printf(" \r Topics : %s\n", s)
		}
	}
	// 改行後の1行が含まれるので-1
	arr1_len := len(arr01) - 1
	// 文字列をoverwrite するためのスペース
	space_str := "                                                                                          "
	// Loading 描写のfor {}
	for i := 1; i <= cnt; i++ {
		if i%(cnt/100) == 0 {
			p := i / (cnt / 100)
			topic_title := string(arr01[rand.Intn(arr1_len)])
			time.Sleep(time.Millisecond * 13)
			fmt.Printf("\r Loading: %s %s %s", mark(p), topic_title, space_str)
		}
	}
	// 終了処理
	fmt.Printf("\r Loading: %s", space_str)
	time.Sleep(time.Millisecond * 600)
	emojiU := "\U0001f389\U0001f389\U0001f389"
	topic_title := string(arr01[rand.Intn(arr1_len)])
	fmt.Printf("\n %s : %s \n", emojiU, topic_title)
	time.Sleep(time.Millisecond * 1000)
	fmt.Println(" Done.  : repo https://github.com/nwiizo/SRETT2021SP")
}
