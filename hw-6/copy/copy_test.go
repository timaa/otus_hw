package copy

import (
	"io/ioutil"
	"os"
	"testing"
)

const readFile = "test_read"
const writeFile = "test_write"

// TestCopy ...
func TestCopy(t *testing.T) {
	fileData := "12345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901"
	fileFrom, err := ioutil.TempFile(".", readFile)

	if err != nil {
		t.Fatalf("Cannot create temporary file for read")
	}

	defer os.Remove(fileFrom.Name())

	_, err = fileFrom.WriteString(fileData)

	if err != nil {
		t.Fatalf(" %v ", err)
	}

	fileTo, err := ioutil.TempFile(".", writeFile)

	if err != nil {
		t.Fatalf("Cannot create temporary file for write")
	}

	defer fileTo.Close()
	defer os.Remove(fileTo.Name())

	err = Copy(fileFrom.Name(), fileTo.Name(), -1, 0)

	if err != nil {
		t.Logf("%v", err)
	}

	fileFromInfo, err := fileFrom.Stat()
	if err != nil {
		t.Logf("%v", err)
	}

	fileToInfo, err := fileTo.Stat()
	if err != nil {
		t.Logf("%v", err)
	}

	if fileFromInfo.Size() != fileToInfo.Size() {
		t.Errorf("written file expected %d bytes got %d", fileFromInfo.Size(), fileToInfo.Size())
	}
}

// TestCopyErrorCase ...
func TestCopyErrorCase(t *testing.T) {
	err := Copy("not_exist_file", "not_exist_fil_2", -1, 0)

	if err == nil {
		t.Errorf("file not exist expect error, got nil")

	}
}
