package unpack

import (
	"regexp"
	"strconv"
	"strings"
)

func Unpack(in string) string {
	re := regexp.MustCompile(`[0-9]`)
	indices := re.FindAllIndex([]byte(in), -1)
	var outSrt strings.Builder
	var runeSlice = []rune(in)

	if len(indices) == len(in) {
		return ""
	}

	if len(indices) > 0 {
		for i := 0; i < len(indices); i++ {
			if i == 0 {
				outSrt.WriteString(string(runeSlice[0 : indices[i][0]-1]))
				repeatCount, _ := strconv.Atoi(string(runeSlice[indices[i][0]]))
				outSrt.WriteString(strings.Repeat(string(runeSlice[indices[i][0]-1]), repeatCount))
			} else if i > 0 && i < len(indices)-1 {
				outSrt.WriteString(in[indices[i-1][0]+1 : indices[i][0]-1])
				repeatCount, _ := strconv.Atoi(string(runeSlice[indices[i][0]]))
				outSrt.WriteString(strings.Repeat(string(runeSlice[indices[i][0]-1]), repeatCount))
			} else if i == len(indices)-1 {
				outSrt.WriteString(string(runeSlice[indices[i-1][0]+1 : indices[i][0]-1]))
				repeatCount, _ := strconv.Atoi(string(runeSlice[indices[i][0]]))
				outSrt.WriteString(strings.Repeat(string(runeSlice[indices[i][0]-1]), repeatCount))
				outSrt.WriteString(string(runeSlice[indices[i][0]+1 :]))

			}
		}
		return outSrt.String()
	}
	return in
}
