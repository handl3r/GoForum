package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Post struct {
	ID        uint64    `json:"id" gorm:"primary_key;auto_increment"`
	Title     string    `json:"title" gorm:"size:255; not null; unique"`
	Content   string    `json:"content" gorm:"text; not null"`
	Author    User      `json:"author"`
	AuthorID  uint32    `json:"author_id" gorm:"not null"`
	UpdatedAt time.Time `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`
	CreatedAt time.Time `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP"`
}

func (post *Post) Prepare() {

}

func (post *Post) Validate(action string) error {

}

func (post *Post) SavePost(db *gorm.DB) (*Post, error) {

}

func (post *Post) FindAllPosts() (*[]Post, error) {

}

func (post *Post) FindPostByID(db *gorm.DB, pid uint64) (*Post, error) {

}

func (post *Post) UpdatePost(db *gorm.DB) (*User, error) {

}

func (post *Post) DeletePost(db *gorm.DB) (int64, error) {

}
