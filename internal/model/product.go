package model

type Product struct {
	ProductID     int     `json:"product_id"            db:"productid"     gorm:"column:productid;primaryKey"`
	ProductName   string  `json:"product_name"          db:"productname"   gorm:"column:productname"`
	Price         float64 `json:"price"                 db:"price"         gorm:"column:price"`
	Description   *string `json:"description,omitempty" db:"description"   gorm:"column:description"`
	ImageURL      *string `json:"image_url,omitempty"   db:"image_url"     gorm:"column:image_url"`
	StockQuantity int     `json:"stock_quantity"        db:"stockquantity" gorm:"column:stockquantity"`
	ShopID        int     `json:"shop_id"               db:"shopid"        gorm:"column:shopid"`
}

func (Product) TableName() string {
	return "product"
}
