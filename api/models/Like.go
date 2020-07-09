package models

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

type Like struct {
	ID        uint64    `json:"id" gorm:"primary_key;auto_increment"`
	UserID    uint32    `json:"user_id" gorm:"not null"`
	PostID    uint64    `json:"post_id" gorm:"not null"`
	UpdatedAt time.Time `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`
	CreatedAt time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
}

func (like *Like) SaveLike(db *gorm.DB) (*Like, error) {
	var err error

	err = db.Debug().Model(&Like{}).Where("user_id=? and post_id=?", like.UserID, like.PostID).Take(&like).Error
	if err != nil {
		if err.Error() == "record not found" {
			err = db.Debug().Model(&Like{}).Create(&like).Error
			if err != nil {
				return &Like{}, err
			}
			return like, nil
		}
	} else {
		return &Like{}, errors.New("doubled like")
	}
	return like, nil
}

func (like *Like) DeleteLike(db *gorm.DB) (*Like, error) {
	var err error
	var deletedLike *Like

	err = db.Debug().Model(&Like{}).Where("id=?", like.ID).Take(&like).Error
	if err != nil {
		return &Like{}, err
	} else {
		deletedLike = like
		err = db.Debug().Model(&Like{}).Where("id=?", like.ID).Take(&Like{}).Delete(&Like{}).Error
		if err != nil {
			fmt.Printf("can not delete like: %v\n", err)
			return &Like{}, err
		}
	}
	return deletedLike, nil
}

// Get likes of a post
func (like *Like) GetLikeInfo(db *gorm.DB, pid uint64) (*[]Like, error) {
	var err error
	var likes []Like

	err = db.Debug().Model(&Like{}).Where("post_id=?", pid).Find(&likes).Error
	if err != nil {
		return &[]Like{}, err
	}
	return &likes, nil
}

// Delete likes of a user
func (like *Like) DeleteLikesOfUser(db *gorm.DB, uid uint32) (int64, error) {
	var err error
	var likes []Like

	err = db.Debug().Model(&Like{}).Where("user_id=?", uid).Find(&likes).Delete(&likes).Error
	if err != nil {
		return 0, err
	}
	return db.RowsAffected, nil
}

// Deleted likes of a post
func DeleteLikesOfPost(db *gorm.DB, pid uint64) (int64, error) {
	var err error
	var likes []Like

	err = db.Debug().Model(&Like{}).Where("post_id=?", pid).Find(&likes).Delete(&likes).Error
	if err != nil {
		return 0, err
	}
	return db.RowsAffected, nil
}
