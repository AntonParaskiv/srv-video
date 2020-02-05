package VideoInfoInteractor

import (
	"github.com/AntonParaskiv/srv-video/domain/MediaInfo"
	"github.com/AntonParaskiv/srv-video/domain/VideoFile"
)

type VideoInfoFileRepository interface {
	GetStreams(file *VideoFile.File) (streamList []*MediaInfo.StreamInfo, err error)
}
