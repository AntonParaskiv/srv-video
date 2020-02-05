package TempFileStorage

import (
	"fmt"
	"io/ioutil"
	"os"
)

type Storage struct {
	tempDir string
	pattern string
}

func New() (s *Storage) {
	s = new(Storage)
	return
}

func (s *Storage) SetTempDir(tempDir string) *Storage {
	s.tempDir = tempDir
	return s
}

func (s *Storage) SetPattern(pattern string) *Storage {
	s.pattern = pattern
	return s
}

func (s *Storage) SaveTempFile(fileData []byte) (fileName string, err error) {
	tempFile, err := ioutil.TempFile(s.tempDir, s.pattern)
	if err != nil {
		err = fmt.Errorf("create tempfile failed: %w", err)
		return
	}
	defer tempFile.Close()

	fileName = tempFile.Name()
	_, err = tempFile.Write(fileData)
	if err != nil {
		err = fmt.Errorf("write tempfile '%s' failed: %w", fileName, err)
		return
	}

	return
}

func (s *Storage) DeleteFile(fileName string) (err error) {
	if len(fileName) == 0 {
		err = fmt.Errorf("empty fileName")
		return
	}
	err = os.Remove(fileName)
	return
}
