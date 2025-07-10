package general

type Conversation struct {
	Peer        Peer `json:"peer"`
	InRead      int  `json:"in_read"`
	OutRead     int  `json:"out_read"`
	UnreadCount int  `json:"unread_count,omitempty"`
	Important   bool `json:"important,omitempty"`
	CanWrite    struct {
		Allowed bool `json:"allowed"`
	} `json:"can_write"`
	// ChatSettings *ChatSettings `json:"chat_settings,omitempty"` // Only for group chats
}

type Peer struct {
	ID      int    `json:"id"`
	Type    string `json:"type"` // user, group, chat, email
	LocalID int    `json:"local_id"`
}
