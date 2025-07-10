package methods

type FriendsAdd struct {
	UserID int    `json:"user_id"`          // ID of the user to add or approve request from
	Text   string `json:"text,omitempty"`   // Optional message text (max 500 chars)
	Follow bool   `json:"follow,omitempty"` // 1 = reject incoming request but follow instead
}
