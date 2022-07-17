package model

type Project struct {
	Model
	Name         string `gorm:"type:varchar(50);unique;notNull" json:"name,omitempty"`
	Pin          bool   `gorm:"type:tinyint(1);default:0" json:"pin,omitempty"`
	PersonId     uint   `gorm:"notNull" json:"person_id,omitempty"`
	Organization string `gorm:"type:varchar(100)" json:"organization,omitempty"`
	Description  string `gorm:"type:varchar(100)" json:"description,omitempty"`
}
