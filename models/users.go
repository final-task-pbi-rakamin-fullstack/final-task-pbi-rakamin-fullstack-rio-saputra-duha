package models

import (
	"github.com/jaswdr/faker/v2"
	"gorm.io/gorm"
	"time"
)

type Users struct {
	Uuid      string `json:"uuid" gorm:"type:varchar(255);primaryKey"`
	Username  string `json:"username" validate:"required"`
	Email     string `json:"email" validate:"email,required" gorm:"unique"`
	Password  string `json:"password" validate:"required,min=6" gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (user *Users) BeforeCreate(tx *gorm.DB) (err error) {
	faker := faker.New()
	user.Uuid = faker.UUID().V4()
	return
}
