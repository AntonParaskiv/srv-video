package TempFileRepository

import (
	"fmt"
	"github.com/AntonParaskiv/srv-video/domain/VideoFile"
)

type Repository struct {
	storage Storage
}

func New() (r *Repository) {
	r = new(Repository)
	return
}

func (r *Repository) SetStorage(storage Storage) *Repository {
	r.storage = storage
	return r
}

func (r *Repository) Save(file *VideoFile.File) (err error) {
	if file == nil {
		err = fmt.Errorf("file is nil")
		return
	}

	filename, err := r.storage.SaveTempFile(file.GetData())
	if err != nil {
		err = fmt.Errorf("save to storage failed: %w", err)
		return
	}

	file.SetFileName(filename)
	return
}

func (r *Repository) Delete(file *VideoFile.File) (err error) {
	if file == nil {
		err = fmt.Errorf("file is nil")
		return
	}

	err = r.storage.DeleteFile(file.GetFileName())
	if err != nil {
		err = fmt.Errorf("delete file from storage failed: %w", err)
		return
	}

	return
}
