package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name      string `gorm:"not null;unique_index"`
	LastName  string
	Roles     []*Role        `gorm:"many2many:user_Roles;"`
	Projects  []*Project     `gorm:"many2many:user_Projects;"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt"`
}

// type User struct {
// 	gorm.Model
// 	Name string
//    }
//   // es igual a
//   type User struct {
// 	ID         uint            `gorm:"primaryKey"`
// 	 CreatedAt time.Time
// 	UpdatedAt time.Time
// 	DeletedAt gorm.DeletedAt `gorm:"index"`
// 	 Name string
//    }
