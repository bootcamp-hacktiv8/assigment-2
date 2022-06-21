package order

import (
	_entities "assigment-2-usamah/entities"
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"time"

	_ "github.com/lib/pq"
)

type OrderRepository struct {
	database *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{
		database: db,
	}
}

func (or *OrderRepository) CreateOrder(order _entities.Order) (int, error) {
	query := "INSERT INTO orders (customer_name, ordered_at) VALUES ($1, $2) RETURNING order_id"
	ctx := context.Background()
	var id int
	err := or.database.QueryRowContext(ctx, query, order.CustomerName, time.Now()).Scan(&id)
	if err != nil {
		return id, err
	}
	return id, nil
}

func (or *OrderRepository) CreateItem(item _entities.Item) error {
	query := "INSERT INTO items (item_code, description, quantity, order_id) VALUES ($1, $2, $3, $4)"
	ctx := context.Background()
	_, err := or.database.ExecContext(ctx, query, item.ItemCode, item.Description, item.Quantity, item.OrderID)
	if err != nil {
		return err
	}
	return nil
}

func (or *OrderRepository) GetAllOrder() ([]_entities.Order, error) {
	query := `SELECT orders.order_id, orders.customer_name, orders.ordered_at, json_agg(json_build_object('id', items.item_id, 'order_id', items.order_id, 'item_code', items.item_code, 'description', items.description, 'quantity', items.quantity)) as items
	FROM orders
	JOIN items ON (orders.order_id = items.order_id)
	GROUP BY orders.order_id`

	ctx := context.Background()

	rows, err := or.database.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []_entities.Order
	for rows.Next() {
		var order _entities.Order
		var itemStr string

		err := rows.Scan(&order.ID, &order.CustomerName, &order.OrderedAt, &itemStr)
		if err != nil {
			return nil, errors.New("error scan")
		}

		var items []_entities.Item
		errUnmarshal := json.Unmarshal([]byte(itemStr), &items)
		if errUnmarshal != nil {
			return nil, errors.New("error when parsing items")
		} else {
			order.Item = append(order.Item, items...)
		}
		orders = append(orders, order)
	}
	return orders, nil
}

func (or *OrderRepository) GetAllItem() ([]_entities.Item, error) {
	query := "SELECT item_id, item_code, description, quantity, order_id FROM items"
	ctx := context.Background()
	rows, err := or.database.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []_entities.Item
	for rows.Next() {
		var item _entities.Item
		rows.Scan(&item.ID, &item.ItemCode, &item.Description, &item.Quantity, item.OrderID)
		items = append(items, item)
	}
	return items, nil
}

func (or *OrderRepository) UpdateOrder(updateOrder _entities.Order, id int) error {
	query := `UPDATE orders SET customer_name = $1, ordered_at = $2
	WHERE order_id = $3`
	ctx := context.Background()

	_, err := or.database.ExecContext(ctx, query, updateOrder.CustomerName, time.Now(), id)
	if err != nil {
		return errors.New("error update order")
	}
	return nil
}

func (or *OrderRepository) UpdateItem(updateItem _entities.Item, id int) error {
	query := `UPDATE items SET item_code = $1, description = $2, quantity = $3
	WHERE item_id = $4`
	ctx := context.Background()

	_, err := or.database.ExecContext(ctx, query, updateItem.ItemCode, updateItem.Description, updateItem.Quantity, id)
	if err != nil {
		return errors.New("error update item")
	}
	return nil
}

func (or *OrderRepository) DeleteItem(id int) error {
	query := `DELETE FROM items WHERE order_id = $1`
	ctx := context.Background()

	_, err := or.database.ExecContext(ctx, query, id)
	if err != nil {
		return errors.New("error delete item")
	}
	return nil
}

func (or *OrderRepository) DeleteOrder(id int) error {
	query := `DELETE FROM orders WHERE order_id = $1`
	ctx := context.Background()

	_, err := or.database.ExecContext(ctx, query, id)
	if err != nil {
		return errors.New("error delete item")
	}
	return nil
}
