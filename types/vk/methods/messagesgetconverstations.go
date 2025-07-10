package methods

type MessagesGetConversations struct {
	Offset         int    `json:"offset,omitempty"`           // Смещение
	Count          int    `json:"count,omitempty"`            // Количество бесед
	Filter         string `json:"filter,omitempty"`           // Типы: all, unread, important, unanswered, archive
	Extended       bool   `json:"extended,omitempty"`         // true – вернуть инфу о пользователях
	StartMessageID int    `json:"start_message_id,omitempty"` // Начать с определенного сообщения
	Fields         string `json:"fields,omitempty"`           // Доп. поля профилей (например, photo_100, sex)
	GroupID        int    `json:"group_id,omitempty"`         // Только для токена сообщества
}
