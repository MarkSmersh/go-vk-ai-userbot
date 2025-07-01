package methods

type MessagesSend struct {
	UserID          int    `json:"user_id,omitempty"`
	PeerID          int    `json:"peer_id,omitempty"`
	RandomID        int    `json:"random_id"`
	Message         string `json:"message,omitempty"`
	Attachment      string `json:"attachment,omitempty"`
	ForwardMessages string `json:"forward_messages,omitempty"`
	ReplyTo         int    `json:"reply_to,omitempty"`
	StickerID       int    `json:"sticker_id,omitempty"`
	GroupID         int    `json:"group_id,omitempty"`
	Keyboard        string `json:"keyboard,omitempty"`
	Template        string `json:"template,omitempty"`
	Payload         string `json:"payload,omitempty"`
	Lat             string `json:"lat,omitempty"`
	Long            string `json:"long,omitempty"`
	Domain          string `json:"domain,omitempty"`
	UserIDs         string `json:"user_ids,omitempty"`
	PeerIDs         string `json:"peer_ids,omitempty"`
	Forward         string `json:"forward,omitempty"`
	ContentSource   string `json:"content_source,omitempty"`
	DontParseLinks  int    `json:"dont_parse_links,omitempty"`
	DisableMentions int    `json:"disable_mentions,omitempty"`
	Intent          string `json:"intent,omitempty"`
	SubscribeID     int    `json:"subscribe_id,omitempty"`
	ChatID          int    `json:"chat_id,omitempty"`
}
