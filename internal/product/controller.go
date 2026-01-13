package product

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	responseBuilder "github.com/raffidevaa/me-commerce/pkg/response-builder"
)

type ProductController struct {
	service *ProductService
}

func NewProductController(service *ProductService) *ProductController {
	return &ProductController{
		service: service,
	}
}

func (c *ProductController) GetProductByID(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		responseBuilder.BadRequest(w, "failed parsing id param")
		return
	}

	product, err := c.service.GetProductByID(r.Context(), uint(id))
	if err != nil {
		responseBuilder.NotFound(w, "product not found")
		return
	}

	res := ProductDTO{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Stock:       product.Stock,
	}

	responseBuilder.OK(w, "product retrieved successfully", res)
}

func (c *ProductController) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	products, err := c.service.GetAllProducts(r.Context())
	if err != nil {
		responseBuilder.InternalError(w)
		return
	}

	responseBuilder.OK(w, "products retrieved successfully", products)
}
