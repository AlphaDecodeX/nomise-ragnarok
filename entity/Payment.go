package entity

import "github.com/alphadecodex/nomise-ragnarok/enums"

type Payment struct {
	paymentId      string
	orderId        string
	amount         float32
	paymentStatus  enums.PaymentMethod
	currency       string
	paymentMethod  string
	dateOfOrder    string
	isRefundNeeded bool
	refundedAmount float32
	merchantId     string
}
