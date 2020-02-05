package MediaFileInfoStorageMock

import (
	"reflect"
	"testing"
)

func TestNewStream(t *testing.T) {
	tests := []struct {
		name  string
		wantS *Stream
	}{
		{
			name:  "Success",
			wantS: &Stream{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotS := NewStream(); !reflect.DeepEqual(gotS, tt.wantS) {
				t.Errorf("NewStream() = %v, want %v", gotS, tt.wantS)
			}
		})
	}
}

func TestStream_SetCodecName(t *testing.T) {
	type fields struct{}
	type args struct {
		codecName string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Stream
	}{
		{
			name:   "Success",
			fields: fields{},
			args: args{
				codecName: "myCodecName",
			},
			want: &Stream{
				codecName: "myCodecName",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Stream{}
			if got := s.SetCodecName(tt.args.codecName); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetCodecName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStream_GetCodecName(t *testing.T) {
	type fields struct {
		codecName string
	}
	tests := []struct {
		name          string
		fields        fields
		wantCodecName string
	}{
		{
			name: "Success",
			fields: fields{
				codecName: "myCodecName",
			},
			wantCodecName: "myCodecName",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Stream{
				codecName: tt.fields.codecName,
			}
			if gotCodecName := s.GetCodecName(); gotCodecName != tt.wantCodecName {
				t.Errorf("GetCodecName() = %v, want %v", gotCodecName, tt.wantCodecName)
			}
		})
	}
}

func TestStream_SetBitRate(t *testing.T) {
	type fields struct{}
	type args struct {
		bitRate string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Stream
	}{
		{
			name:   "Success",
			fields: fields{},
			args: args{
				bitRate: "myBitRate",
			},
			want: &Stream{
				bitRate: "myBitRate",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Stream{}
			if got := s.SetBitRate(tt.args.bitRate); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetBitRate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStream_GetBitRate(t *testing.T) {
	type fields struct {
		bitRate string
	}
	tests := []struct {
		name        string
		fields      fields
		wantBitRate string
	}{
		{
			name: "Success",
			fields: fields{
				bitRate: "myBitRate",
			},
			wantBitRate: "myBitRate",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Stream{
				bitRate: tt.fields.bitRate,
			}
			if gotBitRate := s.GetBitRate(); gotBitRate != tt.wantBitRate {
				t.Errorf("GetBitRate() = %v, want %v", gotBitRate, tt.wantBitRate)
			}
		})
	}
}

func TestStream_SetDuration(t *testing.T) {
	type fields struct{}
	type args struct {
		duration string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Stream
	}{
		{
			name:   "Success",
			fields: fields{},
			args: args{
				duration: "myDuration",
			},
			want: &Stream{
				duration: "myDuration",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Stream{}
			if got := s.SetDuration(tt.args.duration); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetDuration() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStream_GetDuration(t *testing.T) {
	type fields struct {
		duration string
	}
	tests := []struct {
		name         string
		fields       fields
		wantDuration string
	}{
		{
			name: "Success",
			fields: fields{
				duration: "myDuration",
			},
			wantDuration: "myDuration",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Stream{
				duration: tt.fields.duration,
			}
			if gotDuration := s.GetDuration(); gotDuration != tt.wantDuration {
				t.Errorf("GetDuration() = %v, want %v", gotDuration, tt.wantDuration)
			}
		})
	}
}

func TestStream_SetWidth(t *testing.T) {
	type fields struct{}
	type args struct {
		width int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Stream
	}{
		{
			name:   "Success",
			fields: fields{},
			args: args{
				width: 100,
			},
			want: &Stream{
				width: 100,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Stream{}
			if got := s.SetWidth(tt.args.width); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetWidth() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStream_GetWidth(t *testing.T) {
	type fields struct {
		width int
	}
	tests := []struct {
		name      string
		fields    fields
		wantWidth int
	}{
		{
			name: "Success",
			fields: fields{
				width: 100,
			},
			wantWidth: 100,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Stream{
				width: tt.fields.width,
			}
			if gotWidth := s.GetWidth(); gotWidth != tt.wantWidth {
				t.Errorf("GetWidth() = %v, want %v", gotWidth, tt.wantWidth)
			}
		})
	}
}

func TestStream_SetHeight(t *testing.T) {
	type fields struct{}
	type args struct {
		height int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Stream
	}{
		{
			name:   "Success",
			fields: fields{},
			args: args{
				height: 100,
			},
			want: &Stream{
				height: 100,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Stream{}
			if got := s.SetHeight(tt.args.height); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetHeight() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStream_GetHeight(t *testing.T) {
	type fields struct {
		height int
	}
	tests := []struct {
		name       string
		fields     fields
		wantHeight int
	}{
		{
			name: "Success",
			fields: fields{
				height: 100,
			},
			wantHeight: 100,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Stream{
				height: tt.fields.height,
			}
			if gotHeight := s.GetHeight(); gotHeight != tt.wantHeight {
				t.Errorf("GetHeight() = %v, want %v", gotHeight, tt.wantHeight)
			}
		})
	}
}

func TestStream_SetIsVideo(t *testing.T) {
	type fields struct{}
	tests := []struct {
		name   string
		fields fields
		want   *Stream
	}{
		{
			name:   "Success",
			fields: fields{},
			want: &Stream{
				isVideo: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Stream{}
			if got := s.SetIsVideo(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetIsVideo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStream_IsVideo(t *testing.T) {
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
			s := &Stream{
				isVideo: tt.fields.isVideo,
			}
			if gotIsVideo := s.IsVideo(); gotIsVideo != tt.wantIsVideo {
				t.Errorf("IsVideo() = %v, want %v", gotIsVideo, tt.wantIsVideo)
			}
		})
	}
}

func TestStream_SetIsAudio(t *testing.T) {
	type fields struct{}
	tests := []struct {
		name   string
		fields fields
		want   *Stream
	}{
		{
			name:   "Success",
			fields: fields{},
			want: &Stream{
				isAudio: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Stream{}
			if got := s.SetIsAudio(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetIsAudio() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStream_IsAudio(t *testing.T) {
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
			s := &Stream{
				isAudio: tt.fields.isAudio,
			}
			if gotIsAudio := s.IsAudio(); gotIsAudio != tt.wantIsAudio {
				t.Errorf("IsAudio() = %v, want %v", gotIsAudio, tt.wantIsAudio)
			}
		})
	}
}
