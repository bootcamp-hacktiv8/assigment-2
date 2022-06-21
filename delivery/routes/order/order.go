package order

import (
	_orderHandler "assigment-2-usamah/delivery/handler/order"
	_orderRepository "assigment-2-usamah/repository/order"
	_orderService "assigment-2-usamah/service/order"
	"net/http"

	"github.com/gorilla/mux"
)

type OrderResource struct{}

func (ur OrderResource) OrderRoute(orderRepository _orderRepository.OrderRepositoryInterface) *mux.Router {
	orderService := _orderService.NewOrderUseCase(orderRepository)
	orderHandler := _orderHandler.NewOrderHandler(orderService)

	router := mux.NewRouter()
	router.Handle("/", http.HandlerFunc(orderHandler.CreateOrderHandler)).Methods("POST")
	router.Handle("/", http.HandlerFunc(orderHandler.GetAllOrderHandler)).Methods("GET")
	router.Handle("/{id}", http.HandlerFunc(orderHandler.UpdateOrderHandler)).Methods("PUT")
	router.Handle("/{id}", http.HandlerFunc(orderHandler.DeleteOrderHandler)).Methods("DELETE")
	return router
}
