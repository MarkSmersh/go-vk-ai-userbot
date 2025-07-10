package general

type Group struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	ScreenName   string `json:"screen_name"`
	IsClosed     int    `json:"is_closed"`
	Deactivated  string `json:"deactivated,omitempty"`
	IsAdmin      int    `json:"is_admin,omitempty"`
	AdminLevel   int    `json:"admin_level,omitempty"`
	IsMember     int    `json:"is_member,omitempty"`
	IsAdvertiser int    `json:"is_advertiser,omitempty"`
	InvitedBy    int    `json:"invited_by,omitempty"`
	Type         string `json:"type"`
	Photo50      string `json:"photo_50"`
	Photo100     string `json:"photo_100"`
	Photo200     string `json:"photo_200"`

	// Опциональные поля
	Activity          string `json:"activity,omitempty"`
	AgeLimits         int    `json:"age_limits,omitempty"`
	Description       string `json:"description,omitempty"`
	FixedPost         int    `json:"fixed_post,omitempty"`
	HasPhoto          int    `json:"has_photo,omitempty"`
	IsFavorite        int    `json:"is_favorite,omitempty"`
	IsHiddenFromFeed  int    `json:"is_hidden_from_feed,omitempty"`
	IsMessagesBlocked int    `json:"is_messages_blocked,omitempty"`
	MainAlbumID       int    `json:"main_album_id,omitempty"`
	MainSection       int    `json:"main_section,omitempty"`
	MembersCount      int    `json:"members_count,omitempty"`
	PublicDateLabel   string `json:"public_date_label,omitempty"`
	Site              string `json:"site,omitempty"`
	StartDate         int64  `json:"start_date,omitempty"`
	FinishDate        int64  `json:"finish_date,omitempty"`
	Status            string `json:"status,omitempty"`
	Trending          int    `json:"trending,omitempty"`
	Verified          int    `json:"verified"`
	Wall              int    `json:"wall,omitempty"`
	WikiPage          string `json:"wiki_page,omitempty"`

	// Вложенные объекты
	// City               *City            `json:"city,omitempty"`
	// Country            *Country         `json:"country,omitempty"`
	// Cover              *Cover           `json:"cover,omitempty"`
	// Place              *Place           `json:"place,omitempty"`
	// CropPhoto          *CropPhoto       `json:"crop_photo,omitempty"`
	// Market             *Market          `json:"market,omitempty"`
	// Counters           *GroupCounters   `json:"counters,omitempty"`
	// Contacts           []Contact        `json:"contacts,omitempty"`
	// Links              []GroupLink      `json:"links,omitempty"`
}
