package order

type CreateOrderRequest struct {
	CartItemID uint `json:"cart_item_id" binding:"required"`
}

type CreateOrderResponse struct {
	UserID    uint   `json:"user_id"`
	OrderID   uint   `json:"order_id"`
	ProductID uint   `json:"product_id"`
	UnitPrice int64  `json:"unit_price"`
	Quantity  int    `json:"quantity"`
	Total     int64  `json:"total"`
	Status    string `json:"status"`
}
