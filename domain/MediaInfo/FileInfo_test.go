package MediaInfo

import (
	"reflect"
	"testing"
)

func TestNewFileInfo(t *testing.T) {
	tests := []struct {
		name  string
		wantI *FileInfo
	}{
		{
			name: "Success",
			wantI: &FileInfo{
				VideoStreamInfoList: []*StreamInfo{},
				AudioStreamInfoList: []*StreamInfo{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotI := NewFileInfo(); !reflect.DeepEqual(gotI, tt.wantI) {
				t.Errorf("NewFileInfo() = %v, want %v", gotI, tt.wantI)
			}
		})
	}
}

func TestFileInfo_AddStreamInfo(t *testing.T) {
	type fields struct {
		VideoStreamInfoList []*StreamInfo
		AudioStreamInfoList []*StreamInfo
	}
	type args struct {
		streamInfo *StreamInfo
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *FileInfo
	}{
		{
			name: "Add video stream info",
			fields: fields{
				VideoStreamInfoList: []*StreamInfo{
					NewStreamInfo().SetIsVideo().SetName("myName"),
				},
			},
			args: args{
				streamInfo: NewStreamInfo().SetIsVideo().SetName("anotherName"),
			},
			want: &FileInfo{
				VideoStreamInfoList: []*StreamInfo{
					NewStreamInfo().SetIsVideo().SetName("myName"),
					NewStreamInfo().SetIsVideo().SetName("anotherName"),
				},
			},
		},
		{
			name: "Add audio stream info",
			fields: fields{
				AudioStreamInfoList: []*StreamInfo{
					NewStreamInfo().SetIsAudio().SetName("myName"),
				},
			},
			args: args{
				streamInfo: NewStreamInfo().SetIsAudio().SetName("anotherName"),
			},
			want: &FileInfo{
				AudioStreamInfoList: []*StreamInfo{
					NewStreamInfo().SetIsAudio().SetName("myName"),
					NewStreamInfo().SetIsAudio().SetName("anotherName"),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &FileInfo{
				VideoStreamInfoList: tt.fields.VideoStreamInfoList,
				AudioStreamInfoList: tt.fields.AudioStreamInfoList,
			}
			if got := i.AddStreamInfo(tt.args.streamInfo); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddStreamInfo() = %v, want %v", got, tt.want)
			}
		})
	}
}
