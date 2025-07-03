package general

type Attachment struct {
	Type  string `json:"type"`
	Photo *Photo `json:"photo,omitempty"`
	// Video       *Video       `json:"video,omitempty"`
	// Audio       *Audio       `json:"audio,omitempty"`
	// AudioMsg    *AudioMsg    `json:"audio_message,omitempty"`
	// Doc         *Document    `json:"doc,omitempty"`
	// Link        *Link        `json:"link,omitempty"`
	// Market      *Market      `json:"market,omitempty"`
	MarketAlbum *Photo `json:"market_album,omitempty"` // Note: market_album uses Photo structure
	// Wall        *WallPost    `json:"wall,omitempty"`
	// WallReply   *WallComment `json:"wall_reply,omitempty"`
	Sticker *Sticker `json:"sticker,omitempty"` // Stickers also reuse photo struct (AND IT IS A FAKE INFORMATION GOTTEN FROM OFFICIAL VK DOCS)
	// Gift        *Gift        `json:"gift_item,omitempty"`
	// Call        *Call        `json:"call,omitempty"`
	AccessKey string `json:"access_key,omitempty"` // For restricted content
}
