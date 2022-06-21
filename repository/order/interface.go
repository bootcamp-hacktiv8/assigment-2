package order

import (
	_entities "assigment-2-usamah/entities"
)

type OrderRepositoryInterface interface {
	CreateOrder(order _entities.Order) (int, error)
	CreateItem(item _entities.Item) error
	GetAllOrder() ([]_entities.Order, error)
	GetAllItem() ([]_entities.Item, error)
	UpdateOrder(updateOrder _entities.Order, id int) error
	UpdateItem(updateItem _entities.Item, id int) error
	DeleteItem(id int) error
	DeleteOrder(id int) error
}
