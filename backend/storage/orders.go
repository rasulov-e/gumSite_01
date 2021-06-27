package storage

import (
	"github.com/lib/pq"
)

// Order describes a row of gumsite_orders from postgresql
type Order struct {
	ID                int    `json:"id"`
	Customer          string `json:"customer"`
	Address           string `json:"address"`
	Orders            []int  `json:"orders"`
	Cost              int    `json:"cost"`
	Payment           int    `json:"payment"`
	OrderDate         string `json:"order_date"`
	OrderShippingDate string `json:"order_shipping_date"`
}

// TestOrder return an Order struct that
// used for testing functions of orders.go
func TestOrder() Order {
	return Order{
		ID:                1,
		Customer:          "John Doe",
		Address:           "12th street",
		Orders:            []int{1, 2},
		Cost:              10,
		Payment:           20,
		OrderDate:         "10-10-2020",
		OrderShippingDate: "10-10-2020",
	}
}

// SelectAllOrders returns all rows from gumsite_order
// it might return error if it won't find anything
func SelectAllOrders(store Storage) (*[]Order, error) {
	rows, err := store.db.Query("SELECT id, customer, address, cost, payment, order_date, order_date_shipping, orders FROM gumsite_orders")
	if err != nil {
		return nil, err
	}

	var id int
	var customer string
	var address string
	var orders []int
	var cost int
	var payment int
	var orderDate string
	var orderShipping string
	o := []Order{}
	for rows.Next() {
		rows.Scan(
			&id,
			&customer,
			&address,
			&cost,
			&payment,
			&orderDate,
			&orderShipping,
			&orders,
		)
		o = append(o, Order{
			ID:                id,
			Customer:          customer,
			Address:           address,
			Orders:            orders,
			Cost:              cost,
			Payment:           payment,
			OrderDate:         orderDate,
			OrderShippingDate: orderShipping,
		})
	}
	return &o, nil
}

func SelectOrderByName(store Storage, name string) (*Order, error) {
	rows, err := store.db.Query("SELECT id, customer, address, cost, payment, order_date, order_date_shipping, orders FROM gumsite_orders WHERE customer=$1", name)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		o := Order{}
		rows.Scan(
			&o.ID,
			&o.Customer,
			&o.Address,
			&o.Cost,
			&o.Payment,
			&o.OrderDate,
			&o.OrderShippingDate,
			&o.Orders,
		)
		return &o, nil
	}
	return nil, err
}

// InsertOrder inserts a new row into gumsite_order
// by taking an argument of type Order
func InsertOrder(store Storage, order Order) error {
	if _, err := store.db.Query(
		"INSERT INTO gumsite_orders (customer, address, cost, payment, order_date, order_date_shipping, orders) VALUES($1,$2,$3,$4,$5,$6, $7)",
		order.Customer,
		order.Address,
		order.Cost,
		order.Payment,
		order.OrderDate,
		order.OrderShippingDate,
		pq.Array(order.Orders),
	); err != nil {
		return err
	}
	return nil
}

// DeleteOrderByName delete a row from gumsite_orders
func DeleteOrderByName(store Storage, name string) error {
	if _, err := store.db.Query("DELETE FROM gumsite_orders WHERE customer=$1", name); err != nil {
		return err
	}
	return nil
}
