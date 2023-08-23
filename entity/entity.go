package entity

type Inventory struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Item_code   string `json:"item_code"`
	Stock       int    `json:"stock"`
	Description string `json:"description"`
	Status      string `json:"status"`
}
