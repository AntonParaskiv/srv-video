package TempFileRepository

import (
	"fmt"
	"github.com/AntonParaskiv/srv-video/domain/VideoFile"
	"github.com/AntonParaskiv/srv-video/infrastructure/TempFileStorageMock"
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
				storage: TempFileStorageMock.New().SetFileName("myFileName"),
			},
			want: &Repository{
				storage: TempFileStorageMock.New().SetFileName("myFileName"),
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

func TestRepository_Save(t *testing.T) {
	type fields struct {
		storage Storage
	}
	type args struct {
		file *VideoFile.File
	}
	tests := []struct {
		name           string
		fields         fields
		args           args
		wantErr        bool
		wantRepository *Repository
	}{
		{
			name: "Success",
			fields: fields{
				storage: TempFileStorageMock.New().
					SetFileName("myFileName"),
			},
			args: args{
				file: VideoFile.New().SetData([]byte(`myData`)),
			},
			wantErr: false,
			wantRepository: &Repository{
				storage: TempFileStorageMock.New().
					SetFileData([]byte(`myData`)).
					SetFileName("myFileName"),
			},
		},
		{
			name: "Storage Error",
			fields: fields{
				storage: TempFileStorageMock.New().
					SetErr(fmt.Errorf("simulate error")),
			},
			args: args{
				file: VideoFile.New().SetData([]byte(`myData`)),
			},
			wantErr: true,
			wantRepository: &Repository{
				storage: TempFileStorageMock.New().
					SetFileData([]byte(`myData`)).
					SetErr(fmt.Errorf("simulate error")),
			},
		},
		{
			name: "File is nil",
			fields: fields{
				storage: TempFileStorageMock.New(),
			},
			args: args{
				file: nil,
			},
			wantErr: true,
			wantRepository: &Repository{
				storage: TempFileStorageMock.New(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Repository{
				storage: tt.fields.storage,
			}
			if err := r.Save(tt.args.file); (err != nil) != tt.wantErr {
				t.Errorf("Save() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(r, tt.wantRepository) {
				t.Errorf("Repository = %v, want %v", r, tt.wantRepository)
			}
		})
	}
}

func TestRepository_Delete(t *testing.T) {
	type fields struct {
		storage Storage
	}
	type args struct {
		file *VideoFile.File
	}
	tests := []struct {
		name           string
		fields         fields
		args           args
		wantErr        bool
		wantRepository *Repository
	}{
		{
			name: "Success",
			fields: fields{
				storage: TempFileStorageMock.New(),
			},
			args: args{
				file: VideoFile.New().SetFileName("myFileName"),
			},
			wantErr: false,
			wantRepository: &Repository{
				storage: TempFileStorageMock.New().
					SetFileName("myFileName"),
			},
		},
		{
			name: "Storage Error",
			fields: fields{
				storage: TempFileStorageMock.New().
					SetErr(fmt.Errorf("simulate error")),
			},
			args: args{
				file: VideoFile.New().SetFileName("myFileName"),
			},
			wantErr: true,
			wantRepository: &Repository{
				storage: TempFileStorageMock.New().
					SetFileName("myFileName").
					SetErr(fmt.Errorf("simulate error")),
			},
		},
		{
			name: "File is nil",
			fields: fields{
				storage: TempFileStorageMock.New(),
			},
			args: args{
				file: nil,
			},
			wantErr: true,
			wantRepository: &Repository{
				storage: TempFileStorageMock.New(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Repository{
				storage: tt.fields.storage,
			}
			if err := r.Delete(tt.args.file); (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(r, tt.wantRepository) {
				t.Errorf("Repository = %v, want %v", r, tt.wantRepository)
			}
		})
	}
}
