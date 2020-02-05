package UploadVideoHandler

import (
	"github.com/AntonParaskiv/srv-video/domain/MediaInfo"
)

type Response struct {
	MediaInfo.FileInfo
	Errors []string `json:"errors,omitempty"`
}

func NewResponse() (r *Response) {
	r = new(Response)
	r.Errors = make([]string, 0)
	return
}

func (r *Response) AddError(err error) *Response {
	r.Errors = append(r.Errors, err.Error())
	return r
}
