package methods

type GroupsGetMembers struct {
	GroupID string `json:"group_id"`         // ID or screen name of the group
	Sort    string `json:"sort,omitempty"`   // id_asc, id_desc, time_asc, time_desc
	Offset  int    `json:"offset,omitempty"` // For pagination
	Count   int    `json:"count,omitempty"`  // Max 1000
	Fields  string `json:"fields,omitempty"` // Comma-separated user fields
	Filter  string `json:"filter,omitempty"` // friends, managers, donut, etc.
}
