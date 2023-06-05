package entity

import "github.com/alphadecodex/nomise-ragnarok/enums"

type Order struct {
	UserID         string              `json:"user_id" db:"user_id"`
	OrderID        string              `json:"order_id" db:"order_id"`
	Amount         float32             `json:"amount" db:"amount"`
	OrderStatus    enums.OrderStatus   `json:"order_status" db:"order_status"`
	Currency       string              `json:"currency" db:"currency"`
	PaymentMethod  enums.PaymentMethod `json:"payment_method" db:"payment_method"`
	DateOfOrder    string              `json:"date_of_order" db:"date_of_order"`
	Feedback       string              `json:"feedback" db:"feedback"`
	IsRefundNeeded bool                `json:"is_refund_needed" db:"is_refund_needed"`
	RefundedAmount float32             `json:"refunded_amount" db:"refunded_amount"`
	MerchantID     string              `json:"merchant_id" db:"merchant_id"`
}
