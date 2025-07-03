package methods

type StoreGetStickersKeywords struct {
	StickersIDs  int  `json:"stickers_ids,omitempty"`  // Optional: specific sticker IDs
	ProductsIDs  int  `json:"products_ids,omitempty"`  // Optional: specific product IDs
	Aliases      bool `json:"aliases,omitempty"`       // Optional: include keyword synonyms
	AllProducts  bool `json:"all_products,omitempty"`  // Optional: include all store products
	NeedStickers bool `json:"need_stickers,omitempty"` // Optional: return stickers with keywords
}
