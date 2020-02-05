package TempFileStorageMock

import (
	"fmt"
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

func TestStorage_SetFileData(t *testing.T) {
	type fields struct{}
	type args struct {
		fileData []byte
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
				fileData: []byte(`myData`),
			},
			want: &Storage{
				fileData: []byte(`myData`),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Storage{}
			if got := s.SetFileData(tt.args.fileData); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetFileData() = %v, want %v", got, tt.want)
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

func TestStorage_SaveTempFile(t *testing.T) {
	type fields struct {
		fileName string
		err      error
	}
	type args struct {
		fileData []byte
	}
	tests := []struct {
		name         string
		fields       fields
		args         args
		wantFileName string
		wantErr      bool
		wantStorage  *Storage
	}{
		{
			name: "Success",
			fields: fields{
				fileName: "myFileName",
				err:      fmt.Errorf("simulated error"),
			},
			args: args{
				fileData: []byte(`myData`),
			},
			wantFileName: "myFileName",
			wantErr:      true,
			wantStorage: &Storage{
				fileData: []byte(`myData`),
				fileName: "myFileName",
				err:      fmt.Errorf("simulated error"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Storage{
				fileName: tt.fields.fileName,
				err:      tt.fields.err,
			}
			gotFileName, err := s.SaveTempFile(tt.args.fileData)
			if (err != nil) != tt.wantErr {
				t.Errorf("SaveTempFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotFileName != tt.wantFileName {
				t.Errorf("SaveTempFile() gotFileName = %v, want %v", gotFileName, tt.wantFileName)
			}
			if !reflect.DeepEqual(s, tt.wantStorage) {
				t.Errorf("storage = %v, want %v", s, tt.wantStorage)
			}
		})
	}
}

func TestStorage_DeleteFile(t *testing.T) {
	type fields struct {
		err error
	}
	type args struct {
		fileName string
	}
	tests := []struct {
		name        string
		fields      fields
		args        args
		wantErr     bool
		wantStorage *Storage
	}{
		{
			name: "Success",
			fields: fields{
				err: fmt.Errorf("simulated error"),
			},
			args: args{
				fileName: "myFileName",
			},
			wantErr: true,
			wantStorage: &Storage{
				fileName: "myFileName",
				err:      fmt.Errorf("simulated error"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Storage{
				err: tt.fields.err,
			}
			if err := s.DeleteFile(tt.args.fileName); (err != nil) != tt.wantErr {
				t.Errorf("DeleteFile() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(s, tt.wantStorage) {
				t.Errorf("storage = %v, want %v", s, tt.wantStorage)
			}
		})
	}
}
