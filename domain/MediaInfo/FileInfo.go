package MediaInfo

type FileInfo struct {
	VideoStreamInfoList []*StreamInfo `json:"video,omitempty"`
	AudioStreamInfoList []*StreamInfo `json:"audio,omitempty"`
}

func NewFileInfo() (i *FileInfo) {
	i = new(FileInfo)
	i.VideoStreamInfoList = make([]*StreamInfo, 0)
	i.AudioStreamInfoList = make([]*StreamInfo, 0)
	return
}

func (i *FileInfo) AddStreamInfo(streamInfo *StreamInfo) *FileInfo {
	if streamInfo.IsVideo() {
		i.VideoStreamInfoList = append(i.VideoStreamInfoList, streamInfo)
	} else if streamInfo.IsAudio() {
		i.AudioStreamInfoList = append(i.AudioStreamInfoList, streamInfo)
	}
	return i
}
