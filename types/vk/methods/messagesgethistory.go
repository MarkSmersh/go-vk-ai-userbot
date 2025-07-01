package methods

type MessagesGetHistory struct {
	Offset         int      `json:"offset,omitempty"`           // Message offset
	Count          int      `json:"count,omitempty"`            // Max 200
	UserID         int      `json:"user_id,omitempty"`          // Deprecated — use PeerID
	PeerID         int      `json:"peer_id,omitempty"`          // Target dialog ID
	StartMessageID int      `json:"start_message_id,omitempty"` // Start from this message
	Rev            int      `json:"rev,omitempty"`              // 1 = chronological, 0 = reverse
	Extended       bool     `json:"extended,omitempty"`         // Return user info
	Fields         []string `json:"fields,omitempty"`           // Extra user profile fields
	GroupID        int      `json:"group_id,omitempty"`         // Community context
}
