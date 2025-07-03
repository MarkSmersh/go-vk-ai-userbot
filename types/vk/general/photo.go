package general

type Photo struct {
	ID        int         `json:"id"`
	AlbumID   int         `json:"album_id"`
	OwnerID   int         `json:"owner_id"`
	UserID    int         `json:"user_id,omitempty"`
	Text      string      `json:"text"`
	Date      int64       `json:"date"`
	ThumbHash string      `json:"thumb_hash,omitempty"`
	HasTags   bool        `json:"has_tags"`
	Sizes     []PhotoSize `json:"sizes"`
	Width     int         `json:"width,omitempty"`
	Height    int         `json:"height,omitempty"`
}

type PhotoSize struct {
	Type   string `json:"type"` // e.g. "s", "m", "x", "y", etc.
	URL    string `json:"url"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}
