package events

type NewMessage struct {
	MessageId   int
	Flags       int
	MinorId     int
	PeerId      int
	Timestamp   int
	Text        string
	Attachments []Attachment
	RandomId    int
}

type Attachment struct {
	ID          int
	ProductId   int
	Type        string
	Attachments []any
}
