package general

type Message struct {
	ID     int    `json:"id"`
	Date   int    `json:"date"`
	PeerID int    `json:"peer_id"`
	FromID int    `json:"from_id"`
	Text   string `json:"text"`
	// FormatData            *FormatData            `json:"format_data,omitempty"`
	RandomID  int    `json:"random_id,omitempty"`
	Ref       string `json:"ref,omitempty"`
	RefSource string `json:"ref_source,omitempty"`
	// Attachments           []Attachment           `json:"attachments,omitempty"`
	Important bool `json:"important,omitempty"`
	// Geo                   *Geo                   `json:"geo,omitempty"`
	Payload string `json:"payload,omitempty"`
	// Keyboard              *Keyboard              `json:"keyboard,omitempty"`
	FwdMessages  []Message `json:"fwd_messages,omitempty"`
	ReplyMessage *Message  `json:"reply_message,omitempty"`
	// Action       *ChatAction `json:"action,omitempty"`
	// AdminAuthorID         int                    `json:"admin_author_id,omitempty"`
	ConversationMessageID int    `json:"conversation_message_id,omitempty"`
	IsCropped             bool   `json:"is_cropped,omitempty"`
	MembersCount          int    `json:"members_count,omitempty"`
	UpdateTime            int    `json:"update_time,omitempty"`
	WasListened           bool   `json:"was_listened,omitempty"`
	PinnedAt              int    `json:"pinned_at,omitempty"`
	MessageTag            string `json:"message_tag,omitempty"`
}
