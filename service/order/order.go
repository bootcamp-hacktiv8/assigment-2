package order

import (
	_entities "assigment-2-usamah/entities"
	_orderRepository "assigment-2-usamah/repository/order"
)

type OrderService struct {
	orderRepository _orderRepository.OrderRepositoryInterface
}

func NewOrderUseCase(orderRepository _orderRepository.OrderRepositoryInterface) OrderServiceInterface {
	return &OrderService{
		orderRepository: orderRepository,
	}
}

func (os *OrderService) CreateOrder(order _entities.Order) error {
	id, err := os.orderRepository.CreateOrder(order)

	//create item in order
	for i := 0; i < len(order.Item); i++ {
		var item _entities.Item
		item.OrderID = id
		item.ItemCode = order.Item[i].ItemCode
		item.Description = order.Item[i].Description
		item.Quantity = order.Item[i].Quantity
		os.orderRepository.CreateItem(item)
	}
	return err
}

func (os *OrderService) GetAllOrder() ([]_entities.Order, error) {
	orders, err := os.orderRepository.GetAllOrder()
	return orders, err
}

func (os *OrderService) UpdateOrder(order _entities.Order, id int) error {
	err := os.orderRepository.UpdateOrder(order, id)

	items, errGet := os.orderRepository.GetAllItem()
	if errGet != nil {
		return errGet
	}

	for i := 0; i < len(order.Item); i++ {
		exist := false
		for j := 0; j < len(items); j++ {
			if order.Item[i].ID == items[j].ID {
				exist = true
				err := os.orderRepository.UpdateItem(order.Item[i], order.Item[i].ID)
				if err != nil {
					return err
				}
			}
		}
		if !exist {
			var item _entities.Item
			item.OrderID = id
			item.ItemCode = order.Item[i].ItemCode
			item.Description = order.Item[i].Description
			item.Quantity = order.Item[i].Quantity
			err := os.orderRepository.CreateItem(item)
			if err != nil {
				return err
			}
		}
	}
	return err
}

func (os *OrderService) DeleteOrder(id int) error {
	errItem := os.orderRepository.DeleteItem(id)
	if errItem != nil {
		return errItem
	}

	errOrder := os.orderRepository.DeleteOrder(id)
	return errOrder
}
