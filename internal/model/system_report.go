package model

type SystemReport struct {
	ReportID         int     `json:"report_id"                   db:"reportid"         gorm:"column:reportid;primaryKey"`
	Title            string  `json:"title"                       db:"title"            gorm:"column:title"`
	Description      *string `json:"description,omitempty"       db:"description"      gorm:"column:description"`
	Status           string  `json:"status"                      db:"status"           gorm:"column:status"`
	ReportedByUserID int     `json:"reported_by_user_id"         db:"reportedbyuserid" gorm:"column:reportedbyuserid"`
}

func (SystemReport) TableName() string {
	return "systemreport"
}
