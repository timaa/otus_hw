package copy

import (
	"fmt"
	"io"
	"os"
)

func Copy(from string, to string, limit int, offset int) error {

	filePath := "/Users/timur/go/src/github.com/timaa/otus_hw/hw-6/test.txt"
	filePath2 := "/Users/timur/go/src/github.com/timaa/otus_hw/hw-6/test2.txt"
	N := 1024 * 1024
	var file *os.File
	buff := make([]byte, N)

	file, err:= os.Open(filePath)
	file2, _ := os.Create(filePath2)

	if err !=nil {
		fmt.Println("%v", err)
	}
	err = nil

	for ;; {

		read, err := io.NewSectionReader(file, buff, 0)
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
	return nil
}