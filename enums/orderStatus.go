package enums

type OrderStatus string

const (
	PROCESSING OrderStatus = "processing"
	COMPLETED  OrderStatus = "completed"
	FAILED     OrderStatus = "failed"
)
