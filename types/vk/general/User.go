package general

type User struct {
	ID         int    `json:"id"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	ScreenName string `json:"screen_name,omitempty"`
	Photo50    string `json:"photo_50,omitempty"`
	Photo100   string `json:"photo_100,omitempty"`
	Photo200   string `json:"photo_200,omitempty"`
	Sex        int    `json:"sex,omitempty"`
	Bdate      string `json:"bdate,omitempty"`
	// City           *City       `json:"city,omitempty"`
	// Country        *City       `json:"country,omitempty"`
	Online         int `json:"online,omitempty"`
	Verified       int `json:"verified,omitempty"`
	HasPhoto       int `json:"has_photo,omitempty"`
	FriendStatus   int `json:"friend_status,omitempty"`
	FollowersCount int `json:"followers_count,omitempty"`
	// Occupation     *Occupation `json:"occupation,omitempty"`
	// LastSeen       *LastSeen   `json:"last_seen,omitempty"`
}
