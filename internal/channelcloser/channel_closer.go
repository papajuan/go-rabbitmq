package channelcloser

type ChannelCloser interface {
	CloseChannel() error
}
