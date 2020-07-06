package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Comment struct {
	ID        uint64    `json:"id" gorm:"primary_key; auto_increment"`
	UserID    uint32    `json:"user_id" gorm:"not null"`
	PostID    uint64    `json:"post_id" gorm:"not null"`
	Author    User      `json:"author"`
	Body      string    `json:"body" gorm:"text; not null"`
	UpdatedAt time.Time `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`
	CreatedAt time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
}

func (comment *Comment) Prepare() {

}

func (comment *Comment) Validate(action string) error {

}

func (comment *Comment) SaveComment(db *gorm.DB) (*Comment, error) {

}

func (comment *Comment) GetCommentsByPost(db *gorm.DB, pid uint64) (*[]Comment, error) {

}

func (comment *Comment) UpdateComment(db *gorm.DB) (*Comment, error) {

}

func (comment *Comment) DeleteComment(db *gorm.DB, cid uint64) (int64, error) {

}

func (comment *Comment) DeleteCommentsByPost(db *gorm.DB, pid uint64) (int64 , error) {

}

func (comment *Comment) DeleteCommentsByUser(db *gorm.DB, uid uint32) (int64, error) {

}
