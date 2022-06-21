package entities

type Item struct {
	ID          int    `json:"id" form:"id"`
	OrderID     int    `json:"order_id" form:"order_id"`
	ItemCode    string `json:"item_code" form:"item_code"`
	Description string `json:"description" form:"description"`
	Quantity    int    `json:"quantity" form:"quantity"`
}
