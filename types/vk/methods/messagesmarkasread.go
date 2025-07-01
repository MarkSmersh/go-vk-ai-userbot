package methods

type MessagesMarkAsRead struct {
	PeerID                 string `json:"peer_id,omitempty"`                   // Target peer (user, chat, or group)
	StartMessageID         int    `json:"start_message_id,omitempty"`          // Mark all messages from this ID as read
	GroupID                int    `json:"group_id,omitempty"`                  // Optional — group context
	MarkConversationAsRead bool   `json:"mark_conversation_as_read,omitempty"` // Whether to mark conversation read
	UpToCMID               int    `json:"up_to_cmid,omitempty"`                // Mark messages up to this cmid
}
