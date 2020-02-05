package MediaFileInfoStorage

import "github.com/vansante/go-ffprobe"

type Stream struct {
	ffprobeStream ffprobe.Stream
}

func NewStream(ffprobeStream ffprobe.Stream) (s *Stream) {
	s = new(Stream)
	s.ffprobeStream = ffprobeStream
	return
}

func (s *Stream) IsVideo() (isVideo bool) {
	if s.ffprobeStream.CodecType == string(ffprobe.StreamVideo) {
		isVideo = true
	}
	return
}

func (s *Stream) IsAudio() (isAudio bool) {
	if s.ffprobeStream.CodecType == string(ffprobe.StreamAudio) {
		isAudio = true
	}
	return
}

func (s *Stream) GetCodecName() (codecName string) {
	codecName = s.ffprobeStream.CodecName
	return
}

func (s *Stream) GetBitRate() (bitRate string) {
	bitRate = s.ffprobeStream.BitRate
	return
}

func (s *Stream) GetDuration() (duration string) {
	duration = s.ffprobeStream.Duration
	return
}

func (s *Stream) GetWidth() (width int) {
	width = s.ffprobeStream.Width
	return
}

func (s *Stream) GetHeight() (height int) {
	height = s.ffprobeStream.Height
	return
}
