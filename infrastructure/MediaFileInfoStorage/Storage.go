package MediaFileInfoStorage

import (
	"fmt"
	"github.com/AntonParaskiv/srv-video/interfaces/StorageStream"
	"github.com/vansante/go-ffprobe"
	"time"
)

type Storage struct {
}

func New() (s *Storage) {
	s = new(Storage)
	return
}

func (s *Storage) GetStreamsFromFileName(fileName string) (streamList []StorageStream.StorageStream, err error) {
	// get data from file
	data, err := ffprobe.GetProbeData(fileName, 10*time.Second)
	if err != nil {
		err = fmt.Errorf("get probe data failed: %w", err)
		return
	}

	// get stream from data
	ffprobeStreams := data.GetStreams(ffprobe.StreamAny)
	for _, ffprobeStream := range ffprobeStreams {
		stream := NewStream(ffprobeStream)
		streamList = append(streamList, stream)
	}

	return
}
