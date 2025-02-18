package models

type District struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	Regencie_Id string `json:"regencie_id"`
	District    string `json:"district"`
}

func (District) TableName() string {
	return "district"
}
