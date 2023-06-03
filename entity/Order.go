package entity

import "github.com/alphadecodex/nomise-ragnarok/enums"

type Order struct {
	userId         string
	orderId        string
	amount         float32
	orderStatus    enums.OrderStatus
	currency       string
	paymentMethod  enums.PaymentMethod
	dateOfOrder    string
	feedback       string
	isRefundNeeded bool
	refundedAmount float32
	merchantId     string
}
