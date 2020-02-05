package TempFileRepositoryMock

import (
	"github.com/AntonParaskiv/srv-video/domain/VideoFile"
)

type Repository struct {
	file *VideoFile.File
	err  error
}

func New() (r *Repository) {
	r = new(Repository)
	return
}

func (r *Repository) SetFile(file *VideoFile.File) *Repository {
	r.file = file
	return r
}

func (r *Repository) SetErr(err error) *Repository {
	r.err = err
	return r
}

func (r *Repository) Save(file *VideoFile.File) (err error) {
	r.SetFile(file)
	err = r.err
	return
}

func (r *Repository) Delete(file *VideoFile.File) (err error) {
	r.SetFile(file)
	err = r.err
	return
}
