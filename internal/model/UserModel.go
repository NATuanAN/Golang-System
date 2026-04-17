package model

type AccountType string

const (
	AccountTypeAdmin   AccountType = "admin"
	AccountTypeUser    AccountType = "user"
	AccountTypePremium AccountType = "premium"
)

type User struct {
	ID          uint        `json:"id"           db:"userid"      gorm:"column:userid;primaryKey"`
	Name        string      `json:"name"         db:"name"        gorm:"column:name"`
	Email       string      `json:"email"        db:"email"       gorm:"column:email"`
	Password    string      `json:"-"            db:"password"    gorm:"column:password"`
	PhoneNumber string      `json:"phone_number" db:"phonenumber" gorm:"column:phonenumber"`
	AccountType AccountType `json:"account_type" db:"accounttype" gorm:"column:accounttype"`
	Active      string      `json:"active" db:"active" gorm:"column:active"`
}

func (User) TableName() string {
	return "users_table"
}
