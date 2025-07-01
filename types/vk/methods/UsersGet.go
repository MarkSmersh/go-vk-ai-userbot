package methods

type UsersGet struct {
	UserIDs     string `json:"user_ids,omitempty"`      // ID пользователей или screen_name, через запятую
	Fields      string `json:"fields,omitempty"`        // Дополнительные поля профиля
	NameCase    string `json:"name_case,omitempty"`     // Падеж имени и фамилии
	FromGroupID int    `json:"from_group_id,omitempty"` // Устаревший параметр, больше не используется
}
