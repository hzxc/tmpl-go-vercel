package model

type Person struct {
	Model
	Name string `gorm:"type:varchar(50);unique;notNull" json:"name"`
}
