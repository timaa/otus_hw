package main

import (
	"fmt"
	"github.com/timaa/otus_hw/hw-3/frequency"
)

func main() {
	text := `word1 word2 bla word2 , pwel han.
				bla ola ola Ola olba asgs sag han
			sagasg
			asgsag
			`
	fmt.Print(frequency.TopN(text, 10))
}
