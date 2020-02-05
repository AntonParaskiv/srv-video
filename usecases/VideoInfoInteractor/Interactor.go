package VideoInfoInteractor

import (
	"fmt"
	"github.com/AntonParaskiv/srv-video/domain/MediaInfo"
	"github.com/AntonParaskiv/srv-video/domain/VideoFile"
)

type Interactor struct {
	videoInfoFileRepository VideoInfoFileRepository
	fileRepository          FileRepository
}

func New() (i *Interactor) {
	i = new(Interactor)
	return
}

func (i *Interactor) SetVideoInfoFileRepository(videoInfoFileRepository VideoInfoFileRepository) *Interactor {
	i.videoInfoFileRepository = videoInfoFileRepository
	return i
}

func (i *Interactor) SetFileRepository(fileRepository FileRepository) *Interactor {
	i.fileRepository = fileRepository
	return i
}

func (i *Interactor) GetMediaInfo(fileData []byte) (fileInfo *MediaInfo.FileInfo, err error) {
	file := VideoFile.New().SetData(fileData)

	// save fileData
	err = i.fileRepository.Save(file)
	if err != nil {
		err = fmt.Errorf("save file to repository failed: %w", err)
		return
	}
	defer i.fileRepository.Delete(file)

	// get info
	streamList, err := i.videoInfoFileRepository.GetStreams(file)
	if err != nil {
		err = fmt.Errorf("get stream list from file failed: %w", err)
		return
	}

	fileInfo = MediaInfo.NewFileInfo()
	for _, stream := range streamList {
		fileInfo.AddStreamInfo(stream)
	}

	return
}
