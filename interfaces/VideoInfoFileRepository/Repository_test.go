package VideoInfoFileRepository

import (
	"fmt"
	"github.com/AntonParaskiv/srv-video/domain/MediaInfo"
	"github.com/AntonParaskiv/srv-video/domain/VideoFile"
	"github.com/AntonParaskiv/srv-video/infrastructure/MediaFileInfoStorageMock"
	"github.com/AntonParaskiv/srv-video/interfaces/StorageStream"
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name  string
		wantR *Repository
	}{
		{
			name:  "Success",
			wantR: &Repository{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotR := New(); !reflect.DeepEqual(gotR, tt.wantR) {
				t.Errorf("New() = %v, want %v", gotR, tt.wantR)
			}
		})
	}
}

func TestRepository_SetStorage(t *testing.T) {
	type fields struct{}
	type args struct {
		storage Storage
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Repository
	}{
		{
			name:   "Success",
			fields: fields{},
			args: args{
				storage: MediaFileInfoStorageMock.New(),
			},
			want: &Repository{
				fileInfoStorage: MediaFileInfoStorageMock.New(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Repository{}
			if got := r.SetStorage(tt.args.storage); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetStorage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRepository_GetStreams(t *testing.T) {
	type fields struct {
		fileInfoStorage Storage
	}
	type args struct {
		file *VideoFile.File
	}
	tests := []struct {
		name           string
		fields         fields
		args           args
		wantStreamList []*MediaInfo.StreamInfo
		wantErr        bool
	}{
		{
			name: "Success",
			fields: fields{
				fileInfoStorage: MediaFileInfoStorageMock.New().SetStreamList(
					[]StorageStream.StorageStream{
						MediaFileInfoStorageMock.NewStream().
							SetIsVideo().
							SetCodecName("myVideoCodec").
							SetBitRate("123").
							SetDuration("60").
							SetWidth(100).
							SetHeight(200),
						MediaFileInfoStorageMock.NewStream().
							SetIsAudio().
							SetCodecName("myAudioCodec").
							SetBitRate("321").
							SetDuration("90"),
					},
				),
			},
			args: args{
				file: VideoFile.New().SetFileName("myFileName"),
			},
			wantStreamList: []*MediaInfo.StreamInfo{
				MediaInfo.NewStreamInfo().
					SetIsVideo().
					SetName("myVideoCodec").
					SetBitRate(123).
					SetDuration("60s").
					SetWidth(100).
					SetHeight(200),
				MediaInfo.NewStreamInfo().
					SetIsAudio().
					SetName("myAudioCodec").
					SetBitRate(321).
					SetDuration("90s"),
			},
			wantErr: false,
		},
		{
			name: "BitRate parse Error",
			fields: fields{
				fileInfoStorage: MediaFileInfoStorageMock.New().SetStreamList(
					[]StorageStream.StorageStream{
						MediaFileInfoStorageMock.NewStream().
							SetBitRate("broken bitrate"),
					},
				),
			},
			args: args{
				file: VideoFile.New().SetFileName("myFileName"),
			},
			wantStreamList: []*MediaInfo.StreamInfo{},
			wantErr:        true,
		},
		{
			name: "Get streams failed",
			fields: fields{
				fileInfoStorage: MediaFileInfoStorageMock.New().SetErr(fmt.Errorf("simulated error")),
			},
			args: args{
				file: VideoFile.New().SetFileName("myFileName"),
			},
			wantStreamList: []*MediaInfo.StreamInfo{},
			wantErr:        true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Repository{
				fileInfoStorage: tt.fields.fileInfoStorage,
			}
			gotStreamList, err := r.GetStreams(tt.args.file)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetStreams() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotStreamList, tt.wantStreamList) {
				t.Errorf("GetStreams() gotStreamList = %v, want %v", gotStreamList, tt.wantStreamList)
			}
		})
	}
}
