package enum

type ChatMessageType int

const (
	Sent ChatMessageType = iota
	Read
)
