package model

import "time"

type Review struct {
	ReviewID   int       `json:"review_id"         db:"reviewid"   gorm:"column:reviewid;primaryKey"`
	Content    *string   `json:"content,omitempty" db:"content"    gorm:"column:content"`
	Rating     *int      `json:"rating,omitempty"  db:"rating"     gorm:"column:rating"`
	DatePosted time.Time `json:"date_posted"       db:"dateposted" gorm:"column:dateposted"`
	OrderID    string    `json:"order_id"          db:"orderid"    gorm:"column:orderid"`
}

func (Review) TableName() string {
	return "review"
}
