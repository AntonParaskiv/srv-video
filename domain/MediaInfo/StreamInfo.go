package MediaInfo

type StreamInfo struct {
	Name     string `json:"name,omitempty"`
	Width    int64  `json:"width,omitempty"`
	Height   int64  `json:"height,omitempty"`
	BitRate  int64  `json:"bitRate,omitempty"`
	Duration string `json:"duration,omitempty"`
	isVideo  bool
	isAudio  bool
}

func NewStreamInfo() (i *StreamInfo) {
	i = new(StreamInfo)
	return
}

func (i *StreamInfo) SetName(name string) *StreamInfo {
	i.Name = name
	return i
}

func (i *StreamInfo) GetName() (name string) {
	name = i.Name
	return
}

func (i *StreamInfo) SetWidth(width int64) *StreamInfo {
	i.Width = width
	return i
}

func (i *StreamInfo) GetWidth() (width int64) {
	width = i.Width
	return
}

func (i *StreamInfo) SetHeight(height int64) *StreamInfo {
	i.Height = height
	return i
}

func (i *StreamInfo) GetHeight() (height int64) {
	height = i.Height
	return
}

func (i *StreamInfo) SetBitRate(bitRate int64) *StreamInfo {
	i.BitRate = bitRate
	return i
}

func (i *StreamInfo) GetBitRate() (bitRate int64) {
	bitRate = i.BitRate
	return
}

func (i *StreamInfo) SetDuration(duration string) *StreamInfo {
	i.Duration = duration
	return i
}

func (i *StreamInfo) GetDuration() (duration string) {
	duration = i.Duration
	return
}

func (i *StreamInfo) SetIsVideo() *StreamInfo {
	i.isVideo = true
	return i
}

func (i *StreamInfo) IsVideo() (isVideo bool) {
	isVideo = i.isVideo
	return
}

func (i *StreamInfo) SetIsAudio() *StreamInfo {
	i.isAudio = true
	return i
}

func (i *StreamInfo) IsAudio() (isAudio bool) {
	isAudio = i.isAudio
	return
}
