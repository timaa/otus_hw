package copy

import (
	"errors"
	"github.com/cheggaaa/pb/v3"
	"io"
	"os"
)
// Copy ...
func Copy(from string, to string, limit int64, offset int64) error {

	N := 1024*1024
	var file *os.File
	buff := make([]byte, N)

	file, err := os.Open(from)

	if err != nil {
		return err
	}

	file2, err := os.Create(to)
	if err != nil {
		return err
	}
	fileInfo,err := file.Stat()
	if err != nil {
		return err
	}

	if limit == -1 {
		limit=fileInfo.Size()
	}

	if offset > limit {
		return errors.New("offset cant be more than limit")
	}

	bar := pb.StartNew(int(limit - offset))
	bar.Set(pb.Bytes, true)
	sectionReader := io.NewSectionReader(file, offset, limit)

	for ; ; {
		read, err := sectionReader.Read(buff)

		written, err := file2.Write(buff[:read])
		bar.Add(written)
		if err != nil {
			return nil
		}
		if err == io.EOF || read == 0 {
			bar.Finish()
			break
		}
	}

	defer file.Close()
	return nil
}
