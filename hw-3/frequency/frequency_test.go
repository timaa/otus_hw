package frequency

import (
	"reflect"
	"testing"
)

func TestTopN(t *testing.T) {
	testText := "word1 word2 word1 word3  word1 word2"
	correctSlice := []string{"word1", "word2", "word3"}
	resultedSlice := TopN(testText, 3)
	if !reflect.DeepEqual(correctSlice, resultedSlice) {
		t.Errorf("Got %s, expected %s ", resultedSlice, correctSlice)
	}
}
