package main

import (
	"flag"
	"fmt"
	"github.com/timaa/otus_hw/hw-6/copy"
	"log"
)


var From, To string
var Offset, Limit int

func init() {
	flag.StringVar(&From, "From", "", "file dest")
	flag.StringVar(&To, "To", "", "file src")
	flag.IntVar(&Offset, "Offset", 0, "offset dest file")
	flag.IntVar(&Limit, "Limit", 0, "Limit dest file")
}
func main() {
	flag.Parse()
	fmt.Printf("From-%s, To-%s Offset-%d, Limit-%d ", From, To, Offset, Limit)

	err := copy.Copy(From, To, Limit, Offset)
	if err != nil {
		log.Panic(err)
	}
}
