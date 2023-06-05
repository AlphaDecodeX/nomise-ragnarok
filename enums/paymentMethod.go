package enums

type PaymentMethod string

const (
	UPI           PaymentMethod  = "upi"
	DEBIT_CARD    PaymentMethod  = "debit_card"
	CREDIT_CARD   PaymentMethod  = "credit_card"
	BANK_TRANSFER PaymentMethod = "bank_transfer"
)
