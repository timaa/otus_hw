package frequency

import (
	"regexp"
	"sort"
	"strings"
)

// TopN function   comments
func TopN(text string, limit int) []string {
	words := split(text)
	freqMap := make(map[string]int, 100)
	for _, word := range words {
		freqMap[strings.ToLower(word)]++
	}

	freqSlice := mapToSliceStruct(freqMap)

	sort.Slice(freqSlice, func(i, j int) bool {
		return freqSlice[i].freq > freqSlice[j].freq
	})

	if limit > len(freqSlice) {
		limit = len(freqSlice)
	}

	words = make([]string, limit, limit)
	for i := 0; i < limit && i < len(freqSlice); i++ {
		words[i] = freqSlice[i].word
	}

	return words
}

func split(text string) []string {
	re := regexp.MustCompile(`\w+`)
	words := re.FindAllString(text, -1)
	return words
}

func mapToSliceStruct(freqMap map[string]int) []freqStruct {
	var freqSlice = make([]freqStruct, 0, len(freqMap))
	for key, val := range freqMap {
		freqSlice = append(freqSlice, freqStruct{val, key})
	}

	return freqSlice
}

type freqStruct struct {
	freq int
	word string
}
