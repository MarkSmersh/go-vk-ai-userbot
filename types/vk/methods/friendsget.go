package methods

type FriendsGet struct {
	UserID int    `json:"user_id,omitempty"` // ID of the user whose friends list to retrieve
	Order  string `json:"order,omitempty"`   // Sorting order: hints, random, name
	ListID int    `json:"list_id,omitempty"` // ID of the friend list (only for current user)
	Count  int    `json:"count,omitempty"`   // Number of friends to return
	Offset int    `json:"offset,omitempty"`  // Offset for pagination
	// You need to write additiona login the method represented in VK core struct
	// Fields   string `json:"fields,omitempty"`    // Comma-separated list of additional user fields to return
	NameCase string `json:"name_case,omitempty"` // Case for declension of user name: nom, gen, dat, acc, ins, abl
	Ref      string `json:"ref,omitempty"`       // Arbitrary string reference (max 255 characters)
}
