package VideoInfoFileRepositoryMock

import (
	"fmt"
	"github.com/AntonParaskiv/srv-video/domain/MediaInfo"
	"github.com/AntonParaskiv/srv-video/domain/VideoFile"
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

func TestRepository_SetFile(t *testing.T) {
	type fields struct{}
	type args struct {
		file *VideoFile.File
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
				file: VideoFile.New().SetFileName("myFileName"),
			},
			want: &Repository{
				file: VideoFile.New().SetFileName("myFileName"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Repository{}
			if got := r.SetFile(tt.args.file); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRepository_SetStreamList(t *testing.T) {
	type fields struct{}
	type args struct {
		streamList []*MediaInfo.StreamInfo
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
				streamList: []*MediaInfo.StreamInfo{
					MediaInfo.NewStreamInfo().SetName("myName"),
				},
			},
			want: &Repository{
				streamList: []*MediaInfo.StreamInfo{
					MediaInfo.NewStreamInfo().SetName("myName"),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Repository{}
			if got := r.SetStreamList(tt.args.streamList); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetStreamList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRepository_SetErr(t *testing.T) {
	type fields struct{}
	type args struct {
		err error
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
				err: fmt.Errorf("simulated error"),
			},
			want: &Repository{
				err: fmt.Errorf("simulated error"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Repository{}
			if got := r.SetErr(tt.args.err); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetErr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRepository_GetStreams(t *testing.T) {
	type fields struct {
		streamList []*MediaInfo.StreamInfo
		err        error
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
		wantRepository *Repository
	}{
		{
			name: "Success",
			fields: fields{
				streamList: []*MediaInfo.StreamInfo{
					MediaInfo.NewStreamInfo().SetName("myName"),
				},
				err: fmt.Errorf("simulated error"),
			},
			args: args{
				file: VideoFile.New().SetFileName("myFileName"),
			},
			wantStreamList: []*MediaInfo.StreamInfo{
				MediaInfo.NewStreamInfo().SetName("myName"),
			},
			wantErr: true,
			wantRepository: &Repository{
				file: VideoFile.New().SetFileName("myFileName"),
				streamList: []*MediaInfo.StreamInfo{
					MediaInfo.NewStreamInfo().SetName("myName"),
				},
				err: fmt.Errorf("simulated error"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Repository{
				streamList: tt.fields.streamList,
				err:        tt.fields.err,
			}
			gotStreamList, err := r.GetStreams(tt.args.file)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetStreams() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotStreamList, tt.wantStreamList) {
				t.Errorf("GetStreams() gotStreamList = %v, want %v", gotStreamList, tt.wantStreamList)
			}
			if !reflect.DeepEqual(r, tt.wantRepository) {
				t.Errorf("Repository = %v, want %v", r, tt.wantRepository)
			}
		})
	}
}
