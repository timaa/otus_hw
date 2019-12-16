package main

import (
	"fmt"
	"github.com/timaa/otus_hw/hw-3/frequency"
)

func main() {
	text := `Word1 word2 bla Word2 , pwel han.
				bla ola ola Ola olba asgs sag han
			Word1
			word1
			`
	fmt.Print(frequency.TopN(text, 10))
}
