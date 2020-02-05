package VideoInfoInteractor

import "github.com/AntonParaskiv/srv-video/domain/VideoFile"

type FileRepository interface {
	Save(file *VideoFile.File) (err error)
	Delete(file *VideoFile.File) (err error)
}
