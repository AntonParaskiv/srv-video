package MediaFileInfoStorageMock

type Stream struct {
	codecName string
	bitRate   string
	duration  string
	width     int
	height    int
	isVideo   bool
	isAudio   bool
}

func NewStream() (s *Stream) {
	s = new(Stream)
	return
}

func (s *Stream) SetCodecName(codecName string) *Stream {
	s.codecName = codecName
	return s
}

func (s *Stream) GetCodecName() (codecName string) {
	codecName = s.codecName
	return
}

func (s *Stream) SetBitRate(bitRate string) *Stream {
	s.bitRate = bitRate
	return s
}

func (s *Stream) GetBitRate() (bitRate string) {
	bitRate = s.bitRate
	return
}

func (s *Stream) SetDuration(duration string) *Stream {
	s.duration = duration
	return s
}

func (s *Stream) GetDuration() (duration string) {
	duration = s.duration
	return
}

func (s *Stream) SetWidth(width int) *Stream {
	s.width = width
	return s
}

func (s *Stream) GetWidth() (width int) {
	width = s.width
	return
}

func (s *Stream) SetHeight(height int) *Stream {
	s.height = height
	return s
}

func (s *Stream) GetHeight() (height int) {
	height = s.height
	return
}

func (s *Stream) SetIsVideo() *Stream {
	s.isVideo = true
	return s
}

func (s *Stream) IsVideo() (isVideo bool) {
	isVideo = s.isVideo
	return
}

func (s *Stream) SetIsAudio() *Stream {
	s.isAudio = true
	return s
}

func (s *Stream) IsAudio() (isAudio bool) {
	isAudio = s.isAudio
	return
}
