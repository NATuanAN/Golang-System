package model

type Wishlist struct {
	WishlistID int `json:"wishlist_id" db:"wishlistid" gorm:"column:wishlistid;primaryKey"`
	BuyerID    int `json:"buyer_id"    db:"buyerid"    gorm:"column:buyerid"`
}

func (Wishlist) TableName() string {
	return "wishlist"
}

type WishlistItem struct {
	WishlistID int `json:"wishlist_id" db:"wishlistid" gorm:"column:wishlistid;primaryKey"`
	ProductID  int `json:"product_id"  db:"productid"  gorm:"column:productid;primaryKey"`
}

func (WishlistItem) TableName() string {
	return "wishlistitem"
}
