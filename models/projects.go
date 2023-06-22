package models

import "gorm.io/gorm"

type Project struct {
	gorm.Model
	Name        string `gorm:"not null;unique_index"`
	Description string
	Users       []*User `gorm:"many2many:project_users;"`
}
