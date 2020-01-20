package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	filePath := "/Users/timur/go/src/github.com/timaa/otus_hw/hw-6/test.txt"
	filePath2 := "/Users/timur/go/src/github.com/timaa/otus_hw/hw-6/test2.txt"
	N := 100
	var file *os.File
	buff := make([]byte, N)

	file, err:= os.Open(filePath)
	file2, _ := os.Create(filePath2)

	if err !=nil {
		fmt.Println("%v", err)
	}
	err = nil

	for ;; {

		read, err := io.ReadFull(file, buff)
		fmt.Println("%v", read)
		fmt.Println("%s", string(buff[:read]))
		written, err := file2.Write(buff[:read])
		fmt.Println("written %v", written)

		if err != nil {
			fmt.Println(" %v", err)
		}

		if err == io.EOF || read ==0 {

			fmt.Println("END")
			return
		}
	}



	defer file.Close()

}
