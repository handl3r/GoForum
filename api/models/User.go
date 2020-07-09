package models

import (
	"errors"
	"github.com/badoux/checkmail"
	"github.com/handl3r/GoForum/api/security"
	"github.com/jinzhu/gorm"
	"html"
	"log"
	"os"
	"strings"
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
	user.Email = html.EscapeString(strings.TrimSpace(user.Email))
	user.Username = html.EscapeString(strings.TrimSpace(user.Username))
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
}

func (user *User) BeforeSave() error {
	hashedPassword, err := security.Hash(user.Password)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	return nil
}

func (user *User) AfterFind() error {
	if user.AvatarPath != "" {
		user.AvatarPath = os.Getenv("DO_SPACES_URL") + user.AvatarPath
	}
	return nil
}

func (user *User) Validate(action string) map[string]string {
	errorMessages := make(map[string]string)
	var err error
	switch strings.ToLower(action) {
	case "update":
		if user.Email == "" {
			errorMessages["require_email"] = "require email"
		}
		if user.Email != "" {
			err = checkmail.ValidateFormat(user.Email)
			if err != nil {
				errorMessages["invalid_email"] = "invalid mail"
			}
		}
	case "login":
		if user.Email == "" {
			errorMessages["require_email"] = "require email"
		}
		if user.Email != "" {
			err = checkmail.ValidateFormat(user.Email)
			if err != nil {
				errorMessages["invalid_email"] = "invalid email"
			}
		}
		if user.Password == "" {
			errorMessages["require_password"] = "require password"
		}
		if user.Password != "" {
			errorMessages["require_password"] = "require password"
		}
	case "forgotpassword":
		if user.Email == "" {
			errorMessages["require_password"] = "require password"
		}
		if user.Email != "" {
			err = checkmail.ValidateFormat(user.Email)
			if err != nil {
				errorMessages["invalid_email"] = "invalid email"
			}
		}
	default:
		if user.Username == "" {
			errorMessages["require_username"] = "require username"
		}
		if user.Email == "" {
			errorMessages["require_email"] = "require email"
		}
		if user.Email != "" {
			err = checkmail.ValidateFormat(user.Email)
			if err != nil {
				errorMessages["invalid_email"] = "invalid email"
			}
		}
		if user.Password == "" {
			errorMessages["require_password"] = "require password"
		}
		if user.Password != "" && len(user.Password) < 6 {
			errorMessages["invalid_password"] = "invalid password"
		}
	}
	return errorMessages
}

func (user *User) SaveUser(db *gorm.DB) (*User, error) {
	var err error

	err = db.Debug().Model(&User{}).Create(&user).Error
	if err != nil {
		return &User{}, err
	}
	return user, nil
}

func (user *User) FindAllUsers(db *gorm.DB) (*[]User, error) {
	var err error
	var users []User

	err = db.Debug().Model(&User{}).Limit(100).Find(&users).Error
	if err != nil {
		return &[]User{}, err
	}
	return &users, nil
}

func (user *User) FindUserByID(db *gorm.DB, uid uint32) (*User, error) {
	var err error

	err = db.Debug().Model(&User{}).Where("id=?", uid).Take(&user).Error
	if err != nil {
		return &User{}, errors.New("user not found")
	}
	return user, nil
}

func (user *User) UpdateUser(db *gorm.DB, uid uint32) (*User, error) {
	var err error
	if user.Password != "" {
		err = user.BeforeSave()
	}
	if err != nil {
		log.Fatal(err)
	}

	db = db.Debug().Model(&User{}).Where("id=?", uid).Take(&User{}).UpdateColumns(
		map[string]interface{}{
			"email":      user.Email,
			"username":   user.Username,
			"password":   user.Password,
			"updated_at": time.Now(),
		})
	if err = db.Error; err != nil {
		return &User{}, err
	}

	err = db.Debug().Model(&User{}).Where("id=?", uid).Take(&user).Error
	if err != nil {
		return &User{}, err
	}
	return user, nil
}

func (user *User) UpdateUserAvatar(db *gorm.DB, uid uint32) (*User, error) {
	var err error

	db = db.Debug().Model(&User{}).Where("id=?", uid).Take(&User{}).UpdateColumns(
		map[string]interface{}{
			"avatar_path": user.AvatarPath,
			"updated_at":  time.Now(),
		})
	if err = db.Error; err != nil {
		return &User{}, err
	}

	err = db.Debug().Model(&User{}).Where("id=?", uid).Take(&user).Error
	if err != nil {
		return &User{}, err
	}
	return user, nil
}

func (user *User) DeleteUser(db *gorm.DB, uid uint32) (int64, error) {
	var err error
	err = db.Debug().Model(&User{}).Delete(&user).Error
	if err != nil {
		return 0, err
	}

	return db.RowsAffected, nil
}

func (user *User) UpdatePassword(db *gorm.DB) error {
	var err error
	err = user.BeforeSave()
	if err != nil {
		log.Fatal(err)
	}

	db = db.Debug().Model(&User{}).Where("email=?", user.Email).Take(&User{}).UpdateColumns(
		map[string]interface{}{
			"password":   user.Password,
			"updated_at": time.Now(),
		})
	if err = db.Error; err != nil {
		return err
	}
	return nil
}
