package TempFileRepositoryMock

import (
	"fmt"
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

func TestRepository_Save(t *testing.T) {
	type fields struct {
		file *VideoFile.File
		err  error
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
				err: fmt.Errorf("simulated error"),
			},
			args: args{
				file: VideoFile.New().SetFileName("myFileName"),
			},
			wantErr: true,
			wantRepository: &Repository{
				file: VideoFile.New().SetFileName("myFileName"),
				err:  fmt.Errorf("simulated error"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Repository{
				file: tt.fields.file,
				err:  tt.fields.err,
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
		file *VideoFile.File
		err  error
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
				err: fmt.Errorf("simulated error"),
			},
			args: args{
				file: VideoFile.New().SetFileName("myFileName"),
			},
			wantErr: true,
			wantRepository: &Repository{
				file: VideoFile.New().SetFileName("myFileName"),
				err:  fmt.Errorf("simulated error"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Repository{
				file: tt.fields.file,
				err:  tt.fields.err,
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
