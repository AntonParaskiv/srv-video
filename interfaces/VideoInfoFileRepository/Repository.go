package VideoInfoFileRepository

import (
	"fmt"
	"github.com/AntonParaskiv/srv-video/domain/MediaInfo"
	"github.com/AntonParaskiv/srv-video/domain/VideoFile"
	"strconv"
)

type Repository struct {
	fileInfoStorage Storage
}

func New() (r *Repository) {
	r = new(Repository)
	return
}

func (r *Repository) SetStorage(storage Storage) *Repository {
	r.fileInfoStorage = storage
	return r
}

func (r *Repository) GetStreams(file *VideoFile.File) (streamList []*MediaInfo.StreamInfo, err error) {
	storageStreamList, err := r.fileInfoStorage.GetStreamsFromFileName(file.GetFileName())
	if err != nil {
		err = fmt.Errorf("get streams from file storage failed: %w", err)
	}

	streamList = make([]*MediaInfo.StreamInfo, 0, len(storageStreamList))
	for _, storageStream := range storageStreamList {
		var bitRate int64
		bitRateString := storageStream.GetBitRate()

		if len(bitRateString) > 0 {
			bitRate, err = strconv.ParseInt(bitRateString, 10, 64)
			if err != nil {
				err = fmt.Errorf("parse bitrate '%s' failed: %w", bitRateString, err)
				return
			}
		}

		duration := storageStream.GetDuration()
		if len(duration) > 0 {
			duration += "s"
		}

		streamInfo := MediaInfo.NewStreamInfo().
			SetName(storageStream.GetCodecName()).
			SetBitRate(bitRate).
			SetDuration(duration).
			SetWidth(int64(storageStream.GetWidth())).
			SetHeight(int64(storageStream.GetHeight()))

		if storageStream.IsVideo() {
			streamInfo.SetIsVideo()
		} else if storageStream.IsAudio() {
			streamInfo.SetIsAudio()
		}

		streamList = append(streamList, streamInfo)
	}

	return
}
