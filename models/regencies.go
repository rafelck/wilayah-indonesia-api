package models

type Regencie struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	Province_ID string `json:"province_id"`
	Regencie    string `json:"regencie"`
}

func (Regencie) TableName() string {
	return "regencies"
}
