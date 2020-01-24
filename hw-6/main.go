package main

import (
	"flag"
	"github.com/timaa/otus_hw/hw-6/copy"
	"log"
)

var From, To string
var Offset, Limit int64

func init() {
	flag.StringVar(&From, "From", "", "file dest")
	flag.StringVar(&To, "To", "", "file src")
	flag.Int64Var(&Offset, "Offset", 0, "offset dest file")
	flag.Int64Var(&Limit, "Limit", -1, "Limit dest file")
}
func main() {
	flag.Parse()
	err := copy.Copy(From, To, Limit, Offset)
	if err != nil {
		log.Panic(err)
	}
}
