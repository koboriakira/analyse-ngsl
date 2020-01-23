package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"regexp"
	"strings"
	"strconv"
	"encoding/json"
)

type OriginalWord struct {
	word string
	rank int
}

type Words struct {
	words []Word `json:"words"`
}

type Word struct {
	word string `json:"word"`
	original_word string `json:"originalWord"`
	rank int `json:"rank"`
}


func main() {
	// ファイルの扱い
	file, err := os.Open("./ngsl.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	var line []string

	// 単語ごとのMap
	ngsl_map := make(map[string]OriginalWord, 2801)

	// 無限ループでCSVを1行ずつ読み込んで処理
	for {
		line, err = reader.Read()
		if err != nil {
			fmt.Println(err)
			break
		}

		original_word := line[0]
		rank, _ := strconv.Atoi(line[1])

		ngsl_map[original_word] = OriginalWord{word: original_word, rank: rank}
		for _, s := range line[2:] {
			if s == "" {
				continue
			}

			ngsl_map[s] = OriginalWord{word: original_word, rank: rank}
		}
	}

	// 分析対象となる英文
  paragraph := strings.ToLower("Today, Google launched three new experimental apps to help you use your phone less as part of its Digital Wellbeing initiative, including one that actually has you seal up your phone in a phone-sized paper envelope (via Android Police). It sounds similar to the pouches some artists require fans to put their phones into at concerts, except it’s something you make at home — and Google’s envelope should at least let you make a call, if you need to.")

	// すべてのカンマなどを削除
	rep := regexp.MustCompile(`[,().]`)
	removedParagraph := rep.ReplaceAllString(paragraph, "")

	// 半角スペースで分割して、単語を抽出
	slicedParagraph := strings.Split(removedParagraph, " ")

	// 単語Mapと比較してみる
	var words []Word
	for _, word := range slicedParagraph {
		if original_word, exists := ngsl_map[word]; exists {
			word := Word{
				word: word,
				original_word: original_word.word,
				rank: original_word.rank,}
			words = append(words, word)
		} else {
			word := Word{
				word: word,
				original_word: "",
				rank: -1}
			words = append(words, word)
		}
	}

	fmt.Println(words)

	sample_words := Words{words:words}
	fmt.Println(sample_words)
	sample_json, err := json.Marshal(sample_words)
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Println(sample_json)

	fmt.Println(sample_json)
	fmt.Println(string(sample_json))
}
