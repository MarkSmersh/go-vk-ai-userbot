package methods

type MessagesSetActivity struct {
	UserID  string `json:"user_id,omitempty"`  // Optional — user to show typing to
	Type    string `json:"type"`               // typing | audiomessage | photo | video | file | videomessage
	PeerID  int    `json:"peer_id"`            // Target peer (user/chat/group)
	GroupID int    `json:"group_id,omitempty"` // Optional — needed if using group token
}
