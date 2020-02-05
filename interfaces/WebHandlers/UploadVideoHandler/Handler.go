package UploadVideoHandler

import (
	"encoding/json"
	"fmt"
	"github.com/AntonParaskiv/srv-video/domain/MediaInfo"
	"io/ioutil"
	"log"
	"net/http"
)

type WebHandler struct {
	videoInfoInteractor VideoInfoInteractor
}

func New() (h *WebHandler) {
	h = new(WebHandler)
	return
}

func (h *WebHandler) SetVideoInfoInteractor(videoInfoInteractor VideoInfoInteractor) *WebHandler {
	h.videoInfoInteractor = videoInfoInteractor
	return h
}

func (h *WebHandler) Handle(httpResponseWriter http.ResponseWriter, httpRequest *http.Request) {
	log.Printf("new request %s from %s", httpRequest.RequestURI, httpRequest.RemoteAddr)

	file, _, err := httpRequest.FormFile("file")
	if err != nil {
		err = fmt.Errorf("parse file failed: %w", err)
		h.sendError(httpResponseWriter, err, http.StatusInternalServerError)
		return
	}

	fileData, err := ioutil.ReadAll(file)
	if err != nil {
		err = fmt.Errorf("read file from form failed: %w", err)
		h.sendError(httpResponseWriter, err, http.StatusInternalServerError)
		return
	}

	fileInfo, err := h.videoInfoInteractor.GetMediaInfo(fileData)
	if err != nil {
		err = fmt.Errorf("get media info failed: %w", err)
		h.sendError(httpResponseWriter, err, http.StatusInternalServerError)
		return
	}

	err = h.sendSuccess(httpResponseWriter, fileInfo)
	if err != nil {
		err = fmt.Errorf("send success response failed: %w", err)
		h.sendError(httpResponseWriter, err, http.StatusInternalServerError)
		return
	}

	return
}

func (h *WebHandler) sendSuccess(httpResponseWriter http.ResponseWriter, fileInfo *MediaInfo.FileInfo) (err error) {
	if fileInfo == nil {
		err = fmt.Errorf("fileInfo is nil")
		return
	}

	httpResponseWriter.Header().Set("Content-Type", "application/json")
	httpResponseWriter.WriteHeader(http.StatusOK)

	response := Response{
		FileInfo: *fileInfo,
	}

	responseBytes, err := json.Marshal(response)
	if err != nil {
		err = fmt.Errorf("marshal response failed: %w", err)
		log.Println(err)
		return
	}
	h.write(httpResponseWriter, responseBytes)
	return
}

func (h *WebHandler) sendError(httpResponseWriter http.ResponseWriter, err error, statusCode int) {
	httpResponseWriter.Header().Set("Content-Type", "application/json")
	httpResponseWriter.WriteHeader(statusCode)

	response := NewResponse().AddError(err)
	responseBytes, err := json.Marshal(response)
	if err != nil {
		err = fmt.Errorf("marshal response failed: %w", err)
		log.Println(err)
		return
	}
	h.write(httpResponseWriter, responseBytes)
}

func (h *WebHandler) write(httpResponseWriter http.ResponseWriter, responseBytes []byte) {
	_, err := httpResponseWriter.Write(responseBytes)
	if err != nil {
		err = fmt.Errorf("write response failed: %w", err)
		log.Println(err)
		return
	}
}
