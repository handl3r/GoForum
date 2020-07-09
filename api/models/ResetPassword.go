package models

import (
	"github.com/jinzhu/gorm"
	"html"
	"strings"
)

type ResetPassword struct {
	gorm.Model
	Email string `json:"email" gorm:"size:100; not null"`
	Token string `json:"token" gorm:"size:255; not null"`
}

func (resetPassword *ResetPassword) Prepare() {
	resetPassword.Email = html.EscapeString(strings.TrimSpace(resetPassword.Email))
	resetPassword.Token = html.EscapeString(strings.TrimSpace(resetPassword.Token))
}

func (resetPassword *ResetPassword) SaveDetails(db *gorm.DB) (*ResetPassword, error) {
	var err error

	err = db.Debug().Model(&ResetPassword{}).Create(resetPassword).Error
	if err != nil {
		return &ResetPassword{}, err
	}
	return resetPassword, nil
}

func (resetPassword *ResetPassword) DeleteDetails(db *gorm.DB) (int64, error) {
	var err error
	db = db.Debug().Model(&ResetPassword{}).Where("id=?", resetPassword.ID).Take(resetPassword).Delete(resetPassword)
	if err = db.Error; err != nil {
		return 0, err
	}
	return db.RowsAffected, nil
}
