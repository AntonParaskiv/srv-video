package VideoFile

type File struct {
	fileName string
	data     []byte
}

func New() (f *File) {
	f = new(File)
	return
}

func (f *File) SetFileName(fileName string) *File {
	f.fileName = fileName
	return f
}

func (f *File) GetFileName() (fileName string) {
	fileName = f.fileName
	return
}

func (f *File) SetData(data []byte) *File {
	f.data = data
	return f
}

func (f *File) GetData() (data []byte) {
	data = f.data
	return
}
