package UploadVideoHandler

import "github.com/AntonParaskiv/srv-video/domain/MediaInfo"

type VideoInfoInteractor interface {
	GetMediaInfo(fileData []byte) (fileInfo *MediaInfo.FileInfo, err error)
}
