package cart

import (
	"encoding/json"
	"net/http"

	"github.com/raffidevaa/me-commerce/pkg/jwtauth"
	responseBuilder "github.com/raffidevaa/me-commerce/pkg/response-builder"
)

type CartController struct {
	service *CartService
}

func NewCartController(service *CartService) *CartController {
	return &CartController{
		service: service,
	}
}

func (c *CartController) AddItemToCart(w http.ResponseWriter, r *http.Request) {
	claims, ok := jwtauth.GetClaims(r.Context())
	if !ok {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	var req AddItemToCartRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		responseBuilder.BadRequest(w, "invalid request body")
	}

	cartItem, err := c.service.AddItemToCart(r.Context(), req, claims.UserID)
	if err != nil {
		responseBuilder.InternalError(w)
	}

	responseBuilder.OK(w, "item added to cart successfully", cartItem)

}

func (c *CartController) GetCartByUserID(w http.ResponseWriter, r *http.Request) {
	claims, ok := jwtauth.GetClaims(r.Context())
	if !ok {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	cartRes, err := c.service.GetCartByUserID(r.Context(), claims.UserID)
	if err != nil {
		responseBuilder.InternalError(w)
		return
	}

	responseBuilder.OK(w, "cart retrieved successfully", cartRes)
}
