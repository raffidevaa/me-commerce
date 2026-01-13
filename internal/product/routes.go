package product

import "github.com/go-chi/chi/v5"

type ProductRoutes struct {
	ProductController *ProductController
}

func NewProductRoutes(ac *ProductController) *ProductRoutes {
	return &ProductRoutes{
		ProductController: ac,
	}
}

func Routes(r chi.Router, c *ProductController) {
	r.Route("/products", func(r chi.Router) {
		r.Get("/{id}", c.GetProductByID)
		r.Get("/", c.GetAllProducts)
	})

}
