package copy

import (
	"io/ioutil"
	"os"
	"testing"
)

const readFile = "test_read"
const writeFile = "test_write"

func TestCopy(t *testing.T) {
	fileData := "12345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901"
	fileFrom, err := ioutil.TempFile(".", readFile)

	if err != nil {
		t.Fatalf("Cannot create temporary file for read")
	}
	defer os.Remove(fileFrom.Name())

	written, err := fileFrom.WriteString(fileData)

	if err != nil {
		t.Fatalf(" %v ", err)
	}

	fileTo, err := ioutil.TempFile(".", writeFile)


	if err != nil {
		t.Fatalf("Cannot create temporary file for write")
	}

	defer os.Remove(fileTo.Name())

	Copy(fileFrom.Name(), fileTo.Name(), 0, -1)

	if err != nil {

	}

}
