package resources

type Item struct {
	Description string `json:"description"`
	ItemID      uint   `json:"itemId"`
	Quantity    uint   `json:"quantity"`
}
