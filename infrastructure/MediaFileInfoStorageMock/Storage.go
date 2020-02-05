package MediaFileInfoStorageMock

import (
	"github.com/AntonParaskiv/srv-video/interfaces/StorageStream"
)

type Storage struct {
	fileName   string
	streamList []StorageStream.StorageStream
	err        error
}

func New() (s *Storage) {
	s = new(Storage)
	return
}

func (s *Storage) SetFileName(fileName string) *Storage {
	s.fileName = fileName
	return s
}

func (s *Storage) SetStreamList(streamList []StorageStream.StorageStream) *Storage {
	s.streamList = streamList
	return s
}

func (s *Storage) SetErr(err error) *Storage {
	s.err = err
	return s
}

func (s *Storage) GetStreamsFromFileName(fileName string) (streamList []StorageStream.StorageStream, err error) {
	s.SetFileName(fileName)
	streamList = s.streamList
	err = s.err
	return
}
