package order

import (
	"assigment-2-usamah/delivery/helper"
	"assigment-2-usamah/entities"
	_orderService "assigment-2-usamah/service/order"
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type OrderHandler struct {
	orderService _orderService.OrderServiceInterface
}

func NewOrderHandler(orderService _orderService.OrderServiceInterface) *OrderHandler {
	return &OrderHandler{
		orderService: orderService,
	}
}

func (oh *OrderHandler) CreateOrderHandler(w http.ResponseWriter, r *http.Request) {
	var newOrder entities.Order
	errDecode := json.NewDecoder(r.Body).Decode(&newOrder)
	if errDecode != nil {
		w.Write([]byte("error decoding json body"))
		return
	}

	err := oh.orderService.CreateOrder(newOrder)
	if err != nil {
		response, _ := json.Marshal(helper.APIResponseFailed(err.Error()))
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(response)
	} else {
		response, _ := json.Marshal(helper.APIResponseSuccessWithouData("success create order"))
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(response)
	}
}

func (oh *OrderHandler) GetAllOrderHandler(w http.ResponseWriter, r *http.Request) {
	orders, err := oh.orderService.GetAllOrder()
	switch {
	case err == sql.ErrNoRows:
		response, _ := json.Marshal(helper.APIResponseFailed("data not found"))
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		w.Write(response)
	case err != nil:
		response, _ := json.Marshal(helper.APIResponseFailed(err.Error()))
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(response)
	default:
		response, _ := json.Marshal(helper.APIResponseSuccess("success get all order", orders))
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(response)
	}
}

func (oh *OrderHandler) UpdateOrderHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idStr := params["id"]
	id, _ := strconv.Atoi(idStr)

	var updateOrder entities.Order
	errDecode := json.NewDecoder(r.Body).Decode(&updateOrder)
	if errDecode != nil {
		w.Write([]byte("error decoding json body"))
		return
	}

	err := oh.orderService.UpdateOrder(updateOrder, id)
	if err != nil {
		response, _ := json.Marshal(helper.APIResponseFailed(err.Error()))
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(response)
	} else {
		response, _ := json.Marshal(helper.APIResponseSuccessWithouData("success update order"))
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(response)
	}
}

func (oh *OrderHandler) DeleteOrderHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idStr := params["id"]
	id, _ := strconv.Atoi(idStr)

	err := oh.orderService.DeleteOrder(id)
	if err != nil {
		response, _ := json.Marshal(helper.APIResponseFailed(err.Error()))
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(response)
	} else {
		response, _ := json.Marshal(helper.APIResponseSuccessWithouData("success delete order"))
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(response)
	}
}
