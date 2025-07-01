package methods

type MessagesGetLongPollServer struct {
	NeedPts   int `json:"need_pts,omitempty"`   // 1 — возвращать поле pts
	GroupID   int `json:"group_id,omitempty"`   // ID сообщества
	LPVersion int `json:"lp_version,omitempty"` // Версия Long Poll API
}
