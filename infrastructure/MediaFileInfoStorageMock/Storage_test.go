package MediaFileInfoStorageMock

import (
	"fmt"
	"github.com/AntonParaskiv/srv-video/interfaces/StorageStream"
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name  string
		wantS *Storage
	}{
		{
			name:  "Success",
			wantS: &Storage{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotS := New(); !reflect.DeepEqual(gotS, tt.wantS) {
				t.Errorf("New() = %v, want %v", gotS, tt.wantS)
			}
		})
	}
}

func TestStorage_SetFileName(t *testing.T) {
	type fields struct{}
	type args struct {
		fileName string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Storage
	}{
		{
			name:   "Success",
			fields: fields{},
			args: args{
				fileName: "myFileName",
			},
			want: &Storage{
				fileName: "myFileName",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Storage{}
			if got := s.SetFileName(tt.args.fileName); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetFileName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStorage_SetStreamList(t *testing.T) {
	type fields struct{}
	type args struct {
		streamList []StorageStream.StorageStream
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Storage
	}{
		{
			name:   "Success",
			fields: fields{},
			args: args{
				streamList: []StorageStream.StorageStream{
					NewStream().SetCodecName("myCodecName"),
				},
			},
			want: &Storage{
				streamList: []StorageStream.StorageStream{
					NewStream().SetCodecName("myCodecName"),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Storage{}
			if got := s.SetStreamList(tt.args.streamList); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetStreamList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStorage_SetErr(t *testing.T) {
	type fields struct{}
	type args struct {
		err error
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Storage
	}{
		{
			name:   "Success",
			fields: fields{},
			args: args{
				err: fmt.Errorf("simulated error"),
			},
			want: &Storage{
				err: fmt.Errorf("simulated error"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Storage{}
			if got := s.SetErr(tt.args.err); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetErr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStorage_GetStreamsFromFileName(t *testing.T) {
	type fields struct {
		streamList []StorageStream.StorageStream
		err        error
	}
	type args struct {
		fileName string
	}
	tests := []struct {
		name           string
		fields         fields
		args           args
		wantStreamList []StorageStream.StorageStream
		wantErr        bool
		wantStorage    *Storage
	}{
		{
			name: "Success",
			fields: fields{
				streamList: []StorageStream.StorageStream{
					NewStream().SetCodecName("myCodecName"),
				},
				err: fmt.Errorf("simulated error"),
			},
			args: args{
				fileName: "myFileName",
			},
			wantStreamList: []StorageStream.StorageStream{
				NewStream().SetCodecName("myCodecName"),
			},
			wantErr: true,
			wantStorage: &Storage{
				fileName: "myFileName",
				streamList: []StorageStream.StorageStream{
					NewStream().SetCodecName("myCodecName"),
				},
				err: fmt.Errorf("simulated error"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Storage{
				streamList: tt.fields.streamList,
				err:        tt.fields.err,
			}
			gotStreamList, err := s.GetStreamsFromFileName(tt.args.fileName)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetStreamsFromFileName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotStreamList, tt.wantStreamList) {
				t.Errorf("GetStreamsFromFileName() gotStreamList = %v, want %v", gotStreamList, tt.wantStreamList)
			}
			if !reflect.DeepEqual(s, tt.wantStorage) {
				t.Errorf("storage = %v, want %v", s, tt.wantStorage)
			}
		})
	}
}
