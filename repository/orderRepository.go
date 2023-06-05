package repository

import (
	"database/sql"

	"github.com/alphadecodex/nomise-ragnarok/entity"
)

type OrderRepository struct {
	Db *sql.DB
}

func (r *OrderRepository) CreateOrder(order *entity.Order) error {
	_, err := r.Db.Exec("INSERT INTO orders (user_id, order_id, amount, order_status, currency, payment_method, date_of_order, feedback, is_refund_needed, refunded_amount, merchant_id) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
		order.UserID, order.OrderID, order.Amount, order.OrderStatus, order.Currency, order.PaymentMethod, order.DateOfOrder, order.Feedback, order.IsRefundNeeded, order.RefundedAmount, order.MerchantID)
	if err != nil {
		return err
	}

	return nil
}
