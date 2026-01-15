package cart

type AddItemToCartRequest struct {
	ProductID uint `json:"product_id" binding:"required"`
	Quantity  int  `json:"quantity" binding:"required,min=1"`
}

type GetCartsResponse struct {
	CartItems []CartItem `json:"cart_items"`
}
