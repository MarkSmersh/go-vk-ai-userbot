package events

type NewMessage struct {
	MessageId   int
	Flags       int
	MinorId     int
	PeerId      int
	Timestamp   int
	Text        string
	Attachments []any
	RandomId    int
}
