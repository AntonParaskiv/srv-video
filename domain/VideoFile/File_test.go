package VideoFile

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name  string
		wantF *File
	}{
		{
			name:  "Success",
			wantF: &File{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotF := New(); !reflect.DeepEqual(gotF, tt.wantF) {
				t.Errorf("New() = %v, want %v", gotF, tt.wantF)
			}
		})
	}
}

func TestVideoFile_SetData(t *testing.T) {
	type fields struct{}
	type args struct {
		data []byte
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *File
	}{
		{
			name:   "Success",
			fields: fields{},
			args: args{
				data: []byte(`myData`),
			},
			want: &File{
				data: []byte(`myData`),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &File{}
			if got := f.SetData(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetData() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVideoFile_GetData(t *testing.T) {
	type fields struct {
		data []byte
	}
	tests := []struct {
		name     string
		fields   fields
		wantData []byte
	}{
		{
			name: "Success",
			fields: fields{
				data: []byte(`myData`),
			},
			wantData: []byte(`myData`),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &File{
				data: tt.fields.data,
			}
			if gotData := f.GetData(); !reflect.DeepEqual(gotData, tt.wantData) {
				t.Errorf("GetData() = %v, want %v", gotData, tt.wantData)
			}
		})
	}
}

func TestFile_SetFileName(t *testing.T) {
	type fields struct{}
	type args struct {
		fileName string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *File
	}{
		{
			name:   "Success",
			fields: fields{},
			args: args{
				fileName: "myFileName",
			},
			want: &File{
				fileName: "myFileName",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &File{}
			if got := f.SetFileName(tt.args.fileName); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetFileName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFile_GetFileName(t *testing.T) {
	type fields struct {
		fileName string
	}
	tests := []struct {
		name         string
		fields       fields
		wantFileName string
	}{
		{
			name: "Success",
			fields: fields{
				fileName: "myFileName",
			},
			wantFileName: "myFileName",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &File{
				fileName: tt.fields.fileName,
			}
			if gotFileName := f.GetFileName(); gotFileName != tt.wantFileName {
				t.Errorf("GetFileName() = %v, want %v", gotFileName, tt.wantFileName)
			}
		})
	}
}
