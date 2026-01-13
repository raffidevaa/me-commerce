package product

type ProductDTO struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int64  `json:"price"`
	Stock       int    `json:"stock"`
}

type GetAllProducts struct {
	Products []ProductDTO `json:"products"`
}
