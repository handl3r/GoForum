package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type User struct {
	ID         uint32    `json:"id" gorm:"primary_key;auto_increment"`
	Username   string    `json:"username" gorm:"size:255; not null; unique"`
	Email      string    `json:"email" gorm:"size:100; not null; unique"`
	Password   string    `json:"password" gorm:"size:100; not null"`
	AvatarPath string    `json:"avatar_path" gorm:"size:100; null"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`
	CreatedAt  time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
}

func (user *User) Prepare() {

}

func (user *User) BeforeSave() error {

}

func (user *User) AfterFind() error {

}

func (user *User) Validate(action string) map[string]string {

}

func (user *User) SaveUser(db *gorm.DB) (*User, error) {

}

func (user *User) FindAllUsers(db *gorm.DB) (*[]User, error) {

}

func (user *User) FindUserByID(db *gorm.DB, uid uint32) (*User, error) {

}

func (user *User) UpdateUser(db *gorm.DB) (*User, error) {

}

func (user *User) UpdateUserAvatar(db *gorm.DB, uid uint32) (*User, error) {

}

func (user *User) DeleteUser(db *gorm.DB, uid uint32) (int64, error) {

}
