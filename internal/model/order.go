package model

import "time"

type PaymentMethod string
type OrderStatus string

const (
	PaymentMethodMomo  PaymentMethod = "MOMO"
	PaymentMethodVNPay PaymentMethod = "VNPAY"
)

const (
	OrderStatusOrdered        OrderStatus = "ORDERED"
	OrderStatusPaymentPending OrderStatus = "PAYMENT_PENDING"
	OrderStatusPaid           OrderStatus = "PAID"
	OrderStatusFailed         OrderStatus = "FAILED"
)

type Order struct {
	OrderID         string        `json:"order_id"          db:"orderid"         gorm:"column:orderid;primaryKey"`
	ShippingAddress string        `json:"shipping_address"  db:"shippingaddress" gorm:"column:shippingaddress"`
	PaymentMethod   PaymentMethod `json:"payment_method"    db:"paymentmethod"   gorm:"column:paymentmethod"`
	CreatedAt       time.Time     `json:"created_at"        db:"created_at"      gorm:"column:created_at"`
	UpdatedAt       time.Time     `json:"updated_at"        db:"updated_at"      gorm:"column:updated_at"`
	Status          OrderStatus   `json:"status"            db:"status"          gorm:"column:status"`
	BuyerID         int           `json:"buyer_id"          db:"buyerid"         gorm:"column:buyerid"`
}

func (Order) TableName() string {
	return "order_table"
}

type OrderItem struct {
	OrderID   string   `json:"order_id"             db:"orderid"   gorm:"column:orderid;primaryKey"`
	ProductID int      `json:"product_id"           db:"productid" gorm:"column:productid;primaryKey"`
	Quantity  int      `json:"quantity"             db:"quantity"  gorm:"column:quantity"`
	UnitPrice *float64 `json:"unit_price,omitempty" db:"unitprice" gorm:"column:unitprice"`
}

func (OrderItem) TableName() string {
	return "orderitem"
}
