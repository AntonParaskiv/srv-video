package TempFileStorageMock

type Storage struct {
	fileData []byte
	fileName string
	err      error
}

func New() (s *Storage) {
	s = new(Storage)
	return
}

func (s *Storage) SetFileData(fileData []byte) *Storage {
	s.fileData = fileData
	return s
}

func (s *Storage) SetFileName(fileName string) *Storage {
	s.fileName = fileName
	return s
}

func (s *Storage) SetErr(err error) *Storage {
	s.err = err
	return s
}

func (s *Storage) SaveTempFile(fileData []byte) (fileName string, err error) {
	s.SetFileData(fileData)
	fileName = s.fileName
	err = s.err
	return
}

func (s *Storage) DeleteFile(fileName string) (err error) {
	s.SetFileName(fileName)
	err = s.err
	return
}
