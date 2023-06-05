package entity

import "github.com/alphadecodex/nomise-ragnarok/enums"

type Order struct {
	UserID         string              `db:"user_id"`
	OrderID        string              `db:"order_id"`
	Amount         float32             `db:"amount"`
	OrderStatus    enums.OrderStatus   `db:"order_status"`
	Currency       string              `db:"currency"`
	PaymentMethod  enums.PaymentMethod `db:"payment_method"`
	DateOfOrder    string              `db:"date_of_order"`
	Feedback       string              `db:"feedback"`
	IsRefundNeeded bool                `db:"is_refund_needed"`
	RefundedAmount float32             `db:"refunded_amount"`
	MerchantID     string              `db:"merchant_id"`
}
