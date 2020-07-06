package models

import (
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

}

func (like *Like) DeleteLike(db *gorm.DB) (*Like, error) {

}

// Get likes of a post
func (like *Like) GetLikeInfo(db *gorm.DB, pid uint64) (*[]Like, error) {

}

// Delete likes of a user
func (like *Like) DeleteLikesOfUser(db *gorm.DB, uid uint32) (int64, error) {

}

// Deleted likes of a post
func DeleteLikesOfPost(db *gorm.DB, pid uint64) (int64, error) {

}
