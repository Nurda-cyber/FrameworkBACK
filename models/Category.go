package models

type Category struct {
	ID   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
	Toys []Toy  `json:"toys" gorm:"foreignKey:CategoryID"`
}
