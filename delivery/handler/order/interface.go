package order

import "net/http"

type OrderHandlerInterface interface {
	CreateOrderHandler(w http.ResponseWriter, r *http.Request)
	GetAllOrderHandler(w http.ResponseWriter, r *http.Request)
	UpdateOrderHandler(w http.ResponseWriter, r *http.Request)
	DeleteOrderHandler(w http.ResponseWriter, r *http.Request)
}
