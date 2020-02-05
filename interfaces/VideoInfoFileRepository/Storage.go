package VideoInfoFileRepository

import "github.com/AntonParaskiv/srv-video/interfaces/StorageStream"

type Storage interface {
	GetStreamsFromFileName(fileName string) (streamList []StorageStream.StorageStream, err error)
}
