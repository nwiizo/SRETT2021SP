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
	b1, err := fs.ReadFile(static, "topic.txt")
	if err != nil {
		panic(err)
	}
	arr01 := strings.Split(string(b1), "\n")
	for _, s := range arr01 {
		if s != "" {
			fmt.Printf(" \r TOPICS: %s\n", s)
		}
	}
	for i := 1; i <= cnt; i++ {
		if i%(cnt/100) == 0 {
			p := i / (cnt / 100)
			fmt.Printf("\r Loading: %s ", mark(p))
		}
	}
	arr1_len := len(arr01) - 1
	topic_title := string(arr01[rand.Intn(arr1_len)])
	fmt.Println("\n 🏁",topic_title)
	fmt.Println(" Done.")
}
