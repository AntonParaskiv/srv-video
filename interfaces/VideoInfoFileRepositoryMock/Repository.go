package VideoInfoFileRepositoryMock

import (
	"github.com/AntonParaskiv/srv-video/domain/MediaInfo"
	"github.com/AntonParaskiv/srv-video/domain/VideoFile"
)

type Repository struct {
	file       *VideoFile.File
	streamList []*MediaInfo.StreamInfo
	err        error
}

func New() (r *Repository) {
	r = new(Repository)
	return
}

func (r *Repository) SetFile(file *VideoFile.File) *Repository {
	r.file = file
	return r
}

func (r *Repository) SetStreamList(streamList []*MediaInfo.StreamInfo) *Repository {
	r.streamList = streamList
	return r
}

func (r *Repository) SetErr(err error) *Repository {
	r.err = err
	return r
}

func (r *Repository) GetStreams(file *VideoFile.File) (streamList []*MediaInfo.StreamInfo, err error) {
	r.SetFile(file)
	streamList = r.streamList
	err = r.err
	return
}
