package main

import (
	"github.com/beevik/ntp"
	"fmt"
)

func main() {
	time, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		fmt.Errorf("%s",err)
	}
	fmt.Print(time)
}
