package general

type Sticker struct {
	InnerType string `json:"inner_type"`
	StickerId int    `json:"sticker_id"`
	ProductID int    `json:"product_id"`
	IsAllowed bool   `json:"is_allowed"`
}
