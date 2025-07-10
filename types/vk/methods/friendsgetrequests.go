package methods

type FriendsGetRequests struct {
	Offset   int `json:"offset,omitempty"`   // Offset needed to fetch a subset of friend requests
	Count    int `json:"count,omitempty"`    // Max number of requests to return (up to 1000)
	Extended int `json:"extended,omitempty"` // 1 to return request messages and suggesters
	// If you want to get more fields (not only slice of ints), you need to write additional login in the method
	// NeedMutual int    `json:"need_mutual,omitempty"` // 1 to return mutual friends (max 2 requests returned)
	Out        int    `json:"out,omitempty"`         // 0 for received requests, 1 for sent requests
	Sort       int    `json:"sort,omitempty"`        // 0 to sort by date, 1 by mutual friends (ignored if out=1)
	NeedViewed int    `json:"need_viewed,omitempty"` // 1 to return viewed requests (ignored if out=1)
	Suggested  int    `json:"suggested,omitempty"`   // 1 to return friend suggestions instead of requests
	Ref        string `json:"ref,omitempty"`         // Arbitrary string reference (max 255 characters)
	Fields     string `json:"fields,omitempty"`      // List of additional user fields to return
}
