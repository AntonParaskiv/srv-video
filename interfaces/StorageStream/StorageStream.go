package StorageStream

type StorageStream interface {
	IsVideo() (isVideo bool)
	IsAudio() (isAudio bool)
	GetCodecName() (codecName string)
	GetBitRate() (bitRate string)
	GetDuration() (duration string)
	GetWidth() (width int)
	GetHeight() (height int)
}
