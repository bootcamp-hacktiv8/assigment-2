package order

import (
	_entities "assigment-2-usamah/entities"
)

type OrderServiceInterface interface {
	CreateOrder(order _entities.Order) error
	GetAllOrder() ([]_entities.Order, error)
	UpdateOrder(order _entities.Order, id int) error
	DeleteOrder(id int) error
}
