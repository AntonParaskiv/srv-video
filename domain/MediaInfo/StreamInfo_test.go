package MediaInfo

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name  string
		wantI *StreamInfo
	}{
		{
			name:  "Success",
			wantI: &StreamInfo{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotI := NewStreamInfo(); !reflect.DeepEqual(gotI, tt.wantI) {
				t.Errorf("NewStreamInfo() = %v, want %v", gotI, tt.wantI)
			}
		})
	}
}

func TestInfo_SetName(t *testing.T) {
	type fields struct{}
	type args struct {
		name string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *StreamInfo
	}{
		{
			name:   "Success",
			fields: fields{},
			args: args{
				name: "myName",
			},
			want: &StreamInfo{
				Name: "myName",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &StreamInfo{}
			if got := i.SetName(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInfo_GetName(t *testing.T) {
	type fields struct {
		Name string
	}
	tests := []struct {
		name     string
		fields   fields
		wantName string
	}{
		{
			name: "Success",
			fields: fields{
				Name: "myName",
			},
			wantName: "myName",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &StreamInfo{
				Name: tt.fields.Name,
			}
			if gotName := i.GetName(); gotName != tt.wantName {
				t.Errorf("GetName() = %v, want %v", gotName, tt.wantName)
			}
		})
	}
}

func TestInfo_SetWidth(t *testing.T) {
	type fields struct{}
	type args struct {
		width int64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *StreamInfo
	}{
		{
			name:   "Success",
			fields: fields{},
			args: args{
				width: 100,
			},
			want: &StreamInfo{
				Width: 100,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &StreamInfo{}
			if got := i.SetWidth(tt.args.width); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetWidth() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInfo_GetWidth(t *testing.T) {
	type fields struct {
		Width int64
	}
	tests := []struct {
		name      string
		fields    fields
		wantWidth int64
	}{
		{
			name: "Success",
			fields: fields{
				Width: 100,
			},
			wantWidth: 100,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &StreamInfo{
				Width: tt.fields.Width,
			}
			if gotWidth := i.GetWidth(); gotWidth != tt.wantWidth {
				t.Errorf("GetWidth() = %v, want %v", gotWidth, tt.wantWidth)
			}
		})
	}
}

func TestInfo_SetHeight(t *testing.T) {
	type fields struct {
	}
	type args struct {
		height int64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *StreamInfo
	}{
		{
			name:   "Success",
			fields: fields{},
			args: args{
				height: 100,
			},
			want: &StreamInfo{
				Height: 100,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &StreamInfo{}
			if got := i.SetHeight(tt.args.height); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetHeight() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInfo_GetHeight(t *testing.T) {
	type fields struct {
		Height int64
	}
	tests := []struct {
		name       string
		fields     fields
		wantHeight int64
	}{
		{
			name: "Success",
			fields: fields{
				Height: 100,
			},
			wantHeight: 100,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &StreamInfo{
				Height: tt.fields.Height,
			}
			if gotHeight := i.GetHeight(); gotHeight != tt.wantHeight {
				t.Errorf("GetHeight() = %v, want %v", gotHeight, tt.wantHeight)
			}
		})
	}
}

func TestInfo_SetBitRate(t *testing.T) {
	type fields struct{}
	type args struct {
		bitRate int64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *StreamInfo
	}{
		{
			name:   "Success",
			fields: fields{},
			args: args{
				bitRate: 100,
			},
			want: &StreamInfo{
				BitRate: 100,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &StreamInfo{}
			if got := i.SetBitRate(tt.args.bitRate); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetBitRate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInfo_GetBitRate(t *testing.T) {
	type fields struct {
		BitRate int64
	}
	tests := []struct {
		name        string
		fields      fields
		wantBitRate int64
	}{
		{
			name: "Success",
			fields: fields{
				BitRate: 100,
			},
			wantBitRate: 100,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &StreamInfo{
				BitRate: tt.fields.BitRate,
			}
			if gotBitRate := i.GetBitRate(); gotBitRate != tt.wantBitRate {
				t.Errorf("GetBitRate() = %v, want %v", gotBitRate, tt.wantBitRate)
			}
		})
	}
}

func TestInfo_SetDuration(t *testing.T) {
	type fields struct{}
	type args struct {
		duration string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *StreamInfo
	}{
		{
			name:   "Success",
			fields: fields{},
			args: args{
				duration: "myDuration",
			},
			want: &StreamInfo{
				Duration: "myDuration",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &StreamInfo{}
			if got := i.SetDuration(tt.args.duration); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetDuration() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInfo_GetDuration(t *testing.T) {
	type fields struct {
		Duration string
	}
	tests := []struct {
		name         string
		fields       fields
		wantDuration string
	}{
		{
			name: "Success",
			fields: fields{
				Duration: "myDuration",
			},
			wantDuration: "myDuration",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &StreamInfo{
				Duration: tt.fields.Duration,
			}
			if gotDuration := i.GetDuration(); gotDuration != tt.wantDuration {
				t.Errorf("GetDuration() = %v, want %v", gotDuration, tt.wantDuration)
			}
		})
	}
}

func TestStreamInfo_IsAudio(t *testing.T) {
	type fields struct {
		isAudio bool
	}
	tests := []struct {
		name        string
		fields      fields
		wantIsAudio bool
	}{
		{
			name: "True",
			fields: fields{
				isAudio: true,
			},
			wantIsAudio: true,
		},
		{
			name: "False",
			fields: fields{
				isAudio: false,
			},
			wantIsAudio: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &StreamInfo{
				isAudio: tt.fields.isAudio,
			}
			if gotIsAudio := i.IsAudio(); gotIsAudio != tt.wantIsAudio {
				t.Errorf("IsAudio() = %v, want %v", gotIsAudio, tt.wantIsAudio)
			}
		})
	}
}

func TestStreamInfo_IsVideo(t *testing.T) {
	type fields struct {
		isVideo bool
	}
	tests := []struct {
		name        string
		fields      fields
		wantIsVideo bool
	}{
		{
			name: "True",
			fields: fields{
				isVideo: true,
			},
			wantIsVideo: true,
		},
		{
			name: "False",
			fields: fields{
				isVideo: false,
			},
			wantIsVideo: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &StreamInfo{
				isVideo: tt.fields.isVideo,
			}
			if gotIsVideo := i.IsVideo(); gotIsVideo != tt.wantIsVideo {
				t.Errorf("IsVideo() = %v, want %v", gotIsVideo, tt.wantIsVideo)
			}
		})
	}
}

func TestStreamInfo_SetIsAudio(t *testing.T) {
	type fields struct{}
	tests := []struct {
		name   string
		fields fields
		want   *StreamInfo
	}{
		{
			name:   "Success",
			fields: fields{},
			want: &StreamInfo{
				isAudio: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &StreamInfo{}
			if got := i.SetIsAudio(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetIsAudio() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStreamInfo_SetIsVideo(t *testing.T) {
	type fields struct{}
	tests := []struct {
		name   string
		fields fields
		want   *StreamInfo
	}{
		{
			name:   "Success",
			fields: fields{},
			want: &StreamInfo{
				isVideo: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &StreamInfo{}
			if got := i.SetIsVideo(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetIsVideo() = %v, want %v", got, tt.want)
			}
		})
	}
}
