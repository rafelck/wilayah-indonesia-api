package models

type Villages struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	District_ID string `json:"district_id"`
	Village     string `json:"village"`
}

func (Villages) TableName() string {
	return "villages"
}
