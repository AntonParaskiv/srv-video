package VideoInfoInteractor

import (
	"fmt"
	"github.com/AntonParaskiv/srv-video/domain/MediaInfo"
	"github.com/AntonParaskiv/srv-video/domain/VideoFile"
	"github.com/AntonParaskiv/srv-video/interfaces/TempFileRepositoryMock"
	"github.com/AntonParaskiv/srv-video/interfaces/VideoInfoFileRepositoryMock"
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name  string
		wantI *Interactor
	}{
		{
			name:  "Success",
			wantI: &Interactor{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotI := New(); !reflect.DeepEqual(gotI, tt.wantI) {
				t.Errorf("New() = %v, want %v", gotI, tt.wantI)
			}
		})
	}
}

func TestInteractor_SetVideoInfoFileRepository(t *testing.T) {
	type fields struct{}
	type args struct {
		videoInfoFileRepository VideoInfoFileRepository
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Interactor
	}{
		{
			name:   "Success",
			fields: fields{},
			args: args{
				videoInfoFileRepository: VideoInfoFileRepositoryMock.New(),
			},
			want: &Interactor{
				videoInfoFileRepository: VideoInfoFileRepositoryMock.New(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Interactor{}
			if got := i.SetVideoInfoFileRepository(tt.args.videoInfoFileRepository); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetVideoInfoFileRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInteractor_SetFileRepository(t *testing.T) {
	type fields struct{}
	type args struct {
		fileRepository FileRepository
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Interactor
	}{
		{
			name:   "Success",
			fields: fields{},
			args: args{
				fileRepository: TempFileRepositoryMock.New(),
			},
			want: &Interactor{
				fileRepository: TempFileRepositoryMock.New(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Interactor{}
			if got := i.SetFileRepository(tt.args.fileRepository); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetFileRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInteractor_GetMediaInfo(t *testing.T) {
	type fields struct {
		videoInfoFileRepository VideoInfoFileRepository
		fileRepository          FileRepository
	}
	type args struct {
		fileData []byte
	}
	tests := []struct {
		name           string
		fields         fields
		args           args
		wantFileInfo   *MediaInfo.FileInfo
		wantErr        bool
		wantInteractor *Interactor
	}{
		{
			name: "Success",
			fields: fields{
				videoInfoFileRepository: VideoInfoFileRepositoryMock.New().
					SetStreamList([]*MediaInfo.StreamInfo{
						MediaInfo.NewStreamInfo().SetIsAudio().SetName("myAudioName"),
						MediaInfo.NewStreamInfo().SetIsVideo().SetName("myVideoName"),
					}),
				fileRepository: TempFileRepositoryMock.New(),
			},
			args: args{
				fileData: []byte(`myData`),
			},
			wantFileInfo: MediaInfo.NewFileInfo().
				AddStreamInfo(MediaInfo.NewStreamInfo().SetIsAudio().SetName("myAudioName")).
				AddStreamInfo(MediaInfo.NewStreamInfo().SetIsVideo().SetName("myVideoName")),
			wantErr: false,
			wantInteractor: &Interactor{
				videoInfoFileRepository: VideoInfoFileRepositoryMock.New().
					SetStreamList([]*MediaInfo.StreamInfo{
						MediaInfo.NewStreamInfo().SetIsAudio().SetName("myAudioName"),
						MediaInfo.NewStreamInfo().SetIsVideo().SetName("myVideoName"),
					}).
					SetFile(VideoFile.New().SetData([]byte(`myData`))),
				fileRepository: TempFileRepositoryMock.New().
					SetFile(VideoFile.New().SetData([]byte(`myData`))),
			},
		},
		{
			name: "VideoInfoFileRepository Error",
			fields: fields{
				videoInfoFileRepository: VideoInfoFileRepositoryMock.New().SetErr(fmt.Errorf("simulated error")),
				fileRepository:          TempFileRepositoryMock.New(),
			},
			args: args{
				fileData: []byte(`myData`),
			},
			wantFileInfo: nil,
			wantErr:      true,
			wantInteractor: &Interactor{
				videoInfoFileRepository: VideoInfoFileRepositoryMock.New().
					SetErr(fmt.Errorf("simulated error")).
					SetFile(VideoFile.New().SetData([]byte(`myData`))),
				fileRepository: TempFileRepositoryMock.New().
					SetFile(VideoFile.New().SetData([]byte(`myData`))),
			},
		},
		{
			name: "FileRepository Error",
			fields: fields{
				videoInfoFileRepository: VideoInfoFileRepositoryMock.New(),
				fileRepository:          TempFileRepositoryMock.New().SetErr(fmt.Errorf("simulated error")),
			},
			args: args{
				fileData: []byte(`myData`),
			},
			wantFileInfo: nil,
			wantErr:      true,
			wantInteractor: &Interactor{
				videoInfoFileRepository: VideoInfoFileRepositoryMock.New(),
				fileRepository: TempFileRepositoryMock.New().
					SetErr(fmt.Errorf("simulated error")).
					SetFile(VideoFile.New().SetData([]byte(`myData`))),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Interactor{
				videoInfoFileRepository: tt.fields.videoInfoFileRepository,
				fileRepository:          tt.fields.fileRepository,
			}
			gotFileInfo, err := i.GetMediaInfo(tt.args.fileData)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetMediaInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotFileInfo, tt.wantFileInfo) {
				t.Errorf("GetMediaInfo() gotFileInfo = %v, want %v", gotFileInfo, tt.wantFileInfo)
			}
			if !reflect.DeepEqual(i, tt.wantInteractor) {
				t.Errorf("Interactor = %v, want %v", i, tt.wantInteractor)
			}
		})
	}
}
