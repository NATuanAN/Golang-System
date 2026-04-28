package model

type AccountType string

const (
	AccountTypeAdmin  AccountType = "Admin"
	AccountTypeBuyer  AccountType = "Buyer"
	AccountTypeSeller AccountType = "Seller"
)

type User struct {
	ID          uint        `json:"id"           db:"userid"      gorm:"column:userid;primaryKey;autoIncrement"`
	Name        string      `json:"name"         db:"name"        gorm:"column:name" binding:"required"`
	Email       string      `json:"email"        db:"email"       gorm:"column:email" binding:"required"`
	Password    string      `json:"-"            db:"password"    gorm:"column:password" binding:"required,min=6"`
	PhoneNumber string      `json:"phone_number" db:"phonenumber" gorm:"column:phonenumber"`
	AccountType AccountType `json:"account_type" db:"accounttype" gorm:"column:accounttype"`
	Active      string      `json:"active"       db:"active"      gorm:"column:active"`
}

func (User) TableName() string {
	return "users_table"
}
