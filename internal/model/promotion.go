package model

import "time"

type PromotionType string

const (
	PromotionTypeShop    PromotionType = "Shop"
	PromotionTypeProduct PromotionType = "Product"
)

type Promotion struct {
	PromotionID int           `json:"promotion_id"          db:"promotionid" gorm:"column:promotionid;primaryKey"`
	Name        string        `json:"name"                  db:"name"        gorm:"column:name"`
	Type        PromotionType `json:"type"                  db:"type"        gorm:"column:type"`
	Value       float64       `json:"value"                 db:"value"       gorm:"column:value"`
	StartAt     time.Time     `json:"start_at"              db:"startat"     gorm:"column:startat"`
	EndAt       time.Time     `json:"end_at"                db:"endat"       gorm:"column:endat"`
	ShopID      *int          `json:"shop_id,omitempty"     db:"shopid"      gorm:"column:shopid"`
	ProductID   *int          `json:"product_id,omitempty"  db:"productid"   gorm:"column:productid"`
}

func (Promotion) TableName() string {
	return "promotion"
}
