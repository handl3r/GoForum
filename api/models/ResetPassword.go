package models

import (
	"github.com/jinzhu/gorm"
)

type ResetPassword struct {
	gorm.Model
	Email string `json:"email" gorm:"size:100; not null"`
	Token string `json:"token" gorm:"size:255; not null"`
}

func (resetPassword *ResetPassword) SaveDetails(db *gorm.DB) (*ResetPassword, error) {

}

func (resetPassword *ResetPassword) DeleteDetails(db *gorm.DB) (int64, error) {

}
