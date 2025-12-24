package models

type Saha struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Name     string `json:"name"`
	Location string `json:"location"`
	Price    int    `json:"price"`
}
