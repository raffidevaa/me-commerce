package order

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/raffidevaa/me-commerce/pkg/jwtauth"
	responseBuilder "github.com/raffidevaa/me-commerce/pkg/response-builder"
)

type OrderController struct {
	orderService *OrderService
}

func NewOrderController(orderService *OrderService) *OrderController {
	return &OrderController{
		orderService: orderService,
	}
}

func (c *OrderController) CreateOrder(w http.ResponseWriter, r *http.Request) {
	claims, ok := jwtauth.GetClaims(r.Context())
	if !ok {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	var req CreateOrderRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		responseBuilder.BadRequest(w, "invalid request body")
		return
	}

	orderRes, err := c.orderService.CreateOrder(r.Context(), req, claims.UserID)
	if err != nil {
		responseBuilder.InternalError(w)
		return
	}

	responseBuilder.Created(w, "success create order", orderRes)
}

func (c *OrderController) GetOrdersByUserID(w http.ResponseWriter, r *http.Request) {
	claims, ok := jwtauth.GetClaims(r.Context())
	if !ok {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	orders, err := c.orderService.GetOrdersByUserID(r.Context(), claims.UserID)
	if err != nil {
		responseBuilder.InternalError(w)
		return
	}

	responseBuilder.OK(w, "orders retrieved successfully", orders)
}

func (c *OrderController) Payment(w http.ResponseWriter, r *http.Request){
	_, ok := jwtauth.GetClaims(r.Context())
	if !ok {
		http.Error(w, "unauthorized", http.StatusUnauthorized)
		return
	}
	
	orderIDParam := chi.URLParam(r, "id")
	orderID, err := strconv.ParseUint(orderIDParam, 10, 32)
	if err != nil {
		responseBuilder.BadRequest(w, "failed parsing id param")
		return
	}

	order, err := c.orderService.Payment(r.Context(), uint(orderID))
	if err != nil{
		responseBuilder.BadRequest(w, "failed payment")
		return
	} 
	
	responseBuilder.OK(w, "payment succes", order)
}
