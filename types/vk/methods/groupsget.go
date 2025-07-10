package methods

type GroupsGet struct {
	UserID   int    `json:"user_id,omitempty"`  // optional
	Extended int    `json:"extended,omitempty"` // 1 = full info
	Filter   string `json:"filter,omitempty"`   // comma-separated: admin, editor, etc.
	Fields   string `json:"fields,omitempty"`   // comma-separated: description, status, etc.
	Offset   int    `json:"offset,omitempty"`
	Count    int    `json:"count,omitempty"` // max 1000
}
