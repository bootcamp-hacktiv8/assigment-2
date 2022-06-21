package routes

import (
	_orderRouter "assigment-2-usamah/delivery/routes/order"
	_orderRepository "assigment-2-usamah/repository/order"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

func Routes(
	orderRepository _orderRepository.OrderRepositoryInterface,
) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	mount(router, "/orders", _orderRouter.OrderResource{}.OrderRoute(orderRepository))

	return router
}

func mount(r *mux.Router, path string, handler http.Handler) {
	r.PathPrefix(path).Handler(
		http.StripPrefix(
			strings.TrimSuffix(path, "/"),
			handler,
		),
	)
}
