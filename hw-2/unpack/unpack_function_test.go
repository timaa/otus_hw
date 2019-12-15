package unpack

import "testing"

func TestUnpack(t *testing.T) {
	testData := map[string]string{
		"a4bc2d5e": "aaaabccddddde",
		"abcd":     "abcd",
		"45":       "",
	}

	for key, value := range testData {
		resultData := Unpack(key)
		if value != resultData {
			t.Errorf("Got %s, expected %s ", resultData, value)
		}
	}
}
