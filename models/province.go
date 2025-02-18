package models

type Province struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Province string `json:"province"`
}
