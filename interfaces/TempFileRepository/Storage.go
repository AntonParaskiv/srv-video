package TempFileRepository

type Storage interface {
	SaveTempFile(fileData []byte) (fileName string, err error)
	DeleteFile(fileName string) (err error)
}
