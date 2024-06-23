package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

func Top10(input string) []string {
	words := strings.Fields(input)

	type WordsCount struct {
		Word  string
		Count int
	}

	wordsCounter := make(map[string]int)
	for i := range words {
		wordsCounter[words[i]]++
	}

	wordsArr := make([]WordsCount, 0, len(wordsCounter))
	for w, c := range wordsCounter {
		wordsArr = append(wordsArr, WordsCount{Word: w, Count: c})
	}

	sort.Slice(wordsArr, func(i, j int) bool {
		if wordsArr[i].Count == wordsArr[j].Count {
			return wordsArr[i].Word < wordsArr[j].Word
		}
		return wordsArr[i].Count > wordsArr[j].Count
	})
	size := 10
	if len(wordsArr) < size {
		size = len(wordsArr)
	}

	top := make([]string, 0, size)
	for i := range wordsArr[:size] {
		top = append(top, wordsArr[i].Word)
	}

	return top
}
