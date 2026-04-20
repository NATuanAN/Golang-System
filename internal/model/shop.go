package model

type Shop struct {
	ShopID      int     `json:"shop_id"               db:"shopid"      gorm:"column:shopid;primaryKey"`
	ShopName    string  `json:"shop_name"             db:"shopname"    gorm:"column:shopname"`
	Address     string  `json:"address"               db:"address"     gorm:"column:address"`
	Description *string `json:"description,omitempty" db:"description" gorm:"column:description"`
	Logo        *string `json:"logo,omitempty"        db:"logo"        gorm:"column:logo"`
	SellerID    int     `json:"seller_id"             db:"sellerid"    gorm:"column:sellerid"`
}

func (Shop) TableName() string {
	return "shop"
}
