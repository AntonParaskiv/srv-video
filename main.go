package main

import (
	"fmt"
	"github.com/AntonParaskiv/srv-video/infrastructure/MediaFileInfoStorage"
	"github.com/AntonParaskiv/srv-video/infrastructure/TempFileStorage"
	"github.com/AntonParaskiv/srv-video/infrastructure/WebServer"
	"github.com/AntonParaskiv/srv-video/interfaces/TempFileRepository"
	"github.com/AntonParaskiv/srv-video/interfaces/VideoInfoFileRepository"
	"github.com/AntonParaskiv/srv-video/interfaces/WebHandlers/UploadVideoHandler"
	"github.com/AntonParaskiv/srv-video/usecases/VideoInfoInteractor"
	"log"
	"os"
)

func main() {
	// temp file repository
	tempFileStorage := TempFileStorage.New().
		SetTempDir(os.TempDir()).
		SetPattern("srv-Video_")

	fileRepository := TempFileRepository.New().
		SetStorage(tempFileStorage)

	// media file info repository
	mediaFileInfoStorage := MediaFileInfoStorage.New()
	videoInfoFileRepository := VideoInfoFileRepository.New().
		SetStorage(mediaFileInfoStorage)

	// interactor
	videoInfoInteractor := VideoInfoInteractor.New().
		SetFileRepository(fileRepository).
		SetVideoInfoFileRepository(videoInfoFileRepository)

	// webHandlers
	uploadVideoHandler := UploadVideoHandler.New().
		SetVideoInfoInteractor(videoInfoInteractor)

	// webServer
	webServer := WebServer.New().
		SetListenAddr(":4000").
		HandleFunc("/api/uploadVideo", uploadVideoHandler.Handle)

	log.Println("staring web server on", webServer.GetListenAddr())
	err := webServer.Start()
	if err != nil {
		err = fmt.Errorf("web server start failed: %w", err)
		log.Fatal(err)
	}

}
